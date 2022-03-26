// Package workflow contains auto-generated files. DO NOT MODIFY
package workflow

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a workflow resource from the account
// See https://www.twilio.com/docs/taskrouter/api/workflow#delete-a-workflow-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a workflow resource from the account
// See https://www.twilio.com/docs/taskrouter/api/workflow#delete-a-workflow-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Workspaces/{workspaceSid}/Workflows/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
