// Package field_type contains auto-generated files. DO NOT MODIFY
package field_type

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/field_type/field_value"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/field_type/field_values"
)

// Client for managing a specific field type resource
// See https://www.twilio.com/docs/autopilot/api/field-type for more details
type Client struct {
	client *client.Client

	assistantSid string
	sid          string

	FieldValue  func(string) *field_value.Client
	FieldValues *field_values.Client
}

// ClientProperties are the properties required to manage the field type resources
type ClientProperties struct {
	AssistantSid string
	Sid          string
}

// New creates a new instance of the field type client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		sid:          properties.Sid,

		FieldValue: func(fieldValueSid string) *field_value.Client {
			return field_value.New(client, field_value.ClientProperties{
				AssistantSid: properties.AssistantSid,
				FieldTypeSid: properties.Sid,
				Sid:          fieldValueSid,
			})
		},
		FieldValues: field_values.New(client, field_values.ClientProperties{
			AssistantSid: properties.AssistantSid,
			FieldTypeSid: properties.Sid,
		}),
	}
}
