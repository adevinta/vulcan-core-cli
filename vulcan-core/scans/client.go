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

var (

	// ErrUnexpectedStatus is returned when a call to the vulcan core api is
	// made and the returned http status does not match the expected one.
	ErrUnexpectedStatus = errors.New("unexpected response status")
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

// ScanData define the data of a scan retrieved by the scans client.
type ScanData struct {
	Date    string
	Reports []report.Report
}

// ReportClient defines the services needed by the concurrent client.
type ReportClient interface {
	GetReport(date, scanID, checkID string) (*report.Report, error)
}

// Client allows to get the reports of a scan concurrently.
type Client struct {
	coreC   *coreclient.Client
	report  ReportClient
	workers int
}

// NewClient returns a new client scans client.
func NewClient(coreClient *coreclient.Client, reportClient ReportClient, workers int) *Client {
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

	scan, err := c.coreC.DecodeScan(resp)
	if err != nil {
		return ScanData{}, err
	}

	// We need the date of the scan in order to build the path for the reports
	// of the checks belonging to it.

	scanDate := scan.Scan.CreatedAt.Format("2006-01-02")

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
			select {
			case <-ctx.Done():
				break LOOP
			default:
				checks <- check{ID: c.ID.String(), ScanID: ID, Date: scanDate}
			}
		}
		close(checks)
	}()

	// Spawn the workers.
	wg.Add(1)
	go func() {
		cctx, _ := context.WithCancel(ctx)
		defer wg.Done()
		downloadCheckReport(wg, cctx.Done(), checks, rs, c.report)
	}()
	wg.Wait()
	if ctx.Err() != nil {
		// The download process has been canceled.
		return ScanData{}, ctx.Err()
	}
	var result []report.Report
	for _, r := range rs.reports {
		if r.err != nil {
			return ScanData{}, r.err
		}
		if r.R == nil {
			return ScanData{}, errors.New("unexpected nil downloading check report")
		}
		result = append(result, *r.R)
	}
	data := ScanData{
		Date:    scanDate,
		Reports: result,
	}
	return data, nil
}

type reportResult struct {
	R   *report.Report
	err error
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
	ID     string
	ScanID string
	Date   string
}

func downloadCheckReport(wg *sync.WaitGroup, done <-chan struct{}, checks <-chan check, results *reportResults, rclient ReportClient) {
LOOP:
	for c := range checks {
		select {
		case <-done:
			break LOOP
		default:
			report, err := rclient.GetReport(c.Date, c.ScanID, c.ID)
			if err != nil {
				err = fmt.Errorf("getting results for check-id: %s, %w", c.ID, err)
			}
			r := reportResult{report, err}
			results.Add(r)
		}
	}
}
