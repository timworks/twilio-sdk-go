// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/sync/v1/service"
	"github.com/timworks/twilio-sdk-go/service/sync/v1/services"
	"github.com/timworks/twilio-sdk-go/session"
)

// Sync client is used to manage resources for Twilio Sync
// See https://www.twilio.com/docs/sync for more details
type Sync struct {
	client *client.Client

	Service  func(string) *service.Client
	Services *services.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Sync {
	return &Sync{
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
func (s Sync) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Sync {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = false
	config.SubDomain = "sync"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}
