// Package assistant contains auto-generated files. DO NOT MODIFY
package assistant

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateAssistantInput defines the input fields for updating an assistant resource
type UpdateAssistantInput struct {
	CallbackEvents   *string `form:"CallbackEvents,omitempty"`
	CallbackURL      *string `form:"CallbackUrl,omitempty"`
	Defaults         *string `form:"Defaults,omitempty"`
	DevelopmentStage *string `form:"DevelopmentStage,omitempty"`
	FriendlyName     *string `form:"FriendlyName,omitempty"`
	LogQueries       *bool   `form:"LogQueries,omitempty"`
	StyleSheet       *string `form:"StyleSheet,omitempty"`
	UniqueName       *string `form:"UniqueName,omitempty"`
}

// UpdateAssistantResponse defines the response fields for the updated assistant
type UpdateAssistantResponse struct {
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

// Update modifies an assistant resource
// See https://www.twilio.com/docs/autopilot/api/assistant#update-an-assistant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateAssistantInput) (*UpdateAssistantResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an assistant resource
// See https://www.twilio.com/docs/autopilot/api/assistant#update-an-assistant-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateAssistantInput) (*UpdateAssistantResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateAssistantInput{}
	}

	response := &UpdateAssistantResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
