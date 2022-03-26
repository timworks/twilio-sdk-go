package messaging

import (
	"github.com/timworks/twilio-sdk-go/client"
	v1 "github.com/timworks/twilio-sdk-go/service/messaging/v1"
	"github.com/timworks/twilio-sdk-go/session"
)

// Messaging client is used to manage versioned resources for Twilio Messaging
// See https://www.twilio.com/docs/messaging for more details on the API
// See https://www.twilio.com/messaging for more details on the product
type Messaging struct {
	V1 *v1.Messaging
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Messaging {
	return &Messaging{
		V1: v1.New(sess, config),
	}
}
