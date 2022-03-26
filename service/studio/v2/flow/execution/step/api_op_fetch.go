// Package step contains auto-generated files. DO NOT MODIFY
package step

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchStepResponse defines the response fields for the retrieved step
type FetchStepResponse struct {
	AccountSid       string      `json:"account_sid"`
	Context          interface{} `json:"context"`
	DateCreated      time.Time   `json:"date_created"`
	DateUpdated      *time.Time  `json:"date_updated,omitempty"`
	ExecutionSid     string      `json:"execution_sid"`
	FlowSid          string      `json:"flow_sid"`
	Name             string      `json:"name"`
	Sid              string      `json:"sid"`
	TransitionedFrom string      `json:"transitioned_from"`
	TransitionedTo   string      `json:"transitioned_to"`
	URL              string      `json:"url"`
}

// Fetch retrieves a step resource
// See https://www.twilio.com/docs/studio/rest-api/v2/step#fetch-a-step-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchStepResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a step resource
// See https://www.twilio.com/docs/studio/rest-api/v2/step#fetch-a-step-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchStepResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/Executions/{executionSid}/Steps/{sid}",
		PathParams: map[string]string{
			"flowSid":      c.flowSid,
			"executionSid": c.executionSid,
			"sid":          c.sid,
		},
	}

	response := &FetchStepResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
