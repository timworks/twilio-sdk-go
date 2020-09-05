package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/messaging/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service/short_codes"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var messagingSession *v1.Messaging

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	messagingSession = twilio.NewWithCredentials(creds).Messaging.V1
}

func main() {
	resp, err := messagingSession.
		Service("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		ShortCodes.
		Create(&short_codes.CreateShortCodeInput{
			ShortCodeSid: "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
