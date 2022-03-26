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

// SyncListPermissionsPageOptions defines the query options for the api operation
type SyncListPermissionsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
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

type PageSyncListPermissionsResponse struct {
	AccountSid string `json:"account_sid"`
	Identity   string `json:"identity"`
	ListSid    string `json:"list_sid"`
	Manage     bool   `json:"manage"`
	Read       bool   `json:"read"`
	ServiceSid string `json:"service_sid"`
	URL        string `json:"url"`
	Write      bool   `json:"write"`
}

// SyncListPermissionsPageResponse defines the response fields for the list permissions page
type SyncListPermissionsPageResponse struct {
	Meta        PageMetaResponse                  `json:"meta"`
	Permissions []PageSyncListPermissionsResponse `json:"permissions"`
}

// Page retrieves a page of list permissions
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource#read-multiple-sync-list-permission-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SyncListPermissionsPageOptions) (*SyncListPermissionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of list permissions
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource#read-multiple-sync-list-permission-resources for more details
func (c Client) PageWithContext(context context.Context, options *SyncListPermissionsPageOptions) (*SyncListPermissionsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Lists/{syncListSid}/Permissions",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"syncListSid": c.syncListSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &SyncListPermissionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SyncListPermissionsPaginator defines the fields for makings paginated api calls
// Permissions is an array of permissions that have been returned from all of the page calls
type SyncListPermissionsPaginator struct {
	options     *SyncListPermissionsPageOptions
	Page        *SyncListPermissionsPage
	Permissions []PageSyncListPermissionsResponse
}

// NewSyncListPermissionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSyncListPermissionsPaginator() *SyncListPermissionsPaginator {
	return c.NewSyncListPermissionsPaginatorWithOptions(nil)
}

// NewSyncListPermissionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSyncListPermissionsPaginatorWithOptions(options *SyncListPermissionsPageOptions) *SyncListPermissionsPaginator {
	return &SyncListPermissionsPaginator{
		options: options,
		Page: &SyncListPermissionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Permissions: make([]PageSyncListPermissionsResponse, 0),
	}
}

// SyncListPermissionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSyncListPermissionsResponse or error that is returned from the api call(s)
type SyncListPermissionsPage struct {
	client *Client

	CurrentPage *SyncListPermissionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SyncListPermissionsPaginator) CurrentPage() *SyncListPermissionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SyncListPermissionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SyncListPermissionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SyncListPermissionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SyncListPermissionsPageOptions{}
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
