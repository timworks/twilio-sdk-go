// Package reservation contains auto-generated files. DO NOT MODIFY
package reservation

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateReservationInput defines input fields for updating a worker reservation resource
type UpdateReservationInput struct {
	Beep                                    *string   `form:"Beep,omitempty"`
	BeepOnCustomerEntrance                  *bool     `form:"BeepOnCustomerEntrance,omitempty"`
	CallAccept                              *bool     `form:"CallAccept,omitempty"`
	CallFrom                                *string   `form:"CallFrom,omitempty"`
	CallRecord                              *string   `form:"CallRecord,omitempty"`
	CallStatusCallbackURL                   *string   `form:"CallStatusCallbackUrl,omitempty"`
	CallTimeout                             *int      `form:"CallTimeout,omitempty"`
	CallTo                                  *string   `form:"CallTo,omitempty"`
	CallURL                                 *string   `form:"CallUrl,omitempty"`
	ConferenceRecord                        *string   `form:"ConferenceRecord,omitempty"`
	ConferenceRecordingStatusCallback       *string   `form:"ConferenceRecordingStatusCallback,omitempty"`
	ConferenceRecordingStatusCallbackMethod *string   `form:"ConferenceRecordingStatusCallbackMethod,omitempty"`
	ConferenceStatusCallback                *string   `form:"ConferenceStatusCallback,omitempty"`
	ConferenceStatusCallbackEvents          *[]string `form:"ConferenceStatusCallbackEvent,omitempty"`
	ConferenceStatusCallbackMethod          *string   `form:"ConferenceStatusCallbackMethod,omitempty"`
	ConferenceTrim                          *string   `form:"ConferenceTrim,omitempty"`
	DequeueFrom                             *string   `form:"DequeueFrom,omitempty"`
	DequeuePostWorkActivitySid              *string   `form:"DequeuePostWorkActivitySid,omitempty"`
	DequeueRecord                           *string   `form:"DequeueRecord,omitempty"`
	DequeueStatusCallbackEvents             *[]string `form:"DequeueStatusCallbackEvent,omitempty"`
	DequeueStatusCallbackURL                *string   `form:"DequeueStatusCallbackUrl,omitempty"`
	DequeueTimeout                          *int      `form:"DequeueTimeout,omitempty"`
	DequeueTo                               *string   `form:"DequeueTo,omitempty"`
	EarlyMedia                              *bool     `form:"EarlyMedia,omitempty"`
	EndConferenceOnCustomerExit             *bool     `form:"EndConferenceOnCustomerExit,omitempty"`
	EndConferenceOnExit                     *bool     `form:"EndConferenceOnExit,omitempty"`
	From                                    *string   `form:"From,omitempty"`
	Instruction                             *string   `form:"Instruction,omitempty"`
	MaxParticipants                         *int      `form:"MaxParticipants,omitempty"`
	Muted                                   *bool     `form:"Muted,omitempty"`
	PostWorkActivitySid                     *string   `form:"PostWorkActivitySid,omitempty"`
	Record                                  *bool     `form:"Record,omitempty"`
	RecordingChannels                       *string   `form:"RecordingChannels,omitempty"`
	RecordingStatusCallback                 *string   `form:"RecordingStatusCallback,omitempty"`
	RecordingStatusCallbackMethod           *string   `form:"RecordingStatusCallbackMethod,omitempty"`
	RedirectAccept                          *string   `form:"RedirectAccept,omitempty"`
	RedirectCallSid                         *string   `form:"RedirectCallSid,omitempty"`
	RedirectURL                             *string   `form:"RedirectUrl,omitempty"`
	Region                                  *string   `form:"Region,omitempty"`
	ReservationStatus                       string    `validate:"required" form:"ReservationStatus"`
	SipAuthPassword                         *string   `form:"SipAuthPassword,omitempty"`
	SipAuthUsername                         *string   `form:"SipAuthUsername,omitempty"`
	StartConferenceOnEnter                  *bool     `form:"StartConferenceOnEnter,omitempty"`
	StatusCallback                          *string   `form:"StatusCallback,omitempty"`
	StatusCallbackEvents                    *[]string `form:"StatusCallbackEvent,omitempty"`
	StatusCallbackMethod                    *string   `form:"StatusCallbackMethod,omitempty"`
	Timeout                                 *int      `form:"Timeout,omitempty"`
	To                                      *string   `form:"To,omitempty"`
	WaitMethod                              *string   `form:"WaitMethod,omitempty"`
	WaitURL                                 *string   `form:"WaitUrl,omitempty"`
	WorkerActivitySid                       *string   `form:"WorkerActivitySid,omitempty"`
}

// UpdateReservationResponse defines the response fields for the updated worker reservation
type UpdateReservationResponse struct {
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

// Update modifies a worker reservation resource
// See https://www.twilio.com/docs/taskrouter/api/worker-reservation#update-a-workerreservation-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateReservationInput) (*UpdateReservationResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a worker reservation resource
// See https://www.twilio.com/docs/taskrouter/api/worker-reservation#update-a-workerreservation-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateReservationInput) (*UpdateReservationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Workers/{workerSid}/Reservations/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"workerSid":    c.workerSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateReservationInput{}
	}

	response := &UpdateReservationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
