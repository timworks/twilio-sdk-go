package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/api"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/address"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/addresses"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/application"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/applications"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/available_phone_number/local"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/available_phone_number/mobile"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/available_phone_number/toll_free"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/call"
	callFeedback "github.com/timworks/twilio-sdk-go/service/api/v2010/account/call/feedback"
	callFeedbacks "github.com/timworks/twilio-sdk-go/service/api/v2010/account/call/feedbacks"
	callRecordings "github.com/timworks/twilio-sdk-go/service/api/v2010/account/call/recordings"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/calls"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/calls/feedback_summaries"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/conference"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/conference/participant"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/conference/participants"
	conferenceRecordings "github.com/timworks/twilio-sdk-go/service/api/v2010/account/conference/recordings"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/conferences"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/incoming_phone_number"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/incoming_phone_numbers"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/key"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/keys"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/message"
	messageFeedbacks "github.com/timworks/twilio-sdk-go/service/api/v2010/account/message/feedbacks"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/message/media_attachments"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/messages"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/queue"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/queue/member"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/queue/members"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/queues"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/recording"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/recordings"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/credential_list"
	sipCredential "github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/credential_list/credential"
	sipCredentials "github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/credential_list/credentials"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/credential_lists"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/domain"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/calls/credential_list_mappings"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/calls/ip_access_control_list_mappings"
	sipRegistrationsCredentialListMappings "github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/registrations/credential_list_mappings"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/domains"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list/ip_address"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list/ip_addresses"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_lists"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/tokens"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/accounts"
	"github.com/timworks/twilio-sdk-go/session"
	"github.com/timworks/twilio-sdk-go/session/credentials"
	"github.com/timworks/twilio-sdk-go/utils"
)

