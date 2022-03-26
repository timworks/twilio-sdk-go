// Package bindings contains auto-generated files. DO NOT MODIFY
package bindings

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// UserBindingsPageOptions defines the query options for the api operation
type UserBindingsPageOptions struct {
	BindingType *[]string
	PageSize    *int
	Page        *int
	PageToken   *string
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

type PageUserBindingResponse struct {
	AccountSid    string     `json:"account_sid"`
	BindingType   *string    `json:"binding_type,omitempty"`
	CredentialSid string     `json:"credential_sid"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	Endpoint      *string    `json:"endpoint,omitempty"`
	Identity      *string    `json:"identity,omitempty"`
	MessageTypes  *[]string  `json:"message_types,omitempty"`
	ServiceSid    string     `json:"service_sid"`
	Sid           string     `json:"sid"`
	URL           string     `json:"url"`
	UserSid       string     `json:"user_sid"`
}

// UserBindingsPageResponse defines the response fields for the user bindings page
type UserBindingsPageResponse struct {
	Bindings []PageUserBindingResponse `json:"bindings"`
	Meta     PageMetaResponse          `json:"meta"`
}

// Page retrieves a page of user bindings
// See https://www.twilio.com/docs/chat/rest/user-binding-resource#read-multiple-userbinding-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *UserBindingsPageOptions) (*UserBindingsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of user bindings
// See https://www.twilio.com/docs/chat/rest/user-binding-resource#read-multiple-userbinding-resources for more details
func (c Client) PageWithContext(context context.Context, options *UserBindingsPageOptions) (*UserBindingsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Users/{userSid}/Bindings",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"userSid":    c.userSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &UserBindingsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UserBindingsPaginator defines the fields for makings paginated api calls
// Bindings is an array of bindings that have been returned from all of the page calls
type UserBindingsPaginator struct {
	options  *UserBindingsPageOptions
	Page     *UserBindingsPage
	Bindings []PageUserBindingResponse
}

// NewUserBindingsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewUserBindingsPaginator() *UserBindingsPaginator {
	return c.NewUserBindingsPaginatorWithOptions(nil)
}

// NewUserBindingsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewUserBindingsPaginatorWithOptions(options *UserBindingsPageOptions) *UserBindingsPaginator {
	return &UserBindingsPaginator{
		options: options,
		Page: &UserBindingsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Bindings: make([]PageUserBindingResponse, 0),
	}
}

// UserBindingsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageUserBindingResponse or error that is returned from the api call(s)
type UserBindingsPage struct {
	client *Client

	CurrentPage *UserBindingsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *UserBindingsPaginator) CurrentPage() *UserBindingsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *UserBindingsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *UserBindingsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *UserBindingsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &UserBindingsPageOptions{}
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
		p.Bindings = append(p.Bindings, resp.Bindings...)
	}

	return p.Page.Error == nil
}
