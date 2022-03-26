// Package flows contains auto-generated files. DO NOT MODIFY
package flows

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing flow resources
// See https://www.twilio.com/docs/studio/rest-api/v2/flow for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the flows client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
