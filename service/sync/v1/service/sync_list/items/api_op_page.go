// Package items contains auto-generated files. DO NOT MODIFY
package items

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// SyncListItemsPageOptions defines the query options for the api operation
type SyncListItemsPageOptions struct {
	PageSize    *int
	Page        *int
	PageToken   *string
	Order       *string
	From        *string
	Bounds      *string
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

type PageSyncListItemResponse struct {
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

// SyncListItemsPageResponse defines the response fields for the list items page
type SyncListItemsPageResponse struct {
	Items []PageSyncListItemResponse `json:"items"`
	Meta  PageMetaResponse           `json:"meta"`
}

// Page retrieves a page of list items
// See https://www.twilio.com/docs/sync/api/listitem-resource#read-a-listitem-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SyncListItemsPageOptions) (*SyncListItemsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of list items
// See https://www.twilio.com/docs/sync/api/listitem-resource#read-a-listitem-resource for more details
func (c Client) PageWithContext(context context.Context, options *SyncListItemsPageOptions) (*SyncListItemsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Lists/{syncListSid}/Items",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"syncListSid": c.syncListSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &SyncListItemsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SyncListItemsPaginator defines the fields for makings paginated api calls
// Items is an array of items that have been returned from all of the page calls
type SyncListItemsPaginator struct {
	options *SyncListItemsPageOptions
	Page    *SyncListItemsPage
	Items   []PageSyncListItemResponse
}

// NewSyncListItemsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSyncListItemsPaginator() *SyncListItemsPaginator {
	return c.NewSyncListItemsPaginatorWithOptions(nil)
}

// NewSyncListItemsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSyncListItemsPaginatorWithOptions(options *SyncListItemsPageOptions) *SyncListItemsPaginator {
	return &SyncListItemsPaginator{
		options: options,
		Page: &SyncListItemsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Items: make([]PageSyncListItemResponse, 0),
	}
}

// SyncListItemsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSyncListItemResponse or error that is returned from the api call(s)
type SyncListItemsPage struct {
	client *Client

	CurrentPage *SyncListItemsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SyncListItemsPaginator) CurrentPage() *SyncListItemsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SyncListItemsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SyncListItemsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SyncListItemsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SyncListItemsPageOptions{}
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
		p.Items = append(p.Items, resp.Items...)
	}

	return p.Page.Error == nil
}
