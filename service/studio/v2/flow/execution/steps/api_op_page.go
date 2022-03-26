// Package steps contains auto-generated files. DO NOT MODIFY
package steps

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// StepsPageOptions defines the query options for the api operation
type StepsPageOptions struct {
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

type PageStepResponse struct {
	AccountSid       string      `json:"account_sid"`
	Context          interface{} `json:"context"`
	DateCreated      time.Time   `json:"date_created"`
	DateUpdated      *time.Time  `json:"date_updated,omitempty"`
	ExecutionSid     string      `json:"execution_sid"`
	FlowSid          string      `json:"flow_sid"`
	Name             string      `json:"name"`
	Sid              string      `json:"sid"`
	TransitionedFrom string      `json:"transitioned_from"`
	TransitionedTo   string      `json:"transitioned_to"`
	URL              string      `json:"url"`
}

// StepsPageResponse defines the response fields for the steps page
type StepsPageResponse struct {
	Meta  PageMetaResponse   `json:"meta"`
	Steps []PageStepResponse `json:"steps"`
}

// Page retrieves a page of steps
// See https://www.twilio.com/docs/studio/rest-api/v2/step#read-a-list-of-step-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *StepsPageOptions) (*StepsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of steps
// See https://www.twilio.com/docs/studio/rest-api/v2/step#read-a-list-of-step-resources for more details
func (c Client) PageWithContext(context context.Context, options *StepsPageOptions) (*StepsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/Executions/{executionSid}/Steps",
		PathParams: map[string]string{
			"flowSid":      c.flowSid,
			"executionSid": c.executionSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &StepsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// StepsPaginator defines the fields for makings paginated api calls
// Steps is an array of steps that have been returned from all of the page calls
type StepsPaginator struct {
	options *StepsPageOptions
	Page    *StepsPage
	Steps   []PageStepResponse
}

// NewStepsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewStepsPaginator() *StepsPaginator {
	return c.NewStepsPaginatorWithOptions(nil)
}

// NewStepsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewStepsPaginatorWithOptions(options *StepsPageOptions) *StepsPaginator {
	return &StepsPaginator{
		options: options,
		Page: &StepsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Steps: make([]PageStepResponse, 0),
	}
}

// StepsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageStepResponse or error that is returned from the api call(s)
type StepsPage struct {
	client *Client

	CurrentPage *StepsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *StepsPaginator) CurrentPage() *StepsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *StepsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *StepsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *StepsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &StepsPageOptions{}
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
		p.Steps = append(p.Steps, resp.Steps...)
	}

	return p.Page.Error == nil
}
