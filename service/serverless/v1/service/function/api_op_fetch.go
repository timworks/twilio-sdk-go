// Package function contains auto-generated files. DO NOT MODIFY
package function

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchFunctionResponse defines the response fields for the retrieved function
type FetchFunctionResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a function resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function#fetch-a-function-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFunctionResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a function resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function#fetch-a-function-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFunctionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Functions/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchFunctionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
