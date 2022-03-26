// Package invites contains auto-generated files. DO NOT MODIFY
package invites

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing invite resources
// See twilio.com/docs/chat/rest/invite-resource for more details
type Client struct {
	client *client.Client

	channelSid string
	serviceSid string
}

// ClientProperties are the properties required to manage the invites resources
type ClientProperties struct {
	ChannelSid string
	ServiceSid string
}

// New creates a new instance of the invites client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		channelSid: properties.ChannelSid,
		serviceSid: properties.ServiceSid,
	}
}
