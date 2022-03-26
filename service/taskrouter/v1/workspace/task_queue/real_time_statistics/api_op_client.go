// Package real_time_statistics contains auto-generated files. DO NOT MODIFY
package real_time_statistics

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing task queue real time statistics
// See https://www.twilio.com/docs/taskrouter/api/taskqueue-statistics for more details
type Client struct {
	client *client.Client

	taskQueueSid string
	workspaceSid string
}

// ClientProperties are the properties required to manage the real time statistics resources
type ClientProperties struct {
	TaskQueueSid string
	WorkspaceSid string
}

// New creates a new instance of the real time statistics client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		taskQueueSid: properties.TaskQueueSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
