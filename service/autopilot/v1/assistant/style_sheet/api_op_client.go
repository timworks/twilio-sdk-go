// Package style_sheet contains auto-generated files. DO NOT MODIFY
package style_sheet

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing style sheet resources
// See https://www.twilio.com/docs/autopilot/api/assistant/stylesheet for more details
type Client struct {
	client *client.Client

	assistantSid string
}

// ClientProperties are the properties required to manage the stylesheet resources
type ClientProperties struct {
	AssistantSid string
}

// New creates a new instance of the stylesheet client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
	}
}
