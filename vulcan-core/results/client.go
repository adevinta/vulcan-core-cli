/*
Copyright 2019 Adevinta
*/

package results

import (
	"io"
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
	doer Doer
}

// NewClient instantiates and returns a new results client.
func NewClient(d Doer) (*Client, error) {
	return &Client{d}, nil
}

// GetReport retrieves a report of a check stored on vulcan results. The report
// is indexed using three parameters: the id of check, the id of the scan the
// check belongs to and the date the san was run.
func (c *Client) GetReport(u string) (*report.Report, error) {
	url, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	resp, err := c.doer.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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
