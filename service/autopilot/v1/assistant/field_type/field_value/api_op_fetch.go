// Package field_value contains auto-generated files. DO NOT MODIFY
package field_value

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchFieldValueResponse defines the response fields for the retrieved field value
type FetchFieldValueResponse struct {
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FieldTypeSid string     `json:"field_type_sid"`
	Language     string     `json:"language"`
	Sid          string     `json:"sid"`
	SynonymOf    *string    `json:"synonym_of,omitempty"`
	URL          string     `json:"url"`
	Value        string     `json:"value"`
}

// Fetch retrieves a field value resource
// See https://www.twilio.com/docs/autopilot/api/field-value#fetch-a-fieldvalue-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFieldValueResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a field value resource
// See https://www.twilio.com/docs/autopilot/api/field-value#fetch-a-fieldvalue-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFieldValueResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/FieldTypes/{fieldTypeSid}/FieldValues/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"fieldTypeSid": c.fieldTypeSid,
			"sid":          c.sid,
		},
	}

	response := &FetchFieldValueResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
