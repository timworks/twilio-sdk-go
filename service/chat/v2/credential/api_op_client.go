// This is an autogenerated file. DO NOT MODIFY
package credential

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client
	sid    string
}

func New(client *client.Client, sid string) *Client {
	return &Client{
		client: client,
		sid:    sid,
	}
}
