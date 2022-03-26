// Package flex_flows contains auto-generated files. DO NOT MODIFY
package flex_flows

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing flex flow resources
type Client struct {
	client *client.Client
}

// New creates a new instance of the flex flows client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
