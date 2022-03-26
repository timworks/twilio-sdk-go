// Package task contains auto-generated files. DO NOT MODIFY
package task

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchTaskResponse defines the response fields for the retrieved task
type FetchTaskResponse struct {
	AccountSid   string     `json:"account_sid"`
	ActionsURL   string     `json:"actions_url"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Fetch retrieves a task resource
// See https://www.twilio.com/docs/autopilot/api/task#fetch-a-task-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchTaskResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a task resource
// See https://www.twilio.com/docs/autopilot/api/task#fetch-a-task-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchTaskResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &FetchTaskResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
