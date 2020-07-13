// This is an autogenerated file. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetConversationWebhookOutputConfiguration struct {
	Url         *string   `json:"url,omitempty"`
	Method      *string   `json:"method,omitempty"`
	Filters     *[]string `json:"filters,omitempty"`
	Triggers    *[]string `json:"triggers,omitempty"`
	FlowSid     *string   `json:"flow_sid,omitempty"`
	ReplayAfter *int      `json:"replay_after,omitempty"`
}

type GetConversationWebhookOutput struct {
	Sid             string                                    `json:"sid"`
	AccountSid      string                                    `json:"account_sid"`
	ConversationSid string                                    `json:"conversation_sid"`
	Target          string                                    `json:"target"`
	Configuration   GetConversationWebhookOutputConfiguration `json:"configuration"`
	DateCreated     time.Time                                 `json:"date_created"`
	DateUpdated     *time.Time                                `json:"date_updated,omitempty"`
	URL             string                                    `json:"url"`
}

func (c Client) Get() (*GetConversationWebhookOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetConversationWebhookOutput, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Conversations/{conversationSid}/Webhooks/{sid}",
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	output := &GetConversationWebhookOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}
