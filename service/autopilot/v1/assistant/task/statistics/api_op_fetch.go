// Package statistics contains auto-generated files. DO NOT MODIFY
package statistics

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchStatisticsResponse defines the response fields for the retrieved task statistics
type FetchStatisticsResponse struct {
	AccountSid   string `json:"account_sid"`
	AssistantSid string `json:"assistant_sid"`
	FieldsCount  int    `json:"fields_count"`
	SamplesCount int    `json:"samples_count"`
	TaskSid      string `json:"task_sid"`
	URL          string `json:"url"`
}

// Fetch retrieves a task statistic resource
// See https://www.twilio.com/docs/autopilot/api/task-statistics#fetch-a-taskstatistics-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchStatisticsResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a task statistic resource
// See https://www.twilio.com/docs/autopilot/api/task-statistics#fetch-a-taskstatistics-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchStatisticsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks/{taskSid}/Statistics",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
		},
	}

	response := &FetchStatisticsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
