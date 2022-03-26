// Package model_build contains auto-generated files. DO NOT MODIFY
package model_build

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchModelBuildResponse defines the response fields for the retrieved model build
type FetchModelBuildResponse struct {
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

// Fetch retrieves a model build resource
// See https://www.twilio.com/docs/autopilot/api/model-build#fetch-a-modelbuild-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchModelBuildResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a model build resource
// See https://www.twilio.com/docs/autopilot/api/model-build#fetch-a-modelbuild-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchModelBuildResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/ModelBuilds/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &FetchModelBuildResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
