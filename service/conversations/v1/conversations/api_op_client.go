// Package conversations contains auto-generated files. DO NOT MODIFY
package conversations

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing conversation resources
// See https://www.twilio.com/docs/conversations/api/conversation-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the conversations client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
