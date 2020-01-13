// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "vulcan-persistence": Scans Resource Client
//
// Command:
// $ gen

package client

import (
	"bytes"
	"context"
	"fmt"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
)

// AbortScansPath computes a request path to the abort action of Scans.
func AbortScansPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/v1/scans/%s/abort", param0)
}

// Abort a scan
func (c *Client) AbortScans(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAbortScansRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAbortScansRequest create the request corresponding to the abort action endpoint of the Scans resource.
func (c *Client) NewAbortScansRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ChecksScansPath computes a request path to the checks action of Scans.
func ChecksScansPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/v1/scans/%s/checks", param0)
}

// Get checks of the Scan
func (c *Client) ChecksScans(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewChecksScansRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewChecksScansRequest create the request corresponding to the checks action endpoint of the Scans resource.
func (c *Client) NewChecksScansRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CreateScansPath computes a request path to the create action of Scans.
func CreateScansPath() string {

	return fmt.Sprintf("/v1/scans/")
}

// Create a new Scan
func (c *Client) CreateScans(ctx context.Context, path string, payload *ScanPayload) (*http.Response, error) {
	req, err := c.NewCreateScansRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateScansRequest create the request corresponding to the create action endpoint of the Scans resource.
func (c *Client) NewCreateScansRequest(ctx context.Context, path string, payload *ScanPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// IndexScansPath computes a request path to the index action of Scans.
func IndexScansPath() string {

	return fmt.Sprintf("/v1/scans")
}

// Get all scans
func (c *Client) IndexScans(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewIndexScansRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewIndexScansRequest create the request corresponding to the index action endpoint of the Scans resource.
func (c *Client) NewIndexScansRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowScansPath computes a request path to the show action of Scans.
func ShowScansPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/v1/scans/%s", param0)
}

// Get a Scan by its ID
func (c *Client) ShowScans(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowScansRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowScansRequest create the request corresponding to the show action endpoint of the Scans resource.
func (c *Client) NewShowScansRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// StatsScansPath computes a request path to the stats action of Scans.
func StatsScansPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/v1/scans/%s/stats", param0)
}

// Get stats of the Scan
func (c *Client) StatsScans(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewStatsScansRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStatsScansRequest create the request corresponding to the stats action endpoint of the Scans resource.
func (c *Client) NewStatsScansRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
