// This is an autogenerated file. DO NOT MODIFY
package field_value

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetFieldValueResponse struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	FieldTypeSid string     `json:"field_type_sid"`
	Language     string     `json:"language"`
	Value        string     `json:"value"`
	SynonymOf    *string    `json:"synonym_of,omitempty"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	URL          string     `json:"url"`
}

func (c Client) Get() (*GetFieldValueResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetFieldValueResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/FieldTypes/{fieldTypeSid}/FieldValues/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"fieldTypeSid": c.fieldTypeSid,
			"sid":          c.sid,
		},
	}

	response := &GetFieldValueResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
