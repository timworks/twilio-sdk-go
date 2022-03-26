// Package message contains auto-generated files. DO NOT MODIFY
package message

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// UpdateMessageInput defines input fields for updating a message resource
type UpdateMessageInput struct {
	Body string `form:"Body"`
}

// UpdateMessageResponse defines the response fields for the updated message
type UpdateMessageResponse struct {
	APIVersion          string             `json:"api_version"`
	AccountSid          string             `json:"account_sid"`
	Body                string             `json:"body"`
	DateCreated         utils.RFC2822Time  `json:"date_created"`
	DateSent            utils.RFC2822Time  `json:"date_sent"`
	DateUpdated         *utils.RFC2822Time `json:"date_updated,omitempty"`
	Direction           string             `json:"direction"`
	ErrorCode           *int               `json:"error_code,omitempty"`
	ErrorMessage        *string            `json:"error_message,omitempty"`
	From                *string            `json:"from,omitempty"`
	MessagingServiceSid *string            `json:"messaging_service_sid,omitempty"`
	NumMedia            string             `json:"num_media"`
	NumSegments         string             `json:"num_segments"`
	Price               *string            `json:"price,omitempty"`
	PriceUnit           string             `json:"price_unit"`
	Sid                 string             `json:"sid"`
	Status              string             `json:"status"`
	To                  string             `json:"to"`
}

// Update modifies a message resource
// See https://www.twilio.com/docs/sms/api/message-resource#update-a-message-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateMessageInput) (*UpdateMessageResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a message resource
// See https://www.twilio.com/docs/sms/api/message-resource#update-a-message-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateMessageInput) (*UpdateMessageResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Messages/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateMessageInput{}
	}

	response := &UpdateMessageResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
