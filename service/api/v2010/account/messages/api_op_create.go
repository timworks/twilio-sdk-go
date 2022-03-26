// Package messages contains auto-generated files. DO NOT MODIFY
package messages

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// CreateMessageInput defines input fields for creating a new message resource
type CreateMessageInput struct {
	AddressRetention    *string   `form:"AddressRetention,omitempty"`
	ApplicationSid      *string   `form:"ApplicationSid,omitempty"`
	Attempt             *int      `form:"Attempt,omitempty"`
	Body                *string   `form:"Body,omitempty"`
	ContentRetention    *string   `form:"ContentRetention,omitempty"`
	ForceDelivery       *bool     `form:"ForceDelivery,omitempty"`
	From                *string   `form:"From,omitempty"`
	MaxPrice            *string   `form:"MaxPrice,omitempty"`
	MediaURLs           *[]string `form:"MediaUrl,omitempty"`
	MessagingServiceSid *string   `form:"MessagingServiceSid,omitempty"`
	PersistentActions   *[]string `form:"PersistentAction,omitempty"`
	ProvideFeedback     *bool     `form:"ProvideFeedback,omitempty"`
	SmartEncoded        *bool     `form:"SmartEncoded,omitempty"`
	StatusCallback      *string   `form:"StatusCallback,omitempty"`
	To                  string    `validate:"required" form:"To"`
	ValidityPeriod      *int      `form:"ValidityPeriod,omitempty"`
}

// CreateMessageResponse defines the response fields for the created message
type CreateMessageResponse struct {
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

// Create creates a new message resource
// See https://www.twilio.com/docs/sms/api/message-resource#create-a-message-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateMessageInput) (*CreateMessageResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new message resource
// See https://www.twilio.com/docs/sms/api/message-resource#create-a-message-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateMessageInput) (*CreateMessageResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Messages.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateMessageInput{}
	}

	response := &CreateMessageResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
