// This is an autogenerated file. DO NOT MODIFY
package messages

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateChannelMessageInput struct {
	From          *string    `form:"From,omitempty"`
	Attributes    *string    `form:"Attributes,omitempty"`
	DateCreated   *time.Time `form:"DateCreated,omitempty"`
	DateUpdated   *time.Time `form:"DateUpdated,omitempty"`
	LastUpdatedBy *string    `form:"LastUpdatedBy,omitempty"`
	Body          *string    `form:"Body,omitempty"`
	MediaSid      *string    `form:"MediaSid,omitempty"`
}

type CreateChannelMessageResponse struct {
	Sid           string                  `json:"sid"`
	AccountSid    string                  `json:"account_sid"`
	ServiceSid    string                  `json:"service_sid"`
	ChannelSid    string                  `json:"channel_sid"`
	To            *string                 `json:"to,omitempty"`
	Attributes    *string                 `json:"attributes,omitempty"`
	LastUpdatedBy *string                 `json:"last_updated_by,omitempty"`
	WasEdited     *bool                   `json:"was_edited,omitempty"`
	From          *string                 `json:"from,omitempty"`
	Body          *string                 `json:"body,omitempty"`
	Type          *string                 `json:"type,omitempty"`
	Index         *int                    `json:"index,omitempty"`
	Media         *map[string]interface{} `json:"media,omitempty"`
	DateCreated   time.Time               `json:"date_created"`
	DateUpdated   *time.Time              `json:"date_updated,omitempty"`
	URL           string                  `json:"url"`
}

func (c Client) Create(input *CreateChannelMessageInput) (*CreateChannelMessageResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateChannelMessageInput) (*CreateChannelMessageResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels/{channelSid}/Messages",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
	}

	response := &CreateChannelMessageResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
