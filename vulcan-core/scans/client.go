package scans

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/goadesign/goa/uuid"

	coreclient "github.com/adevinta/vulcan-core-cli/vulcan-core/client"
	report "github.com/adevinta/vulcan-report"
)

// ErrorUnexpectedStatus is returned when one of calls made by the concurrent
// client to one endpoint of vulcan core returns an unexpected http status.
type ErrorUnexpectedStatus struct {
	Expected int
	Got      int
}

func (e *ErrorUnexpectedStatus) Error() string {
	return fmt.Sprintf("unexpected http status, expected %d, got %d", e.Expected, e.Got)
}

var (
	// ErrCheckNotFinished is returned by the client when a check of a scan does
	// not have report because is not FINISHED.
	ErrCheckNotFinished = errors.New("Check status is not FINISHED")
)

// CheckData contains the info regarding a check of a scan.
type CheckData struct {
	coreclient.Checkdata
	Report ReportData
}

// ReportData contanis either the info regarding the report of a check or the error
// that the client got trying to dowload it.
type ReportData struct {
	Err error
	report.Report
}

// ScanData define the data of a scan retrieved by the scans client.
type ScanData struct {
	CreationDate string
	ChecksData   map[string]CheckData
}

// CheckReportClient defines the services needed by the concurrent client.
type CheckReportClient interface {
	GetReport(url string) (*report.Report, error)
}

// Client allows to get the reports of a scan concurrently.
type Client struct {
	coreC   *coreclient.Client
	report  CheckReportClient
	workers int
}

// NewClient returns a new client scans client.
func NewClient(coreClient *coreclient.Client, reportClient CheckReportClient, workers int) *Client {
	return &Client{coreClient, reportClient, workers}
}

// Data downloads all the reports of a given scan using, at most, the
// specified number of goroutines in the client.
func (c *Client) Data(ctx context.Context, ID string) (ScanData, error) {
	scanID, err := uuid.FromString(ID)
	if err != nil {
		return ScanData{}, err
	}

	resp, err := c.coreC.ShowScans(ctx, coreclient.ShowScansPath(scanID))
	if err != nil {
		return ScanData{}, err
	}

	if resp.StatusCode != http.StatusOK {
		err = &ErrorUnexpectedStatus{http.StatusOK, resp.StatusCode}
		return ScanData{}, fmt.Errorf("error querying scan in vulcan core %w", err)
	}

	/*scan, err := c.coreC.DecodeScandata(resp)
	if err != nil {
		return ScanData{}, err
	}*/

	resp, err = c.coreC.ChecksScans(ctx, coreclient.ChecksScansPath(scanID))
	if resp.StatusCode != http.StatusOK {
		err = &ErrorUnexpectedStatus{http.StatusOK, resp.StatusCode}
		return ScanData{}, fmt.Errorf("error querying scan checks in vulcan core %w", err)
	}
	checksData, err := c.coreC.DecodeChecks(resp)
	if err != nil {
		return ScanData{}, err
	}
	if checksData == nil {
		return ScanData{}, errors.New("unexpected nil getting checks for a scan")
	}

	rs := &reportResults{}
	rs.reports = make([]reportResult, 0, len(checksData.Checks))
	wg := &sync.WaitGroup{}
	checks := make(chan check, c.workers)

	// Spawn a goroutine to spawn to the channel that the workers will read from.
	wg.Add(1)
	go func() {
		defer wg.Done()
	LOOP:
		for _, c := range checksData.Checks {
			c := c
			select {
			case <-ctx.Done():
				break LOOP
			default:
				if c == nil {
					// Skip the a nil check.
					continue
				}
				checks <- check{*c, ID}
			}
		}
		close(checks)
	}()

	// Spawn the workers.
	for i := 0; i < c.workers; i++ {
		wg.Add(1)
		go func() {
			cctx, _ := context.WithCancel(ctx)
			defer wg.Done()
			downloadCheckReport(cctx.Done(), checks, rs, c.report)
		}()
	}
	wg.Wait()
	if ctx.Err() != nil {
		// The download process has been canceled.
		return ScanData{}, ctx.Err()
	}
	var results = make(map[string]CheckData)
	for _, r := range rs.reports {
		cid := r.CheckData.ID.String()
		// We add the check data in any case.
		data := CheckData{Checkdata: r.CheckData}
		// If we got an error downloading a report we store that error.
		if r.Err != nil {
			data.Report.Err = r.Err
		} else {
			// In this case the report of the check cannot be nil.
			if r.R == nil {
				return ScanData{}, errors.New("unexpected nil downloading check report")
			}
			data.Report.Report = *r.R
		}
		results[cid] = data
	}
	data := ScanData{
		ChecksData: results,
	}
	return data, nil
}

type reportResult struct {
	CheckData coreclient.Checkdata
	R         *report.Report
	Err       error
}

type reportResults struct {
	reports []reportResult
	sync.Mutex
}

func (r *reportResults) Add(result reportResult) {
	r.Lock()
	defer r.Unlock()
	r.reports = append(r.reports, result)
}

type check struct {
	Check  coreclient.Checkdata
	ScanID string
}

func downloadCheckReport(done <-chan struct{}, checks <-chan check, results *reportResults, rclient CheckReportClient) {
LOOP:
	for c := range checks {
		select {
		case <-done:
			break LOOP
		default:
			checkID := c.Check.ID.String()
			if c.Check.Status != "FINISHED" {
				err := fmt.Errorf("%w, check status %s", ErrCheckNotFinished, c.Check.Status)
				r := reportResult{c.Check, nil, err}
				results.Add(r)
				continue
			}
			reportURL := *c.Check.Report
			report, err := rclient.GetReport(reportURL)
			if err != nil {
				err = fmt.Errorf("getting results for check-id: %s, %w", checkID, err)
			}
			r := reportResult{c.Check, report, err}
			results.Add(r)
		}
	}
}