var _ = Describe("API V2010", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	apiClient := api.New(session.New(creds), &client.Config{
		RetryAttempts: utils.Int(0),
	}).V2010

	httpmock.ActivateNonDefault(apiClient.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the accounts client", func() {
		accountsClient := apiClient.Accounts

		Describe("When the assistant is successfully created", func() {
			createInput := &accounts.CreateAccountInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := accountsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create assistant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.Status).To(Equal("active"))
				Expect(resp.Type).To(Equal("Trial"))
				Expect(resp.AuthToken).To(Equal("TestToken"))
				Expect(resp.OwnerAccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create account api returns a 500 response", func() {
			createInput := &accounts.CreateAccountInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := accountsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create account response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of accounts are successfully retrieved", func() {
			pageOptions := &accounts.AccountsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := accountsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the accounts page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				accounts := resp.Accounts
				Expect(accounts).ToNot(BeNil())
				Expect(len(accounts)).To(Equal(1))

				Expect(accounts[0].Sid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(accounts[0].FriendlyName).To(Equal("Test"))
				Expect(accounts[0].Status).To(Equal("active"))
				Expect(accounts[0].Type).To(Equal("Trial"))
				Expect(accounts[0].AuthToken).To(Equal("TestToken"))
				Expect(accounts[0].OwnerAccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(accounts[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(accounts[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of accounts api returns a 500 response", func() {
			pageOptions := &accounts.AccountsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := accountsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the accounts page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated accounts are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := accountsClient.NewAccountsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated accounts current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated accounts results should be returned", func() {
				Expect(len(paginator.Accounts)).To(Equal(3))
			})
		})

		Describe("When the accounts api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := accountsClient.NewAccountsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated accounts current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a account sid", func() {
		accountClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the account is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := accountClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get account response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.Status).To(Equal("active"))
				Expect(resp.Type).To(Equal("Trial"))
				Expect(resp.AuthToken).To(Equal("TestToken"))
				Expect(resp.OwnerAccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the account api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/AC71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("AC71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get account response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the account is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateAccountResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &account.UpdateAccountInput{
				Status: utils.String("closed"),
			}

			resp, err := accountClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update account response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.Status).To(Equal("closed"))
				Expect(resp.Type).To(Equal("Trial"))
				Expect(resp.AuthToken).To(Equal("TestToken"))
				Expect(resp.OwnerAccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the account api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/AC71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &account.UpdateAccountInput{
				Status: utils.String("closed"),
			}

			resp, err := apiClient.Account("AC71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update account response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the keys client", func() {
		keysClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Keys

		Describe("When the key is successfully created", func() {
			createInput := &keys.CreateKeyInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keyResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := keysClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create key response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Secret).To(Equal("SecretValue"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create key api returns a 500 response", func() {
			createInput := &keys.CreateKeyInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := keysClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create key response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of keys are successfully retrieved", func() {
			pageOptions := &keys.KeysPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keysPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := keysClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the keys page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				keys := resp.Keys
				Expect(keys).ToNot(BeNil())
				Expect(len(keys)).To(Equal(1))

				Expect(keys[0].Sid).To(Equal("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(keys[0].FriendlyName).To(Equal("Test"))
				Expect(keys[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(keys[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of keys api returns a 500 response", func() {
			pageOptions := &keys.KeysPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := keysClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the keys page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated keys are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keysPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keysPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := keysClient.NewKeysPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated keys current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated keys results should be returned", func() {
				Expect(len(paginator.Keys)).To(Equal(3))
			})
		})

		Describe("When the keys api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keysPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := keysClient.NewKeysPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated keys current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a key sid", func() {
		keyClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Key("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the key is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/getKeyResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := keyClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get key response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the key api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SK71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Key("SK71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get key response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the key is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateKeyResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &key.UpdateKeyInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := keyClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update key response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the key api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SK71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &key.UpdateKeyInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Key("SK71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update key response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the key is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := keyClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the key api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SK71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Key("SK71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the messages client", func() {
		messagesClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Messages

		Describe("When the message is successfully created", func() {
			createInput := &messages.CreateMessageInput{
				To:                  "+10123456789",
				MessagingServiceSid: utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				Body:                utils.String("Hello World"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := messagesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Body).To(Equal("Hello World"))
				Expect(resp.NumSegments).To(Equal("1"))
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.From).To(Equal(utils.String("")))
				Expect(resp.Price).To(BeNil())
				Expect(resp.ErrorMessage).To(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.NumMedia).To(Equal("0"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.Status).To(Equal("failed"))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ErrorCode).To(Equal(utils.Int(21704)))
				Expect(resp.PriceUnit).To(Equal("GBP"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateSent.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the message request does not contain a to", func() {
			createInput := &messages.CreateMessageInput{
				MessagingServiceSid: utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				Body:                utils.String("Hello World"),
			}

			resp, err := messagesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create message api returns a 500 response", func() {
			createInput := &messages.CreateMessageInput{
				To:                  "+10123456789",
				MessagingServiceSid: utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				Body:                utils.String("Hello World"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of messages are successfully retrieved", func() {
			pageOptions := &messages.MessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messagesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the messages page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				messages := resp.Messages
				Expect(messages).ToNot(BeNil())
				Expect(len(messages)).To(Equal(1))

				Expect(messages[0].Sid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].Body).To(Equal("Hello World"))
				Expect(messages[0].NumSegments).To(Equal("1"))
				Expect(messages[0].Direction).To(Equal("outbound-api"))
				Expect(messages[0].From).To(Equal(utils.String("")))
				Expect(messages[0].Price).To(BeNil())
				Expect(messages[0].ErrorMessage).To(BeNil())
				Expect(messages[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].NumMedia).To(Equal("0"))
				Expect(messages[0].To).To(Equal("+10123456789"))
				Expect(messages[0].Status).To(Equal("failed"))
				Expect(messages[0].MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(messages[0].ErrorCode).To(Equal(utils.Int(21704)))
				Expect(messages[0].PriceUnit).To(Equal("GBP"))
				Expect(messages[0].APIVersion).To(Equal("2010-04-01"))
				Expect(messages[0].DateSent.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(messages[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(messages[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of messages api returns a 500 response", func() {
			pageOptions := &messages.MessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the messages page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated messages are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := messagesClient.NewMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated messages current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated messages results should be returned", func() {
				Expect(len(paginator.Messages)).To(Equal(3))
			})
		})

		Describe("When the messages api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := messagesClient.NewMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated messages current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a message sid", func() {
		messageClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the message is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messageClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Body).To(Equal("Hello World"))
				Expect(resp.NumSegments).To(Equal("1"))
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.From).To(Equal(utils.String("")))
				Expect(resp.Price).To(BeNil())
				Expect(resp.ErrorMessage).To(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.NumMedia).To(Equal("0"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.Status).To(Equal("failed"))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ErrorCode).To(Equal(utils.Int(21704)))
				Expect(resp.PriceUnit).To(Equal("GBP"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateSent.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SM71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SM71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the message is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &message.UpdateMessageInput{
				Body: "Test",
			}

			resp, err := messageClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Body).To(Equal("Test"))
				Expect(resp.NumSegments).To(Equal("1"))
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.From).To(Equal(utils.String("")))
				Expect(resp.Price).To(BeNil())
				Expect(resp.ErrorMessage).To(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.NumMedia).To(Equal("0"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.Status).To(Equal("failed"))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ErrorCode).To(Equal(utils.Int(21704)))
				Expect(resp.PriceUnit).To(Equal("GBP"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateSent.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SM71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &message.UpdateMessageInput{
				Body: "Test",
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SM71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the message is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := messageClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SM71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SM71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the feedbacks client", func() {
		feedbacksClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("MMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Feedbacks

		Describe("When the feedback is successfully created", func() {
			createInput := &messageFeedbacks.CreateFeedbackInput{
				Outcome: utils.String("confirmed"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/MMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/feedbackResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := feedbacksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create feedback response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MessageSid).To(Equal("MMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Outcome).To(Equal("confirmed"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create feedbacks api returns a 500 response", func() {
			createInput := &messageFeedbacks.CreateFeedbackInput{
				Outcome: utils.String("confirmed"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/MMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := feedbacksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the media attachments client", func() {
		mediaAttachmentsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MediaAttachments

		Describe("When the page of media are successfully retrieved", func() {
			pageOptions := &media_attachments.MediaPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := mediaAttachmentsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the media page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				media := resp.Media
				Expect(media).ToNot(BeNil())
				Expect(len(media)).To(Equal(1))

				Expect(media[0].Sid).To(Equal("MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(media[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(media[0].ParentSid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(media[0].ContentType).To(Equal("image/jpeg"))
				Expect(media[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(media[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of media api returns a 500 response", func() {
			pageOptions := &media_attachments.MediaPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := mediaAttachmentsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the media page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated media are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := mediaAttachmentsClient.NewMediaPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated media current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated media results should be returned", func() {
				Expect(len(paginator.Media)).To(Equal(3))
			})
		})

		Describe("When the media api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := mediaAttachmentsClient.NewMediaPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated media current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a media attachment sid", func() {
		mediaAttachmentClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MediaAttachment("MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the media is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := mediaAttachmentClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get media response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParentSid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ContentType).To(Equal("image/jpeg"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the media api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/ME71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MediaAttachment("ME71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get media response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the media is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := mediaAttachmentClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the media api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/ME71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MediaAttachment("ME71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})

		Describe("Given I have a balance client", func() {
			balanceClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Balance()

			Describe("When the balance is successfully retrieved", func() {
				httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Balance.json",
					func(req *http.Request) (*http.Response, error) {
						fixture, _ := ioutil.ReadFile("testdata/balanceResponse.json")
						resp := make(map[string]interface{})
						json.Unmarshal(fixture, &resp)
						return httpmock.NewJsonResponse(200, resp)
					},
				)

				resp, err := balanceClient.Fetch()
				It("Then no error should be returned", func() {
					Expect(err).To(BeNil())
				})

				It("Then the get balance response should be returned", func() {
					Expect(resp).ToNot(BeNil())
					Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
					Expect(resp.Balance).To(Equal("1.00000"))
					Expect(resp.Currency).To(Equal("GBP"))
				})
			})

			Describe("When the balance api returns a 404", func() {
				httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Balance.json",
					func(req *http.Request) (*http.Response, error) {
						fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
						resp := make(map[string]interface{})
						json.Unmarshal(fixture, &resp)
						return httpmock.NewJsonResponse(500, resp)
					},
				)

				resp, err := balanceClient.Fetch()
				It("Then an error should be returned", func() {
					ExpectInternalServerError(err)
				})

				It("Then the get balance response should be nil", func() {
					Expect(resp).To(BeNil())
				})
			})
		})

		Describe("Given I have a tokens client", func() {
			tokensClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Tokens

			Describe("When the token is successfully created", func() {
				httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tokens.json",
					func(req *http.Request) (*http.Response, error) {
						fixture, _ := ioutil.ReadFile("testdata/tokenResponse.json")
						resp := make(map[string]interface{})
						json.Unmarshal(fixture, &resp)
						return httpmock.NewJsonResponse(200, resp)
					},
				)

				resp, err := tokensClient.Create(&tokens.CreateTokenInput{})
				It("Then no error should be returned", func() {
					Expect(err).To(BeNil())
				})

				It("Then the get token response should be returned", func() {
					Expect(resp).ToNot(BeNil())
					Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
					Expect(resp.Username).To(Equal("username"))
					Expect(resp.Password).To(Equal("password"))
					Expect(resp.IceServers).To(Equal([]tokens.CreateIceServerResponse{{
						URL:        "stun:global.stun.twilio.com:3478?transport=udp",
						URLs:       "stun:global.stun.twilio.com:3478?transport=udp",
						Username:   nil,
						Credential: nil,
					}, {
						URL:        "turn:global.turn.twilio.com:3478?transport=udp",
						URLs:       "turn:global.turn.twilio.com:3478?transport=udp",
						Username:   utils.String("username"),
						Credential: utils.String("password"),
					}, {
						URL:        "turn:global.turn.twilio.com:3478?transport=tcp",
						URLs:       "turn:global.turn.twilio.com:3478?transport=tcp",
						Username:   utils.String("username"),
						Credential: utils.String("password"),
					}, {
						URL:        "turn:global.turn.twilio.com:443?transport=tcp",
						URLs:       "turn:global.turn.twilio.com:443?transport=tcp",
						Username:   utils.String("username"),
						Credential: utils.String("password"),
					}}))
					Expect(resp.Ttl).To(Equal("1"))
					Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
					Expect(resp.DateUpdated).To(BeNil())
				})
			})

			Describe("When the token api returns a 404", func() {
				httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tokens.json",
					func(req *http.Request) (*http.Response, error) {
						fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
						resp := make(map[string]interface{})
						json.Unmarshal(fixture, &resp)
						return httpmock.NewJsonResponse(500, resp)
					},
				)

				resp, err := tokensClient.Create(&tokens.CreateTokenInput{})
				It("Then an error should be returned", func() {
					ExpectInternalServerError(err)
				})

				It("Then the get token response should be nil", func() {
					Expect(resp).To(BeNil())
				})
			})
		})
	})

	Describe("Given the queues client", func() {
		queuesClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queues

		Describe("When the queue is successfully created", func() {
			createInput := &queues.CreateQueueInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := queuesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.CurrentSize).To(Equal(0))
				Expect(resp.AverageWaitTime).To(Equal(0))
				Expect(resp.MaxSize).To(Equal(100))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the queue request does not contain a to", func() {
			createInput := &queues.CreateQueueInput{}

			resp, err := queuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create queue api returns a 500 response", func() {
			createInput := &queues.CreateQueueInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := queuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of queues are successfully retrieved", func() {
			pageOptions := &queues.QueuesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queuesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := queuesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the queues page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				queues := resp.Queues
				Expect(queues).ToNot(BeNil())
				Expect(len(queues)).To(Equal(1))

				Expect(queues[0].Sid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(queues[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(queues[0].FriendlyName).To(Equal("Test"))
				Expect(queues[0].CurrentSize).To(Equal(0))
				Expect(queues[0].AverageWaitTime).To(Equal(0))
				Expect(queues[0].MaxSize).To(Equal(100))
				Expect(queues[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(queues[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of queues api returns a 500 response", func() {
			pageOptions := &queues.QueuesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := queuesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the queues page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated queues are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queuesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queuesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := queuesClient.NewQueuesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated queues current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated queues results should be returned", func() {
				Expect(len(paginator.Queues)).To(Equal(3))
			})
		})

		Describe("When the queues api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queuesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := queuesClient.NewQueuesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated queue current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a queue sid", func() {
		queueClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the queue is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := queueClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.CurrentSize).To(Equal(0))
				Expect(resp.AverageWaitTime).To(Equal(0))
				Expect(resp.MaxSize).To(Equal(100))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the queue api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QU71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QU71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the queue is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &queue.UpdateQueueInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := queueClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CurrentSize).To(Equal(0))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.AverageWaitTime).To(Equal(0))
				Expect(resp.MaxSize).To(Equal(100))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the queue api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QU71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &queue.UpdateQueueInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QU71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the queue is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := queueClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the queue api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QU71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QU71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the members client", func() {
		membersClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Members

		Describe("When the page of members are successfully retrieved", func() {
			pageOptions := &members.MembersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/membersPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := membersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the members page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				members := resp.Members
				Expect(members).ToNot(BeNil())
				Expect(len(members)).To(Equal(1))

				Expect(members[0].QueueSid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(members[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(members[0].Position).To(Equal(1))
				Expect(members[0].WaitTime).To(Equal(100))
				Expect(members[0].DateEnqueued.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
			})
		})

		Describe("When the page of members api returns a 500 response", func() {
			pageOptions := &members.MembersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := membersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the members page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated members are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/membersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/membersPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := membersClient.NewMembersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated members current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated members results should be returned", func() {
				Expect(len(paginator.Members)).To(Equal(3))
			})
		})

		Describe("When the members api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/membersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := membersClient.NewMembersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated member current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a member sid", func() {
		memberClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the member is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/memberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := memberClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get member response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.QueueSid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Position).To(Equal(1))
				Expect(resp.WaitTime).To(Equal(100))
				Expect(resp.DateEnqueued.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
			})
		})

		Describe("When the member api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("CA71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the member is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/memberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &member.UpdateMemberInput{
				URL: "http://localhost",
			}

			resp, err := memberClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update member response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.QueueSid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Position).To(Equal(1))
				Expect(resp.WaitTime).To(Equal(100))
				Expect(resp.DateEnqueued.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
			})
		})

		Describe("When the member request does not contain a url", func() {
			updateInput := &member.UpdateMemberInput{}

			resp, err := memberClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the member api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &member.UpdateMemberInput{
				URL: "http://localhost",
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("CA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the calls client", func() {
		callsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls

		Describe("When the call is successfully created", func() {
			createInput := &calls.CreateCallInput{
				To:    "+10123456789",
				From:  "+19876543210",
				TwiML: utils.String(`<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello World</Say></Response>`),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := callsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create call response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AnsweredBy).To(BeNil())
				Expect(resp.CallerName).To(BeNil())
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.Duration).To(Equal("0"))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.ForwardedFrom).To(BeNil())
				Expect(resp.From).To(Equal("+19876543210"))
				Expect(resp.FromFormatted).To(Equal("+19876543210"))
				Expect(resp.GroupSid).To(BeNil())
				Expect(resp.ParentCallSid).To(BeNil())
				Expect(resp.PhoneNumberSid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.QueueTime).To(Equal("0"))
				Expect(resp.StartTime).To(BeNil())
				Expect(resp.Status).To(Equal("ringing"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.ToFormatted).To(Equal("+10123456789"))
				Expect(resp.TrunkSid).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the call request does not contain a to", func() {
			createInput := &calls.CreateCallInput{
				From: "+1987654321",
			}

			resp, err := callsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the call request does not contain a from", func() {
			createInput := &calls.CreateCallInput{
				To: "+10123456789",
			}

			resp, err := callsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create call api returns a 500 response", func() {
			createInput := &calls.CreateCallInput{
				To:    "+10123456789",
				From:  "+1987654321",
				TwiML: utils.String(`<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello World</Say></Response>`),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := callsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of calls are successfully retrieved", func() {
			pageOptions := &calls.CallsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := callsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the calls page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				calls := resp.Calls
				Expect(calls).ToNot(BeNil())
				Expect(len(calls)).To(Equal(1))

				Expect(calls[0].Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(calls[0].APIVersion).To(Equal("2010-04-01"))
				Expect(calls[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(calls[0].AnsweredBy).To(BeNil())
				Expect(calls[0].CallerName).To(BeNil())
				Expect(calls[0].Direction).To(Equal("outbound-api"))
				Expect(calls[0].Duration).To(Equal("0"))
				Expect(calls[0].EndTime).To(BeNil())
				Expect(calls[0].ForwardedFrom).To(BeNil())
				Expect(calls[0].From).To(Equal("+19876543210"))
				Expect(calls[0].FromFormatted).To(Equal("+19876543210"))
				Expect(calls[0].GroupSid).To(BeNil())
				Expect(calls[0].ParentCallSid).To(BeNil())
				Expect(calls[0].PhoneNumberSid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(calls[0].Price).To(BeNil())
				Expect(calls[0].PriceUnit).To(Equal(utils.String("GBP")))
				Expect(calls[0].QueueTime).To(Equal("0"))
				Expect(calls[0].StartTime).To(BeNil())
				Expect(calls[0].Status).To(Equal("ringing"))
				Expect(calls[0].To).To(Equal("+10123456789"))
				Expect(calls[0].ToFormatted).To(Equal("+10123456789"))
				Expect(calls[0].TrunkSid).To(BeNil())
				Expect(calls[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(calls[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of calls api returns a 500 response", func() {
			pageOptions := &calls.CallsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := callsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the calls page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated calls are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := callsClient.NewCallsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated calls current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated calls results should be returned", func() {
				Expect(len(paginator.Calls)).To(Equal(3))
			})
		})

		Describe("When the calls api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := callsClient.NewCallsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated calls current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a call sid", func() {
		callClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the call is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := callClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get call response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AnsweredBy).To(BeNil())
				Expect(resp.CallerName).To(BeNil())
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.Duration).To(Equal("0"))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.ForwardedFrom).To(BeNil())
				Expect(resp.From).To(Equal("+19876543210"))
				Expect(resp.FromFormatted).To(Equal("+19876543210"))
				Expect(resp.GroupSid).To(BeNil())
				Expect(resp.ParentCallSid).To(BeNil())
				Expect(resp.PhoneNumberSid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.QueueTime).To(Equal("0"))
				Expect(resp.StartTime).To(BeNil())
				Expect(resp.Status).To(Equal("ringing"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.ToFormatted).To(Equal("+10123456789"))
				Expect(resp.TrunkSid).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the call api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CA71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the call is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateCallResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &call.UpdateCallInput{
				Status: utils.String("Completed"),
			}

			resp, err := callClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update call response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AnsweredBy).To(BeNil())
				Expect(resp.CallerName).To(BeNil())
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.Duration).To(Equal("0"))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.ForwardedFrom).To(BeNil())
				Expect(resp.From).To(Equal("+19876543210"))
				Expect(resp.FromFormatted).To(Equal("+19876543210"))
				Expect(resp.GroupSid).To(BeNil())
				Expect(resp.ParentCallSid).To(BeNil())
				Expect(resp.PhoneNumberSid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.QueueTime).To(Equal("0"))
				Expect(resp.StartTime).To(BeNil())
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.ToFormatted).To(Equal("+10123456789"))
				Expect(resp.TrunkSid).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the call api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &call.UpdateCallInput{
				Status: utils.String("Completed"),
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the call is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := callClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the call api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CA71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the conferences client", func() {
		conferencesClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conferences

		Describe("When the page of conferences are successfully retrieved", func() {
			pageOptions := &conferences.ConferencesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferencesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conferencesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the conferences page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				conferences := resp.Conferences
				Expect(conferences).ToNot(BeNil())
				Expect(len(conferences)).To(Equal(1))

				Expect(conferences[0].Sid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(conferences[0].Status).To(Equal("in-progress"))
				Expect(conferences[0].ReasonConferenceEnded).To(BeNil())
				Expect(conferences[0].Region).To(Equal("us1"))
				Expect(conferences[0].FriendlyName).To(Equal("Test"))
				Expect(conferences[0].CallSidEndingConference).To(BeNil())
				Expect(conferences[0].APIVersion).To(Equal("2010-04-01"))
				Expect(conferences[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(conferences[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of conferences api returns a 500 response", func() {
			pageOptions := &conferences.ConferencesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conferencesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the conferences page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated conferences are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferencesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferencesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := conferencesClient.NewConferencesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated conferences current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated conferences results should be returned", func() {
				Expect(len(paginator.Conferences)).To(Equal(3))
			})
		})

		Describe("When the conferences api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferencesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := conferencesClient.NewConferencesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated conferences current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a conference sid", func() {
		conferenceClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the conference is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferenceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conferenceClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get conference response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.ReasonConferenceEnded).To(BeNil())
				Expect(resp.Region).To(Equal("us1"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.CallSidEndingConference).To(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the conference api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CF71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CF71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get conference response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conference is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateConferenceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &conference.UpdateConferenceInput{
				Status: utils.String("Completed"),
			}

			resp, err := conferenceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update conference response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.ReasonConferenceEnded).To(BeNil())
				Expect(resp.Region).To(Equal("us1"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.CallSidEndingConference).To(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the conference api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CF71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &conference.UpdateConferenceInput{
				Status: utils.String("Completed"),
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CF71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update conference response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the participants client", func() {
		participantsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participants

		Describe("When the participant is successfully created", func() {
			createInput := &participants.CreateParticipantInput{
				From: "+19876543210",
				To:   "+10123456789",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := participantsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Status).To(Equal("connected"))
				Expect(resp.ConferenceSid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Hold).To(Equal(false))
				Expect(resp.EndConferenceOnExit).To(Equal(false))
				Expect(resp.Label).To(BeNil())
				Expect(resp.Muted).To(Equal(false))
				Expect(resp.Coaching).To(Equal(false))
				Expect(resp.StartConferenceOnEnter).To(Equal(true))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSidToCoach).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the participant request does not contain a From", func() {
			createInput := &participants.CreateParticipantInput{
				To: "+10123456789",
			}

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create participants response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant request does not contain a To", func() {
			createInput := &participants.CreateParticipantInput{
				From: "+19876543210",
			}

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create participants response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create participant api returns a 500 response", func() {
			createInput := &participants.CreateParticipantInput{
				From: "+19876543210",
				To:   "+10123456789",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of participants are successfully retrieved", func() {
			pageOptions := &participants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := participantsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the participants page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				participants := resp.Participants
				Expect(participants).ToNot(BeNil())
				Expect(len(participants)).To(Equal(1))

				Expect(participants[0].Status).To(Equal("connected"))
				Expect(participants[0].ConferenceSid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].Hold).To(Equal(false))
				Expect(participants[0].EndConferenceOnExit).To(Equal(false))
				Expect(participants[0].Label).To(BeNil())
				Expect(participants[0].Muted).To(Equal(false))
				Expect(participants[0].Coaching).To(Equal(false))
				Expect(participants[0].StartConferenceOnEnter).To(Equal(true))
				Expect(participants[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].CallSidToCoach).To(BeNil())
				Expect(participants[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(participants[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of participants api returns a 500 response", func() {
			pageOptions := &participants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := participantsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the participants page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated participants are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := participantsClient.NewParticipantsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated participants current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated participants results should be returned", func() {
				Expect(len(paginator.Participants)).To(Equal(3))
			})
		})

		Describe("When the participants api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := participantsClient.NewParticipantsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated participants current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a participant sid", func() {
		participantClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the participant is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := participantClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Status).To(Equal("connected"))
				Expect(resp.ConferenceSid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Hold).To(Equal(false))
				Expect(resp.EndConferenceOnExit).To(Equal(false))
				Expect(resp.Label).To(BeNil())
				Expect(resp.Muted).To(Equal(false))
				Expect(resp.Coaching).To(Equal(false))
				Expect(resp.StartConferenceOnEnter).To(Equal(true))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSidToCoach).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("CA71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateParticipantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &participant.UpdateParticipantInput{
				Muted: utils.Bool(true),
			}

			resp, err := participantClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Status).To(Equal("connected"))
				Expect(resp.ConferenceSid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Hold).To(Equal(false))
				Expect(resp.EndConferenceOnExit).To(Equal(false))
				Expect(resp.Label).To(BeNil())
				Expect(resp.Muted).To(Equal(true))
				Expect(resp.Coaching).To(Equal(false))
				Expect(resp.StartConferenceOnEnter).To(Equal(true))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSidToCoach).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &participant.UpdateParticipantInput{
				Muted: utils.Bool(true),
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("CA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := participantClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("CA71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the addresses client", func() {
		addressesClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Addresses

		Describe("When the address is successfully created", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := addressesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create addresses response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.CustomerName).To(Equal("Test User"))
				Expect(resp.Street).To(Equal("123 Fake Street"))
				Expect(resp.StreetSecondary).To(BeNil())
				Expect(resp.City).To(Equal("Fake City"))
				Expect(resp.Region).To(Equal("Fake Region"))
				Expect(resp.PostalCode).To(Equal("AB12CD"))
				Expect(resp.IsoCountry).To(Equal("GB"))
				Expect(resp.EmergencyEnabled).To(Equal(false))
				Expect(resp.Validated).To(Equal(false))
				Expect(resp.Verified).To(Equal(false))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the address request does not contain a customer name", func() {
			createInput := &addresses.CreateAddressInput{
				Street:     "123 Fake Street",
				City:       "Fake City",
				Region:     "Fake Region",
				PostalCode: "AB12DC",
				IsoCountry: "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a street", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				City:         "Fake City",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a city", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a region", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a postal code", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				Region:       "Fake Region",
				IsoCountry:   "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a iso country", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create address api returns a 500 response", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of addresses are successfully retrieved", func() {
			pageOptions := &addresses.AddressesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := addressesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the addresses page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				addresses := resp.Addresses
				Expect(addresses).ToNot(BeNil())
				Expect(len(addresses)).To(Equal(1))

				Expect(addresses[0].Sid).To(Equal("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(addresses[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(addresses[0].FriendlyName).To(BeNil())
				Expect(addresses[0].CustomerName).To(Equal("Test User"))
				Expect(addresses[0].Street).To(Equal("123 Fake Street"))
				Expect(addresses[0].StreetSecondary).To(BeNil())
				Expect(addresses[0].City).To(Equal("Fake City"))
				Expect(addresses[0].Region).To(Equal("Fake Region"))
				Expect(addresses[0].PostalCode).To(Equal("AB12CD"))
				Expect(addresses[0].IsoCountry).To(Equal("GB"))
				Expect(addresses[0].EmergencyEnabled).To(Equal(false))
				Expect(addresses[0].Validated).To(Equal(false))
				Expect(addresses[0].Verified).To(Equal(false))
				Expect(addresses[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(addresses[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of addresses api returns a 500 response", func() {
			pageOptions := &addresses.AddressesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := addressesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the addresses page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated addresses are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := addressesClient.NewAddressesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated addresses current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated addresses results should be returned", func() {
				Expect(len(paginator.Addresses)).To(Equal(3))
			})
		})

		Describe("When the addresses api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := addressesClient.NewAddressesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated addresses current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a address sid", func() {
		addressClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Address("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the address is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := addressClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get address response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.CustomerName).To(Equal("Test User"))
				Expect(resp.Street).To(Equal("123 Fake Street"))
				Expect(resp.StreetSecondary).To(BeNil())
				Expect(resp.City).To(Equal("Fake City"))
				Expect(resp.Region).To(Equal("Fake Region"))
				Expect(resp.PostalCode).To(Equal("AB12CD"))
				Expect(resp.IsoCountry).To(Equal("GB"))
				Expect(resp.EmergencyEnabled).To(Equal(false))
				Expect(resp.Validated).To(Equal(false))
				Expect(resp.Verified).To(Equal(false))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the address api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/AD71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Address("AD71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateAddressResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &address.UpdateAddressInput{
				PostalCode: utils.String("Fake Postal Code"),
			}

			resp, err := addressClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update address response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.CustomerName).To(Equal("Test User"))
				Expect(resp.Street).To(Equal("123 Fake Street"))
				Expect(resp.StreetSecondary).To(BeNil())
				Expect(resp.City).To(Equal("Fake City"))
				Expect(resp.Region).To(Equal("Fake Region"))
				Expect(resp.PostalCode).To(Equal("Fake Postal Code"))
				Expect(resp.IsoCountry).To(Equal("GB"))
				Expect(resp.EmergencyEnabled).To(Equal(false))
				Expect(resp.Validated).To(Equal(false))
				Expect(resp.Verified).To(Equal(false))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the address api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/AD71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &address.UpdateAddressInput{
				PostalCode: utils.String("Fake Postal Code"),
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Address("AD71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := addressClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the address api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/AD71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Address("AD71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the recordings client", func() {
		recordingsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recordings

		Describe("When the page of recordings are successfully retrieved", func() {
			pageOptions := &recordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the recordings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].APIVersion).To(Equal("2010-04-01"))
				Expect(recordings[0].Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].ConferenceSid).To(BeNil())
				Expect(recordings[0].Duration).To(BeNil())
				Expect(recordings[0].Price).To(BeNil())
				Expect(recordings[0].PriceUnit).To(Equal(utils.String("GBP")))
				Expect(recordings[0].Status).To(Equal("in-progress"))
				Expect(recordings[0].Channels).To(Equal(1))
				Expect(recordings[0].Source).To(Equal("OutboundAPI"))
				Expect(recordings[0].ErrorCode).To(BeNil())
				Expect(recordings[0].EncryptionDetails).To(BeNil())
				Expect(recordings[0].StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of recordings api returns a 500 response", func() {
			pageOptions := &recordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the recordings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated recordings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated recordings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated recordings results should be returned", func() {
				Expect(len(paginator.Recordings)).To(Equal(3))
			})
		})

		Describe("When the recordings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated recordings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a recording sid", func() {
		recordingClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the recording is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get recording response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the recording is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateRecordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &recording.UpdateRecordingInput{
				Status: "completed",
			}

			resp, err := recordingClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update recording response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the recording update request does not contain a status", func() {
			updateInput := &recording.UpdateRecordingInput{}

			resp, err := recordingClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &recording.UpdateRecordingInput{
				Status: "completed",
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the recording is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := recordingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the call recordings client", func() {
		recordingsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recordings

		Describe("When the recordings is successfully created", func() {
			createInput := &callRecordings.CreateRecordingInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := recordingsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create addresses response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create recording api returns a 500 response", func() {
			createInput := &callRecordings.CreateRecordingInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of recordings are successfully retrieved", func() {
			pageOptions := &callRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the recordings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].APIVersion).To(Equal("2010-04-01"))
				Expect(recordings[0].Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].ConferenceSid).To(BeNil())
				Expect(recordings[0].Duration).To(BeNil())
				Expect(recordings[0].Price).To(BeNil())
				Expect(recordings[0].PriceUnit).To(Equal(utils.String("GBP")))
				Expect(recordings[0].Status).To(Equal("in-progress"))
				Expect(recordings[0].Channels).To(Equal(1))
				Expect(recordings[0].Source).To(Equal("OutboundAPI"))
				Expect(recordings[0].ErrorCode).To(BeNil())
				Expect(recordings[0].EncryptionDetails).To(BeNil())
				Expect(recordings[0].StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of recordings api returns a 500 response", func() {
			pageOptions := &callRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the recordings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated recordings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated recordings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated recordings results should be returned", func() {
				Expect(len(paginator.Recordings)).To(Equal(3))
			})
		})

		Describe("When the recordings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated recordings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a call recording sid", func() {
		recordingClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the recording is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get recording response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the conference recordings client", func() {
		recordingsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recordings

		Describe("When the recordings is successfully created", func() {
			createInput := &conferenceRecordings.CreateRecordingInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := recordingsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create addresses response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create recording api returns a 500 response", func() {
			createInput := &conferenceRecordings.CreateRecordingInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of recordings are successfully retrieved", func() {
			pageOptions := &conferenceRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the recordings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].APIVersion).To(Equal("2010-04-01"))
				Expect(recordings[0].Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].ConferenceSid).To(BeNil())
				Expect(recordings[0].Duration).To(BeNil())
				Expect(recordings[0].Price).To(BeNil())
				Expect(recordings[0].PriceUnit).To(Equal(utils.String("GBP")))
				Expect(recordings[0].Status).To(Equal("in-progress"))
				Expect(recordings[0].Channels).To(Equal(1))
				Expect(recordings[0].Source).To(Equal("OutboundAPI"))
				Expect(recordings[0].ErrorCode).To(BeNil())
				Expect(recordings[0].EncryptionDetails).To(BeNil())
				Expect(recordings[0].StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of recordings api returns a 500 response", func() {
			pageOptions := &conferenceRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the recordings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated recordings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated recordings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated recordings results should be returned", func() {
				Expect(len(paginator.Recordings)).To(Equal(3))
			})
		})

		Describe("When the recordings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated recordings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a conference recording sid", func() {
		recordingClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the recording is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get recording response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the feedback summaries client", func() {
		feedbackSummariesClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls.FeedbackSummaries

		Describe("When the feedback summary is successfully created", func() {
			createInput := &feedback_summaries.CreateFeedbackSummaryInput{
				StartDate: "2019-10-03",
				EndDate:   "2020-10-03",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/feedbackSummaryResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := feedbackSummariesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create feedback summary response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.StartDate).To(Equal("2019-10-03"))
				Expect(resp.EndDate).To(Equal("2020-10-03"))
				Expect(resp.IncludeSubaccounts).To(Equal(false))
				Expect(resp.Status).To(Equal("queued"))
				Expect(resp.CallCount).To(Equal(0))
				Expect(resp.CallFeedbackCount).To(Equal(0))
				Expect(resp.QualityScoreAverage).To(BeNil())
				Expect(resp.QualityScoreMedian).To(BeNil())
				Expect(resp.QualityScoreStandardDeviation).To(BeNil())
				Expect(resp.Issues).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the feedback summary request does not contain a start date", func() {
			createInput := &feedback_summaries.CreateFeedbackSummaryInput{
				EndDate: "2020-10-03",
			}

			resp, err := feedbackSummariesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create feedback summary  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the feedback summary request does not contain a end date", func() {
			createInput := &feedback_summaries.CreateFeedbackSummaryInput{
				StartDate: "2019-10-03",
			}

			resp, err := feedbackSummariesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create feedback summary response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create feedback summary api returns a 500 response", func() {
			createInput := &feedback_summaries.CreateFeedbackSummaryInput{
				StartDate: "2019-10-03",
				EndDate:   "2020-10-03",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := feedbackSummariesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create feedback summary response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a feedback summary sid", func() {
		feedbackSummaryClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls.FeedbackSummary("FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the feedback summary is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary/FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/feedbackSummaryResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := feedbackSummaryClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get feedback summary response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.StartDate).To(Equal("2019-10-03"))
				Expect(resp.EndDate).To(Equal("2020-10-03"))
				Expect(resp.IncludeSubaccounts).To(Equal(false))
				Expect(resp.Status).To(Equal("queued"))
				Expect(resp.CallCount).To(Equal(0))
				Expect(resp.CallFeedbackCount).To(Equal(0))
				Expect(resp.QualityScoreAverage).To(BeNil())
				Expect(resp.QualityScoreMedian).To(BeNil())
				Expect(resp.QualityScoreStandardDeviation).To(BeNil())
				Expect(resp.Issues).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the feedback summary api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary/FS71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls.FeedbackSummary("FS71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get feedback summary response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the feedback summary is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary/FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := feedbackSummaryClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the feedback summary api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary/FS71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls.FeedbackSummary("FS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the feedbacks client", func() {
		feedbacksClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Feedbacks

		Describe("When the feedback is successfully created", func() {
			createInput := &callFeedbacks.CreateFeedbackInput{
				QualityScore: 5,
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callFeedbackResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := feedbacksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create feedback response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.QualityScore).To(Equal(5))
				Expect(resp.Issues).To(Equal([]string{}))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the feedback request does not contain a quality score", func() {
			createInput := &callFeedbacks.CreateFeedbackInput{}

			resp, err := feedbacksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create feedback api returns a 500 response", func() {
			createInput := &callFeedbacks.CreateFeedbackInput{
				QualityScore: 5,
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := feedbacksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a feedback sid", func() {
		feedbackClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Feedback()

		Describe("When the feedback is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callFeedbackResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := feedbackClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get feedback response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.QualityScore).To(Equal(5))
				Expect(resp.Issues).To(Equal([]string{}))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the feedback api returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := feedbackClient.Fetch()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the feedback is successfully updated", func() {
			updateInput := &callFeedback.UpdateFeedbackInput{
				QualityScore: 4,
				Issues:       &[]string{"audio-latency"},
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateCallFeedbackResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := feedbackClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update feedback response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.QualityScore).To(Equal(4))
				Expect(resp.Issues).To(Equal([]string{"audio-latency"}))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the feedback request does not contain a quality score", func() {
			updateInput := &callFeedback.UpdateFeedbackInput{}

			resp, err := feedbackClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the update feedback api returns a 500 response", func() {
			updateInput := &callFeedback.UpdateFeedbackInput{
				QualityScore: 5,
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := feedbackClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the update feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the applications client", func() {
		applicationsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Applications

		Describe("When the application is successfully created", func() {
			createInput := &applications.CreateApplicationInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := applicationsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create applications response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.MessageStatusCallback).To(BeNil())
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsStatusCallback).To(BeNil())
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create applications api returns a 500 response", func() {
			createInput := &applications.CreateApplicationInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := applicationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create application response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of applications are successfully retrieved", func() {
			pageOptions := &applications.ApplicationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := applicationsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the applications page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				applications := resp.Applications
				Expect(applications).ToNot(BeNil())
				Expect(len(applications)).To(Equal(1))

				Expect(applications[0].Sid).To(Equal("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(applications[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(applications[0].APIVersion).To(Equal("2010-04-01"))
				Expect(applications[0].FriendlyName).To(BeNil())
				Expect(applications[0].MessageStatusCallback).To(BeNil())
				Expect(applications[0].SmsFallbackMethod).To(Equal("POST"))
				Expect(applications[0].SmsFallbackURL).To(BeNil())
				Expect(applications[0].SmsMethod).To(Equal("POST"))
				Expect(applications[0].SmsStatusCallback).To(BeNil())
				Expect(applications[0].SmsURL).To(BeNil())
				Expect(applications[0].StatusCallback).To(BeNil())
				Expect(applications[0].StatusCallbackMethod).To(Equal("POST"))
				Expect(applications[0].VoiceCallerIDLookup).To(Equal(false))
				Expect(applications[0].VoiceFallbackMethod).To(Equal("POST"))
				Expect(applications[0].VoiceFallbackURL).To(BeNil())
				Expect(applications[0].VoiceMethod).To(Equal("POST"))
				Expect(applications[0].VoiceURL).To(BeNil())
				Expect(applications[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(applications[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of applications api returns a 500 response", func() {
			pageOptions := &applications.ApplicationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := applicationsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the applications page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated applications are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := applicationsClient.NewApplicationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated applications current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated applications results should be returned", func() {
				Expect(len(paginator.Applications)).To(Equal(3))
			})
		})

		Describe("When the applications api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := applicationsClient.NewApplicationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated applications current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a application sid", func() {
		applicationClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Application("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the application is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := applicationClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get application response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.MessageStatusCallback).To(BeNil())
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsStatusCallback).To(BeNil())
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the application api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/AP71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Application("AP71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get application response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the application is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateApplicationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &application.UpdateApplicationInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := applicationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update application response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.MessageStatusCallback).To(BeNil())
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsStatusCallback).To(BeNil())
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the application api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/AP71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &application.UpdateApplicationInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Application("AP71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update application response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the application is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := applicationClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the application api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/AP71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Application("AP71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the available phone number countries client", func() {
		availablePhoneNumbersClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").AvailablePhoneNumbers

		Describe("When the page of available phone number countries are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/countriesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := availablePhoneNumbersClient.Page()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the available phone number countries page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers.json"))

				countries := resp.Countries
				Expect(countries).ToNot(BeNil())
				Expect(len(countries)).To(Equal(1))

				Expect(countries[0].CountryCode).To(Equal("GB"))
				Expect(countries[0].Country).To(Equal("United Kingdom"))
				Expect(countries[0].Beta).To(Equal(false))
			})
		})

		Describe("When the page of available phone number countries api returns a 500 response", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := availablePhoneNumbersClient.Page()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the available phone number countries page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have an available phone number country code", func() {
		availablePhoneNumberClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").AvailablePhoneNumber("GB")

		Describe("When the available phone number country is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/countryResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := availablePhoneNumberClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get available phone number country response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.CountryCode).To(Equal("GB"))
				Expect(resp.Country).To(Equal("United Kingdom"))
				Expect(resp.Beta).To(Equal(false))
			})
		})

		Describe("When the available phone number country api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/New.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").AvailablePhoneNumber("New").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get available phone number country response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the available local phone number countries client", func() {
		availableLocalPhoneNumbersClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").AvailablePhoneNumber("GB").Local

		Describe("When the page of available local phone number countries are successfully retrieved", func() {
			pageOptions := &local.AvailablePhoneNumbersPageOptions{
				PageSize: utils.Int(50),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB/Local.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/availableLocalPhoneNumberPageReponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := availableLocalPhoneNumbersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the available local phone number countries page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB/Local.json"))

				availablePhoneNumbers := resp.AvailablePhoneNumbers
				Expect(availablePhoneNumbers).ToNot(BeNil())
				Expect(len(availablePhoneNumbers)).To(Equal(1))

				Expect(availablePhoneNumbers[0].FriendlyName).To(Equal("+441234567890"))
				Expect(availablePhoneNumbers[0].PhoneNumber).To(Equal("+441234567890"))
				Expect(availablePhoneNumbers[0].Lata).To(BeNil())
				Expect(availablePhoneNumbers[0].RateCenter).To(BeNil())
				Expect(availablePhoneNumbers[0].Latitude).To(Equal("0.000000"))
				Expect(availablePhoneNumbers[0].Longitude).To(Equal("0.000000"))
				Expect(availablePhoneNumbers[0].Locality).To(BeNil())
				Expect(availablePhoneNumbers[0].Region).To(BeNil())
				Expect(availablePhoneNumbers[0].PostalCode).To(BeNil())
				Expect(availablePhoneNumbers[0].IsoCountry).To(Equal("GB"))
				Expect(availablePhoneNumbers[0].AddressRequirements).To(Equal("none"))
				Expect(availablePhoneNumbers[0].Beta).To(Equal(false))
				Expect(availablePhoneNumbers[0].Capabilities).To(Equal(local.PageAvailablePhoneNumberCapabilitiesResponse{
					Fax:   utils.Bool(true),
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
			})
		})

		Describe("When the page of available local phone number countries api returns a 500 response", func() {
			pageOptions := &local.AvailablePhoneNumbersPageOptions{
				PageSize: utils.Int(50),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB/Local.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := availableLocalPhoneNumbersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the available local phone number countries page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the available toll free phone number countries client", func() {
		availableTollFreePhoneNumbersClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").AvailablePhoneNumber("GB").TollFree

		Describe("When the page of available toll free phone number countries are successfully retrieved", func() {
			pageOptions := &toll_free.AvailablePhoneNumbersPageOptions{
				PageSize: utils.Int(50),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB/TollFree.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/availableTollFreePhoneNumberPageReponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := availableTollFreePhoneNumbersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the available toll free phone number countries page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB/TollFree.json"))

				availablePhoneNumbers := resp.AvailablePhoneNumbers
				Expect(availablePhoneNumbers).ToNot(BeNil())
				Expect(len(availablePhoneNumbers)).To(Equal(1))

				Expect(availablePhoneNumbers[0].FriendlyName).To(Equal("+441234567890"))
				Expect(availablePhoneNumbers[0].PhoneNumber).To(Equal("+441234567890"))
				Expect(availablePhoneNumbers[0].Lata).To(BeNil())
				Expect(availablePhoneNumbers[0].RateCenter).To(BeNil())
				Expect(availablePhoneNumbers[0].Latitude).To(Equal("0.000000"))
				Expect(availablePhoneNumbers[0].Longitude).To(Equal("0.000000"))
				Expect(availablePhoneNumbers[0].Locality).To(BeNil())
				Expect(availablePhoneNumbers[0].Region).To(BeNil())
				Expect(availablePhoneNumbers[0].PostalCode).To(BeNil())
				Expect(availablePhoneNumbers[0].IsoCountry).To(Equal("GB"))
				Expect(availablePhoneNumbers[0].AddressRequirements).To(Equal("none"))
				Expect(availablePhoneNumbers[0].Beta).To(Equal(false))
				Expect(availablePhoneNumbers[0].Capabilities).To(Equal(toll_free.PageAvailablePhoneNumberCapabilitiesResponse{
					Fax:   utils.Bool(true),
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
			})
		})

		Describe("When the page of available toll free phone number countries api returns a 500 response", func() {
			pageOptions := &toll_free.AvailablePhoneNumbersPageOptions{
				PageSize: utils.Int(50),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB/TollFree.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := availableTollFreePhoneNumbersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the available toll free phone number countries page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the available mobile phone number countries client", func() {
		availableMobilePhoneNumbersClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").AvailablePhoneNumber("GB").Mobile

		Describe("When the page of available mobile phone number countries are successfully retrieved", func() {
			pageOptions := &mobile.AvailablePhoneNumbersPageOptions{
				PageSize: utils.Int(50),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB/Mobile.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/availableMobilePhoneNumberPageReponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := availableMobilePhoneNumbersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the available mobile phone number countries page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB/Mobile.json"))

				availablePhoneNumbers := resp.AvailablePhoneNumbers
				Expect(availablePhoneNumbers).ToNot(BeNil())
				Expect(len(availablePhoneNumbers)).To(Equal(1))

				Expect(availablePhoneNumbers[0].FriendlyName).To(Equal("+441234567890"))
				Expect(availablePhoneNumbers[0].PhoneNumber).To(Equal("+441234567890"))
				Expect(availablePhoneNumbers[0].Lata).To(BeNil())
				Expect(availablePhoneNumbers[0].RateCenter).To(BeNil())
				Expect(availablePhoneNumbers[0].Latitude).To(Equal("0.000000"))
				Expect(availablePhoneNumbers[0].Longitude).To(Equal("0.000000"))
				Expect(availablePhoneNumbers[0].Locality).To(BeNil())
				Expect(availablePhoneNumbers[0].Region).To(BeNil())
				Expect(availablePhoneNumbers[0].PostalCode).To(BeNil())
				Expect(availablePhoneNumbers[0].IsoCountry).To(Equal("GB"))
				Expect(availablePhoneNumbers[0].AddressRequirements).To(Equal("none"))
				Expect(availablePhoneNumbers[0].Beta).To(Equal(false))
				Expect(availablePhoneNumbers[0].Capabilities).To(Equal(mobile.PageAvailablePhoneNumberCapabilitiesResponse{
					Fax:   utils.Bool(true),
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
			})
		})

		Describe("When the page of available mobile phone number countries api returns a 500 response", func() {
			pageOptions := &mobile.AvailablePhoneNumbersPageOptions{
				PageSize: utils.Int(50),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AvailablePhoneNumbers/GB/Mobile.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := availableMobilePhoneNumbersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the available mobile phone number countries page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the incoming phone numbers client", func() {
		incomingPhoneNumbersClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IncomingPhoneNumbers

		Describe("When the incoming phone number is successfully created", func() {
			createInput := &incoming_phone_numbers.CreateIncomingPhoneNumberInput{
				PhoneNumber: utils.String("+441234567890"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/incomingPhoneNumberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := incomingPhoneNumbersClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create incoming phone number response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("441234567890")))
				Expect(resp.PhoneNumber).To(Equal("+441234567890"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.SmsApplicationSid).To(BeNil())
				Expect(resp.VoiceApplicationSid).To(BeNil())
				Expect(resp.VoiceReceiveMode).To(Equal(utils.String("voice")))
				Expect(resp.AddressSid).To(BeNil())
				Expect(resp.IdentitySid).To(BeNil())
				Expect(resp.BundleSid).To(BeNil())
				Expect(resp.TrunkSid).To(BeNil())
				Expect(resp.EmergencyAddressSid).To(BeNil())
				Expect(resp.EmergencyStatus).To(Equal("Inactive"))
				Expect(resp.Origin).To(Equal("twilio"))
				Expect(resp.Beta).To(Equal(false))
				Expect(resp.Status).To(Equal("in-use"))
				Expect(resp.Capabilities).To(Equal(incoming_phone_numbers.CreateIncomingPhoneNumberCapabilitiesResponse{
					Fax:   utils.Bool(true),
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create incoming phone number api returns a 500 response", func() {
			createInput := &incoming_phone_numbers.CreateIncomingPhoneNumberInput{
				PhoneNumber: utils.String("+441234567890"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := incomingPhoneNumbersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create incoming phone number response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of incoming phone numbers are successfully retrieved", func() {
			pageOptions := &incoming_phone_numbers.IncomingPhoneNumbersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/incomingPhoneNumbersPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := incomingPhoneNumbersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the incoming phone numbers page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				phoneNumbers := resp.PhoneNumbers
				Expect(phoneNumbers).ToNot(BeNil())
				Expect(len(phoneNumbers)).To(Equal(1))

				Expect(phoneNumbers[0].Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(phoneNumbers[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(phoneNumbers[0].FriendlyName).To(Equal(utils.String("441234567890")))
				Expect(phoneNumbers[0].PhoneNumber).To(Equal("+441234567890"))
				Expect(phoneNumbers[0].APIVersion).To(Equal("2010-04-01"))
				Expect(phoneNumbers[0].SmsFallbackMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].SmsFallbackURL).To(BeNil())
				Expect(phoneNumbers[0].SmsMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].SmsURL).To(BeNil())
				Expect(phoneNumbers[0].SmsApplicationSid).To(BeNil())
				Expect(phoneNumbers[0].VoiceApplicationSid).To(BeNil())
				Expect(phoneNumbers[0].VoiceReceiveMode).To(Equal(utils.String("voice")))
				Expect(phoneNumbers[0].AddressSid).To(BeNil())
				Expect(phoneNumbers[0].IdentitySid).To(BeNil())
				Expect(phoneNumbers[0].BundleSid).To(BeNil())
				Expect(phoneNumbers[0].TrunkSid).To(BeNil())
				Expect(phoneNumbers[0].EmergencyAddressSid).To(BeNil())
				Expect(phoneNumbers[0].EmergencyStatus).To(Equal("Inactive"))
				Expect(phoneNumbers[0].Origin).To(Equal("twilio"))
				Expect(phoneNumbers[0].Beta).To(Equal(false))
				Expect(phoneNumbers[0].Status).To(Equal("in-use"))
				Expect(phoneNumbers[0].Capabilities).To(Equal(incoming_phone_numbers.PageIncomingPhoneNumberCapabilitiesResponse{
					Fax:   utils.Bool(true),
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
				Expect(phoneNumbers[0].StatusCallback).To(BeNil())
				Expect(phoneNumbers[0].StatusCallbackMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].VoiceCallerIDLookup).To(Equal(false))
				Expect(phoneNumbers[0].VoiceFallbackMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].VoiceFallbackURL).To(BeNil())
				Expect(phoneNumbers[0].VoiceMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].VoiceURL).To(BeNil())
				Expect(phoneNumbers[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(phoneNumbers[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of incoming phone numbers api returns a 500 response", func() {
			pageOptions := &incoming_phone_numbers.IncomingPhoneNumbersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := incomingPhoneNumbersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the incoming phone numbers page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated incoming phone numbers are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/incomingPhoneNumbersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/incomingPhoneNumbersPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := incomingPhoneNumbersClient.NewIncomingPhoneNumbersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated incoming phone numbers current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated incoming phone numbers results should be returned", func() {
				Expect(len(paginator.PhoneNumbers)).To(Equal(3))
			})
		})

		Describe("When the incoming phone numbers api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/incomingPhoneNumbersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := incomingPhoneNumbersClient.NewIncomingPhoneNumbersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated incoming phone numbers current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a incoming phone number sid", func() {
		incomingPhoneNumberClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IncomingPhoneNumber("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the incoming phone number is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/incomingPhoneNumberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := incomingPhoneNumberClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get incoming phone number response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("441234567890")))
				Expect(resp.PhoneNumber).To(Equal("+441234567890"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.SmsApplicationSid).To(BeNil())
				Expect(resp.VoiceApplicationSid).To(BeNil())
				Expect(resp.VoiceReceiveMode).To(Equal(utils.String("voice")))
				Expect(resp.AddressSid).To(BeNil())
				Expect(resp.IdentitySid).To(BeNil())
				Expect(resp.BundleSid).To(BeNil())
				Expect(resp.TrunkSid).To(BeNil())
				Expect(resp.EmergencyAddressSid).To(BeNil())
				Expect(resp.EmergencyStatus).To(Equal("Inactive"))
				Expect(resp.Origin).To(Equal("twilio"))
				Expect(resp.Beta).To(Equal(false))
				Expect(resp.Status).To(Equal("in-use"))
				Expect(resp.Capabilities).To(Equal(incoming_phone_number.FetchIncomingPhoneNumberCapabilitiesResponse{
					Fax:   utils.Bool(true),
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the incoming phone number api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers/PN71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IncomingPhoneNumber("PN71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get incoming phone number response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the incoming phone number is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateIncomingPhoneNumberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &incoming_phone_number.UpdateIncomingPhoneNumberInput{
				AddressSid: utils.String("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			resp, err := incomingPhoneNumberClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update incoming phone number response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("441234567890")))
				Expect(resp.PhoneNumber).To(Equal("+441234567890"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.SmsApplicationSid).To(BeNil())
				Expect(resp.VoiceApplicationSid).To(BeNil())
				Expect(resp.VoiceReceiveMode).To(Equal(utils.String("voice")))
				Expect(resp.AddressSid).To(Equal(utils.String("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.IdentitySid).To(BeNil())
				Expect(resp.BundleSid).To(BeNil())
				Expect(resp.TrunkSid).To(BeNil())
				Expect(resp.EmergencyAddressSid).To(BeNil())
				Expect(resp.EmergencyStatus).To(Equal("Inactive"))
				Expect(resp.Origin).To(Equal("twilio"))
				Expect(resp.Beta).To(Equal(false))
				Expect(resp.Status).To(Equal("in-use"))
				Expect(resp.Capabilities).To(Equal(incoming_phone_number.UpdateIncomingPhoneNumberCapabilitiesResponse{
					Fax:   utils.Bool(true),
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the incoming phone number api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers/PN71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &incoming_phone_number.UpdateIncomingPhoneNumberInput{
				AddressSid: utils.String("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IncomingPhoneNumber("PN71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update incoming phone number response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the incoming phone number is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := incomingPhoneNumberClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the incoming phone number api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers/PN71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IncomingPhoneNumber("PN71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the credential lists client", func() {
		credentialListsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialLists

		Describe("When the credential list is successfully created", func() {
			createInput := &credential_lists.CreateCredentialListInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := credentialListsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create credential list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the credential list request does not contain a friendly name", func() {
			createInput := &credential_lists.CreateCredentialListInput{}

			resp, err := credentialListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create credential list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create credential list api returns a 500 response", func() {
			createInput := &credential_lists.CreateCredentialListInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create credential list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of credential lists are successfully retrieved", func() {
			pageOptions := &credential_lists.CredentialListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialListsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the credential lists page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				credentialLists := resp.CredentialLists
				Expect(credentialLists).ToNot(BeNil())
				Expect(len(credentialLists)).To(Equal(1))

				Expect(credentialLists[0].Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentialLists[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentialLists[0].FriendlyName).To(Equal("Test"))
				Expect(credentialLists[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(credentialLists[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of credential lists api returns a 500 response", func() {
			pageOptions := &credential_lists.CredentialListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialListsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the credential lists page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated credential lists are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := credentialListsClient.NewCredentialListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated credential lists current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated credential lists results should be returned", func() {
				Expect(len(paginator.CredentialLists)).To(Equal(3))
			})
		})

		Describe("When the credential lists api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := credentialListsClient.NewCredentialListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated credential lists current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a credential list sid", func() {
		credentialListClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialList("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the credential list is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialListClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get credential list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the credential list api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialList("CL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get credential list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential list is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateCredentialListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &credential_list.UpdateCredentialListInput{
				FriendlyName: "Test 2",
			}

			resp, err := credentialListClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update credential list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the credential list request does not contain a friendly name", func() {
			updateInput := &credential_list.UpdateCredentialListInput{}

			resp, err := credentialListClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update credential list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential list api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &credential_list.UpdateCredentialListInput{
				FriendlyName: "Test 2",
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialList("CL71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update credential list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential list is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := credentialListClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the credential list api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialList("CL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the Credentials client", func() {
		credentialsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialList("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Credentials

		Describe("When the credential is successfully created", func() {
			createInput := &sipCredentials.CreateCredentialInput{
				Username: "Test",
				Password: "test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := credentialsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create credential response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CredentialListSid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Username).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the credential request does not contain an password", func() {
			createInput := &sipCredentials.CreateCredentialInput{
				Username: "Test",
			}

			resp, err := credentialsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential request does not contain a username", func() {
			createInput := &sipCredentials.CreateCredentialInput{
				Password: "test",
			}

			resp, err := credentialsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create credential api returns a 500 response", func() {
			createInput := &sipCredentials.CreateCredentialInput{
				Username: "Test",
				Password: "test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of credentials are successfully retrieved", func() {
			pageOptions := &sipCredentials.CredentialsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the credentials page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				credentials := resp.Credentials
				Expect(credentials).ToNot(BeNil())
				Expect(len(credentials)).To(Equal(1))

				Expect(credentials[0].Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentials[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentials[0].CredentialListSid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentials[0].Username).To(Equal("Test"))
				Expect(credentials[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(credentials[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of credentials api returns a 500 response", func() {
			pageOptions := &sipCredentials.CredentialsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the credentials page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated credentials are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := credentialsClient.NewCredentialsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated Credentials current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated Credentials results should be returned", func() {
				Expect(len(paginator.Credentials)).To(Equal(3))
			})
		})

		Describe("When the credentials api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := credentialsClient.NewCredentialsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated Credentials current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a credential sid", func() {
		credentialClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialList("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Credential("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the credential is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get credential response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CredentialListSid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Username).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the credential api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials/CR71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialList("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Credential("CR71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateCredentialResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &sipCredential.UpdateCredentialInput{
				Password: "Test 2",
			}

			resp, err := credentialClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update credential response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CredentialListSid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Username).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the credential request does not contain an password", func() {
			updateInput := &sipCredential.UpdateCredentialInput{}

			resp, err := credentialClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update Credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials/CR71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &sipCredential.UpdateCredentialInput{
				Password: "Test 2",
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialList("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Credential("CR71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := credentialClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the credential api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials/CR71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.CredentialList("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Credential("CR71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the IP access control lists client", func() {
		ipAccessControlListsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlLists

		Describe("When the IP access control list is successfully created", func() {
			createInput := &ip_access_control_lists.CreateIpAccessControlListInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create IP access control list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the IP access control list request does not contain a friendly name", func() {
			createInput := &ip_access_control_lists.CreateIpAccessControlListInput{}

			resp, err := ipAccessControlListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create IP access control list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create IP access control list api returns a 500 response", func() {
			createInput := &ip_access_control_lists.CreateIpAccessControlListInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create IP access control list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of IP access control lists are successfully retrieved", func() {
			pageOptions := &ip_access_control_lists.IpAccessControlListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the IP access control lists page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				ipAccessControlLists := resp.IpAccessControlLists
				Expect(ipAccessControlLists).ToNot(BeNil())
				Expect(len(ipAccessControlLists)).To(Equal(1))

				Expect(ipAccessControlLists[0].Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAccessControlLists[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAccessControlLists[0].FriendlyName).To(Equal("Test"))
				Expect(ipAccessControlLists[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(ipAccessControlLists[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of IP access control lists api returns a 500 response", func() {
			pageOptions := &ip_access_control_lists.IpAccessControlListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the IP access control lists page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated IP access control lists are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := ipAccessControlListsClient.NewIpAccessControlListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated IP access control lists current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated IP access control lists results should be returned", func() {
				Expect(len(paginator.IpAccessControlLists)).To(Equal(3))
			})
		})

		Describe("When the IP access control lists api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := ipAccessControlListsClient.NewIpAccessControlListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated IP access control lists current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a IP access control list sid", func() {
		ipAccessControlListClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlList("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the IP access control list is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := ipAccessControlListClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get IP access control list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the IP access control list api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/AL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlList("AL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get IP access control list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the IP access control list is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateIpAccessControlListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &ip_access_control_list.UpdateIpAccessControlListInput{
				FriendlyName: "Test 2",
			}

			resp, err := ipAccessControlListClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update IP access control list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the IP access control list request does not contain a friendly name", func() {
			updateInput := &ip_access_control_list.UpdateIpAccessControlListInput{}

			resp, err := ipAccessControlListClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update IP access control list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the IP access control list api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/AL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &ip_access_control_list.UpdateIpAccessControlListInput{
				FriendlyName: "Test 2",
			}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlList("AL71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update IP access control list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the IP access control list is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := ipAccessControlListClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the IP access control list api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/AL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlList("AL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the IP Addresses client", func() {
		ipAddressesClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlList("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IpAddresses

		Describe("When the IP Address is successfully created", func() {
			createInput := &ip_addresses.CreateIpAddressInput{
				FriendlyName: "Test",
				IpAddress:    "127.0.0.1",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAddressResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := ipAddressesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create IP Address response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.IpAccessControlListSid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.IpAddress).To(Equal("127.0.0.1"))
				Expect(resp.CidrPrefixLength).To(Equal(32))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the IP Address request does not contain an IP Address", func() {
			createInput := &ip_addresses.CreateIpAddressInput{
				FriendlyName: "Test",
			}

			resp, err := ipAddressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create IP Address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the IP Address request does not contain a friendly name", func() {
			createInput := &ip_addresses.CreateIpAddressInput{
				IpAddress: "127.0.0.1",
			}

			resp, err := ipAddressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create IP Address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create IP Address api returns a 500 response", func() {
			createInput := &ip_addresses.CreateIpAddressInput{
				FriendlyName: "Test",
				IpAddress:    "127.0.0.1",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := ipAddressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create IP Address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of IP Addresses are successfully retrieved", func() {
			pageOptions := &ip_addresses.IpAddressesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAddressesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := ipAddressesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the IP Addresses page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				ipAddresses := resp.IpAddresses
				Expect(ipAddresses).ToNot(BeNil())
				Expect(len(ipAddresses)).To(Equal(1))

				Expect(ipAddresses[0].Sid).To(Equal("IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAddresses[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAddresses[0].IpAccessControlListSid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAddresses[0].FriendlyName).To(Equal("Test"))
				Expect(ipAddresses[0].IpAddress).To(Equal("127.0.0.1"))
				Expect(ipAddresses[0].CidrPrefixLength).To(Equal(32))
				Expect(ipAddresses[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(ipAddresses[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of IP Addresses api returns a 500 response", func() {
			pageOptions := &ip_addresses.IpAddressesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := ipAddressesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the IP Addresses page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated IP Addresses are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAddressesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAddressesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := ipAddressesClient.NewIpAddressesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated IP Addresses current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated IP Addresses results should be returned", func() {
				Expect(len(paginator.IpAddresses)).To(Equal(3))
			})
		})

		Describe("When the IP Addresses api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := ipAddressesClient.NewIpAddressesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated IP Addresses current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a IP Address sid", func() {
		ipAddressClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlList("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IpAddress("IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the IP Address is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses/IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAddressResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := ipAddressClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get IP Address response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.IpAccessControlListSid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.IpAddress).To(Equal("127.0.0.1"))
				Expect(resp.CidrPrefixLength).To(Equal(32))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the IP Address api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses/IP71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlList("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IpAddress("IP71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get IP Address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the IP Address is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses/IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateIpAddressResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &ip_address.UpdateIpAddressInput{}

			resp, err := ipAddressClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update IP Address response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.IpAccessControlListSid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.IpAddress).To(Equal("127.0.0.1"))
				Expect(resp.CidrPrefixLength).To(Equal(32))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the IP Address api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses/IP71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &ip_address.UpdateIpAddressInput{}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlList("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IpAddress("IP71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update IP Address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the IP Address is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses/IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := ipAddressClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the IP Address api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses/IP71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.IpAccessControlList("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IpAddress("IP71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the domains client", func() {
		domainsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domains

		Describe("When the domains is successfully created", func() {
			createInput := &domains.CreateDomainInput{
				DomainName: "test" + ".sip.twilio.com",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/domainResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := domainsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create domain response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ApiVersion).To(Equal("2010-04-01"))
				Expect(resp.AuthType).To(BeNil())
				Expect(resp.ByocTrunkSid).To(BeNil())
				Expect(resp.DomainName).To(Equal("test.sip.twilio.com"))
				Expect(resp.EmergencyCallerSid).To(BeNil())
				Expect(resp.EmergencyCallingEnabled).To(Equal(false))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.Secure).To(Equal(false))
				Expect(resp.SipRegistration).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(BeNil())
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(BeNil())
				Expect(resp.VoiceStatusCallbackMethod).To(BeNil())
				Expect(resp.VoiceStatusCallbackURL).To(BeNil())
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the domain request does not contain a domain", func() {
			createInput := &domains.CreateDomainInput{}

			resp, err := domainsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create domain response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create domain api returns a 500 response", func() {
			createInput := &domains.CreateDomainInput{
				DomainName: "test" + ".sip.twilio.com",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := domainsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create domains response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of domains are successfully retrieved", func() {
			pageOptions := &domains.DomainsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/domainsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := domainsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the domains page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				domains := resp.Domains
				Expect(domains).ToNot(BeNil())
				Expect(len(domains)).To(Equal(1))

				Expect(domains[0].Sid).To(Equal("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(domains[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(domains[0].ApiVersion).To(Equal("2010-04-01"))
				Expect(domains[0].AuthType).To(BeNil())
				Expect(domains[0].ByocTrunkSid).To(BeNil())
				Expect(domains[0].DomainName).To(Equal("test.sip.twilio.com"))
				Expect(domains[0].EmergencyCallerSid).To(BeNil())
				Expect(domains[0].EmergencyCallingEnabled).To(Equal(false))
				Expect(domains[0].FriendlyName).To(BeNil())
				Expect(domains[0].Secure).To(Equal(false))
				Expect(domains[0].SipRegistration).To(Equal(false))
				Expect(domains[0].VoiceFallbackMethod).To(BeNil())
				Expect(domains[0].VoiceFallbackURL).To(BeNil())
				Expect(domains[0].VoiceMethod).To(BeNil())
				Expect(domains[0].VoiceStatusCallbackMethod).To(BeNil())
				Expect(domains[0].VoiceStatusCallbackURL).To(BeNil())
				Expect(domains[0].VoiceURL).To(BeNil())
				Expect(domains[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(domains[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of domains api returns a 500 response", func() {
			pageOptions := &domains.DomainsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := domainsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the domains page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated domains are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/domainsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/domainsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := domainsClient.NewDomainsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated domains current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated domains results should be returned", func() {
				Expect(len(paginator.Domains)).To(Equal(3))
			})
		})

		Describe("When the domains api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/domainsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := domainsClient.NewDomainsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated domains current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a domain sid", func() {
		domainClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the domain is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/domainResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := domainClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get domains response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ApiVersion).To(Equal("2010-04-01"))
				Expect(resp.AuthType).To(BeNil())
				Expect(resp.ByocTrunkSid).To(BeNil())
				Expect(resp.DomainName).To(Equal("test.sip.twilio.com"))
				Expect(resp.EmergencyCallerSid).To(BeNil())
				Expect(resp.EmergencyCallingEnabled).To(Equal(false))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.Secure).To(Equal(false))
				Expect(resp.SipRegistration).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(BeNil())
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(BeNil())
				Expect(resp.VoiceStatusCallbackMethod).To(BeNil())
				Expect(resp.VoiceStatusCallbackURL).To(BeNil())
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the domain api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SD71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SD71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get domain response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the domain is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateDomainResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &domain.UpdateDomainInput{}

			resp, err := domainClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update domain response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ApiVersion).To(Equal("2010-04-01"))
				Expect(resp.AuthType).To(BeNil())
				Expect(resp.ByocTrunkSid).To(BeNil())
				Expect(resp.DomainName).To(Equal("test.sip.twilio.com"))
				Expect(resp.EmergencyCallerSid).To(BeNil())
				Expect(resp.EmergencyCallingEnabled).To(Equal(false))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.Secure).To(Equal(false))
				Expect(resp.SipRegistration).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(BeNil())
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(BeNil())
				Expect(resp.VoiceStatusCallbackMethod).To(BeNil())
				Expect(resp.VoiceStatusCallbackURL).To(BeNil())
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the domain api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SD71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &domain.UpdateDomainInput{}

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SD71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update domain response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the domain is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := domainClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the domain api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SD71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SD71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the credential list mappings client", func() {
		credentialListMappingsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Calls.CredentialListMappings

		Describe("When the credential list mapping is successfully created", func() {
			createInput := &credential_list_mappings.CreateCredentialListMappingInput{
				CredentialListSid: "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListMappingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := credentialListMappingsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create credential list mapping response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the credential list mapping request does not contain a credential list sid", func() {
			createInput := &credential_list_mappings.CreateCredentialListMappingInput{}

			resp, err := credentialListMappingsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create credential list mapping response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create credential list mapping api returns a 500 response", func() {
			createInput := &credential_list_mappings.CreateCredentialListMappingInput{
				CredentialListSid: "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialListMappingsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create credential list mapping response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of credential list mappings are successfully retrieved", func() {
			pageOptions := &credential_list_mappings.CredentialListMappingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListMappingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialListMappingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the credential list mappings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				credentialListMappings := resp.CredentialListMappings
				Expect(credentialListMappings).ToNot(BeNil())
				Expect(len(credentialListMappings)).To(Equal(1))

				Expect(credentialListMappings[0].Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentialListMappings[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentialListMappings[0].FriendlyName).To(Equal("Test"))
				Expect(credentialListMappings[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(credentialListMappings[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of credential list mappings api returns a 500 response", func() {
			pageOptions := &credential_list_mappings.CredentialListMappingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialListMappingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the credential list mappings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated credential list mappings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListMappingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListMappingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := credentialListMappingsClient.NewCredentialListMappingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated credential list mappings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated credential list mappings results should be returned", func() {
				Expect(len(paginator.CredentialListMappings)).To(Equal(3))
			})
		})

		Describe("When the credential list mappings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListMappingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := credentialListMappingsClient.NewCredentialListMappingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated credential list mappings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a credential list mapping sid", func() {
		credentialListMappingClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Calls.CredentialListMapping("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the credential list mapping is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListMappingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialListMappingClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get credential list mapping response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the credential list mapping api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings/CL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Calls.CredentialListMapping("CL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get credential list mapping response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential list mapping is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := credentialListMappingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the credential list mapping api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings/CL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Calls.CredentialListMapping("CL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the SIP registration credential list mappings client", func() {
		credentialListMappingsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Registrations.CredentialListMappings

		Describe("When the credential list mapping is successfully created", func() {
			createInput := &sipRegistrationsCredentialListMappings.CreateCredentialListMappingInput{
				CredentialListSid: "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sipRegistrationCredentialListMappingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := credentialListMappingsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create credential list mapping response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the credential list mapping request does not contain a credential list sid", func() {
			createInput := &sipRegistrationsCredentialListMappings.CreateCredentialListMappingInput{}

			resp, err := credentialListMappingsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create credential list mapping response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create credential list mapping api returns a 500 response", func() {
			createInput := &sipRegistrationsCredentialListMappings.CreateCredentialListMappingInput{
				CredentialListSid: "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialListMappingsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create credential list mapping response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of credential list mappings are successfully retrieved", func() {
			pageOptions := &sipRegistrationsCredentialListMappings.CredentialListMappingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sipRegistrationCredentialListMappingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialListMappingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the credential list mappings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				credentialListMappings := resp.CredentialListMappings
				Expect(credentialListMappings).ToNot(BeNil())
				Expect(len(credentialListMappings)).To(Equal(1))

				Expect(credentialListMappings[0].Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentialListMappings[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentialListMappings[0].FriendlyName).To(Equal("Test"))
				Expect(credentialListMappings[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(credentialListMappings[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of credential list mappings api returns a 500 response", func() {
			pageOptions := &sipRegistrationsCredentialListMappings.CredentialListMappingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialListMappingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the credential list mappings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated credential list mappings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sipRegistrationCredentialListMappingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sipRegistrationCredentialListMappingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := credentialListMappingsClient.NewCredentialListMappingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated credential list mappings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated credential list mappings results should be returned", func() {
				Expect(len(paginator.CredentialListMappings)).To(Equal(3))
			})
		})

		Describe("When the credential list mappings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sipRegistrationCredentialListMappingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := credentialListMappingsClient.NewCredentialListMappingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated credential list mappings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a SIP registration credential list mapping sid", func() {
		credentialListMappingClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Registrations.CredentialListMapping("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the credential list mapping is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sipRegistrationCredentialListMappingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialListMappingClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get credential list mapping response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the credential list mapping api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings/CL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Registrations.CredentialListMapping("CL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get credential list mapping response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential list mapping is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := credentialListMappingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the credential list mapping api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Registrations/CredentialListMappings/CL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Registrations.CredentialListMapping("CL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the IP access control list mappings client", func() {
		ipAccessControlListsClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Calls.IpAccessControlListMappings

		Describe("When the IP access control list mapping is successfully created", func() {
			createInput := &ip_access_control_list_mappings.CreateIpAccessControlListMappingInput{
				IpAccessControlListSid: "ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListMappingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create IP access control list mapping response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the IP access control list mapping request does not contain a IP access control list sid", func() {
			createInput := &ip_access_control_list_mappings.CreateIpAccessControlListMappingInput{}

			resp, err := ipAccessControlListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create IP access control list mapping response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create IP access control list mapping api returns a 500 response", func() {
			createInput := &ip_access_control_list_mappings.CreateIpAccessControlListMappingInput{
				IpAccessControlListSid: "ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create IP access control list mapping response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of IP access control list mappings are successfully retrieved", func() {
			pageOptions := &ip_access_control_list_mappings.IpAccessControlListMappingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListMappingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the IP access control list mappings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				ipAccessControlLists := resp.IpAccessControlListMappings
				Expect(ipAccessControlLists).ToNot(BeNil())
				Expect(len(ipAccessControlLists)).To(Equal(1))

				Expect(ipAccessControlLists[0].Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAccessControlLists[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAccessControlLists[0].FriendlyName).To(Equal("Test"))
				Expect(ipAccessControlLists[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(ipAccessControlLists[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of IP access control list mappings api returns a 500 response", func() {
			pageOptions := &ip_access_control_list_mappings.IpAccessControlListMappingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the IP access control list mappings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated IP access control list mappings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListMappingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListMappingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := ipAccessControlListsClient.NewIpAccessControlListMappingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated IP access control list mappings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated IP access control list mappings results should be returned", func() {
				Expect(len(paginator.IpAccessControlListMappings)).To(Equal(3))
			})
		})

		Describe("When the IP access control list mappings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListMappingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := ipAccessControlListsClient.NewIpAccessControlListMappingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated IP access control list mappings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a IP access control list mapping sid", func() {
		ipAccessControlListClient := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Calls.IpAccessControlListMapping("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the IP access control list mapping is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListMappingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := ipAccessControlListClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get IP access control list mapping response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the IP access control list mapping api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings/AL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Calls.IpAccessControlListMapping("AL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get IP access control list mapping response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the IP access control list mapping is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := ipAccessControlListClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the IP access control list mapping api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/SIP/Domains/SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings/AL71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiClient.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sip.Domain("SDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Auth.Calls.IpAccessControlListMapping("AL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})
})

func ExpectInternalServerError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))
	Expect(twilioErr.Code).To(BeNil())
	Expect(twilioErr.Message).To(Equal("An error occurred"))
	Expect(twilioErr.MoreInfo).To(BeNil())
	Expect(twilioErr.Status).To(Equal(500))
}

func ExpectInvalidInputError(err error) {
	ExpectErrorToNotBeATwilioError(err)
	Expect(err.Error()).To(Equal("Invalid input supplied"))
}

func ExpectNotFoundError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))

	code := 20404
	Expect(twilioErr.Code).To(Equal(&code))
	Expect(twilioErr.Message).To(Equal("The requested resource /Account/AC71.json was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
