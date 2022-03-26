// Package assistants contains auto-generated files. DO NOT MODIFY
package assistants

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// AssistantsPageOptions defines the query options for the api operation
type AssistantsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageAssistantResponse struct {
	AccountSid          string     `json:"account_sid"`
	CallbackEvents      *string    `json:"callback_events,omitempty"`
	CallbackURL         *string    `json:"callback_url,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	DevelopmentStage    string     `json:"development_stage"`
	FriendlyName        *string    `json:"friendly_name,omitempty"`
	LatestModelBuildSid *string    `json:"latest_model_build_sid,omitempty"`
	LogQueries          bool       `json:"log_queries"`
	NeedsModelBuild     *bool      `json:"needs_model_build,omitempty"`
	Sid                 string     `json:"sid"`
	URL                 string     `json:"url"`
	UniqueName          string     `json:"unique_name"`
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

// AssistantsPageResponse defines the response fields for the assistants page
type AssistantsPageResponse struct {
	Assistants []PageAssistantResponse `json:"assistants"`
	Meta       PageMetaResponse        `json:"meta"`
}

// Page retrieves a page of assistants
// See https://www.twilio.com/docs/autopilot/api/assistant#read-multiple-assistant-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *AssistantsPageOptions) (*AssistantsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of assistants
// See https://www.twilio.com/docs/autopilot/api/assistant#read-multiple-assistant-resources for more details
func (c Client) PageWithContext(context context.Context, options *AssistantsPageOptions) (*AssistantsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Assistants",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &AssistantsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// AssistantsPaginator defines the fields for makings paginated api calls
// Assistants is an array of assistants that have been returned from all of the page calls
type AssistantsPaginator struct {
	options    *AssistantsPageOptions
	Page       *AssistantsPage
	Assistants []PageAssistantResponse
}

// NewAssistantsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewAssistantsPaginator() *AssistantsPaginator {
	return c.NewAssistantsPaginatorWithOptions(nil)
}

// NewAssistantsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewAssistantsPaginatorWithOptions(options *AssistantsPageOptions) *AssistantsPaginator {
	return &AssistantsPaginator{
		options: options,
		Page: &AssistantsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Assistants: make([]PageAssistantResponse, 0),
	}
}

// AssistantsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageAssistantResponse or error that is returned from the api call(s)
type AssistantsPage struct {
	client *Client

	CurrentPage *AssistantsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *AssistantsPaginator) CurrentPage() *AssistantsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *AssistantsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *AssistantsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *AssistantsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &AssistantsPageOptions{}
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
		p.Assistants = append(p.Assistants, resp.Assistants...)
	}

	return p.Page.Error == nil
}
