// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "vulcan-core": Checks Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
)

// ShowChecksPath computes a request path to the show action of Checks.
func ShowChecksPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/v1/checks/%s", param0)
}

// Get a Check by its ID
func (c *Client) ShowChecks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowChecksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowChecksRequest create the request corresponding to the show action endpoint of the Checks resource.
func (c *Client) NewShowChecksRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
