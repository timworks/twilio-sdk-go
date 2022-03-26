// Package sample contains auto-generated files. DO NOT MODIFY
package sample

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchSampleResponse defines the response fields for the retrieved task sample
type FetchSampleResponse struct {
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

// Fetch retrieves a task sample resource
// See https://www.twilio.com/docs/autopilot/api/task-sample#fetch-a-sample-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchSampleResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a task sample resource
// See https://www.twilio.com/docs/autopilot/api/task-sample#fetch-a-sample-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchSampleResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks/{taskSid}/Samples/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
			"sid":          c.sid,
		},
	}

	response := &FetchSampleResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
