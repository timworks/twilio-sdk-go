// Package binding contains auto-generated files. DO NOT MODIFY
package binding

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific user binding resource
// See https://www.twilio.com/docs/chat/rest/user-binding-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string
	userSid    string
}

// ClientProperties are the properties required to manage the binding resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
	UserSid    string
}

// New creates a new instance of the binding client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
		userSid:    properties.UserSid,
	}
}
