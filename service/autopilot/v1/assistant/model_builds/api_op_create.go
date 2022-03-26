// Package model_builds contains auto-generated files. DO NOT MODIFY
package model_builds

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateModelBuildInput defines the input fields for creating a new model build resource
type CreateModelBuildInput struct {
	StatusCallback *string `form:"StatusCallback,omitempty"`
	UniqueName     *string `form:"UniqueName,omitempty"`
}

// CreateModelBuildResponse defines the response fields for the created model build
type CreateModelBuildResponse struct {
	AccountSid    string     `json:"account_sid"`
	AssistantSid  string     `json:"assistant_sid"`
	BuildDuration *int       `json:"build_duration,omitempty"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	ErrorCode     *int       `json:"error_code,omitempty"`
	Sid           string     `json:"sid"`
	Status        string     `json:"status"`
	URL           string     `json:"url"`
	UniqueName    string     `json:"unique_name"`
}

// Create creates a new model build
// See https://www.twilio.com/docs/autopilot/api/model-build#create-a-modelbuild-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateModelBuildInput) (*CreateModelBuildResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new model build
// See https://www.twilio.com/docs/autopilot/api/model-build#create-a-modelbuild-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateModelBuildInput) (*CreateModelBuildResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/ModelBuilds",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	if input == nil {
		input = &CreateModelBuildInput{}
	}

	response := &CreateModelBuildResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
