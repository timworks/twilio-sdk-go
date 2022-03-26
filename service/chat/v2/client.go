// Package v2 contains auto-generated files. DO NOT MODIFY
package v2

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/credential"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/credentials"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/services"
	"github.com/timworks/twilio-sdk-go/session"
)

// Chat client is used to manage resources for Programmable Chat
// See https://www.twilio.com/docs/chat for more details
type Chat struct {
	client *client.Client

	Credential  func(string) *credential.Client
	Credentials *credentials.Client
	Service     func(string) *service.Client
	Services    *services.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Chat {
	return &Chat{
		client: client,

		Credential: func(credentialSid string) *credential.Client {
			return credential.New(client, credential.ClientProperties{
				Sid: credentialSid,
			})
		},
		Credentials: credentials.New(client),
		Service: func(serviceSid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: serviceSid,
			})
		},
		Services: services.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Chat) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Chat {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = false
	config.SubDomain = "chat"
	config.APIVersion = "v2"

	return NewWithClient(client.New(sess, config))
}
