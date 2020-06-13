// This is an autogenerated file. DO NOT MODIFY
package versions

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
)

type Client struct {
	client     *client.Client
	assetSid   string
	serviceSid string
}

func New(client *client.Client, assetSid string, serviceSid string) *Client {
	return &Client{
		client:     client,
		assetSid:   assetSid,
		serviceSid: serviceSid,
	}
}
