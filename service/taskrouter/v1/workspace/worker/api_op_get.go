// This is an autogenerated file. DO NOT MODIFY
package worker

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetWorkerOutput struct {
	Sid               string      `json:"sid"`
	AccountSid        string      `json:"account_sid"`
	WorkspaceSid      string      `json:"workspace_sid"`
	ActivitySid       string      `json:"activity_sid"`
	FriendlyName      string      `json:"friendly_name"`
	ActivityName      string      `json:"activity_name"`
	Attributes        interface{} `json:"attributes"`
	Available         bool        `json:"available"`
	DateCreated       time.Time   `json:"date_created"`
	DateUpdated       *time.Time  `json:"date_updated,omitempty"`
	DateStatusChanged *time.Time  `json:"date_status_changed,omitempty"`
	URL               string      `json:"url"`
}

func (c Client) Get() (*GetWorkerOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetWorkerOutput, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workers/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	output := &GetWorkerOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}
