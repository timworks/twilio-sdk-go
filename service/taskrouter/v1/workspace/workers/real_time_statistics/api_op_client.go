// Package real_time_statistics contains auto-generated files. DO NOT MODIFY
package real_time_statistics

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing all workers real time statistics
// See https://www.twilio.com/docs/taskrouter/api/worker/statistics for more details
type Client struct {
	client *client.Client

	workspaceSid string
}

// ClientProperties are the properties required to manage the real time statistics resources
type ClientProperties struct {
	WorkspaceSid string
}

// New creates a new instance of the real time statistics client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workspaceSid: properties.WorkspaceSid,
	}
}
