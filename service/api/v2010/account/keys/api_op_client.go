// Package keys contains auto-generated files. DO NOT MODIFY
package keys

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing api key resources
// See https://www.twilio.com/docs/iam/keys/api-key-resource for more details
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the keys resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the keys client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}
