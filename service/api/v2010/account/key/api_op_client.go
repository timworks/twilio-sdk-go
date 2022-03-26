// Package key contains auto-generated files. DO NOT MODIFY
package key

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific api key resource
// See https://www.twilio.com/docs/iam/keys/api-key-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string
}

// ClientProperties are the properties required to manage the key resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the key client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,
	}
}
