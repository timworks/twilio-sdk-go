// Package media_attachments contains auto-generated files. DO NOT MODIFY
package media_attachments

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing message media attachments resources
// See https://www.twilio.com/docs/sms/api/media-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	messageSid string
}

// ClientProperties are the properties required to manage the media attachments resources
type ClientProperties struct {
	AccountSid string
	MessageSid string
}

// New creates a new instance of the media attachments client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		messageSid: properties.MessageSid,
	}
}
