// Package sample contains auto-generated files. DO NOT MODIFY
package sample

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateSampleInput defines the input fields for updating a task sample
type UpdateSampleInput struct {
	Language      *string `form:"Language,omitempty"`
	SourceChannel *string `form:"SourceChannel,omitempty"`
	TaggedText    *string `form:"TaggedText,omitempty"`
}

// UpdateSampleResponse defines the response fields for the updated task sample
type UpdateSampleResponse struct {
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

// Update modifies a task sample resource
// See https://www.twilio.com/docs/autopilot/api/task-sample#update-a-sample-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSampleInput) (*UpdateSampleResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a task sample resource
// See https://www.twilio.com/docs/autopilot/api/task-sample#update-a-sample-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSampleInput) (*UpdateSampleResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Tasks/{taskSid}/Samples/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateSampleInput{}
	}

	response := &UpdateSampleResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
