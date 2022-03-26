// Package task_queues contains auto-generated files. DO NOT MODIFY
package task_queues

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing task queue resources
// See https://www.twilio.com/docs/taskrouter/api/task-queue for more details
type Client struct {
	client *client.Client

	workspaceSid string
}

// ClientProperties are the properties required to manage the task queues resources
type ClientProperties struct {
	WorkspaceSid string
}

// New creates a new instance of the task queues client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workspaceSid: properties.WorkspaceSid,
	}
}
