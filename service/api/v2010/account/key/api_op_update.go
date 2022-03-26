// Package key contains auto-generated files. DO NOT MODIFY
package key

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// UpdateKeyInput defines input fields for updating a API Key
type UpdateKeyInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

// UpdateKeyResponse defines the response fields for the updated API Key
type UpdateKeyResponse struct {
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Update modifies a key resource. The secret is not returned for security reasons
// See https://www.twilio.com/docs/iam/keys/api-key-resource#update-a-key-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateKeyInput) (*UpdateKeyResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a key resource. The secret is not returned for security reasons
// See https://www.twilio.com/docs/iam/keys/api-key-resource#update-a-key-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateKeyInput) (*UpdateKeyResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Keys/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateKeyInput{}
	}

	response := &UpdateKeyResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
