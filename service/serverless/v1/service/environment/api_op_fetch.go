// Package environment contains auto-generated files. DO NOT MODIFY
package environment

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchEnvironmentResponse defines the response fields for the retrieved environment
type FetchEnvironmentResponse struct {
	AccountSid   string     `json:"account_sid"`
	BuildSid     *string    `json:"build_sid,omitempty"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	DomainName   string     `json:"domain_name"`
	DomainSuffix *string    `json:"domain_suffix,omitempty"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Fetch retrieves a environment resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/environment#fetch-an-environment-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchEnvironmentResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a environment resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/environment#fetch-an-environment-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchEnvironmentResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Environments/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchEnvironmentResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
