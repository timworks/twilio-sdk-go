// Package channel contains auto-generated files. DO NOT MODIFY
package channel

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific user channel resource
// See https://www.twilio.com/docs/chat/rest/user-channel-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string
	userSid    string
}

// ClientProperties are the properties required to manage the channel resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
	UserSid    string
}

// New creates a new instance of the channel client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
		userSid:    properties.UserSid,
	}
}
