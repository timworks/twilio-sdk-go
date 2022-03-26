// Package document contains auto-generated files. DO NOT MODIFY
package document

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/sync/v1/service/document/permission"
	"github.com/timworks/twilio-sdk-go/service/sync/v1/service/document/permissions"
)

// Client for managing a specific document resource
// See https://www.twilio.com/docs/sync/api/document-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Permission  func(string) *permission.Client
	Permissions *permissions.Client
}

// ClientProperties are the properties required to manage the document resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the document client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Permission: func(identity string) *permission.Client {
			return permission.New(client, permission.ClientProperties{
				DocumentSid: properties.Sid,
				Identity:    identity,
				ServiceSid:  properties.ServiceSid,
			})
		},
		Permissions: permissions.New(client, permissions.ClientProperties{
			DocumentSid: properties.Sid,
			ServiceSid:  properties.ServiceSid,
		}),
	}
}
