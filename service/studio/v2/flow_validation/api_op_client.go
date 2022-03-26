// Package flow_validation contains auto-generated files. DO NOT MODIFY
package flow_validation

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing flow validation resources
// See https://www.twilio.com/docs/studio/rest-api/v2/flow-validate for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the flow validation client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
