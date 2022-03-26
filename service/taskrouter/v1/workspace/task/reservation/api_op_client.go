// Package reservation contains auto-generated files. DO NOT MODIFY
package reservation

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific task reservation resource
// See https://www.twilio.com/docs/taskrouter/api/reservations for more details
type Client struct {
	client *client.Client

	sid          string
	taskSid      string
	workspaceSid string
}

// ClientProperties are the properties required to manage the reservation resources
type ClientProperties struct {
	Sid          string
	TaskSid      string
	WorkspaceSid string
}

// New creates a new instance of the reservation client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		taskSid:      properties.TaskSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
