// Package permissions contains auto-generated files. DO NOT MODIFY
package permissions

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing document permission resources
// See https://www.twilio.com/docs/sync/api/document-permission-resource for more details
type Client struct {
	client *client.Client

	documentSid string
	serviceSid  string
}

// ClientProperties are the properties required to manage the permissions resources
type ClientProperties struct {
	DocumentSid string
	ServiceSid  string
}

// New creates a new instance of the permissions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		documentSid: properties.DocumentSid,
		serviceSid:  properties.ServiceSid,
	}
}
