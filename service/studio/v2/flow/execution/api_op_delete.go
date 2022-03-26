// Package execution contains auto-generated files. DO NOT MODIFY
package execution

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a execution resource from the account
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#delete-an-execution for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a execution resource from the account
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#delete-an-execution for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Flows/{flowSid}/Executions/{sid}",
		PathParams: map[string]string{
			"flowSid": c.flowSid,
			"sid":     c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
