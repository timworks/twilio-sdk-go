// Package field contains auto-generated files. DO NOT MODIFY
package field

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a task field resource from the account
// See https://www.twilio.com/docs/autopilot/api/task-field#delete-a-field-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a task field resource from the account
// See https://www.twilio.com/docs/autopilot/api/task-field#delete-a-field-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Assistants/{assistantSid}/Tasks/{taskSid}/Fields/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
			"sid":          c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
