// Package field contains auto-generated files. DO NOT MODIFY
package field

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific task field resource
// See https://www.twilio.com/docs/autopilot/api/task-field for more details
type Client struct {
	client *client.Client

	assistantSid string
	sid          string
	taskSid      string
}

// ClientProperties are the properties required to manage the field resources
type ClientProperties struct {
	AssistantSid string
	Sid          string
	TaskSid      string
}

// New creates a new instance of the field client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		sid:          properties.Sid,
		taskSid:      properties.TaskSid,
	}
}
