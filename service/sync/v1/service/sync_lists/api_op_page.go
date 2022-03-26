// Package sync_lists contains auto-generated files. DO NOT MODIFY
package sync_lists

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// SyncListsPageOptions defines the query options for the api operation
type SyncListsPageOptions struct {
	PageSize    *int
	Page        *int
	PageToken   *string
	HideExpired *bool
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

type PageSyncListResponse struct {
	AccountSid  string     `json:"account_sid"`
	CreatedBy   string     `json:"created_by"`
	DateCreated time.Time  `json:"date_created"`
	DateExpires *time.Time `json:"date_expires,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Revision    string     `json:"revision"`
	ServiceSid  string     `json:"service_Sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  *string    `json:"unique_name,omitempty"`
}

// SyncListsPageResponse defines the response fields for the lists page
type SyncListsPageResponse struct {
	Meta      PageMetaResponse       `json:"meta"`
	SyncLists []PageSyncListResponse `json:"lists"`
}

// Page retrieves a page of lists
// See https://www.twilio.com/docs/sync/api/list-resource#read-a-list-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SyncListsPageOptions) (*SyncListsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of lists
// See https://www.twilio.com/docs/sync/api/list-resource#read-a-list-resource for more details
func (c Client) PageWithContext(context context.Context, options *SyncListsPageOptions) (*SyncListsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Lists",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &SyncListsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SyncListsPaginator defines the fields for makings paginated api calls
// SyncLists is an array of synclists that have been returned from all of the page calls
type SyncListsPaginator struct {
	options   *SyncListsPageOptions
	Page      *SyncListsPage
	SyncLists []PageSyncListResponse
}

// NewSyncListsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSyncListsPaginator() *SyncListsPaginator {
	return c.NewSyncListsPaginatorWithOptions(nil)
}

// NewSyncListsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSyncListsPaginatorWithOptions(options *SyncListsPageOptions) *SyncListsPaginator {
	return &SyncListsPaginator{
		options: options,
		Page: &SyncListsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		SyncLists: make([]PageSyncListResponse, 0),
	}
}

// SyncListsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSyncListResponse or error that is returned from the api call(s)
type SyncListsPage struct {
	client *Client

	CurrentPage *SyncListsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SyncListsPaginator) CurrentPage() *SyncListsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SyncListsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SyncListsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SyncListsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SyncListsPageOptions{}
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
		p.SyncLists = append(p.SyncLists, resp.SyncLists...)
	}

	return p.Page.Error == nil
}
