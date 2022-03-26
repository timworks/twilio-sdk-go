// Package cumulative_statistics contains auto-generated files. DO NOT MODIFY
package cumulative_statistics

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing task queue cumulative statistics
// See https://www.twilio.com/docs/taskrouter/api/taskqueue-statistics for more details
type Client struct {
	client *client.Client

	taskQueueSid string
	workspaceSid string
}

// ClientProperties are the properties required to manage the cumulative statistics resources
type ClientProperties struct {
	TaskQueueSid string
	WorkspaceSid string
}

// New creates a new instance of the cumulative statistics client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		taskQueueSid: properties.TaskQueueSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
