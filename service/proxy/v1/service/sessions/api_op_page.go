// Package sessions contains auto-generated files. DO NOT MODIFY
package sessions

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// SessionsPageOptions defines the query options for the api operation
type SessionsPageOptions struct {
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

type PageSessionResponse struct {
	AccountSid          string     `json:"account_sid"`
	ClosedReason        *string    `json:"closed_reason,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateEnded           *time.Time `json:"date_ended,omitempty"`
	DateExpiry          *time.Time `json:"date_expiry,omitempty"`
	DateLastInteraction *time.Time `json:"date_last_interaction,omitempty"`
	DateStarted         *time.Time `json:"date_started,omitempty"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	Mode                *string    `json:"mode,omitempty"`
	ServiceSid          string     `json:"service_sid"`
	Sid                 string     `json:"sid"`
	Status              *string    `json:"status,omitempty"`
	Ttl                 *int       `json:"ttl,omitempty"`
	URL                 string     `json:"url"`
	UniqueName          string     `json:"unique_name"`
}

// SessionsPageResponse defines the response fields for the sessions page
type SessionsPageResponse struct {
	Meta     PageMetaResponse      `json:"meta"`
	Sessions []PageSessionResponse `json:"sessions"`
}

// Page retrieves a page of sessions
// See https://www.twilio.com/docs/proxy/api/session#read-multiple-session-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SessionsPageOptions) (*SessionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of sessions
// See https://www.twilio.com/docs/proxy/api/session#read-multiple-session-resources for more details
func (c Client) PageWithContext(context context.Context, options *SessionsPageOptions) (*SessionsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Sessions",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &SessionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SessionsPaginator defines the fields for makings paginated api calls
// Sessions is an array of sessions that have been returned from all of the page calls
type SessionsPaginator struct {
	options  *SessionsPageOptions
	Page     *SessionsPage
	Sessions []PageSessionResponse
}

// NewSessionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSessionsPaginator() *SessionsPaginator {
	return c.NewSessionsPaginatorWithOptions(nil)
}

// NewSessionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSessionsPaginatorWithOptions(options *SessionsPageOptions) *SessionsPaginator {
	return &SessionsPaginator{
		options: options,
		Page: &SessionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Sessions: make([]PageSessionResponse, 0),
	}
}

// SessionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSessionResponse or error that is returned from the api call(s)
type SessionsPage struct {
	client *Client

	CurrentPage *SessionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SessionsPaginator) CurrentPage() *SessionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SessionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SessionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SessionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SessionsPageOptions{}
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
		p.Sessions = append(p.Sessions, resp.Sessions...)
	}

	return p.Page.Error == nil
}
