// Package feedbacks contains auto-generated files. DO NOT MODIFY
package feedbacks

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// CreateFeedbackInput defines input fields for creating a new feedback resource
type CreateFeedbackInput struct {
	Outcome *string `form:"Outcome,omitempty"`
}

// CreateFeedbackResponse defines the response fields for the created feedback
type CreateFeedbackResponse struct {
	AccountSid  string             `json:"account_sid"`
	DateCreated utils.RFC2822Time  `json:"date_created"`
	DateUpdated *utils.RFC2822Time `json:"date_updated,omitempty"`
	MessageSid  string             `json:"message_sid"`
	Outcome     string             `json:"outcome"`
	Sid         string             `json:"sid"`
}

// Create creates a new feedback resource
// See https://www.twilio.com/docs/sms/api/message-feedback-resource#create-a-messagefeedback-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFeedbackInput) (*CreateFeedbackResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new feedback resource
// See https://www.twilio.com/docs/sms/api/message-feedback-resource#create-a-messagefeedback-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFeedbackInput) (*CreateFeedbackResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Messages/{messageSid}/Feedback.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"messageSid": c.messageSid,
		},
	}

	if input == nil {
		input = &CreateFeedbackInput{}
	}

	response := &CreateFeedbackResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
