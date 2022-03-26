// Package sample contains auto-generated files. DO NOT MODIFY
package sample

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific task sample resource
// See https://www.twilio.com/docs/autopilot/api/task-sample for more details
type Client struct {
	client *client.Client

	assistantSid string
	sid          string
	taskSid      string
}

// ClientProperties are the properties required to manage the sample resources
type ClientProperties struct {
	AssistantSid string
	Sid          string
	TaskSid      string
}

// New creates a new instance of the sample client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		sid:          properties.Sid,
		taskSid:      properties.TaskSid,
	}
}
