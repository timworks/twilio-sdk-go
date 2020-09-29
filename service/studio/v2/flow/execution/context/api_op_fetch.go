// Package context contains auto-generated files. DO NOT MODIFY
package context

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchContextResponse defines the response fields for the retrieved execution context
type FetchContextResponse struct {
	AccountSid   string      `json:"account_sid"`
	Context      interface{} `json:"context"`
	ExecutionSid string      `json:"execution_sid"`
	FlowSid      string      `json:"flow_sid"`
	URL          string      `json:"url"`
}

// Fetch retrieves a execution context resource
// See https://www.twilio.com/docs/studio/rest-api/v2/execution-context#fetch-a-single-execution-context for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchContextResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a execution context resource
// See https://www.twilio.com/docs/studio/rest-api/v2/execution-context#fetch-a-single-execution-context for more details
func (c Client) FetchWithContext(context context.Context) (*FetchContextResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/Executions/{executionSid}/Context",
		PathParams: map[string]string{
			"flowSid":      c.flowSid,
			"executionSid": c.executionSid,
		},
	}

	response := &FetchContextResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
