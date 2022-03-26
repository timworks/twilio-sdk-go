// Package samples contains auto-generated files. DO NOT MODIFY
package samples

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// SamplesPageOptions defines the query options for the api operation
type SamplesPageOptions struct {
	Language  *string
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

type PageSampleResponse struct {
	AccountSid    string     `json:"account_sid"`
	AssistantSid  string     `json:"assistant_sid"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	Language      string     `json:"language"`
	Sid           string     `json:"sid"`
	SourceChannel *string    `json:"source_channel,omitempty"`
	TaggedText    string     `json:"tagged_text"`
	TaskSid       string     `json:"task_sid"`
	URL           string     `json:"url"`
}

// SamplesPageResponse defines the response fields for the samples page
type SamplesPageResponse struct {
	Meta    PageMetaResponse     `json:"meta"`
	Samples []PageSampleResponse `json:"samples"`
}

// Page retrieves a page of samples
// See https://www.twilio.com/docs/autopilot/api/task-sample#read-multiple-sample-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SamplesPageOptions) (*SamplesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of samples
// See https://www.twilio.com/docs/autopilot/api/task-sample#read-multiple-sample-resources for more details
func (c Client) PageWithContext(context context.Context, options *SamplesPageOptions) (*SamplesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks/{taskSid}/Samples",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &SamplesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SamplesPaginator defines the fields for makings paginated api calls
// Samples is an array of samples that have been returned from all of the page calls
type SamplesPaginator struct {
	options *SamplesPageOptions
	Page    *SamplesPage
	Samples []PageSampleResponse
}

// NewSamplesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSamplesPaginator() *SamplesPaginator {
	return c.NewSamplesPaginatorWithOptions(nil)
}

// NewSamplesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSamplesPaginatorWithOptions(options *SamplesPageOptions) *SamplesPaginator {
	return &SamplesPaginator{
		options: options,
		Page: &SamplesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Samples: make([]PageSampleResponse, 0),
	}
}

// SamplesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSampleResponse or error that is returned from the api call(s)
type SamplesPage struct {
	client *Client

	CurrentPage *SamplesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SamplesPaginator) CurrentPage() *SamplesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SamplesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SamplesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SamplesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SamplesPageOptions{}
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
		p.Samples = append(p.Samples, resp.Samples...)
	}

	return p.Page.Error == nil
}
