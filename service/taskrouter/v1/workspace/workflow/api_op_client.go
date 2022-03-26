// Package workflow contains auto-generated files. DO NOT MODIFY
package workflow

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/workflow/cumulative_statistics"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/workflow/real_time_statistics"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/workflow/statistics"
)

// Client for managing a specific workflow resource
// See https://www.twilio.com/docs/taskrouter/api/workflow for more details
type Client struct {
	client *client.Client

	sid          string
	workspaceSid string

	CumulativeStatistics func() *cumulative_statistics.Client
	RealTimeStatistics   func() *real_time_statistics.Client
	Statistics           func() *statistics.Client
}

// ClientProperties are the properties required to manage the workflow resources
type ClientProperties struct {
	Sid          string
	WorkspaceSid string
}

// New creates a new instance of the workflow client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		workspaceSid: properties.WorkspaceSid,

		CumulativeStatistics: func() *cumulative_statistics.Client {
			return cumulative_statistics.New(client, cumulative_statistics.ClientProperties{
				WorkflowSid:  properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
		RealTimeStatistics: func() *real_time_statistics.Client {
			return real_time_statistics.New(client, real_time_statistics.ClientProperties{
				WorkflowSid:  properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
		Statistics: func() *statistics.Client {
			return statistics.New(client, statistics.ClientProperties{
				WorkflowSid:  properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
	}
}
