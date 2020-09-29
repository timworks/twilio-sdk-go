// Package flex_flow contains auto-generated files. DO NOT MODIFY
package flex_flow

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateFlexFlowInput defines input fields for updating a flex flow resource
type UpdateFlexFlowInput struct {
	ChannelType                  *string `form:"ChannelType,omitempty"`
	ChatServiceSid               *string `form:"ChatServiceSid,omitempty"`
	ContactIdentity              *string `form:"ContactIdentity,omitempty"`
	Enabled                      *bool   `form:"Enabled,omitempty"`
	FriendlyName                 *string `form:"FriendlyName,omitempty"`
	IntegrationChannel           *string `form:"Integration.Channel,omitempty"`
	IntegrationCreationOnMessage *bool   `form:"Integration.CreationOnMessage,omitempty"`
	IntegrationFlowSid           *string `form:"Integration.FlowSid,omitempty"`
	IntegrationPriority          *int    `form:"Integration.Priority,omitempty"`
	IntegrationRetryCount        *int    `form:"Integration.RetryCount,omitempty"`
	IntegrationTimeout           *int    `form:"Integration.Timeout,omitempty"`
	IntegrationType              *string `form:"IntegrationType,omitempty"`
	IntegrationURL               *string `form:"Integration.Url,omitempty"`
	IntegrationWorkflowSid       *string `form:"Integration.WorkflowSid,omitempty"`
	IntegrationWorkspaceSid      *string `form:"Integration.WorkspaceSid,omitempty"`
	JanitorEnabled               *bool   `form:"JanitorEnabled,omitempty"`
	LongLived                    *bool   `form:"LongLived,omitempty"`
}

type UpdateFlexFlowResponseIntegration struct {
	Channel           *string `json:"channel,omitempty"`
	CreationOnMessage *bool   `json:"creation_on_message,omitempty"`
	FlowSid           *string `json:"flow_sid,omitempty"`
	Priority          *int    `json:"priority,omitempty"`
	RetryCount        *int    `json:"retry_count,omitempty"`
	Timeout           *int    `json:"timeout,omitempty"`
	URL               *string `json:"url,omitempty"`
	WorkflowSid       *string `json:"workflow_sid,omitempty"`
	WorkspaceSid      *string `json:"workspace_sid,omitempty"`
}

// UpdateFlexFlowResponse defines the response fields for the updated flex flow
type UpdateFlexFlowResponse struct {
	AccountSid      string                             `json:"account_sid"`
	ChannelType     string                             `json:"channel_type"`
	ChatServiceSid  string                             `json:"chat_service_sid"`
	ContactIdentity *string                            `json:"contact_identity,omitempty"`
	DateCreated     time.Time                          `json:"date_created"`
	DateUpdated     *time.Time                         `json:"date_updated,omitempty"`
	Enabled         bool                               `json:"enabled"`
	FriendlyName    string                             `json:"friendly_name"`
	Integration     *UpdateFlexFlowResponseIntegration `json:"integration,omitempty"`
	IntegrationType *string                            `json:"integration_type,omitempty"`
	JanitorEnabled  *bool                              `json:"janitor_enabled,omitempty"`
	LongLived       *bool                              `json:"long_lived,omitempty"`
	Sid             string                             `json:"sid"`
	URL             string                             `json:"url"`
}

// Update modifies a flex flow resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateFlexFlowInput) (*UpdateFlexFlowResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a flex flow resource
func (c Client) UpdateWithContext(context context.Context, input *UpdateFlexFlowInput) (*UpdateFlexFlowResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/FlexFlows/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateFlexFlowInput{}
	}

	response := &UpdateFlexFlowResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
