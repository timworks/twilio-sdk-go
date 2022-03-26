// Package model_build contains auto-generated files. DO NOT MODIFY
package model_build

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a model build resource from the account
// See https://www.twilio.com/docs/autopilot/api/model-build#delete-a-modelbuild-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a model build resource from the account
// See https://www.twilio.com/docs/autopilot/api/model-build#delete-a-modelbuild-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Assistants/{assistantSid}/ModelBuilds/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
