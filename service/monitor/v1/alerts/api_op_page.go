// Package alerts contains auto-generated files. DO NOT MODIFY
package alerts

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// AlertsPageOptions defines the query options for the api operation
type AlertsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
	LogLevel  *string
	StartDate *time.Time
	EndDate   *time.Time
}

type PageAlertResponse struct {
	APIVersion       *string    `json:"api_version,omitempty"`
	AccountSid       string     `json:"account_sid"`
	AlertText        *string    `json:"alert_text,omitempty"`
	DateCreated      time.Time  `json:"date_created"`
	DateGenerated    time.Time  `json:"date_generated"`
	DateUpdated      *time.Time `json:"date_updated,omitempty"`
	ErrorCode        string     `json:"error_code"`
	LogLevel         string     `json:"log_level"`
	MoreInfo         string     `json:"more_info"`
	RequestHeaders   *string    `json:"request_headers,omitempty"`
	RequestMethod    *string    `json:"request_method,omitempty"`
	RequestURL       *string    `json:"request_url,omitempty"`
	RequestVariables *string    `json:"request_variables,omitempty"`
	ResourceSid      string     `json:"resource_sid"`
	ResponseBody     *string    `json:"response_body,omitempty"`
	ResponseHeaders  *string    `json:"response_headers,omitempty"`
	ServiceSid       string     `json:"service_sid"`
	Sid              string     `json:"sid"`
	URL              string     `json:"url"`
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

// AlertsPageResponse defines the response fields for the alerts page
type AlertsPageResponse struct {
	Alerts []PageAlertResponse `json:"alerts"`
	Meta   PageMetaResponse    `json:"meta"`
}

// Page retrieves a page of alerts
// See https://www.twilio.com/docs/usage/monitor-alert#read-multiple-alert-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *AlertsPageOptions) (*AlertsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of alerts
// See https://www.twilio.com/docs/usage/monitor-alert#read-multiple-alert-resources for more details
func (c Client) PageWithContext(context context.Context, options *AlertsPageOptions) (*AlertsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Alerts",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &AlertsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// AlertsPaginator defines the fields for makings paginated api calls
// Alerts is an array of alerts that have been returned from all of the page calls
type AlertsPaginator struct {
	options *AlertsPageOptions
	Page    *AlertsPage
	Alerts  []PageAlertResponse
}

// NewAlertsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewAlertsPaginator() *AlertsPaginator {
	return c.NewAlertsPaginatorWithOptions(nil)
}

// NewAlertsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewAlertsPaginatorWithOptions(options *AlertsPageOptions) *AlertsPaginator {
	return &AlertsPaginator{
		options: options,
		Page: &AlertsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Alerts: make([]PageAlertResponse, 0),
	}
}

// AlertsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageAlertResponse or error that is returned from the api call(s)
type AlertsPage struct {
	client *Client

	CurrentPage *AlertsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *AlertsPaginator) CurrentPage() *AlertsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *AlertsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *AlertsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *AlertsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &AlertsPageOptions{}
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
		p.Alerts = append(p.Alerts, resp.Alerts...)
	}

	return p.Page.Error == nil
}
