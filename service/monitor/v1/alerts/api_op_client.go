// Package alerts contains auto-generated files. DO NOT MODIFY
package alerts

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing alert resources
// See https://www.twilio.com/docs/usage/monitor-alert for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the alerts client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
