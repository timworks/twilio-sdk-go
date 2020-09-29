// Package participants contains auto-generated files. DO NOT MODIFY
package participants

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateParticipantInput defines the input fields for creating a new participant resource
type CreateParticipantInput struct {
	Attributes                       *string            `form:"Attributes,omitempty"`
	DateCreated                      *utils.RFC2822Time `form:"DateCreated,omitempty"`
	DateUpdated                      *utils.RFC2822Time `form:"DateUpdated,omitempty"`
	Identity                         *string            `form:"Identity,omitempty"`
	MessagingBindingAddress          *string            `form:"MessagingBinding.Address,omitempty"`
	MessagingBindingProjectedAddress *string            `form:"MessagingBinding.ProjectedAddress,omitempty"`
	MessagingBindingProxyAddress     *string            `form:"MessagingBinding.ProxyAddress,omitempty"`
	RoleSid                          *string            `form:"RoleSid,omitempty"`
}

type CreateParticipantResponseMessageBinding struct {
	Address          string  `json:"address"`
	ProjectedAddress *string `json:"projected_address,omitempty"`
	ProxyAddress     string  `json:"proxy_address"`
	Type             string  `json:"type"`
}

// CreateParticipantResponse defines the response fields for the created participant
type CreateParticipantResponse struct {
	AccountSid       string                                   `json:"account_sid"`
	Attributes       string                                   `json:"attributes"`
	ConversationSid  string                                   `json:"conversation_sid"`
	DateCreated      time.Time                                `json:"date_created"`
	DateUpdated      *time.Time                               `json:"date_updated,omitempty"`
	Identity         *string                                  `json:"identity,omitempty"`
	MessagingBinding *CreateParticipantResponseMessageBinding `json:"messaging_binding,omitempty"`
	RoleSid          *string                                  `json:"role_sid,omitempty"`
	Sid              string                                   `json:"sid"`
	URL              string                                   `json:"url"`
}

// Create creates a new participant
// See https://www.twilio.com/docs/conversations/api/conversation-participant-resource#add-a-conversation-participant-sms for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateParticipantInput) (*CreateParticipantResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new participant
// See https://www.twilio.com/docs/conversations/api/conversation-participant-resource#add-a-conversation-participant-sms for more details
func (c Client) CreateWithContext(context context.Context, input *CreateParticipantInput) (*CreateParticipantResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Conversations/{conversationSid}/Participants",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
		},
	}

	if input == nil {
		input = &CreateParticipantInput{}
	}

	response := &CreateParticipantResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
