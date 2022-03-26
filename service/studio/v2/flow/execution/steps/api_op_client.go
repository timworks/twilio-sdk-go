// Package steps contains auto-generated files. DO NOT MODIFY
package steps

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing step resources
// See https://www.twilio.com/docs/studio/rest-api/v2/step for more details
type Client struct {
	client *client.Client

	executionSid string
	flowSid      string
}

// ClientProperties are the properties required to manage the steps resources
type ClientProperties struct {
	ExecutionSid string
	FlowSid      string
}

// New creates a new instance of the steps client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		executionSid: properties.ExecutionSid,
		flowSid:      properties.FlowSid,
	}
}
