// This is an autogenerated file. DO NOT MODIFY
package versions

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	functionSid string
	serviceSid  string
}

// The properties required to manage the versions resources
type ClientProperties struct {
	FunctionSid string
	ServiceSid  string
}

// Create a new instance of the client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		functionSid: properties.FunctionSid,
		serviceSid:  properties.ServiceSid,
	}
}
