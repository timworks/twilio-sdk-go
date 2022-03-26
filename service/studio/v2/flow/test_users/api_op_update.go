// Package test_users contains auto-generated files. DO NOT MODIFY
package test_users

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateTestUsersInput defines input fields for updating a test users resource
type UpdateTestUsersInput struct {
	TestUsers []string `validate:"required" form:"TestUsers"`
}

// UpdateTestUsersResponse defines the response fields for the updated test users
type UpdateTestUsersResponse struct {
	Sid       string   `json:"sid"`
	TestUsers []string `json:"test_users"`
	URL       string   `json:"url"`
}

// Update modifies a test users resource
// See https://www.twilio.com/docs/studio/rest-api/v2/test-user#update-a-testuser-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateTestUsersInput) (*UpdateTestUsersResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a test users resource
// See https://www.twilio.com/docs/studio/rest-api/v2/test-user#update-a-testuser-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateTestUsersInput) (*UpdateTestUsersResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Flows/{flowSid}/TestUsers",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"flowSid": c.flowSid,
		},
	}

	if input == nil {
		input = &UpdateTestUsersInput{}
	}

	response := &UpdateTestUsersResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
