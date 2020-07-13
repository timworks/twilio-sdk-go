// This is an autogenerated file. DO NOT MODIFY
package message

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type UpdateMessageInput struct {
	Author      *string            `form:"Author,omitempty"`
	Body        *string            `form:"Body,omitempty"`
	DateCreated *utils.RFC2822Time `form:"DateCreated,omitempty"`
	DateUpdated *utils.RFC2822Time `form:"DateUpdated,omitempty"`
	Attributes  *string            `form:"Attributes.Filters,omitempty"`
}

type UpdateMessageOutputMedia struct {
	Sid         string `json:"sid"`
	ContentType string `json:"content_type"`
	Filename    string `json:"filename"`
	Size        int    `json:"size"`
}

type UpdateMessageOutput struct {
	Sid             string                      `json:"sid"`
	AccountSid      string                      `json:"account_sid"`
	ConversationSid string                      `json:"conversation_sid"`
	ParticipantSid  *string                     `json:"participant_sid,omitempty"`
	Body            *string                     `json:"body,omitempty"`
	Index           int                         `json:"index"`
	Author          string                      `json:"author"`
	Attributes      string                      `json:"attributes"`
	Media           *[]UpdateMessageOutputMedia `json:"media,omitempty"`
	DateCreated     time.Time                   `json:"date_created"`
	DateUpdated     *time.Time                  `json:"date_updated,omitempty"`
	URL             string                      `json:"url"`
}

func (c Client) Update(input *UpdateMessageInput) (*UpdateMessageOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateMessageInput) (*UpdateMessageOutput, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Conversations/{conversationSid}/Messages/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	output := &UpdateMessageOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
