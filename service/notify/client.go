package notify

import (
	"github.com/timworks/twilio-sdk-go/client"
	v1 "github.com/timworks/twilio-sdk-go/service/notify/v1"
	"github.com/timworks/twilio-sdk-go/session"
)

// Notify client is used to manage versioned resources for Twilio Notify
type Notify struct {
	V1 *v1.Notify
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Notify {
	return &Notify{
		V1: v1.New(sess, config),
	}
}
