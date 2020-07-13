// This is an autogenerated file. DO NOT MODIFY
package message

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	serviceSid string
	sid        string
	channelSid string
}

type ClientProperties struct {
	ServiceSid string
	Sid        string
	ChannelSid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
		channelSid: properties.ChannelSid,
	}
}
