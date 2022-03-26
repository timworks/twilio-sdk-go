// Package flow_validation contains auto-generated files. DO NOT MODIFY
package flow_validation

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// ValidateFlowInput defines the input fields for validating a flow
type ValidateFlowInput struct {
	CommitMessage *string `form:"CommitMessage,omitempty"`
	Definition    string  `validate:"required" form:"Definition"`
	FriendlyName  string  `validate:"required" form:"FriendlyName"`
	Status        string  `validate:"required" form:"Status"`
}

// ValidateFlowResponse defines the response fields for the validating a flow
type ValidateFlowResponse struct {
	Valid bool `json:"valid"`
}

// Validate validate a flow
// See https://www.twilio.com/docs/studio/rest-api/v2/flow-validate#validate-flow for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Validate(input *ValidateFlowInput) (*ValidateFlowResponse, error) {
	return c.ValidateWithContext(context.Background(), input)
}

// ValidateWithContext validate a flow
// See https://www.twilio.com/docs/studio/rest-api/v2/flow-validate#validate-flow for more details
func (c Client) ValidateWithContext(context context.Context, input *ValidateFlowInput) (*ValidateFlowResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Flows/Validate",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &ValidateFlowInput{}
	}

	response := &ValidateFlowResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
