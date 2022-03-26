// Package flow contains auto-generated files. DO NOT MODIFY
package flow

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/studio/v2/flow/execution"
	"github.com/timworks/twilio-sdk-go/service/studio/v2/flow/executions"
	"github.com/timworks/twilio-sdk-go/service/studio/v2/flow/revision"
	"github.com/timworks/twilio-sdk-go/service/studio/v2/flow/revisions"
	"github.com/timworks/twilio-sdk-go/service/studio/v2/flow/test_users"
)

// Client for managing a specific flow resource
// See https://www.twilio.com/docs/studio/rest-api/v2/flow for more details
type Client struct {
	client *client.Client

	sid string

	Execution  func(string) *execution.Client
	Executions *executions.Client
	Revision   func(int) *revision.Client
	Revisions  *revisions.Client
	TestUsers  func() *test_users.Client
}

// ClientProperties are the properties required to manage the flow resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the flow client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Execution: func(executionSid string) *execution.Client {
			return execution.New(client, execution.ClientProperties{
				FlowSid: properties.Sid,
				Sid:     executionSid,
			})
		},
		Executions: executions.New(client, executions.ClientProperties{
			FlowSid: properties.Sid,
		}),
		Revision: func(revisionNumber int) *revision.Client {
			return revision.New(client, revision.ClientProperties{
				FlowSid:        properties.Sid,
				RevisionNumber: revisionNumber,
			})
		},
		Revisions: revisions.New(client, revisions.ClientProperties{
			FlowSid: properties.Sid,
		}),
		TestUsers: func() *test_users.Client {
			return test_users.New(client, test_users.ClientProperties{
				FlowSid: properties.Sid,
			})
		},
	}
}
