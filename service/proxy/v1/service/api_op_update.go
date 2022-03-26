// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateServiceInput defines input fields for updating a service resource
type UpdateServiceInput struct {
	CallbackURL             *string `form:"CallbackUrl,omitempty"`
	ChatInstanceSid         *string `form:"ChatInstanceSid,omitempty"`
	DefaultTtl              *int    `form:"DefaultTtl,omitempty"`
	GeoMatchLevel           *string `form:"GeoMatchLevel,omitempty"`
	InterceptCallbackURL    *string `form:"InterceptCallbackUrl,omitempty"`
	NumberSelectionBehavior *string `form:"NumberSelectionBehavior,omitempty"`
	OutOfSessionCallbackURL *string `form:"OutOfSessionCallbackUrl,omitempty"`
	UniqueName              *string `form:"UniqueName,omitempty"`
}

// UpdateServiceResponse defines the response fields for the updated service
type UpdateServiceResponse struct {
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

// Update modifies a service resource
// See https://www.twilio.com/docs/proxy/api/service#update-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a service resource
// See https://www.twilio.com/docs/proxy/api/service#update-a-service-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateServiceInput{}
	}

	response := &UpdateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
