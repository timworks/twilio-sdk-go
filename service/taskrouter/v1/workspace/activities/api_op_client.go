// Package activities contains auto-generated files. DO NOT MODIFY
package activities

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing activity resources
// See https://www.twilio.com/docs/taskrouter/api/activity for more details
type Client struct {
	client *client.Client

	workspaceSid string
}

// ClientProperties are the properties required to manage the activities resources
type ClientProperties struct {
	WorkspaceSid string
}

// New creates a new instance of the activities client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workspaceSid: properties.WorkspaceSid,
	}
}
