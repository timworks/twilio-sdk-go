package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/timworks/twilio-sdk-go/service/video/v1/composition_hook"
	"github.com/timworks/twilio-sdk-go/service/video/v1/composition_hooks"
	"github.com/timworks/twilio-sdk-go/service/video/v1/composition_settings"
	"github.com/timworks/twilio-sdk-go/service/video/v1/recording_settings"
	"github.com/timworks/twilio-sdk-go/service/video/v1/room"
	"github.com/timworks/twilio-sdk-go/service/video/v1/rooms"
	"github.com/timworks/twilio-sdk-go/session/credentials"
	"github.com/timworks/twilio-sdk-go/utils"
)

var _ = Describe("Video Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	videoSession := twilio.NewWithCredentials(creds).Video.V1

	Describe("Given the video room clients", func() {
		It("Then the room is created, fetched and updated", func() {
			roomsClient := videoSession.Rooms

			createResp, createErr := roomsClient.Create(&rooms.CreateRoomInput{
				Type: utils.String("go"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := roomsClient.Page(&rooms.RoomsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Rooms)).Should(BeNumerically(">=", 1))

			paginator := roomsClient.NewRoomsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Rooms)).Should(BeNumerically(">=", 1))

			roomClient := videoSession.Room(createResp.Sid)

			fetchResp, fetchErr := roomClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := roomClient.Update(&room.UpdateRoomInput{
				Status: "completed",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the video composition hook clients", func() {
		It("Then the composition hook is created, fetched, updated and deleted", func() {
			compositionHooksClient := videoSession.CompositionHooks

			friendlyName := uuid.New().String()

			createResp, createErr := compositionHooksClient.Create(&composition_hooks.CreateCompositionHookInput{
				FriendlyName: friendlyName,
				AudioSources: &[]string{"*"},
				Format:       utils.String("mp4"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := compositionHooksClient.Page(&composition_hooks.CompositionHooksPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.CompositionHooks)).Should(BeNumerically(">=", 1))

			paginator := compositionHooksClient.NewCompositionHooksPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.CompositionHooks)).Should(BeNumerically(">=", 1))

			compositionHookClient := videoSession.CompositionHook(createResp.Sid)

			fetchResp, fetchErr := compositionHookClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := compositionHookClient.Update(&composition_hook.UpdateCompositionHookInput{
				FriendlyName: friendlyName,
				AudioSources: &[]string{"*"},
				Format:       utils.String("mp4"),
				Enabled:      utils.Bool(false),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := compositionHookClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the video recording settings clients", func() {
		It("Then the recording settings is created, fetched and updated", func() {
			recordingSettingsClient := videoSession.RecordingSettings()

			fetchResp, fetchErr := recordingSettingsClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := recordingSettingsClient.Update(&recording_settings.UpdateRecordingSettingsInput{
				FriendlyName: "Basic Recording Setting",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the video composition settings clients", func() {
		It("Then the composition settings is created, fetched and updated", func() {
			compositionSettingsClient := videoSession.CompositionSettings()

			fetchResp, fetchErr := compositionSettingsClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := compositionSettingsClient.Update(&composition_settings.UpdateCompositionSettingsInput{
				FriendlyName: "Basic Composition Setting",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	// TODO Add Recording tests
	// TODO Add Room Recording tests
	// TODO Add Composition tests
	// TODO Add Participant tests
	// TODO Add Published Track tests
	// TODO Add Subscribed Track tests
	// TODO Add Subscribed Rules tests
})
