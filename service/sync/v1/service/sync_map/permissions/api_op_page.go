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

// SyncMapPermissionsPageOptions defines the query options for the api operation
type SyncMapPermissionsPageOptions struct {
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

type PageSyncMapPermissionsResponse struct {
	AccountSid string `json:"account_sid"`
	Identity   string `json:"identity"`
	Manage     bool   `json:"manage"`
	MapSid     string `json:"map_sid"`
	Read       bool   `json:"read"`
	ServiceSid string `json:"service_sid"`
	URL        string `json:"url"`
	Write      bool   `json:"write"`
}

// SyncMapPermissionsPageResponse defines the response fields for the map permissions page
type SyncMapPermissionsPageResponse struct {
	Meta        PageMetaResponse                 `json:"meta"`
	Permissions []PageSyncMapPermissionsResponse `json:"permissions"`
}

// Page retrieves a page of map permissions
// See https://www.twilio.com/docs/sync/api/sync-map-permission-resource#read-multiple-sync-map-permission-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SyncMapPermissionsPageOptions) (*SyncMapPermissionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of map permissions
// See https://www.twilio.com/docs/sync/api/sync-map-permission-resource#read-multiple-sync-map-permission-resources for more details
func (c Client) PageWithContext(context context.Context, options *SyncMapPermissionsPageOptions) (*SyncMapPermissionsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Maps/{syncMapSid}/Permissions",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"syncMapSid": c.syncMapSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &SyncMapPermissionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SyncMapPermissionsPaginator defines the fields for makings paginated api calls
// Permissions is an array of permissions that have been returned from all of the page calls
type SyncMapPermissionsPaginator struct {
	options     *SyncMapPermissionsPageOptions
	Page        *SyncMapPermissionsPage
	Permissions []PageSyncMapPermissionsResponse
}

// NewSyncMapPermissionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSyncMapPermissionsPaginator() *SyncMapPermissionsPaginator {
	return c.NewSyncMapPermissionsPaginatorWithOptions(nil)
}

// NewSyncMapPermissionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSyncMapPermissionsPaginatorWithOptions(options *SyncMapPermissionsPageOptions) *SyncMapPermissionsPaginator {
	return &SyncMapPermissionsPaginator{
		options: options,
		Page: &SyncMapPermissionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Permissions: make([]PageSyncMapPermissionsResponse, 0),
	}
}

// SyncMapPermissionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSyncMapPermissionsResponse or error that is returned from the api call(s)
type SyncMapPermissionsPage struct {
	client *Client

	CurrentPage *SyncMapPermissionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SyncMapPermissionsPaginator) CurrentPage() *SyncMapPermissionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SyncMapPermissionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SyncMapPermissionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SyncMapPermissionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SyncMapPermissionsPageOptions{}
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
