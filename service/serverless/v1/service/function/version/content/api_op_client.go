// Package content contains auto-generated files. DO NOT MODIFY
package content

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing function version content resources
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function-version/function-version-content for more details
type Client struct {
	client *client.Client

	functionSid string
	serviceSid  string
	versionSid  string
}

// ClientProperties are the properties required to manage the content resources
type ClientProperties struct {
	FunctionSid string
	ServiceSid  string
	VersionSid  string
}

// New creates a new instance of the content client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		functionSid: properties.FunctionSid,
		serviceSid:  properties.ServiceSid,
		versionSid:  properties.VersionSid,
	}
}
