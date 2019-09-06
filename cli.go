package cli

import (
	"bufio"
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/goadesign/goa/uuid"

	"github.com/adevinta/vulcan-core-cli/vulcan-core/client"
)

// ScanStatus contains the list of possible status of a check, and indicates
// whether the status is a final state or not.
var ScanStatus = map[string]bool{
	"ABORTED":   true,
	"TIMEOUT":   true,
	"CLEANUP":   false,
	"CREATED":   false,
	"FAILED":    true,
	"FINISHED":  true,
	"KILLED":    true,
	"MALFORMED": true,
	"QUEUED":    false,
	"RUNNING":   false,
}

const fileScanAPIPath = "/v1/filescan"

// Scan defines a set of checks that have been executed together.
type Scan struct {
	ID           uuid.UUID
	ExpectedSize int
	StartTime    time.Time
	Stats        *Stats
}

func newScan(id uuid.UUID, e int) *Scan {
	return &Scan{
		ID:           id,
		ExpectedSize: e,
		StartTime:    time.Now(),
		Stats:        newStats(),
	}
}

func LoadState(stateFile string) (*Scan, error) {
	b, err := ioutil.ReadFile(stateFile)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)

	var scan Scan
	err = dec.Decode(&scan)
	return &scan, err
}

func (s *Scan) SaveState(stateFile string) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(s); err != nil {
		return err
	}

	f, err := os.Create(stateFile)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

func (s *Scan) IsFinished() bool {
	// TODO: if one of the checks that were initially sent when creating the scan
	// fails to be created, the number will never coincide. Probably we need to
	// implement a Scan status in the Persistence.
	var size int
	for status, total := range s.Stats.GroupByStatus {
		if !ScanStatus[status] {
			return false
		}

		size += total
	}

	return size == s.ExpectedSize
}

func (s *Scan) updateStats(cliStats *client.Stats) error {
	gStatus := make(map[string]int)

	for _, stat := range cliStats.Checks {
		if stat == nil {
			return errors.New("no stat reported")
		}

		gStatus[stat.Status] = stat.Total
	}

	s.Stats.GroupByStatus = gStatus

	return nil
}

func (s *Scan) ToString() string {
	return toString(s)
}

type Stats struct {
	// GroupByStatus keys are different status (e.g. FINISHED), and values
	// are the total number of checks in this state.
	GroupByStatus map[string]int
}

func newStats() *Stats {
	gStatus := make(map[string]int)

	return &Stats{
		GroupByStatus: gStatus,
	}
}

func (sts *Stats) PrintToFile(fp string) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	sts.Print(w)
	w.Flush()

	return nil
}

