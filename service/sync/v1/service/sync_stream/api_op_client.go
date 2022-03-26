// Package sync_stream contains auto-generated files. DO NOT MODIFY
package sync_stream

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/sync/v1/service/sync_stream/messages"
)

// Client for managing a specific stream resource
// See https://www.twilio.com/docs/sync/api/stream-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Messages *messages.Client
}

// ClientProperties are the properties required to manage the syncstream resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the syncstream client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Messages: messages.New(client, messages.ClientProperties{
			ServiceSid:    properties.ServiceSid,
			SyncStreamSid: properties.Sid,
		}),
	}
}
