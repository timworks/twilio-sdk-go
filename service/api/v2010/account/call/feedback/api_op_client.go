// Package feedback contains auto-generated files. DO NOT MODIFY
package feedback

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific feedback resource
// See https://www.twilio.com/docs/voice/api/feedback-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	callSid    string
}

// ClientProperties are the properties required to manage the feedback resources
type ClientProperties struct {
	AccountSid string
	CallSid    string
}

// New creates a new instance of the feedback client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		callSid:    properties.CallSid,
	}
}