func (sts *Stats) Print(w io.Writer) {
	var keys []string
	for k, _ := range sts.GroupByStatus {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Fprintln(w, "STATUS COUNT")
	fmt.Fprintln(w, "------------")

	for _, k := range keys {
		fmt.Fprintf(w, "%v: %v\n", k, sts.GroupByStatus[k])
	}
}

// Checks is a set of checktypes with the associated Assets to be checked.
type Checks map[string][]Asset

func (c Checks) ToString() string {
	return toString(c)
}

// Asset defines the asset where a concrete checktype will be executed
// with a set of options and a concrete queue.
type Asset struct {
	Target  string
	Options string
	QueueID string
}

type CLI struct {
	logger goa.LogAdapter
	c      *client.Client
	ctx    context.Context
}

func NewCLI(scheme, host string) (*CLI, error) {
	httpClient := newHTTPClient()
	c := client.New(goaclient.HTTPClientDoer(httpClient))
	c.Client.Scheme = scheme
	c.Client.Host = host

	var buf bytes.Buffer
	logger := goa.NewLogger(log.New(&buf, "", log.LstdFlags))

	ctx := goa.WithLogger(context.Background(), logger)

	return &CLI{
		logger: logger,
		c:      c,
		ctx:    ctx,
	}, nil
}

// RunScan starts execution, or simulates, the execution of a scan composed by a set of checks.
// Uses the rest api or the multipart file upload api depending on the params.
func (cli *CLI) RunScan(checks Checks, dryRun string, useMultipart bool) (*Scan, error) {
	var checksPayload []*client.CheckPayload

	for checktypeName, assets := range checks {
		// NOTE: this is a big gotcha. If you use &checktypeName directly instead,
		// the chektypes will be the same for all the checks in checksPayload.
		checktypeName := checktypeName

		for _, asset := range assets {
			// NOTE: fix same gotcha specified above.
			opts := asset.Options
			var queueID *uuid.UUID
			if asset.QueueID != "" {
				id, err := uuid.FromString(asset.QueueID)
				if err != nil {
					return nil, fmt.Errorf("error wrong queue id, %v", err)
				}
				queueID = &id
			}
			checkPayload := &client.CheckPayload{
				Check: &client.CheckData{
					ChecktypeName: &checktypeName,
					Target:        asset.Target,
					Options:       &opts,
					JobqueueID:    queueID,
				},
			}
			checksPayload = append(checksPayload, checkPayload)
		}
	}

	return cli.runScan(checksPayload, dryRun, useMultipart)
}

func (cli *CLI) runScan(checks []*client.CheckPayload, dryRun string, useFileScan bool) (*Scan, error) {

	if dryRun != "" && useFileScan {
		return nil, errors.New("both dryRun and useFileScan can not be specified")
	}

	payload := &client.ScanPayload{
		Scan: &client.ScanChecksPayload{
			Checks: checks,
		},
	}
	var err error
	var scanFile *os.File
	if dryRun != "" || useFileScan {
		content, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}

		if dryRun != "" {
			scanFile, err = os.Create(dryRun)
		} else {
			scanFile, err = ioutil.TempFile("", "scan_")
		}
		if err != nil {
			return nil, err
		}
		defer scanFile.Close()
		if _, err = scanFile.Write(content); err != nil {
			return nil, err
		}
		if dryRun != "" {
			return nil, nil
		}
	}
	var scan *client.Scan
	if useFileScan {
		scan, err = cli.startScanByMultipartUpload(scanFile.Name())
	} else {
		scan, err = cli.startScanByRestAPI(payload)
	}
	if err != nil {
		return nil, err
	}
	if err := scan.Validate(); err != nil {
		return nil, err
	}

	s := newScan(scan.Scan.ID, len(checks))
	return s, nil
}

func (cli *CLI) startScanByRestAPI(payload *client.ScanPayload) (*client.Scan, error) {

	c := cli.c
	ctx := cli.ctx
	res, err := c.CreateScans(ctx, client.CreateScansPath(), payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return c.DecodeScan(res)
}

func (cli *CLI) startScanByMultipartUpload(file string) (*client.Scan, error) {
	payload := client.FileScanPayload{
		Upload: file,
	}
	resp, err := cli.c.UploadFileScans(context.Background(), client.UploadFileScansPath(), &payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint
	return cli.c.DecodeScan(resp)
}

func (cli *CLI) UpdateStats(scan *Scan) error {
	c := cli.c
	ctx := cli.ctx

	res, err := c.StatsScans(ctx, client.StatsScansPath(scan.ID))
	if err != nil {
		return err
	}

	stats, err := c.DecodeStats(res)
	if err != nil {
		return err
	}

	if err := stats.Validate(); err != nil {
		return err
	}

	return scan.updateStats(stats)
}

// AbortScan sends an abort scan request to vulcan core for the given id.
func (cli *CLI) AbortScan(ID uuid.UUID) error {
	c := cli.c
	ctx := cli.ctx

	res, err := c.AbortScans(ctx, client.AbortScansPath(ID))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusAccepted {
		if res.StatusCode == http.StatusConflict {
			return errors.New("scan already aborted")
		}
		return fmt.Errorf("error aborting scan, vulcan core returned http status: %s", res.Status)
	}
	return nil
}

// newHTTPClient returns the HTTP client used by the API client to make requests to the service.
func newHTTPClient() *http.Client {
	// TBD: Change as needed (e.g. to use a different transport to control redirection policy or
	// disable cert validation or...)
	return http.DefaultClient
}

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func toString(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		return err.Error()
	}

	return string(b)
}
