// Package statistics contains auto-generated files. DO NOT MODIFY
package statistics

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing task statistics resources
// See https://www.twilio.com/docs/autopilot/api/task-statistics for more details
type Client struct {
	client *client.Client

	assistantSid string
	taskSid      string
}

// ClientProperties are the properties required to manage the statistics resources
type ClientProperties struct {
	AssistantSid string
	TaskSid      string
}

// New creates a new instance of the statistics client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		taskSid:      properties.TaskSid,
	}
}
