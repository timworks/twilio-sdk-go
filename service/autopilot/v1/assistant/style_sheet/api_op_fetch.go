// Package style_sheet contains auto-generated files. DO NOT MODIFY
package style_sheet

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchStyleSheetResponse defines the response fields for the retrieved style sheet
type FetchStyleSheetResponse struct {
	AccountSid   string                 `json:"account_sid"`
	AssistantSid string                 `json:"assistant_sid"`
	Data         map[string]interface{} `json:"data"`
	URL          string                 `json:"url"`
}

// Fetch retrieves a style sheet resource
// See https://www.twilio.com/docs/autopilot/api/assistant/stylesheet#fetch-a-stylesheet-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchStyleSheetResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a style sheet resource
// See https://www.twilio.com/docs/autopilot/api/assistant/stylesheet#fetch-a-stylesheet-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchStyleSheetResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/StyleSheet",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	response := &FetchStyleSheetResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
