// Package webhooks contains auto-generated files. DO NOT MODIFY
package webhooks

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing webhook resources
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource for more details
type Client struct {
	client *client.Client

	conversationSid string
}

// ClientProperties are the properties required to manage the webhooks resources
type ClientProperties struct {
	ConversationSid string
}

// New creates a new instance of the webhooks client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		conversationSid: properties.ConversationSid,
	}
}
