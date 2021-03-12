/*
Copyright 2021 Adevinta
*/

// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "vulcan-core": Assettypes Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// IndexAssettypesPath computes a request path to the index action of Assettypes.
func IndexAssettypesPath() string {

	return fmt.Sprintf("/v1/assettypes")
}

// Get all assettypes and per each one the checktypes that are accepting that concrete assettype
func (c *Client) IndexAssettypes(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewIndexAssettypesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewIndexAssettypesRequest create the request corresponding to the index action endpoint of the Assettypes resource.
func (c *Client) NewIndexAssettypesRequest(ctx context.Context, path string) (*http.Request, error) {
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
