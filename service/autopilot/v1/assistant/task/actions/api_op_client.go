// Package actions contains auto-generated files. DO NOT MODIFY
package actions

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing task action resources
// See https://www.twilio.com/docs/autopilot/api/task-action for more details
type Client struct {
	client *client.Client

	assistantSid string
	taskSid      string
}

// ClientProperties are the properties required to manage the actions resources
type ClientProperties struct {
	AssistantSid string
	TaskSid      string
}

// New creates a new instance of the actions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		taskSid:      properties.TaskSid,
	}
}
