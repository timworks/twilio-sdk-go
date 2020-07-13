// This is an autogenerated file. DO NOT MODIFY
package messages

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	channelSid string
	serviceSid string
}

type ClientProperties struct {
	ChannelSid string
	ServiceSid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		channelSid: properties.ChannelSid,
		serviceSid: properties.ServiceSid,
	}
}
