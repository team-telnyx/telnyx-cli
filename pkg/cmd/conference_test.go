// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestConferencesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "conferences", "create",
			"--api-key", "string",
			"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--name", "Business",
			"--beep-enabled", "always",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--comfort-noise=false",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--duration-minutes", "5",
			"--hold-audio-url", "http://www.example.com/audio.wav",
			"--hold-media-name", "my_media_uploaded_to_media_storage_api",
			"--max-participants", "250",
			"--region", "US",
			"--start-conference-on-create=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_id: v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"name: Business\n" +
			"beep_enabled: always\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"comfort_noise: false\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"duration_minutes: 5\n" +
			"hold_audio_url: http://www.example.com/audio.wav\n" +
			"hold_media_name: my_media_uploaded_to_media_storage_api\n" +
			"max_participants: 250\n" +
			"region: US\n" +
			"start_conference_on_create: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "conferences", "create",
			"--api-key", "string",
		)
	})
}

func TestConferencesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "conferences", "retrieve",
			"--api-key", "string",
			"--id", "id",
			"--region", "Australia",
		)
	})
}

func TestConferencesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "conferences", "list",
			"--api-key", "string",
			"--filter", "{application_name: {contains: contains}, application_session_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, connection_id: connection_id, failed: false, from: '+12025550142', leg_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, name: name, occurred_at: {eq: '2019-03-29T11:10:00Z', gt: '2019-03-29T11:10:00Z', gte: '2019-03-29T11:10:00Z', lt: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}, outbound.outbound_voice_profile_id: '1293384261075731499', product: texml, status: init, to: '+12025550142', type: webhook}",
			"--page-number", "0",
			"--page-size", "0",
			"--region", "Australia",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(conferencesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "conferences", "list",
			"--api-key", "string",
			"--filter.application-name", "{contains: contains}",
			"--filter.application-session-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter.connection-id", "connection_id",
			"--filter.failed=false",
			"--filter.from", "+12025550142",
			"--filter.leg-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter.name", "name",
			"--filter.occurred-at", "{eq: '2019-03-29T11:10:00Z', gt: '2019-03-29T11:10:00Z', gte: '2019-03-29T11:10:00Z', lt: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}",
			"--filter.outbound-outbound-voice-profile-id", "1293384261075731499",
			"--filter.product", "texml",
			"--filter.status", "init",
			"--filter.to", "+12025550142",
			"--filter.type", "webhook",
			"--page-number", "0",
			"--page-size", "0",
			"--region", "Australia",
		)
	})
}

func TestConferencesListParticipants(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "conferences", "list-participants",
			"--api-key", "string",
			"--conference-id", "conference_id",
			"--filter", "{muted: true, on_hold: true, whispering: true}",
			"--page-number", "0",
			"--page-size", "0",
			"--region", "Australia",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(conferencesListParticipants)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "conferences", "list-participants",
			"--api-key", "string",
			"--conference-id", "conference_id",
			"--filter.muted=true",
			"--filter.on-hold=true",
			"--filter.whispering=true",
			"--page-number", "0",
			"--page-size", "0",
			"--region", "Australia",
		)
	})
}

func TestConferencesRetrieveParticipant(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "conferences", "retrieve-participant",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--participant-id", "participant_id",
		)
	})
}

func TestConferencesUpdateParticipant(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "conferences", "update-participant",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--participant-id", "participant_id",
			"--beep-enabled", "never",
			"--end-conference-on-exit=true",
			"--soft-end-conference-on-exit=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"beep_enabled: never\n" +
			"end_conference_on_exit: true\n" +
			"soft_end_conference_on_exit: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "conferences", "update-participant",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--participant-id", "participant_id",
		)
	})
}
