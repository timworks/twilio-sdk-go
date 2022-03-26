// Package channels contains auto-generated files. DO NOT MODIFY
package channels

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing channel resources
type Client struct {
	client *client.Client
}

// New creates a new instance of the channels client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
