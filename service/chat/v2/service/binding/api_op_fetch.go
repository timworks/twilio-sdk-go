// Package binding contains auto-generated files. DO NOT MODIFY
package binding

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchBindingResponse defines the response fields for the retrieved service binding
type FetchBindingResponse struct {
	AccountSid    string     `json:"account_sid"`
	BindingType   *string    `json:"binding_type,omitempty"`
	CredentialSid *string    `json:"credential_sid,omitempty"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	Endpoint      *string    `json:"endpoint,omitempty"`
	Identity      *string    `json:"identity,omitempty"`
	MessageTypes  *[]string  `json:"message_types,omitempty"`
	ServiceSid    string     `json:"service_sid"`
	Sid           string     `json:"sid"`
	URL           string     `json:"url"`
}

// Fetch retrieves an service binding resource
// See https://www.twilio.com/docs/chat/rest/binding-resource#fetch-a-binding-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchBindingResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an service binding resource
// See https://www.twilio.com/docs/chat/rest/binding-resource#fetch-a-binding-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchBindingResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Bindings/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchBindingResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
