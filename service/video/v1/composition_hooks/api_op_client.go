// Package composition_hooks contains auto-generated files. DO NOT MODIFY
package composition_hooks

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing composition hook resources
// See https://www.twilio.com/docs/video/api/composition-hooks for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the composition hooks client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
