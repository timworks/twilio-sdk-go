// Package services contains auto-generated files. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateServiceInput defines the input fields for creating a new service resource
type CreateServiceInput struct {
	CallbackURL             *string `form:"CallbackUrl,omitempty"`
	ChatInstanceSid         *string `form:"ChatInstanceSid,omitempty"`
	DefaultTtl              *int    `form:"DefaultTtl,omitempty"`
	GeoMatchLevel           *string `form:"GeoMatchLevel,omitempty"`
	InterceptCallbackURL    *string `form:"InterceptCallbackUrl,omitempty"`
	NumberSelectionBehavior *string `form:"NumberSelectionBehavior,omitempty"`
	OutOfSessionCallbackURL *string `form:"OutOfSessionCallbackUrl,omitempty"`
	UniqueName              string  `validate:"required" form:"UniqueName"`
}

// CreateServiceResponse defines the response fields for the created service
type CreateServiceResponse struct {
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

// Create creates a new service
// See https://www.twilio.com/docs/proxy/api/service#create-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateServiceInput) (*CreateServiceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new service
// See https://www.twilio.com/docs/proxy/api/service#create-a-service-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateServiceInput{}
	}

	response := &CreateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
