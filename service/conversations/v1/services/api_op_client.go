// Package services contains auto-generated files. DO NOT MODIFY
package services

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing service resources
// See https://www.twilio.com/docs/conversations/api/service-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the services client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
