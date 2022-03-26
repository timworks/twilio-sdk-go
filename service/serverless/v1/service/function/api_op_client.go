// Package function contains auto-generated files. DO NOT MODIFY
package function

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/function/version"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/function/versions"
)

// Client for managing a specific function resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Version  func(string) *version.Client
	Versions *versions.Client
}

// ClientProperties are the properties required to manage the function resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the function client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Version: func(versionSid string) *version.Client {
			return version.New(client, version.ClientProperties{
				FunctionSid: properties.Sid,
				ServiceSid:  properties.ServiceSid,
				Sid:         versionSid,
			})
		},
		Versions: versions.New(client, versions.ClientProperties{
			FunctionSid: properties.Sid,
			ServiceSid:  properties.ServiceSid,
		}),
	}
}
