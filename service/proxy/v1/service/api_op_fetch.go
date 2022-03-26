// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchServiceResponse defines the response fields for the retrieved service
type FetchServiceResponse struct {
	AccountSid              string     `json:"account_sid"`
	CallbackURL             *string    `json:"callback_url,omitempty"`
	ChatInstanceSid         *string    `json:"chat_instance_sid,omitempty"`
	ChatServiceSid          string     `json:"chat_service_sid"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	DefaultTtl              *int       `json:"default_ttl,omitempty"`
	GeoMatchLevel           *string    `json:"geo_match_level,omitempty"`
	InterceptCallbackURL    *string    `json:"intercept_callback_url,omitempty"`
	NumberSelectionBehavior *string    `json:"number_selection_behavior,omitempty"`
	OutOfSessionCallbackURL *string    `json:"out_of_session_callback_url,omitempty"`
	Sid                     string     `json:"sid"`
	URL                     string     `json:"url"`
	UniqueName              string     `json:"unique_name"`
}

// Fetch retrieves a service resource
// See https://www.twilio.com/docs/proxy/api/service#fetch-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchServiceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a service resource
// See https://www.twilio.com/docs/proxy/api/service#fetch-a-service-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchServiceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchServiceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
