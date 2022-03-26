// Package media_files contains auto-generated files. DO NOT MODIFY
package media_files

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// MediaPageOptions defines the query options for the api operation
type MediaPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageMediaResponse struct {
	AccountSid  string     `json:"account_sid"`
	ContentType string     `json:"content_type"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FaxSid      string     `json:"fax_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
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

// MediaPageResponse defines the response fields for the fax media page
type MediaPageResponse struct {
	Media []PageMediaResponse `json:"media"`
	Meta  PageMetaResponse    `json:"meta"`
}

// Page retrieves a page of fax media files
// See https://www.twilio.com/docs/fax/api/fax-media-resource#read-multiple-fax-media-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *MediaPageOptions) (*MediaPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of fax media files
// See https://www.twilio.com/docs/fax/api/fax-media-resource#read-multiple-fax-media-resources for more details
func (c Client) PageWithContext(context context.Context, options *MediaPageOptions) (*MediaPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Faxes/{faxSid}/Media",
		PathParams: map[string]string{
			"faxSid": c.faxSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &MediaPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// MediaPaginator defines the fields for makings paginated api calls
// Media is an array of media that have been returned from all of the page calls
type MediaPaginator struct {
	options *MediaPageOptions
	Page    *MediaPage
	Media   []PageMediaResponse
}

// NewMediaPaginator creates a new instance of the paginator for Page.
func (c *Client) NewMediaPaginator() *MediaPaginator {
	return c.NewMediaPaginatorWithOptions(nil)
}

// NewMediaPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewMediaPaginatorWithOptions(options *MediaPageOptions) *MediaPaginator {
	return &MediaPaginator{
		options: options,
		Page: &MediaPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Media: make([]PageMediaResponse, 0),
	}
}

// MediaPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageMediaResponse or error that is returned from the api call(s)
type MediaPage struct {
	client *Client

	CurrentPage *MediaPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *MediaPaginator) CurrentPage() *MediaPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *MediaPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *MediaPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *MediaPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &MediaPageOptions{}
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
		p.Media = append(p.Media, resp.Media...)
	}

	return p.Page.Error == nil
}
