// Package credentials contains auto-generated files. DO NOT MODIFY
package credentials

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing credential resources
// See https://www.twilio.com/docs/conversations/api/credential-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the credentials client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
