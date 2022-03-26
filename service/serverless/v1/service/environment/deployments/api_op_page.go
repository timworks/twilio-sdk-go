// Package deployments contains auto-generated files. DO NOT MODIFY
package deployments

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// DeploymentsPageOptions defines the query options for the api operation
type DeploymentsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageDeploymentResponse struct {
	AccountSid     string     `json:"account_sid"`
	BuildSid       string     `json:"build_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	EnvironmentSid string     `json:"environment_sid"`
	ServiceSid     string     `json:"service_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
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

// DeploymentsPageResponse defines the response fields for the deployments page
type DeploymentsPageResponse struct {
	Deployments []PageDeploymentResponse `json:"deployments"`
	Meta        PageMetaResponse         `json:"meta"`
}

// Page retrieves a page of deployments
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment#read-multiple-deployment-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *DeploymentsPageOptions) (*DeploymentsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of deployments
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment#read-multiple-deployment-resources for more details
func (c Client) PageWithContext(context context.Context, options *DeploymentsPageOptions) (*DeploymentsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Environments/{environmentSid}/Deployments",
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &DeploymentsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeploymentsPaginator defines the fields for makings paginated api calls
// Deployments is an array of deployments that have been returned from all of the page calls
type DeploymentsPaginator struct {
	options     *DeploymentsPageOptions
	Page        *DeploymentsPage
	Deployments []PageDeploymentResponse
}

// NewDeploymentsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewDeploymentsPaginator() *DeploymentsPaginator {
	return c.NewDeploymentsPaginatorWithOptions(nil)
}

// NewDeploymentsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewDeploymentsPaginatorWithOptions(options *DeploymentsPageOptions) *DeploymentsPaginator {
	return &DeploymentsPaginator{
		options: options,
		Page: &DeploymentsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Deployments: make([]PageDeploymentResponse, 0),
	}
}

// DeploymentsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageDeploymentResponse or error that is returned from the api call(s)
type DeploymentsPage struct {
	client *Client

	CurrentPage *DeploymentsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *DeploymentsPaginator) CurrentPage() *DeploymentsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *DeploymentsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *DeploymentsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *DeploymentsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &DeploymentsPageOptions{}
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
		p.Deployments = append(p.Deployments, resp.Deployments...)
	}

	return p.Page.Error == nil
}
