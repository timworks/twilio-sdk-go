// This is an autogenerated file. DO NOT MODIFY
package asset

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset/version"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset/versions"
)

type Client struct {
	client     *client.Client
	serviceSid string
	sid        string
	Versions   *versions.Client
	Version    func(string) *version.Client
}

func New(client *client.Client, serviceSid string, sid string) *Client {
	return &Client{
		client:     client,
		serviceSid: serviceSid,
		sid:        sid,
		Versions:   versions.New(client, sid, serviceSid),
		Version:    func(versionSid string) *version.Client { return version.New(client, sid, serviceSid, versionSid) },
	}
}
