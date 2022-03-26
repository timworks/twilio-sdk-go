// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing conversation configuration webhook resources
type Client struct {
	client *client.Client
}

// New creates a new instance of the webhook client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
