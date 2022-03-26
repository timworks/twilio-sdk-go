package client

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/go-playground/form"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/timworks/twilio-sdk-go/session"
	"github.com/timworks/twilio-sdk-go/utils"
)

const (
	JSON       = "application/json"
	URLEncoded = "application/x-www-form-urlencoded"
	FormData   = "multipart/form-data"
)

var encoder = form.NewEncoder()

type Client struct {
	baseURL string
	client  *resty.Client
}

// Used for testing purposes only
func (c Client) GetRestyClient() *resty.Client {
	return c.client
}

func New(sess *session.Session, config *APIClientConfig) *Client {
	retryConfig := config.RetryConfig
	credentials := sess.Credentials

	restyClient := resty.New()
	restyClient.SetDebug(config.DebugEnabled).
		SetRetryCount(retryConfig.Attempts).
		SetRetryWaitTime(retryConfig.WaitTime).
		SetBasicAuth(credentials.Username, credentials.Password).
		SetHeader("User-Agent", "go-sdk").
		SetHeader("Accept", "application/json")

	return &Client{
		baseURL: CreateBaseURL(config.SubDomain, config.APIVersion, config.Edge, config.Region),
		client:  restyClient,
	}
}

type Operation struct {
	OverrideBaseURL *string
	Method          string
	URI             string
	ContentType     string
	PathParams      map[string]string
	QueryParams     *url.Values
}

func (c Client) Send(context context.Context, op Operation, input interface{}, output interface{}) error {
	req, err := configureRequest(context, c.client, op, input, output)
	if err != nil {
		return err
	}

	var baseURL = c.baseURL
	if op.OverrideBaseURL != nil {
		baseURL = *op.OverrideBaseURL
	}

	resp, err := req.Execute(op.Method, baseURL+op.URI)
	if err != nil {
		return err
	}

	if resp.IsError() {
		output = nil
		return resp.Error().(*utils.TwilioError)
	}
	return nil
}

func CreateBaseURL(subDomain string, apiVersion string, edge *string, region *string) string {
	if edge != nil {
		requestRegion := region
		if region == nil {
			requestRegion = utils.String("us1")
		}
		return fmt.Sprintf("https://%s.%s.%s.twilio.com/%s", subDomain, *edge, *requestRegion, apiVersion)
	}
	if region != nil {
		return fmt.Sprintf("https://%s.%s.twilio.com/%s", subDomain, *region, apiVersion)
	}
	return fmt.Sprintf("https://%s.twilio.com/%s", subDomain, apiVersion)
}

func configureRequest(context context.Context, client *resty.Client, op Operation, input interface{}, output interface{}) (*resty.Request, error) {
	req := client.R().
		SetError(&utils.TwilioError{}).
		SetContext(context)

	if op.PathParams != nil {
		req = req.SetPathParams(op.PathParams)
	}

	if op.QueryParams != nil {
		req = req.SetQueryParamsFromValues(*op.QueryParams)
	}

	if input != nil {
		inputReq, err := createInput(req, op.ContentType, input)
		if err != nil {
			return nil, err
		}

		req = inputReq
	}

	if output != nil {
		req = req.SetResult(output)
	}

	return req, nil
}

func createInput(baseRequest *resty.Request, contentType string, input interface{}) (*resty.Request, error) {
	if err := utils.ValidateInput(input); err != nil {
		return nil, fmt.Errorf("Invalid input supplied")
	}

	if contentType == URLEncoded {
		values, err := encoder.Encode(&input)
		if err != nil {
			return nil, err
		}

		return baseRequest.
			SetContentLength(true).
			SetFormDataFromValues(values), nil
	}

	if contentType == FormData {
		values := make(map[string]interface{}, 0)
		if err := mapstructure.Decode(input, &values); err != nil {
			return nil, err
		}

		for key, value := range values {
			fileName, contentType, content := getMultipartFieldDetails(value)
			baseRequest = baseRequest.SetMultipartField(key, fileName, contentType, content)
		}

		return baseRequest, nil
	}

	if contentType == JSON {
		return baseRequest.
			SetContentLength(true).
			SetBody(input), nil
	}

	return nil, fmt.Errorf("%s is not a supported content type", contentType)
}

func getMultipartFieldDetails(value interface{}) (string, string, io.ReadSeeker) {
	fileDetails, ok := value.(map[string]interface{})
	if ok {
		return fileDetails["FileName"].(string), fileDetails["ContentType"].(string), fileDetails["Body"].(io.ReadSeeker)
	}

	return "", "", strings.NewReader(value.(string))
}
