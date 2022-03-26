// Package defaults contains auto-generated files. DO NOT MODIFY
package defaults

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchDefaultResponse defines the response fields for the retrieved defaults
type FetchDefaultResponse struct {
	AccountSid   string                 `json:"account_sid"`
	AssistantSid string                 `json:"assistant_sid"`
	Data         map[string]interface{} `json:"data"`
	URL          string                 `json:"url"`
}

// Fetch retrieves a defaults resource
// See https://www.twilio.com/docs/autopilot/api/assistant/defaults#fetch-a-defaults-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchDefaultResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a defaults resource
// See https://www.twilio.com/docs/autopilot/api/assistant/defaults#fetch-a-defaults-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchDefaultResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Defaults",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	response := &FetchDefaultResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
