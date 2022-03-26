// Package samples contains auto-generated files. DO NOT MODIFY
package samples

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateSampleInput defines the input fields for creating a new task sample resource
type CreateSampleInput struct {
	Language      string  `validate:"required" form:"Language"`
	SourceChannel *string `form:"SourceChannel,omitempty"`
	TaggedText    string  `validate:"required" form:"TaggedText"`
}

// CreateSampleResponse defines the response fields for the created task sample
type CreateSampleResponse struct {
	AccountSid    string     `json:"account_sid"`
	AssistantSid  string     `json:"assistant_sid"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	Language      string     `json:"language"`
	Sid           string     `json:"sid"`
	SourceChannel *string    `json:"source_channel,omitempty"`
	TaggedText    string     `json:"tagged_text"`
	TaskSid       string     `json:"task_sid"`
	URL           string     `json:"url"`
}

// Create creates a new task sample
// See https://www.twilio.com/docs/autopilot/api/task-sample#create-a-sample-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateSampleInput) (*CreateSampleResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new task sample
// See https://www.twilio.com/docs/autopilot/api/task-sample#create-a-sample-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateSampleInput) (*CreateSampleResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Tasks/{taskSid}/Samples",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
		},
	}

	if input == nil {
		input = &CreateSampleInput{}
	}

	response := &CreateSampleResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
