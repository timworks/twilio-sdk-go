// Package model_builds contains auto-generated files. DO NOT MODIFY
package model_builds

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// ModelBuildsPageOptions defines the query options for the api operation
type ModelBuildsPageOptions struct {
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

type PageModelBuildResponse struct {
	AccountSid    string     `json:"account_sid"`
	AssistantSid  string     `json:"assistant_sid"`
	BuildDuration *int       `json:"build_duration,omitempty"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	ErrorCode     *int       `json:"error_code,omitempty"`
	Sid           string     `json:"sid"`
	Status        string     `json:"status"`
	URL           string     `json:"url"`
	UniqueName    string     `json:"unique_name"`
}

// ModelBuildsPageResponse defines the response fields for the model build page
type ModelBuildsPageResponse struct {
	Meta        PageMetaResponse         `json:"meta"`
	ModelBuilds []PageModelBuildResponse `json:"model_builds"`
}

// Page retrieves a page of model builds
// See https://www.twilio.com/docs/autopilot/api/model-build#read-multiple-modelbuild-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ModelBuildsPageOptions) (*ModelBuildsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of model builds
// See https://www.twilio.com/docs/autopilot/api/model-build#read-multiple-modelbuild-resources for more details
func (c Client) PageWithContext(context context.Context, options *ModelBuildsPageOptions) (*ModelBuildsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/ModelBuilds",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ModelBuildsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ModelBuildsPaginator defines the fields for makings paginated api calls
// ModelBuilds is an array of modelbuilds that have been returned from all of the page calls
type ModelBuildsPaginator struct {
	options     *ModelBuildsPageOptions
	Page        *ModelBuildsPage
	ModelBuilds []PageModelBuildResponse
}

// NewModelBuildsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewModelBuildsPaginator() *ModelBuildsPaginator {
	return c.NewModelBuildsPaginatorWithOptions(nil)
}

// NewModelBuildsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewModelBuildsPaginatorWithOptions(options *ModelBuildsPageOptions) *ModelBuildsPaginator {
	return &ModelBuildsPaginator{
		options: options,
		Page: &ModelBuildsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		ModelBuilds: make([]PageModelBuildResponse, 0),
	}
}

// ModelBuildsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageModelBuildResponse or error that is returned from the api call(s)
type ModelBuildsPage struct {
	client *Client

	CurrentPage *ModelBuildsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ModelBuildsPaginator) CurrentPage() *ModelBuildsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ModelBuildsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ModelBuildsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ModelBuildsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ModelBuildsPageOptions{}
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
		p.ModelBuilds = append(p.ModelBuilds, resp.ModelBuilds...)
	}

	return p.Page.Error == nil
}
