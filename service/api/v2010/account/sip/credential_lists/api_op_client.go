// Package credential_lists contains auto-generated files. DO NOT MODIFY
package credential_lists

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing credential list resources
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the credential lists resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the credential lists client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}
