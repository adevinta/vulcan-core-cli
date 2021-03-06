// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "vulcan-persistence": Checktypes Resource Client
//
// Command:
// $ gen

package client

import (
	"context"
	"fmt"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
)

// IndexChecktypesPath computes a request path to the index action of Checktypes.
func IndexChecktypesPath() string {

	return fmt.Sprintf("/v1/checktypes")
}

// Get all checktypes
func (c *Client) IndexChecktypes(ctx context.Context, path string, enabled *string) (*http.Response, error) {
	req, err := c.NewIndexChecktypesRequest(ctx, path, enabled)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewIndexChecktypesRequest create the request corresponding to the index action endpoint of the Checktypes resource.
func (c *Client) NewIndexChecktypesRequest(ctx context.Context, path string, enabled *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if enabled != nil {
		values.Set("enabled", *enabled)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowChecktypesPath computes a request path to the show action of Checktypes.
func ShowChecktypesPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/v1/checktypes/%s", param0)
}

// Get a Checktype by its ID
func (c *Client) ShowChecktypes(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowChecktypesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowChecktypesRequest create the request corresponding to the show action endpoint of the Checktypes resource.
func (c *Client) NewShowChecktypesRequest(ctx context.Context, path string) (*http.Request, error) {
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
