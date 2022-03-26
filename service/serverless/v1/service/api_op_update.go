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
	FriendlyName       *string `form:"FriendlyName,omitempty"`
	IncludeCredentials *bool   `form:"IncludeCredentials,omitempty"`
	UiEditable         *bool   `form:"UiEditable,omitempty"`
}

// UpdateServiceResponse defines the response fields for the updated service
type UpdateServiceResponse struct {
	AccountSid         string     `json:"account_sid"`
	DateCreated        time.Time  `json:"date_created"`
	DateUpdated        *time.Time `json:"date_updated,omitempty"`
	FriendlyName       string     `json:"friendly_name"`
	IncludeCredentials bool       `json:"include_credentials"`
	Sid                string     `json:"sid"`
	URL                string     `json:"url"`
	UiEditable         bool       `json:"ui_editable"`
	UniqueName         string     `json:"unique_name"`
}

// Update modifies a service resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/service#update-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a service resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/service#update-a-service-resource for more details
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
