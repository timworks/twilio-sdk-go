// Package members contains auto-generated files. DO NOT MODIFY
package members

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing member resources
// See https://www.twilio.com/docs/chat/rest/member-resource for more details
type Client struct {
	client *client.Client

	channelSid string
	serviceSid string
}

// ClientProperties are the properties required to manage the members resources
type ClientProperties struct {
	ChannelSid string
	ServiceSid string
}

// New creates a new instance of the members client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		channelSid: properties.ChannelSid,
		serviceSid: properties.ServiceSid,
	}
}
