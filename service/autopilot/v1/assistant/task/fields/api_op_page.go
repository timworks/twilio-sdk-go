// Package fields contains auto-generated files. DO NOT MODIFY
package fields

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FieldsPageOptions defines the query options for the api operation
type FieldsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageFieldResponse struct {
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FieldType    string     `json:"field_type"`
	Sid          string     `json:"sid"`
	TaskSid      string     `json:"task_sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

type PageMetaResponse struct {
	FirstPageURL    string  `json:"first_page_url"`
	Key             string  `json:"key"`
	NextPageURL     *string `json:"next_page_url,omitempty"`
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url,omitempty"`
	URL             string  `json:"url"`
}

// FieldsPageResponse defines the response fields for the fields page
type FieldsPageResponse struct {
	Fields []PageFieldResponse `json:"fields"`
	Meta   PageMetaResponse    `json:"meta"`
}

// Page retrieves a page of fields
// See https://www.twilio.com/docs/autopilot/api/task-field#read-multiple-field-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *FieldsPageOptions) (*FieldsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of fields
// See https://www.twilio.com/docs/autopilot/api/task-field#read-multiple-field-resources for more details
func (c Client) PageWithContext(context context.Context, options *FieldsPageOptions) (*FieldsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks/{taskSid}/Fields",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FieldsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// FieldsPaginator defines the fields for makings paginated api calls
// Fields is an array of fields that have been returned from all of the page calls
type FieldsPaginator struct {
	options *FieldsPageOptions
	Page    *FieldsPage
	Fields  []PageFieldResponse
}

// NewFieldsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewFieldsPaginator() *FieldsPaginator {
	return c.NewFieldsPaginatorWithOptions(nil)
}

// NewFieldsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewFieldsPaginatorWithOptions(options *FieldsPageOptions) *FieldsPaginator {
	return &FieldsPaginator{
		options: options,
		Page: &FieldsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Fields: make([]PageFieldResponse, 0),
	}
}

// FieldsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageFieldResponse or error that is returned from the api call(s)
type FieldsPage struct {
	client *Client

	CurrentPage *FieldsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *FieldsPaginator) CurrentPage() *FieldsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *FieldsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *FieldsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *FieldsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &FieldsPageOptions{}
	}

	if p.CurrentPage() != nil {
		nextPage := p.CurrentPage().Meta.NextPageURL

		if nextPage == nil {
			return false
		}

		parsedURL, err := url.Parse(*nextPage)
		if err != nil {
			p.Page.Error = err
			return false
		}

		options.PageToken = utils.String(parsedURL.Query().Get("PageToken"))

		page, pageErr := strconv.Atoi(parsedURL.Query().Get("Page"))
		if pageErr != nil {
			p.Page.Error = pageErr
			return false
		}
		options.Page = utils.Int(page)

		pageSize, pageSizeErr := strconv.Atoi(parsedURL.Query().Get("PageSize"))
		if pageSizeErr != nil {
			p.Page.Error = pageSizeErr
			return false
		}
		options.PageSize = utils.Int(pageSize)
	}

	resp, err := p.Page.client.PageWithContext(context, options)
	p.Page.CurrentPage = resp
	p.Page.Error = err

	if p.Page.Error == nil {
		p.Fields = append(p.Fields, resp.Fields...)
	}

	return p.Page.Error == nil
}
