// Package users contains auto-generated files. DO NOT MODIFY
package users

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// UsersPageOptions defines the query options for the api operation
type UsersPageOptions struct {
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

type PageUserResponse struct {
	AccountSid     string     `json:"account_sid"`
	Attributes     string     `json:"attributes"`
	ChatServiceSid string     `json:"chat_service_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	FriendlyName   *string    `json:"friendly_name,omitempty"`
	Identity       string     `json:"identity"`
	IsNotifiable   *bool      `json:"is_notifiable,omitempty"`
	IsOnline       *bool      `json:"is_online,omitempty"`
	RoleSid        string     `json:"role_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
}

// UsersPageResponse defines the response fields for the users page
type UsersPageResponse struct {
	Meta  PageMetaResponse   `json:"meta"`
	Users []PageUserResponse `json:"users"`
}

// Page retrieves a page of users
// See https://www.twilio.com/docs/conversations/api/user-resource#read-multiple-conversationuser-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *UsersPageOptions) (*UsersPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of users
// See https://www.twilio.com/docs/conversations/api/user-resource#read-multiple-conversationuser-resources for more details
func (c Client) PageWithContext(context context.Context, options *UsersPageOptions) (*UsersPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Users",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &UsersPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UsersPaginator defines the fields for makings paginated api calls
// Users is an array of users that have been returned from all of the page calls
type UsersPaginator struct {
	options *UsersPageOptions
	Page    *UsersPage
	Users   []PageUserResponse
}

// NewUsersPaginator creates a new instance of the paginator for Page.
func (c *Client) NewUsersPaginator() *UsersPaginator {
	return c.NewUsersPaginatorWithOptions(nil)
}

// NewUsersPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewUsersPaginatorWithOptions(options *UsersPageOptions) *UsersPaginator {
	return &UsersPaginator{
		options: options,
		Page: &UsersPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Users: make([]PageUserResponse, 0),
	}
}

// UsersPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageUserResponse or error that is returned from the api call(s)
type UsersPage struct {
	client *Client

	CurrentPage *UsersPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *UsersPaginator) CurrentPage() *UsersPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *UsersPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *UsersPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *UsersPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &UsersPageOptions{}
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
		p.Users = append(p.Users, resp.Users...)
	}

	return p.Page.Error == nil
}
