// Package permissions contains auto-generated files. DO NOT MODIFY
package permissions

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// DocumentPermissionsPageOptions defines the query options for the api operation
type DocumentPermissionsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageDocumentPermissionsResponse struct {
	AccountSid  string `json:"account_sid"`
	DocumentSid string `json:"document_sid"`
	Identity    string `json:"identity"`
	Manage      bool   `json:"manage"`
	Read        bool   `json:"read"`
	ServiceSid  string `json:"service_sid"`
	URL         string `json:"url"`
	Write       bool   `json:"write"`
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

// DocumentPermissionsPageResponse defines the response fields for the document permissions page
type DocumentPermissionsPageResponse struct {
	Meta        PageMetaResponse                  `json:"meta"`
	Permissions []PageDocumentPermissionsResponse `json:"permissions"`
}

// Page retrieves a page of document permissions
// See https://www.twilio.com/docs/sync/api/document-permission-resource#read-multiple-document-permission-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *DocumentPermissionsPageOptions) (*DocumentPermissionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of document permissions
// See https://www.twilio.com/docs/sync/api/document-permission-resource#read-multiple-document-permission-resources for more details
func (c Client) PageWithContext(context context.Context, options *DocumentPermissionsPageOptions) (*DocumentPermissionsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Documents/{documentSid}/Permissions",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"documentSid": c.documentSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &DocumentPermissionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DocumentPermissionsPaginator defines the fields for makings paginated api calls
// Permissions is an array of permissions that have been returned from all of the page calls
type DocumentPermissionsPaginator struct {
	options     *DocumentPermissionsPageOptions
	Page        *DocumentPermissionsPage
	Permissions []PageDocumentPermissionsResponse
}

// NewDocumentPermissionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewDocumentPermissionsPaginator() *DocumentPermissionsPaginator {
	return c.NewDocumentPermissionsPaginatorWithOptions(nil)
}

// NewDocumentPermissionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewDocumentPermissionsPaginatorWithOptions(options *DocumentPermissionsPageOptions) *DocumentPermissionsPaginator {
	return &DocumentPermissionsPaginator{
		options: options,
		Page: &DocumentPermissionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Permissions: make([]PageDocumentPermissionsResponse, 0),
	}
}

// DocumentPermissionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageDocumentPermissionsResponse or error that is returned from the api call(s)
type DocumentPermissionsPage struct {
	client *Client

	CurrentPage *DocumentPermissionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *DocumentPermissionsPaginator) CurrentPage() *DocumentPermissionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *DocumentPermissionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *DocumentPermissionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *DocumentPermissionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &DocumentPermissionsPageOptions{}
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
		p.Permissions = append(p.Permissions, resp.Permissions...)
	}

	return p.Page.Error == nil
}
