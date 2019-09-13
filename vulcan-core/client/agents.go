// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "vulcan-persistence": Agents Resource Client
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

// IndexAgentsPath computes a request path to the index action of Agents.
func IndexAgentsPath() string {

	return fmt.Sprintf("/v1/agents")
}

// Get all agents
func (c *Client) IndexAgents(ctx context.Context, path string, status *string) (*http.Response, error) {
	req, err := c.NewIndexAgentsRequest(ctx, path, status)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewIndexAgentsRequest create the request corresponding to the index action endpoint of the Agents resource.
func (c *Client) NewIndexAgentsRequest(ctx context.Context, path string, status *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if status != nil {
		values.Set("status", *status)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PauseAgentsPath computes a request path to the pause action of Agents.
func PauseAgentsPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/v1/agents/%s/pause", param0)
}

// Pause an agent
func (c *Client) PauseAgents(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewPauseAgentsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPauseAgentsRequest create the request corresponding to the pause action endpoint of the Agents resource.
func (c *Client) NewPauseAgentsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowAgentsPath computes a request path to the show action of Agents.
func ShowAgentsPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/v1/agents/%s", param0)
}

// Get an Agent by its ID
func (c *Client) ShowAgents(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowAgentsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAgentsRequest create the request corresponding to the show action endpoint of the Agents resource.
func (c *Client) NewShowAgentsRequest(ctx context.Context, path string) (*http.Request, error) {
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