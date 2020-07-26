// This is an autogenerated file. DO NOT MODIFY
package item

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetListItemResponse struct {
	AccountSid  string                 `json:"account_sid"`
	CreatedBy   string                 `json:"created_by"`
	Data        map[string]interface{} `json:"data"`
	DateCreated time.Time              `json:"date_created"`
	DateExpires *time.Time             `json:"date_expires,omitempty"`
	DateUpdated *time.Time             `json:"date_updated,omitempty"`
	Index       int                    `json:"index"`
	ListSid     string                 `json:"list_sid"`
	Revision    string                 `json:"revision"`
	ServiceSid  string                 `json:"service_Sid"`
	URL         string                 `json:"url"`
}

func (c Client) Get() (*GetListItemResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetListItemResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Lists/{listSid}/Items/{index}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"listSid":    c.listSid,
			"index":      strconv.Itoa(c.index),
		},
	}

	response := &GetListItemResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
