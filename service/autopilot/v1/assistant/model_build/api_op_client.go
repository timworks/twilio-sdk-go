// Package model_build contains auto-generated files. DO NOT MODIFY
package model_build

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific model build resource
// See https://www.twilio.com/docs/autopilot/api/model-build for more details
type Client struct {
	client *client.Client

	assistantSid string
	sid          string
}

// ClientProperties are the properties required to manage the model build resources
type ClientProperties struct {
	AssistantSid string
	Sid          string
}

// New creates a new instance of the model build client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		sid:          properties.Sid,
	}
}
