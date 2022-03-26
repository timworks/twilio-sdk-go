// Package field_value contains auto-generated files. DO NOT MODIFY
package field_value

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific field value resource
// See https://www.twilio.com/docs/autopilot/api/field-value for more details
type Client struct {
	client *client.Client

	assistantSid string
	fieldTypeSid string
	sid          string
}

// ClientProperties are the properties required to manage the field value resources
type ClientProperties struct {
	AssistantSid string
	FieldTypeSid string
	Sid          string
}

// New creates a new instance of the field value client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		fieldTypeSid: properties.FieldTypeSid,
		sid:          properties.Sid,
	}
}
