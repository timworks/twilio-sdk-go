// Package assistants contains auto-generated files. DO NOT MODIFY
package assistants

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateAssistantInput defines the input fields for creating a new assistant resource
type CreateAssistantInput struct {
	CallbackEvents *string `form:"CallbackEvents,omitempty"`
	CallbackURL    *string `form:"CallbackUrl,omitempty"`
	Defaults       *string `form:"Defaults,omitempty"`
	FriendlyName   *string `form:"FriendlyName,omitempty"`
	LogQueries     *bool   `form:"LogQueries,omitempty"`
	StyleSheet     *string `form:"StyleSheet,omitempty"`
	UniqueName     *string `form:"UniqueName,omitempty"`
}

// CreateAssistantResponse defines the response fields for the created assistant
type CreateAssistantResponse struct {
	AccountSid          string     `json:"account_sid"`
	CallbackEvents      *string    `json:"callback_events,omitempty"`
	CallbackURL         *string    `json:"callback_url,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	DevelopmentStage    string     `json:"development_stage"`
	FriendlyName        *string    `json:"friendly_name,omitempty"`
	LatestModelBuildSid *string    `json:"latest_model_build_sid,omitempty"`
	LogQueries          bool       `json:"log_queries"`
	NeedsModelBuild     *bool      `json:"needs_model_build,omitempty"`
	Sid                 string     `json:"sid"`
	URL                 string     `json:"url"`
	UniqueName          string     `json:"unique_name"`
}

// Create creates a new assistant
// See https://www.twilio.com/docs/autopilot/api/assistant#create-an-assistant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateAssistantInput) (*CreateAssistantResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new assistant
// See https://www.twilio.com/docs/autopilot/api/assistant#create-an-assistant-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateAssistantInput) (*CreateAssistantResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateAssistantInput{}
	}

	response := &CreateAssistantResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
