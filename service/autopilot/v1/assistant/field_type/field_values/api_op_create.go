// Package field_values contains auto-generated files. DO NOT MODIFY
package field_values

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateFieldValueInput defines the input fields for creating a new field value resource
type CreateFieldValueInput struct {
	Language  string  `validate:"required" form:"Language"`
	SynonymOf *string `form:"SynonymOf,omitempty"`
	Value     string  `validate:"required" form:"Value"`
}

// CreateFieldValueResponse defines the response fields for the created field value
type CreateFieldValueResponse struct {
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

// Create creates a new field value
// See https://www.twilio.com/docs/autopilot/api/field-value#create-a-fieldvalue-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFieldValueInput) (*CreateFieldValueResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new field value
// See https://www.twilio.com/docs/autopilot/api/field-value#create-a-fieldvalue-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFieldValueInput) (*CreateFieldValueResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/FieldTypes/{fieldTypeSid}/FieldValues",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"fieldTypeSid": c.fieldTypeSid,
		},
	}

	if input == nil {
		input = &CreateFieldValueInput{}
	}

	response := &CreateFieldValueResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
