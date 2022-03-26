// Package execution contains auto-generated files. DO NOT MODIFY
package execution

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/studio/v2/flow/execution/context"
	"github.com/timworks/twilio-sdk-go/service/studio/v2/flow/execution/step"
	"github.com/timworks/twilio-sdk-go/service/studio/v2/flow/execution/steps"
)

// Client for managing a specific execution resource
// See https://www.twilio.com/docs/studio/rest-api/v2/execution for more details
type Client struct {
	client *client.Client

	flowSid string
	sid     string

	Context func() *context.Client
	Step    func(string) *step.Client
	Steps   *steps.Client
}

// ClientProperties are the properties required to manage the execution resources
type ClientProperties struct {
	FlowSid string
	Sid     string
}

// New creates a new instance of the execution client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		flowSid: properties.FlowSid,
		sid:     properties.Sid,

		Context: func() *context.Client {
			return context.New(client, context.ClientProperties{
				ExecutionSid: properties.Sid,
				FlowSid:      properties.FlowSid,
			})
		},
		Step: func(stepSid string) *step.Client {
			return step.New(client, step.ClientProperties{
				ExecutionSid: properties.Sid,
				FlowSid:      properties.FlowSid,
				Sid:          stepSid,
			})
		},
		Steps: steps.New(client, steps.ClientProperties{
			ExecutionSid: properties.Sid,
			FlowSid:      properties.FlowSid,
		}),
	}
}
