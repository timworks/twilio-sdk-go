// Package statistics contains auto-generated files. DO NOT MODIFY
package statistics

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing workflow statistics
// See https://www.twilio.com/docs/taskrouter/api/workflow-statistics for more details
type Client struct {
	client *client.Client

	workflowSid  string
	workspaceSid string
}

// ClientProperties are the properties required to manage the statistics resources
type ClientProperties struct {
	WorkflowSid  string
	WorkspaceSid string
}

// New creates a new instance of the statistics client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workflowSid:  properties.WorkflowSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
