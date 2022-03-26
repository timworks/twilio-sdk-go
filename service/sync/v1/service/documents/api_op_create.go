// Package documents contains auto-generated files. DO NOT MODIFY
package documents

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateDocumentInput defines the input fields for creating a new document resource
type CreateDocumentInput struct {
	Data       *string `form:"Data,omitempty"`
	Ttl        *int    `form:"Ttl,omitempty"`
	UniqueName *string `form:"UniqueName,omitempty"`
}

// CreateDocumentResponse defines the response fields for the created document
type CreateDocumentResponse struct {
	AccountSid  string                 `json:"account_sid"`
	CreatedBy   string                 `json:"created_by"`
	Data        map[string]interface{} `json:"data"`
	DateCreated time.Time              `json:"date_created"`
	DateExpires *time.Time             `json:"date_expires,omitempty"`
	DateUpdated *time.Time             `json:"date_updated,omitempty"`
	Revision    string                 `json:"revision"`
	ServiceSid  string                 `json:"service_Sid"`
	Sid         string                 `json:"sid"`
	URL         string                 `json:"url"`
	UniqueName  *string                `json:"unique_name,omitempty"`
}

// Create creates a new document
// See https://www.twilio.com/docs/sync/api/document-resource#create-a-document-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateDocumentInput) (*CreateDocumentResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new document
// See https://www.twilio.com/docs/sync/api/document-resource#create-a-document-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateDocumentInput) (*CreateDocumentResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Documents",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateDocumentInput{}
	}

	response := &CreateDocumentResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
