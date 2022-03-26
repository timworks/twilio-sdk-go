// Package step contains auto-generated files. DO NOT MODIFY
package step

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/studio/v2/flow/execution/step/context"
)

// Client for managing a specific step resource
// See https://www.twilio.com/docs/studio/rest-api/v2/step for more details
type Client struct {
	client *client.Client

	executionSid string
	flowSid      string
	sid          string

	Context func() *context.Client
}

// ClientProperties are the properties required to manage the step resources
type ClientProperties struct {
	ExecutionSid string
	FlowSid      string
	Sid          string
}

// New creates a new instance of the step client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		executionSid: properties.ExecutionSid,
		flowSid:      properties.FlowSid,
		sid:          properties.Sid,

		Context: func() *context.Client {
			return context.New(client, context.ClientProperties{
				ExecutionSid: properties.ExecutionSid,
				FlowSid:      properties.FlowSid,
				StepSid:      properties.Sid,
			})
		},
	}
}
