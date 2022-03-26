// Package field_types contains auto-generated files. DO NOT MODIFY
package field_types

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing field type resources
// See https://www.twilio.com/docs/autopilot/api/field-type for more details
type Client struct {
	client *client.Client

	assistantSid string
}

// ClientProperties are the properties required to manage the field types resources
type ClientProperties struct {
	AssistantSid string
}

// New creates a new instance of the field types client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
	}
}
