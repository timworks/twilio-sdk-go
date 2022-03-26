// Package messages contains auto-generated files. DO NOT MODIFY
package messages

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing message resources
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource for more details
type Client struct {
	client *client.Client

	conversationSid string
}

// ClientProperties are the properties required to manage the messages resources
type ClientProperties struct {
	ConversationSid string
}

// New creates a new instance of the messages client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		conversationSid: properties.ConversationSid,
	}
}
