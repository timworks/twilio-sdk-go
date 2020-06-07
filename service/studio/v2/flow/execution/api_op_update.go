// This is an autogenerated file. DO NOT MODIFY
package execution

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateExecutionInput struct {
	Status string `validate:"required" form:"Status"`
}

type UpdateExecutionOutput struct {
	Sid                   string      `json:"sid"`
	AccountSid            string      `json:"account_sid"`
	FlowSid               string      `json:"flow_sid"`
	Context               interface{} `json:"context"`
	ContactChannelAddress string      `json:"contact_channel_address"`
	Status                string      `json:"status"`
	DateCreated           time.Time   `json:"date_created"`
	DateUpdated           *time.Time  `json:"date_updated,omitempty"`
	URL                   string      `json:"url"`
}

func (c Client) Update(input *UpdateExecutionInput) (*UpdateExecutionOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateExecutionInput) (*UpdateExecutionOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Flows/{flowSid}/Executions/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"flowSid": c.flowSid,
			"sid":     c.sid,
		},
	}

	output := &UpdateExecutionOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
