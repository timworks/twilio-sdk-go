// Package plugins contains auto-generated files. DO NOT MODIFY
package plugins

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing plugin resources
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin for more details
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client
}

// New creates a new instance of the plugins client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
