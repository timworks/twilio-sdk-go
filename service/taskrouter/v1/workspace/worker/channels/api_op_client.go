// Package channels contains auto-generated files. DO NOT MODIFY
package channels

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing worker channel resources
// See https://www.twilio.com/docs/taskrouter/api/worker-channel for more details
type Client struct {
	client *client.Client

	workerSid    string
	workspaceSid string
}

// ClientProperties are the properties required to manage the channels resources
type ClientProperties struct {
	WorkerSid    string
	WorkspaceSid string
}

// New creates a new instance of the channels client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workerSid:    properties.WorkerSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
