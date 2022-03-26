// Package workspaces contains auto-generated files. DO NOT MODIFY
package workspaces

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing workspace resources
// See https://www.twilio.com/docs/taskrouter/api/workspace for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the workspaces client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
