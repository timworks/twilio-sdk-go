// Package versions contains auto-generated files. DO NOT MODIFY
package versions

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing function version resources
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function-version for more details
type Client struct {
	client *client.Client

	functionSid string
	serviceSid  string
}

// ClientProperties are the properties required to manage the versions resources
type ClientProperties struct {
	FunctionSid string
	ServiceSid  string
}

// New creates a new instance of the versions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		functionSid: properties.FunctionSid,
		serviceSid:  properties.ServiceSid,
	}
}
