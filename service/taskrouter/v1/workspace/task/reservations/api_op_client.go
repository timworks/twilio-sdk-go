// Package reservations contains auto-generated files. DO NOT MODIFY
package reservations

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing task reservation resources
// See https://www.twilio.com/docs/taskrouter/api/reservations for more details
type Client struct {
	client *client.Client

	taskSid      string
	workspaceSid string
}

// ClientProperties are the properties required to manage the reservations resources
type ClientProperties struct {
	TaskSid      string
	WorkspaceSid string
}

// New creates a new instance of the reservations client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		taskSid:      properties.TaskSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
