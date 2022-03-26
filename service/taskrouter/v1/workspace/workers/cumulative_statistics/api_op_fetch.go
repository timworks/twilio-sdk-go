// Package cumulative_statistics contains auto-generated files. DO NOT MODIFY
package cumulative_statistics

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FetchCumulativeStatisticsOptions defines the query options for the api operation
type FetchCumulativeStatisticsOptions struct {
	Minutes     *int
	StartDate   *time.Time
	EndDate     *time.Time
	TaskChannel *string
}

type FetchActivityDuration struct {
	Avg          int    `json:"avg"`
	FriendlyName string `json:"friendly_name"`
	Max          int    `json:"max"`
	Min          int    `json:"min"`
	Sid          string `json:"sid"`
	Total        int    `json:"total"`
}

// FetchCumulativeStatisticsResponse defines the response fields for the retrieved cumulative statistics
type FetchCumulativeStatisticsResponse struct {
	AccountSid            string                  `json:"account_sid"`
	ActivityDurations     []FetchActivityDuration `json:"activity_durations"`
	EndTime               time.Time               `json:"end_time"`
	ReservationsAccepted  int                     `json:"reservations_accepted"`
	ReservationsCanceled  int                     `json:"reservations_canceled"`
	ReservationsCreated   int                     `json:"reservations_created"`
	ReservationsRejected  int                     `json:"reservations_rejected"`
	ReservationsRescinded int                     `json:"reservations_rescinded"`
	ReservationsTimedOut  int                     `json:"reservations_timed_out"`
	StartTime             time.Time               `json:"start_time"`
	URL                   string                  `json:"url"`
	WorkspaceSid          string                  `json:"workspace_sid"`
}

// Fetch retrieves cumulative statistics
// See https://www.twilio.com/docs/taskrouter/api/worker/statistics#fetch-cumulative-worker-statistics for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch(options *FetchCumulativeStatisticsOptions) (*FetchCumulativeStatisticsResponse, error) {
	return c.FetchWithContext(context.Background(), options)
}

// FetchWithContext retrieves cumulative statistics
// See https://www.twilio.com/docs/taskrouter/api/worker/statistics#fetch-cumulative-worker-statistics for more details
func (c Client) FetchWithContext(context context.Context, options *FetchCumulativeStatisticsOptions) (*FetchCumulativeStatisticsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workers/CumulativeStatistics",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FetchCumulativeStatisticsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
