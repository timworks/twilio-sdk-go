// Package faxes contains auto-generated files. DO NOT MODIFY
package faxes

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing fax resources
// See https://www.twilio.com/docs/fax/api/fax-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the faxes client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
