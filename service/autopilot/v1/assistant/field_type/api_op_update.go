// Package field_type contains auto-generated files. DO NOT MODIFY
package field_type

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateFieldTypeInput defines the input fields for updating a field type
type UpdateFieldTypeInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
	UniqueName   *string `form:"UniqueName,omitempty"`
}

// UpdateFieldTypeResponse defines the response fields for the updated field type
type UpdateFieldTypeResponse struct {
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Update modifies a field type resource
// See https://www.twilio.com/docs/autopilot/api/field-type#update-a-fieldtype-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateFieldTypeInput) (*UpdateFieldTypeResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a field type resource
// See https://www.twilio.com/docs/autopilot/api/field-type#update-a-fieldtype-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateFieldTypeInput) (*UpdateFieldTypeResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/FieldTypes/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateFieldTypeInput{}
	}

	response := &UpdateFieldTypeResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
