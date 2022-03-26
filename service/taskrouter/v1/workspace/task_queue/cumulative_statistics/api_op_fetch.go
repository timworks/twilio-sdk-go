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
	Minutes         *int
	StartDate       *time.Time
	EndDate         *time.Time
	TaskChannel     *string
	SplitByWaitTime *string
}

type FetchStatisticsBreakdown struct {
	Avg   int `json:"avg"`
	Max   int `json:"max"`
	Min   int `json:"min"`
	Total int `json:"total"`
}

type FetchWaitTime struct {
	Above FetchWaitTimeTasks `json:"above"`
	Below FetchWaitTimeTasks `json:"below"`
}

type FetchWaitTimeTasks struct {
	ReservationsAccepted int `json:"reservations_accepted"`
	TasksCanceled        int `json:"tasks_canceled"`
}

// FetchCumulativeStatisticsResponse defines the response fields for the retrieved cumulative statistics
type FetchCumulativeStatisticsResponse struct {
	AccountSid                       string                    `json:"account_sid"`
	AvgTaskAcceptanceTime            int                       `json:"avg_task_acceptance_time"`
	EndTime                          time.Time                 `json:"end_time"`
	ReservationsAccepted             int                       `json:"reservations_accepted"`
	ReservationsCanceled             int                       `json:"reservations_canceled"`
	ReservationsCreated              int                       `json:"reservations_created"`
	ReservationsRejected             int                       `json:"reservations_rejected"`
	ReservationsRescinded            int                       `json:"reservations_rescinded"`
	ReservationsTimedOut             int                       `json:"reservations_timed_out"`
	SplitByWaitTime                  *map[string]FetchWaitTime `json:"split_by_wait_time,omitempty"`
	StartTime                        time.Time                 `json:"start_time"`
	TaskQueueSid                     string                    `json:"task_queue_sid"`
	TasksCanceled                    int                       `json:"tasks_canceled"`
	TasksCompleted                   int                       `json:"tasks_completed"`
	TasksDeleted                     int                       `json:"tasks_deleted"`
	TasksEntered                     int                       `json:"tasks_entered"`
	TasksMoved                       int                       `json:"tasks_moved"`
	URL                              string                    `json:"url"`
	WaitDurationInQueueUntilAccepted FetchStatisticsBreakdown  `json:"wait_duration_in_queue_until_accepted"`
	WaitDurationUntilAccepted        FetchStatisticsBreakdown  `json:"wait_duration_until_accepted"`
	WaitDurationUntilCanceled        FetchStatisticsBreakdown  `json:"wait_duration_until_canceled"`
	WorkspaceSid                     string                    `json:"workspace_sid"`
}

// Fetch retrieves cumulative statistics
// See https://www.twilio.com/docs/taskrouter/api/taskqueue-statistics#taskqueue-cumulative-statistics for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch(options *FetchCumulativeStatisticsOptions) (*FetchCumulativeStatisticsResponse, error) {
	return c.FetchWithContext(context.Background(), options)
}

// FetchWithContext retrieves cumulative statistics
// See https://www.twilio.com/docs/taskrouter/api/taskqueue-statistics#taskqueue-cumulative-statistics for more details
func (c Client) FetchWithContext(context context.Context, options *FetchCumulativeStatisticsOptions) (*FetchCumulativeStatisticsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/TaskQueues/{taskQueueSid}/CumulativeStatistics",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"taskQueueSid": c.taskQueueSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FetchCumulativeStatisticsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
