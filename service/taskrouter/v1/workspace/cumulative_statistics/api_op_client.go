// Package cumulative_statistics contains auto-generated files. DO NOT MODIFY
package cumulative_statistics

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing workspace cumulative statistics
// See https://www.twilio.com/docs/taskrouter/api/workspace-statistics for more details
type Client struct {
	client *client.Client

	workspaceSid string
}

// ClientProperties are the properties required to manage the cumulative statistics resources
type ClientProperties struct {
	WorkspaceSid string
}

// New creates a new instance of the cumulative statistics client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workspaceSid: properties.WorkspaceSid,
	}
}
