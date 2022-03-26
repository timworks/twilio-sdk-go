// Package reservations contains auto-generated files. DO NOT MODIFY
package reservations

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// ReservationsPageOptions defines the query options for the api operation
type ReservationsPageOptions struct {
	PageSize          *int
	Page              *int
	PageToken         *string
	WorkerSid         *string
	ReservationStatus *string
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

type PageReservationResponse struct {
	AccountSid        string     `json:"account_sid"`
	DateCreated       time.Time  `json:"date_created"`
	DateUpdated       *time.Time `json:"date_updated,omitempty"`
	ReservationStatus string     `json:"reservation_status"`
	Sid               string     `json:"sid"`
	TaskSid           string     `json:"task_sid"`
	URL               string     `json:"url"`
	WorkerName        string     `json:"worker_name"`
	WorkerSid         string     `json:"worker_sid"`
	WorkspaceSid      string     `json:"workspace_sid"`
}

// ReservationsPageResponse defines the response fields for the task reservations page
type ReservationsPageResponse struct {
	Meta         PageMetaResponse          `json:"meta"`
	Reservations []PageReservationResponse `json:"reservations"`
}

// Page retrieves a page of task reservations
// See https://www.twilio.com/docs/taskrouter/api/reservations#reservations-list-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ReservationsPageOptions) (*ReservationsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of task reservations
// See https://www.twilio.com/docs/taskrouter/api/reservations#reservations-list-resource for more details
func (c Client) PageWithContext(context context.Context, options *ReservationsPageOptions) (*ReservationsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Tasks/{taskSid}/Reservations",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"taskSid":      c.taskSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ReservationsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ReservationsPaginator defines the fields for makings paginated api calls
// Reservations is an array of reservations that have been returned from all of the page calls
type ReservationsPaginator struct {
	options      *ReservationsPageOptions
	Page         *ReservationsPage
	Reservations []PageReservationResponse
}

// NewReservationsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewReservationsPaginator() *ReservationsPaginator {
	return c.NewReservationsPaginatorWithOptions(nil)
}

// NewReservationsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewReservationsPaginatorWithOptions(options *ReservationsPageOptions) *ReservationsPaginator {
	return &ReservationsPaginator{
		options: options,
		Page: &ReservationsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Reservations: make([]PageReservationResponse, 0),
	}
}

// ReservationsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageReservationResponse or error that is returned from the api call(s)
type ReservationsPage struct {
	client *Client

	CurrentPage *ReservationsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ReservationsPaginator) CurrentPage() *ReservationsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ReservationsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ReservationsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ReservationsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ReservationsPageOptions{}
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
		p.Reservations = append(p.Reservations, resp.Reservations...)
	}

	return p.Page.Error == nil
}
