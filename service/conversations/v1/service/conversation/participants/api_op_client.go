// Package participants contains auto-generated files. DO NOT MODIFY
package participants

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing participant resources
// See https://www.twilio.com/docs/conversations/api/conversation-participant-resource for more details
type Client struct {
	client *client.Client

	conversationSid string
	serviceSid      string
}

// ClientProperties are the properties required to manage the participants resources
type ClientProperties struct {
	ConversationSid string
	ServiceSid      string
}

// New creates a new instance of the participants client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		conversationSid: properties.ConversationSid,
		serviceSid:      properties.ServiceSid,
	}
}
