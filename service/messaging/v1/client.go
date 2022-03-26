// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/messaging/v1/service"
	"github.com/timworks/twilio-sdk-go/service/messaging/v1/services"
	"github.com/timworks/twilio-sdk-go/session"
)

// Messaging client is used to manage resources for Twilio Messaging
// See https://www.twilio.com/docs/messaging for more details
// This client is currently in beta and subject to change. Please use with caution
type Messaging struct {
	client *client.Client

	Service  func(string) *service.Client
	Services *services.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Messaging {
	return &Messaging{
		client: client,

		Service: func(serviceSid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: serviceSid,
			})
		},
		Services: services.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Messaging) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Messaging {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = true
	config.SubDomain = "messaging"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}
