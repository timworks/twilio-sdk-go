package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2010 "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/mobile"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var apiSession *v2010.V2010

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	apiSession = twilio.NewWithCredentials(creds).API.V2010
}

func main() {
	resp, err := apiSession.
		Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		AvailablePhoneNumber("GB").
		Mobile.
		Page(&mobile.AvailablePhoneNumbersPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v available mobile phone number(s) found on page", len(resp.AvailablePhoneNumbers))
}
