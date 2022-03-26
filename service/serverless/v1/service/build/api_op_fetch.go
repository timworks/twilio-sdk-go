// Package build contains auto-generated files. DO NOT MODIFY
package build

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

type FetchAssetVersion struct {
	AccountSid  string    `json:"account_sid"`
	AssetSid    string    `json:"asset_sid"`
	DateCreated time.Time `json:"date_created"`
	Path        string    `json:"path"`
	ServiceSid  string    `json:"service_sid"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Visibility  string    `json:"visibility"`
}

type FetchDependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type FetchFunctionVersion struct {
	AccountSid  string    `json:"account_sid"`
	DateCreated time.Time `json:"date_created"`
	FunctionSid string    `json:"function_sid"`
	Path        string    `json:"path"`
	ServiceSid  string    `json:"service_sid"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Visibility  string    `json:"visibility"`
}

// FetchBuildResponse defines the response fields for the retrieved build
type FetchBuildResponse struct {
	AccountSid       string                  `json:"account_sid"`
	AssetVersions    *[]FetchAssetVersion    `json:"asset_versions,omitempty"`
	DateCreated      time.Time               `json:"date_created"`
	DateUpdated      *time.Time              `json:"date_updated,omitempty"`
	Dependencies     *[]FetchDependency      `json:"dependencies,omitempty"`
	FunctionVersions *[]FetchFunctionVersion `json:"function_versions,omitempty"`
	Runtime          string                  `json:"runtime"`
	ServiceSid       string                  `json:"service_sid"`
	Sid              string                  `json:"sid"`
	Status           string                  `json:"status"`
	URL              string                  `json:"url"`
}

// Fetch retrieves a build resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/build#fetch-a-build-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchBuildResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a build resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/build#fetch-a-build-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchBuildResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Builds/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchBuildResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
