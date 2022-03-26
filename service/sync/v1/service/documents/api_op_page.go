// Package documents contains auto-generated files. DO NOT MODIFY
package documents

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// DocumentsPageOptions defines the query options for the api operation
type DocumentsPageOptions struct {
	PageSize    *int
	Page        *int
	PageToken   *string
	HideExpired *bool
}

type PageDocumentResponse struct {
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

type PageMetaResponse struct {
	FirstPageURL    string  `json:"first_page_url"`
	Key             string  `json:"key"`
	NextPageURL     *string `json:"next_page_url,omitempty"`
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url,omitempty"`
	URL             string  `json:"url"`
}

// DocumentsPageResponse defines the response fields for the documents page
type DocumentsPageResponse struct {
	Documents []PageDocumentResponse `json:"documents"`
	Meta      PageMetaResponse       `json:"meta"`
}

// Page retrieves a page of documents
// See https://www.twilio.com/docs/sync/api/document-resource#read-multiple-document-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *DocumentsPageOptions) (*DocumentsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of documents
// See https://www.twilio.com/docs/sync/api/document-resource#read-multiple-document-resources for more details
func (c Client) PageWithContext(context context.Context, options *DocumentsPageOptions) (*DocumentsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Documents",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &DocumentsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DocumentsPaginator defines the fields for makings paginated api calls
// Documents is an array of documents that have been returned from all of the page calls
type DocumentsPaginator struct {
	options   *DocumentsPageOptions
	Page      *DocumentsPage
	Documents []PageDocumentResponse
}

// NewDocumentsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewDocumentsPaginator() *DocumentsPaginator {
	return c.NewDocumentsPaginatorWithOptions(nil)
}

// NewDocumentsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewDocumentsPaginatorWithOptions(options *DocumentsPageOptions) *DocumentsPaginator {
	return &DocumentsPaginator{
		options: options,
		Page: &DocumentsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Documents: make([]PageDocumentResponse, 0),
	}
}

// DocumentsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageDocumentResponse or error that is returned from the api call(s)
type DocumentsPage struct {
	client *Client

	CurrentPage *DocumentsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *DocumentsPaginator) CurrentPage() *DocumentsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *DocumentsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *DocumentsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *DocumentsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &DocumentsPageOptions{}
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
		p.Documents = append(p.Documents, resp.Documents...)
	}

	return p.Page.Error == nil
}
