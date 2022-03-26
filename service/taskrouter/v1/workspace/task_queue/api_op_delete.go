// Package task_queue contains auto-generated files. DO NOT MODIFY
package task_queue

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a task queue resource from the account
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-delete for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a task queue resource from the account
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-delete for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Workspaces/{workspaceSid}/TaskQueues/{sid}",
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
