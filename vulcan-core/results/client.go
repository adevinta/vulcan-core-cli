package results

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	report "github.com/adevinta/vulcan-report"
)

// Doer defines the methods needed by the http client used by
// the results client.
type Doer interface {
	Get(url string) (resp *http.Response, err error)
}

// Client implements methods for downloading a report. The client is safe to use
// concurrently.
type Client struct {
	doer     Doer
	baseAddr *url.URL
}

// NewClient instantiates and returns a new results client.
func NewClient(baseURL string, d Doer) (*Client, error) {
	addr, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &Client{d, addr}, nil
}

// GetReport retrieves a report of a check stored on vulcan results. The report
// is indexed using three parameters: the id of check, the id of the scan the
// check belongs to and the date the san was run.
func (c *Client) GetReport(date, scanID, checkID string) (*report.Report, error) {
	// Copy the address of the base url to avoid race condition.
	u := *c.baseAddr
	u.Path = fmt.Sprintf("/v1/reports/dt=%s/scan=%s/%s.json", date, scanID, checkID)
	resp, err := c.doer.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &report.Report{}

	err = r.UnmarshalJSONTimeAsString(body)
	if err != nil {
		return nil, err
	}

	return r, nil
}
