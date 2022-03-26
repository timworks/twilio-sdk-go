// Package content contains auto-generated files. DO NOT MODIFY
package content

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchContentResponse defines the response fields for the retrieved function version content
type FetchContentResponse struct {
	AccountSid  string `json:"account_sid"`
	Content     string `json:"content"`
	FunctionSid string `json:"function_sid"`
	ServiceSid  string `json:"service_sid"`
	Sid         string `json:"sid"`
	URL         string `json:"url"`
}

// Fetch retrieves a function version content resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function-version/function-version-content#fetch-a-functionversioncontent-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchContentResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a function version content resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function-version/function-version-content#fetch-a-functionversioncontent-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchContentResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Functions/{functionSid}/Versions/{versionSid}/Content",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"functionSid": c.functionSid,
			"versionSid":  c.versionSid,
		},
	}

	response := &FetchContentResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
