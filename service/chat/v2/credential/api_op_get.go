// This is an autogenerated file. DO NOT MODIFY
package credential

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetCredentialResponse struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Type         string     `json:"type"`
	Sandbox      *string    `json:"sandbox,omitempty"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	URL          string     `json:"url"`
}

func (c Client) Get() (*GetCredentialResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetCredentialResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Credentials/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &GetCredentialResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
