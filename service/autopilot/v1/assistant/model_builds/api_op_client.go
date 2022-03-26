// Package model_builds contains auto-generated files. DO NOT MODIFY
package model_builds

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing model build resources
// See https://www.twilio.com/docs/autopilot/api/model-build for more details
type Client struct {
	client *client.Client

	assistantSid string
}

// ClientProperties are the properties required to manage the model builds resources
type ClientProperties struct {
	AssistantSid string
}

// New creates a new instance of the model builds client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
	}
}
