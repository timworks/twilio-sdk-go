// Package field_types contains auto-generated files. DO NOT MODIFY
package field_types

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateFieldTypeInput defines the input fields for creating a new field type resource
type CreateFieldTypeInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
	UniqueName   string  `validate:"required" form:"UniqueName"`
}

// CreateFieldTypeResponse defines the response fields for the created field type
type CreateFieldTypeResponse struct {
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Create creates a new field type
// See https://www.twilio.com/docs/autopilot/api/field-type#create-a-fieldtype-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFieldTypeInput) (*CreateFieldTypeResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new field type
// See https://www.twilio.com/docs/autopilot/api/field-type#create-a-fieldtype-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFieldTypeInput) (*CreateFieldTypeResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/FieldTypes",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	if input == nil {
		input = &CreateFieldTypeInput{}
	}

	response := &CreateFieldTypeResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
