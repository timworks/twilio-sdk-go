// Package web_channels contains auto-generated files. DO NOT MODIFY
package web_channels

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing web channel resources
type Client struct {
	client *client.Client
}

// New creates a new instance of the web channels client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
