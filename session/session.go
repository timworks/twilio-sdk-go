package session

import (
	"github.com/timworks/twilio-sdk-go/session/credentials"
)

// Session represents a session object that can be used to make requests against the Twilio APIs
type Session struct {
	*credentials.Credentials
}

// New creates a new session instance using the credentials supplied
func New(creds *credentials.Credentials) *Session {
	return &Session{
		Credentials: creds,
	}
}
