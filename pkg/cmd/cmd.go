// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/team-telnyx/telnyx-cli/internal/autocomplete"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command *cli.Command
)

func init() {
	Command = &cli.Command{
		Name:    "telnyx",
		Usage:   "CLI for the telnyx API",
		Suggest: true,
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "legacy:reporting:batch-detail-records:messaging",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&legacyReportingBatchDetailRecordsMessagingCreate,
					&legacyReportingBatchDetailRecordsMessagingRetrieve,
					&legacyReportingBatchDetailRecordsMessagingList,
					&legacyReportingBatchDetailRecordsMessagingDelete,
				},
			},
			{
				Name:     "legacy:reporting:batch-detail-records:speech-to-text",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&legacyReportingBatchDetailRecordsSpeechToTextCreate,
					&legacyReportingBatchDetailRecordsSpeechToTextRetrieve,
					&legacyReportingBatchDetailRecordsSpeechToTextList,
					&legacyReportingBatchDetailRecordsSpeechToTextDelete,
				},
			},
			{
				Name:     "legacy:reporting:batch-detail-records:voice",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&legacyReportingBatchDetailRecordsVoiceCreate,
					&legacyReportingBatchDetailRecordsVoiceRetrieve,
					&legacyReportingBatchDetailRecordsVoiceList,
					&legacyReportingBatchDetailRecordsVoiceDelete,
					&legacyReportingBatchDetailRecordsVoiceRetrieveFields,
				},
			},
			{
				Name:     "legacy:reporting:usage-reports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&legacyReportingUsageReportsRetrieveSpeechToText,
				},
			},
			{
				Name:     "legacy:reporting:usage-reports:messaging",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&legacyReportingUsageReportsMessagingCreate,
					&legacyReportingUsageReportsMessagingRetrieve,
					&legacyReportingUsageReportsMessagingList,
					&legacyReportingUsageReportsMessagingDelete,
				},
			},
			{
				Name:     "legacy:reporting:usage-reports:number-lookup",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&legacyReportingUsageReportsNumberLookupCreate,
					&legacyReportingUsageReportsNumberLookupRetrieve,
					&legacyReportingUsageReportsNumberLookupList,
					&legacyReportingUsageReportsNumberLookupDelete,
				},
			},
			{
				Name:     "legacy:reporting:usage-reports:voice",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&legacyReportingUsageReportsVoiceCreate,
					&legacyReportingUsageReportsVoiceRetrieve,
					&legacyReportingUsageReportsVoiceList,
					&legacyReportingUsageReportsVoiceDelete,
				},
			},
			{
				Name:     "oauth",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&oauthRetrieve,
					&oauthGrants,
					&oauthIntrospect,
					&oauthRegister,
					&oauthRetrieveAuthorize,
					&oauthRetrieveJwks,
					&oauthToken,
				},
			},
			{
				Name:     "oauth-clients",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&oauthClientsCreate,
					&oauthClientsRetrieve,
					&oauthClientsUpdate,
					&oauthClientsList,
					&oauthClientsDelete,
				},
			},
			{
				Name:     "oauth-grants",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&oauthGrantsRetrieve,
					&oauthGrantsList,
					&oauthGrantsDelete,
				},
			},
			{
				Name:     "access-ip-address",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accessIPAddressCreate,
					&accessIPAddressRetrieve,
					&accessIPAddressList,
					&accessIPAddressDelete,
				},
			},
			{
				Name:     "access-ip-ranges",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accessIPRangesCreate,
					&accessIPRangesList,
					&accessIPRangesDelete,
				},
			},
			{
				Name:     "actions:purchase",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&actionsPurchaseCreate,
				},
			},
			{
				Name:     "actions:register",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&actionsRegisterCreate,
				},
			},
			{
				Name:     "addresses",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&addressesCreate,
					&addressesRetrieve,
					&addressesList,
					&addressesDelete,
				},
			},
			{
				Name:     "addresses:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&addressesActionsAcceptSuggestions,
					&addressesActionsValidate,
				},
			},
			{
				Name:     "advanced-orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&advancedOrdersCreate,
					&advancedOrdersRetrieve,
					&advancedOrdersList,
					&advancedOrdersUpdateRequirementGroup,
				},
			},
			{
				Name:     "ai",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiRetrieveModels,
					&aiSummarize,
				},
			},
			{
				Name:     "ai:assistants",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAssistantsCreate,
					&aiAssistantsRetrieve,
					&aiAssistantsUpdate,
					&aiAssistantsList,
					&aiAssistantsDelete,
					&aiAssistantsChat,
					&aiAssistantsClone,
					&aiAssistantsGetTexml,
					&aiAssistantsImports,
					&aiAssistantsSendSMS,
				},
			},
			{
				Name:     "ai:assistants:tests",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAssistantsTestsCreate,
					&aiAssistantsTestsRetrieve,
					&aiAssistantsTestsUpdate,
					&aiAssistantsTestsList,
					&aiAssistantsTestsDelete,
				},
			},
			{
				Name:     "ai:assistants:tests:test-suites",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAssistantsTestsTestSuitesList,
				},
			},
			{
				Name:     "ai:assistants:tests:test-suites:runs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAssistantsTestsTestSuitesRunsList,
					&aiAssistantsTestsTestSuitesRunsTrigger,
				},
			},
			{
				Name:     "ai:assistants:tests:runs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAssistantsTestsRunsRetrieve,
					&aiAssistantsTestsRunsList,
					&aiAssistantsTestsRunsTrigger,
				},
			},
			{
				Name:     "ai:assistants:canary-deploys",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAssistantsCanaryDeploysCreate,
					&aiAssistantsCanaryDeploysRetrieve,
					&aiAssistantsCanaryDeploysUpdate,
					&aiAssistantsCanaryDeploysDelete,
				},
			},
			{
				Name:     "ai:assistants:scheduled-events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAssistantsScheduledEventsCreate,
					&aiAssistantsScheduledEventsRetrieve,
					&aiAssistantsScheduledEventsList,
					&aiAssistantsScheduledEventsDelete,
				},
			},
			{
				Name:     "ai:assistants:tools",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAssistantsToolsTest,
				},
			},
			{
				Name:     "ai:assistants:versions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAssistantsVersionsRetrieve,
					&aiAssistantsVersionsUpdate,
					&aiAssistantsVersionsList,
					&aiAssistantsVersionsDelete,
					&aiAssistantsVersionsPromote,
				},
			},
			{
				Name:     "ai:audio",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiAudioTranscribe,
				},
			},
			{
				Name:     "ai:chat",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiChatCreateCompletion,
				},
			},
			{
				Name:     "ai:clusters",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiClustersRetrieve,
					&aiClustersList,
					&aiClustersDelete,
					&aiClustersCompute,
				},
			},
			{
				Name:     "ai:conversations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiConversationsCreate,
					&aiConversationsRetrieve,
					&aiConversationsUpdate,
					&aiConversationsList,
					&aiConversationsDelete,
					&aiConversationsAddMessage,
					&aiConversationsRetrieveConversationsInsights,
				},
			},
			{
				Name:     "ai:conversations:insight-groups",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiConversationsInsightGroupsRetrieve,
					&aiConversationsInsightGroupsUpdate,
					&aiConversationsInsightGroupsDelete,
					&aiConversationsInsightGroupsInsightGroups,
					&aiConversationsInsightGroupsRetrieveInsightGroups,
				},
			},
			{
				Name:     "ai:conversations:insight-groups:insights",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiConversationsInsightGroupsInsightsAssign,
					&aiConversationsInsightGroupsInsightsDeleteUnassign,
				},
			},
			{
				Name:     "ai:conversations:insights",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiConversationsInsightsCreate,
					&aiConversationsInsightsRetrieve,
					&aiConversationsInsightsUpdate,
					&aiConversationsInsightsList,
					&aiConversationsInsightsDelete,
				},
			},
			{
				Name:     "ai:conversations:messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiConversationsMessagesList,
				},
			},
			{
				Name:     "ai:embeddings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiEmbeddingsCreate,
					&aiEmbeddingsRetrieve,
					&aiEmbeddingsList,
					&aiEmbeddingsSimilaritySearch,
					&aiEmbeddingsURL,
				},
			},
			{
				Name:     "ai:embeddings:buckets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiEmbeddingsBucketsRetrieve,
					&aiEmbeddingsBucketsList,
					&aiEmbeddingsBucketsDelete,
				},
			},
			{
				Name:     "ai:fine-tuning:jobs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiFineTuningJobsCreate,
					&aiFineTuningJobsRetrieve,
					&aiFineTuningJobsList,
					&aiFineTuningJobsCancel,
				},
			},
			{
				Name:     "ai:integrations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiIntegrationsRetrieve,
					&aiIntegrationsList,
				},
			},
			{
				Name:     "ai:integrations:connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiIntegrationsConnectionsRetrieve,
					&aiIntegrationsConnectionsList,
					&aiIntegrationsConnectionsDelete,
				},
			},
			{
				Name:     "ai:mcp-servers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiMcpServersCreate,
					&aiMcpServersRetrieve,
					&aiMcpServersUpdate,
					&aiMcpServersList,
					&aiMcpServersDelete,
				},
			},
			{
				Name:     "ai:missions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiMissionsCreate,
					&aiMissionsRetrieve,
					&aiMissionsList,
					&aiMissionsCloneMission,
					&aiMissionsDeleteMission,
					&aiMissionsListEvents,
					&aiMissionsUpdateMission,
				},
			},
			{
				Name:     "ai:missions:runs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiMissionsRunsCreate,
					&aiMissionsRunsRetrieve,
					&aiMissionsRunsUpdate,
					&aiMissionsRunsList,
					&aiMissionsRunsCancelRun,
					&aiMissionsRunsListRuns,
					&aiMissionsRunsPauseRun,
					&aiMissionsRunsResumeRun,
				},
			},
			{
				Name:     "ai:missions:runs:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiMissionsRunsEventsList,
					&aiMissionsRunsEventsGetEventDetails,
					&aiMissionsRunsEventsLog,
				},
			},
			{
				Name:     "ai:missions:runs:plan",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiMissionsRunsPlanCreate,
					&aiMissionsRunsPlanRetrieve,
					&aiMissionsRunsPlanAddStepsToPlan,
					&aiMissionsRunsPlanGetStepDetails,
					&aiMissionsRunsPlanUpdateStep,
				},
			},
			{
				Name:     "ai:missions:runs:telnyx-agents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiMissionsRunsTelnyxAgentsList,
					&aiMissionsRunsTelnyxAgentsLink,
					&aiMissionsRunsTelnyxAgentsUnlink,
				},
			},
			{
				Name:     "ai:missions:knowledge-bases",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiMissionsKnowledgeBasesCreateKnowledgeBase,
					&aiMissionsKnowledgeBasesDeleteKnowledgeBase,
					&aiMissionsKnowledgeBasesGetKnowledgeBase,
					&aiMissionsKnowledgeBasesListKnowledgeBases,
					&aiMissionsKnowledgeBasesUpdateKnowledgeBase,
				},
			},
			{
				Name:     "ai:missions:mcp-servers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiMissionsMcpServersCreateMcpServer,
					&aiMissionsMcpServersDeleteMcpServer,
					&aiMissionsMcpServersGetMcpServer,
					&aiMissionsMcpServersListMcpServers,
					&aiMissionsMcpServersUpdateMcpServer,
				},
			},
			{
				Name:     "ai:missions:tools",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiMissionsToolsCreateTool,
					&aiMissionsToolsDeleteTool,
					&aiMissionsToolsGetTool,
					&aiMissionsToolsListTools,
					&aiMissionsToolsUpdateTool,
				},
			},
			{
				Name:     "ai:openai:embeddings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&aiOpenAIEmbeddingsCreateEmbeddings,
					&aiOpenAIEmbeddingsListEmbeddingModels,
				},
			},
			{
				Name:     "audit-events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&auditEventsList,
				},
			},
			{
				Name:     "authentication-providers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&authenticationProvidersCreate,
					&authenticationProvidersRetrieve,
					&authenticationProvidersUpdate,
					&authenticationProvidersList,
					&authenticationProvidersDelete,
				},
			},
			{
				Name:     "available-phone-number-blocks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&availablePhoneNumberBlocksList,
				},
			},
			{
				Name:     "available-phone-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&availablePhoneNumbersList,
				},
			},
			{
				Name:     "balance",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&balanceRetrieve,
				},
			},
			{
				Name:     "billing-groups",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&billingGroupsCreate,
					&billingGroupsRetrieve,
					&billingGroupsUpdate,
					&billingGroupsList,
					&billingGroupsDelete,
				},
			},
			{
				Name:     "bulk-sim-card-actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&bulkSimCardActionsRetrieve,
					&bulkSimCardActionsList,
				},
			},
			{
				Name:     "bundle-pricing:billing-bundles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&bundlePricingBillingBundlesRetrieve,
					&bundlePricingBillingBundlesList,
				},
			},
			{
				Name:     "bundle-pricing:user-bundles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&bundlePricingUserBundlesCreate,
					&bundlePricingUserBundlesRetrieve,
					&bundlePricingUserBundlesList,
					&bundlePricingUserBundlesDeactivate,
					&bundlePricingUserBundlesListResources,
					&bundlePricingUserBundlesListUnused,
				},
			},
			{
				Name:     "call-control-applications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&callControlApplicationsCreate,
					&callControlApplicationsRetrieve,
					&callControlApplicationsUpdate,
					&callControlApplicationsList,
					&callControlApplicationsDelete,
				},
			},
			{
				Name:     "call-events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&callEventsList,
				},
			},
			{
				Name:     "calls",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&callsDial,
					&callsRetrieveStatus,
				},
			},
			{
				Name:     "calls:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&callsActionsAddAIAssistantMessages,
					&callsActionsAnswer,
					&callsActionsBridge,
					&callsActionsEnqueue,
					&callsActionsGather,
					&callsActionsGatherUsingAI,
					&callsActionsGatherUsingAudio,
					&callsActionsGatherUsingSpeak,
					&callsActionsHangup,
					&callsActionsLeaveQueue,
					&callsActionsPauseRecording,
					&callsActionsRefer,
					&callsActionsReject,
					&callsActionsResumeRecording,
					&callsActionsSendDtmf,
					&callsActionsSendSipInfo,
					&callsActionsSpeak,
					&callsActionsStartAIAssistant,
					&callsActionsStartForking,
					&callsActionsStartNoiseSuppression,
					&callsActionsStartPlayback,
					&callsActionsStartRecording,
					&callsActionsStartSiprec,
					&callsActionsStartStreaming,
					&callsActionsStartTranscription,
					&callsActionsStopAIAssistant,
					&callsActionsStopForking,
					&callsActionsStopGather,
					&callsActionsStopNoiseSuppression,
					&callsActionsStopPlayback,
					&callsActionsStopRecording,
					&callsActionsStopSiprec,
					&callsActionsStopStreaming,
					&callsActionsStopTranscription,
					&callsActionsSwitchSupervisorRole,
					&callsActionsTransfer,
					&callsActionsUpdateClientState,
				},
			},
			{
				Name:     "channel-zones",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&channelZonesUpdate,
					&channelZonesList,
				},
			},
			{
				Name:     "charges-breakdown",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&chargesBreakdownRetrieve,
				},
			},
			{
				Name:     "charges-summary",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&chargesSummaryRetrieve,
				},
			},
			{
				Name:     "comments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&commentsCreate,
					&commentsRetrieve,
					&commentsList,
					&commentsMarkAsRead,
				},
			},
			{
				Name:     "conferences",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&conferencesCreate,
					&conferencesRetrieve,
					&conferencesList,
					&conferencesListParticipants,
				},
			},
			{
				Name:     "conferences:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&conferencesActionsUpdate,
					&conferencesActionsHold,
					&conferencesActionsJoin,
					&conferencesActionsLeave,
					&conferencesActionsMute,
					&conferencesActionsPlay,
					&conferencesActionsRecordPause,
					&conferencesActionsRecordResume,
					&conferencesActionsRecordStart,
					&conferencesActionsRecordStop,
					&conferencesActionsSpeak,
					&conferencesActionsStop,
					&conferencesActionsUnhold,
					&conferencesActionsUnmute,
				},
			},
			{
				Name:     "connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&connectionsRetrieve,
					&connectionsList,
					&connectionsListActiveCalls,
				},
			},
			{
				Name:     "country-coverage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&countryCoverageRetrieve,
					&countryCoverageRetrieveCountry,
				},
			},
			{
				Name:     "credential-connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&credentialConnectionsCreate,
					&credentialConnectionsRetrieve,
					&credentialConnectionsUpdate,
					&credentialConnectionsList,
					&credentialConnectionsDelete,
				},
			},
			{
				Name:     "credential-connections:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&credentialConnectionsActionsCheckRegistrationStatus,
				},
			},
			{
				Name:     "custom-storage-credentials",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&customStorageCredentialsCreate,
					&customStorageCredentialsRetrieve,
					&customStorageCredentialsUpdate,
					&customStorageCredentialsDelete,
				},
			},
			{
				Name:     "customer-service-records",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&customerServiceRecordsCreate,
					&customerServiceRecordsRetrieve,
					&customerServiceRecordsList,
					&customerServiceRecordsVerifyPhoneNumberCoverage,
				},
			},
			{
				Name:     "detail-records",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&detailRecordsList,
				},
			},
			{
				Name:     "dialogflow-connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&dialogflowConnectionsCreate,
					&dialogflowConnectionsRetrieve,
					&dialogflowConnectionsUpdate,
					&dialogflowConnectionsDelete,
				},
			},
			{
				Name:     "document-links",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&documentLinksList,
				},
			},
			{
				Name:     "documents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&documentsRetrieve,
					&documentsUpdate,
					&documentsList,
					&documentsDelete,
					&documentsGenerateDownloadLink,
					&documentsUpload,
					&documentsUploadJson,
				},
			},
			{
				Name:     "dynamic-emergency-addresses",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&dynamicEmergencyAddressesCreate,
					&dynamicEmergencyAddressesRetrieve,
					&dynamicEmergencyAddressesList,
					&dynamicEmergencyAddressesDelete,
				},
			},
			{
				Name:     "dynamic-emergency-endpoints",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&dynamicEmergencyEndpointsCreate,
					&dynamicEmergencyEndpointsRetrieve,
					&dynamicEmergencyEndpointsList,
					&dynamicEmergencyEndpointsDelete,
				},
			},
			{
				Name:     "external-connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&externalConnectionsCreate,
					&externalConnectionsRetrieve,
					&externalConnectionsUpdate,
					&externalConnectionsList,
					&externalConnectionsDelete,
					&externalConnectionsUpdateLocation,
				},
			},
			{
				Name:     "external-connections:log-messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&externalConnectionsLogMessagesRetrieve,
					&externalConnectionsLogMessagesList,
					&externalConnectionsLogMessagesDismiss,
				},
			},
			{
				Name:     "external-connections:civic-addresses",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&externalConnectionsCivicAddressesRetrieve,
					&externalConnectionsCivicAddressesList,
				},
			},
			{
				Name:     "external-connections:phone-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&externalConnectionsPhoneNumbersRetrieve,
					&externalConnectionsPhoneNumbersUpdate,
					&externalConnectionsPhoneNumbersList,
				},
			},
			{
				Name:     "external-connections:releases",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&externalConnectionsReleasesRetrieve,
					&externalConnectionsReleasesList,
				},
			},
			{
				Name:     "external-connections:uploads",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&externalConnectionsUploadsCreate,
					&externalConnectionsUploadsRetrieve,
					&externalConnectionsUploadsList,
					&externalConnectionsUploadsPendingCount,
					&externalConnectionsUploadsRefreshStatus,
					&externalConnectionsUploadsRetry,
				},
			},
			{
				Name:     "fax-applications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&faxApplicationsCreate,
					&faxApplicationsRetrieve,
					&faxApplicationsUpdate,
					&faxApplicationsList,
					&faxApplicationsDelete,
				},
			},
			{
				Name:     "faxes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&faxesCreate,
					&faxesRetrieve,
					&faxesList,
					&faxesDelete,
				},
			},
			{
				Name:     "faxes:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&faxesActionsCancel,
					&faxesActionsRefresh,
				},
			},
			{
				Name:     "fqdn-connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&fqdnConnectionsCreate,
					&fqdnConnectionsRetrieve,
					&fqdnConnectionsUpdate,
					&fqdnConnectionsList,
					&fqdnConnectionsDelete,
				},
			},
			{
				Name:     "fqdns",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&fqdnsCreate,
					&fqdnsRetrieve,
					&fqdnsUpdate,
					&fqdnsList,
					&fqdnsDelete,
				},
			},
			{
				Name:     "global-ip-allowed-ports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPAllowedPortsList,
				},
			},
			{
				Name:     "global-ip-assignment-health",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPAssignmentHealthRetrieve,
				},
			},
			{
				Name:     "global-ip-assignments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPAssignmentsCreate,
					&globalIPAssignmentsRetrieve,
					&globalIPAssignmentsUpdate,
					&globalIPAssignmentsList,
					&globalIPAssignmentsDelete,
				},
			},
			{
				Name:     "global-ip-assignments-usage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPAssignmentsUsageRetrieve,
				},
			},
			{
				Name:     "global-ip-health-check-types",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPHealthCheckTypesList,
				},
			},
			{
				Name:     "global-ip-health-checks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPHealthChecksCreate,
					&globalIPHealthChecksRetrieve,
					&globalIPHealthChecksList,
					&globalIPHealthChecksDelete,
				},
			},
			{
				Name:     "global-ip-latency",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPLatencyRetrieve,
				},
			},
			{
				Name:     "global-ip-protocols",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPProtocolsList,
				},
			},
			{
				Name:     "global-ip-usage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPUsageRetrieve,
				},
			},
			{
				Name:     "global-ips",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalIPsCreate,
					&globalIPsRetrieve,
					&globalIPsList,
					&globalIPsDelete,
				},
			},
			{
				Name:     "inbound-channels",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inboundChannelsUpdate,
					&inboundChannelsList,
				},
			},
			{
				Name:     "integration-secrets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&integrationSecretsCreate,
					&integrationSecretsList,
					&integrationSecretsDelete,
				},
			},
			{
				Name:     "inventory-coverage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inventoryCoverageList,
				},
			},
			{
				Name:     "invoices",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&invoicesRetrieve,
					&invoicesList,
				},
			},
			{
				Name:     "ip-connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&ipConnectionsCreate,
					&ipConnectionsRetrieve,
					&ipConnectionsUpdate,
					&ipConnectionsList,
					&ipConnectionsDelete,
				},
			},
			{
				Name:     "ips",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&ipsCreate,
					&ipsRetrieve,
					&ipsUpdate,
					&ipsList,
					&ipsDelete,
				},
			},
			{
				Name:     "ledger-billing-group-reports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&ledgerBillingGroupReportsCreate,
					&ledgerBillingGroupReportsRetrieve,
				},
			},
			{
				Name:     "list",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&listRetrieveAll,
					&listRetrieveByZone,
				},
			},
			{
				Name:     "managed-accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&managedAccountsCreate,
					&managedAccountsRetrieve,
					&managedAccountsUpdate,
					&managedAccountsList,
					&managedAccountsGetAllocatableGlobalOutboundChannels,
					&managedAccountsUpdateGlobalChannelLimit,
				},
			},
			{
				Name:     "managed-accounts:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&managedAccountsActionsDisable,
					&managedAccountsActionsEnable,
				},
			},
			{
				Name:     "media",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mediaRetrieve,
					&mediaUpdate,
					&mediaList,
					&mediaDelete,
					&mediaUpload,
				},
			},
			{
				Name:     "messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagesRetrieve,
					&messagesCancelScheduled,
					&messagesSchedule,
					&messagesSend,
					&messagesSendGroupMms,
					&messagesSendLongCode,
					&messagesSendNumberPool,
					&messagesSendShortCode,
					&messagesSendWhatsapp,
				},
			},
			{
				Name:     "messages:rcs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagesRcsGenerateDeeplink,
					&messagesRcsSend,
				},
			},
			{
				Name:     "messaging:rcs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingRcsInviteTestNumber,
					&messagingRcsListBulkCapabilities,
					&messagingRcsRetrieveCapabilities,
				},
			},
			{
				Name:     "messaging:rcs:agents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingRcsAgentsRetrieve,
					&messagingRcsAgentsUpdate,
					&messagingRcsAgentsList,
				},
			},
			{
				Name:     "messaging-hosted-number-orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingHostedNumberOrdersCreate,
					&messagingHostedNumberOrdersRetrieve,
					&messagingHostedNumberOrdersList,
					&messagingHostedNumberOrdersDelete,
					&messagingHostedNumberOrdersCheckEligibility,
					&messagingHostedNumberOrdersCreateVerificationCodes,
					&messagingHostedNumberOrdersValidateCodes,
				},
			},
			{
				Name:     "messaging-hosted-number-orders:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingHostedNumberOrdersActionsUploadFile,
				},
			},
			{
				Name:     "messaging-hosted-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingHostedNumbersDelete,
				},
			},
			{
				Name:     "messaging-numbers-bulk-updates",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingNumbersBulkUpdatesCreate,
					&messagingNumbersBulkUpdatesRetrieve,
				},
			},
			{
				Name:     "messaging-optouts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingOptoutsList,
				},
			},
			{
				Name:     "messaging-profiles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingProfilesCreate,
					&messagingProfilesRetrieve,
					&messagingProfilesUpdate,
					&messagingProfilesList,
					&messagingProfilesDelete,
					&messagingProfilesListPhoneNumbers,
					&messagingProfilesListShortCodes,
				},
			},
			{
				Name:     "messaging-profiles:autoresp-configs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingProfilesAutorespConfigsCreate,
					&messagingProfilesAutorespConfigsRetrieve,
					&messagingProfilesAutorespConfigsUpdate,
					&messagingProfilesAutorespConfigsList,
					&messagingProfilesAutorespConfigsDelete,
				},
			},
			{
				Name:     "messaging-tollfree:verification:requests",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingTollfreeVerificationRequestsCreate,
					&messagingTollfreeVerificationRequestsRetrieve,
					&messagingTollfreeVerificationRequestsUpdate,
					&messagingTollfreeVerificationRequestsList,
					&messagingTollfreeVerificationRequestsDelete,
				},
			},
			{
				Name:     "messaging-url-domains",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagingURLDomainsList,
				},
			},
			{
				Name:     "mobile-network-operators",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mobileNetworkOperatorsList,
				},
			},
			{
				Name:     "mobile-push-credentials",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mobilePushCredentialsCreate,
					&mobilePushCredentialsRetrieve,
					&mobilePushCredentialsList,
					&mobilePushCredentialsDelete,
				},
			},
			{
				Name:     "network-coverage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&networkCoverageList,
				},
			},
			{
				Name:     "networks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&networksCreate,
					&networksRetrieve,
					&networksUpdate,
					&networksList,
					&networksDelete,
					&networksListInterfaces,
				},
			},
			{
				Name:     "networks:default-gateway",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&networksDefaultGatewayCreate,
					&networksDefaultGatewayRetrieve,
					&networksDefaultGatewayDelete,
				},
			},
			{
				Name:     "notification-channels",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationChannelsCreate,
					&notificationChannelsRetrieve,
					&notificationChannelsUpdate,
					&notificationChannelsList,
					&notificationChannelsDelete,
				},
			},
			{
				Name:     "notification-event-conditions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationEventConditionsList,
				},
			},
			{
				Name:     "notification-events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationEventsList,
				},
			},
			{
				Name:     "notification-profiles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationProfilesCreate,
					&notificationProfilesRetrieve,
					&notificationProfilesUpdate,
					&notificationProfilesList,
					&notificationProfilesDelete,
				},
			},
			{
				Name:     "notification-settings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationSettingsCreate,
					&notificationSettingsRetrieve,
					&notificationSettingsList,
					&notificationSettingsDelete,
				},
			},
			{
				Name:     "number-block-orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&numberBlockOrdersCreate,
					&numberBlockOrdersRetrieve,
					&numberBlockOrdersList,
				},
			},
			{
				Name:     "number-lookup",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&numberLookupRetrieve,
				},
			},
			{
				Name:     "number-order-phone-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&numberOrderPhoneNumbersRetrieve,
					&numberOrderPhoneNumbersList,
					&numberOrderPhoneNumbersUpdateRequirementGroup,
					&numberOrderPhoneNumbersUpdateRequirements,
				},
			},
			{
				Name:     "number-orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&numberOrdersCreate,
					&numberOrdersRetrieve,
					&numberOrdersUpdate,
					&numberOrdersList,
				},
			},
			{
				Name:     "number-reservations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&numberReservationsCreate,
					&numberReservationsRetrieve,
					&numberReservationsList,
				},
			},
			{
				Name:     "number-reservations:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&numberReservationsActionsExtend,
				},
			},
			{
				Name:     "numbers-features",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&numbersFeaturesCreate,
				},
			},
			{
				Name:     "operator-connect:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&operatorConnectActionsRefresh,
				},
			},
			{
				Name:     "ota-updates",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&otaUpdatesRetrieve,
					&otaUpdatesList,
				},
			},
			{
				Name:     "outbound-voice-profiles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&outboundVoiceProfilesCreate,
					&outboundVoiceProfilesRetrieve,
					&outboundVoiceProfilesUpdate,
					&outboundVoiceProfilesList,
					&outboundVoiceProfilesDelete,
				},
			},
			{
				Name:     "payment:auto-recharge-prefs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&paymentAutoRechargePrefsUpdate,
					&paymentAutoRechargePrefsList,
				},
			},
			{
				Name:     "phone-number-blocks:jobs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumberBlocksJobsRetrieve,
					&phoneNumberBlocksJobsList,
					&phoneNumberBlocksJobsDeletePhoneNumberBlock,
				},
			},
			{
				Name:     "phone-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumbersRetrieve,
					&phoneNumbersUpdate,
					&phoneNumbersList,
					&phoneNumbersDelete,
					&phoneNumbersSlimList,
				},
			},
			{
				Name:     "phone-numbers:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumbersActionsChangeBundleStatus,
					&phoneNumbersActionsEnableEmergency,
					&phoneNumbersActionsVerifyOwnership,
				},
			},
			{
				Name:     "phone-numbers:csv-downloads",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumbersCsvDownloadsCreate,
					&phoneNumbersCsvDownloadsRetrieve,
					&phoneNumbersCsvDownloadsList,
				},
			},
			{
				Name:     "phone-numbers:jobs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumbersJobsRetrieve,
					&phoneNumbersJobsList,
					&phoneNumbersJobsDeleteBatch,
					&phoneNumbersJobsUpdateBatch,
					&phoneNumbersJobsUpdateEmergencySettingsBatch,
				},
			},
			{
				Name:     "phone-numbers:messaging",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumbersMessagingRetrieve,
					&phoneNumbersMessagingUpdate,
					&phoneNumbersMessagingList,
				},
			},
			{
				Name:     "phone-numbers:voice",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumbersVoiceRetrieve,
					&phoneNumbersVoiceUpdate,
					&phoneNumbersVoiceList,
				},
			},
			{
				Name:     "phone-numbers:voicemail",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumbersVoicemailCreate,
					&phoneNumbersVoicemailRetrieve,
					&phoneNumbersVoicemailUpdate,
				},
			},
			{
				Name:     "phone-numbers-regulatory-requirements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumbersRegulatoryRequirementsRetrieve,
				},
			},
			{
				Name:     "portability-checks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portabilityChecksRun,
				},
			},
			{
				Name:     "porting",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingListUkCarriers,
				},
			},
			{
				Name:     "porting:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingEventsRetrieve,
					&portingEventsList,
					&portingEventsRepublish,
				},
			},
			{
				Name:     "porting:reports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingReportsCreate,
					&portingReportsRetrieve,
					&portingReportsList,
				},
			},
			{
				Name:     "porting:loa-configurations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingLoaConfigurationsCreate,
					&portingLoaConfigurationsRetrieve,
					&portingLoaConfigurationsUpdate,
					&portingLoaConfigurationsList,
					&portingLoaConfigurationsDelete,
				},
			},
			{
				Name:     "porting-orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersCreate,
					&portingOrdersRetrieve,
					&portingOrdersUpdate,
					&portingOrdersList,
					&portingOrdersDelete,
					&portingOrdersRetrieveAllowedFocWindows,
					&portingOrdersRetrieveExceptionTypes,
					&portingOrdersRetrieveRequirements,
					&portingOrdersRetrieveSubRequest,
				},
			},
			{
				Name:     "porting-orders:phone-number-configurations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersPhoneNumberConfigurationsCreate,
					&portingOrdersPhoneNumberConfigurationsList,
				},
			},
			{
				Name:     "porting-orders:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersActionsActivate,
					&portingOrdersActionsCancel,
					&portingOrdersActionsConfirm,
					&portingOrdersActionsShare,
				},
			},
			{
				Name:     "porting-orders:activation-jobs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersActivationJobsRetrieve,
					&portingOrdersActivationJobsUpdate,
					&portingOrdersActivationJobsList,
				},
			},
			{
				Name:     "porting-orders:additional-documents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersAdditionalDocumentsCreate,
					&portingOrdersAdditionalDocumentsList,
					&portingOrdersAdditionalDocumentsDelete,
				},
			},
			{
				Name:     "porting-orders:comments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersCommentsCreate,
					&portingOrdersCommentsList,
				},
			},
			{
				Name:     "porting-orders:verification-codes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersVerificationCodesList,
					&portingOrdersVerificationCodesSend,
					&portingOrdersVerificationCodesVerify,
				},
			},
			{
				Name:     "porting-orders:action-requirements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersActionRequirementsList,
					&portingOrdersActionRequirementsInitiate,
				},
			},
			{
				Name:     "porting-orders:associated-phone-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersAssociatedPhoneNumbersCreate,
					&portingOrdersAssociatedPhoneNumbersList,
					&portingOrdersAssociatedPhoneNumbersDelete,
				},
			},
			{
				Name:     "porting-orders:phone-number-blocks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersPhoneNumberBlocksCreate,
					&portingOrdersPhoneNumberBlocksList,
					&portingOrdersPhoneNumberBlocksDelete,
				},
			},
			{
				Name:     "porting-orders:phone-number-extensions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingOrdersPhoneNumberExtensionsCreate,
					&portingOrdersPhoneNumberExtensionsList,
					&portingOrdersPhoneNumberExtensionsDelete,
				},
			},
			{
				Name:     "porting-phone-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portingPhoneNumbersList,
				},
			},
			{
				Name:     "portouts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portoutsRetrieve,
					&portoutsList,
					&portoutsListRejectionCodes,
					&portoutsUpdateStatus,
				},
			},
			{
				Name:     "portouts:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portoutsEventsRetrieve,
					&portoutsEventsList,
					&portoutsEventsRepublish,
				},
			},
			{
				Name:     "portouts:reports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portoutsReportsCreate,
					&portoutsReportsRetrieve,
					&portoutsReportsList,
				},
			},
			{
				Name:     "portouts:comments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portoutsCommentsCreate,
					&portoutsCommentsList,
				},
			},
			{
				Name:     "portouts:supporting-documents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&portoutsSupportingDocumentsCreate,
					&portoutsSupportingDocumentsList,
				},
			},
			{
				Name:     "private-wireless-gateways",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&privateWirelessGatewaysCreate,
					&privateWirelessGatewaysRetrieve,
					&privateWirelessGatewaysList,
					&privateWirelessGatewaysDelete,
				},
			},
			{
				Name:     "public-internet-gateways",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&publicInternetGatewaysCreate,
					&publicInternetGatewaysRetrieve,
					&publicInternetGatewaysList,
					&publicInternetGatewaysDelete,
				},
			},
			{
				Name:     "queues",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&queuesRetrieve,
				},
			},
			{
				Name:     "queues:calls",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&queuesCallsRetrieve,
					&queuesCallsUpdate,
					&queuesCallsList,
					&queuesCallsRemove,
				},
			},
			{
				Name:     "recording-transcriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&recordingTranscriptionsRetrieve,
					&recordingTranscriptionsList,
					&recordingTranscriptionsDelete,
				},
			},
			{
				Name:     "recordings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&recordingsRetrieve,
					&recordingsList,
					&recordingsDelete,
				},
			},
			{
				Name:     "recordings:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&recordingsActionsDelete,
				},
			},
			{
				Name:     "regions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&regionsList,
				},
			},
			{
				Name:     "regulatory-requirements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&regulatoryRequirementsRetrieve,
				},
			},
			{
				Name:     "reports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&reportsListMdrs,
					&reportsListWdrs,
				},
			},
			{
				Name:     "reports:cdr-usage-reports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&reportsCdrUsageReportsFetchSync,
				},
			},
			{
				Name:     "reports:mdr-usage-reports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&reportsMdrUsageReportsCreate,
					&reportsMdrUsageReportsRetrieve,
					&reportsMdrUsageReportsList,
					&reportsMdrUsageReportsDelete,
					&reportsMdrUsageReportsFetchSync,
				},
			},
			{
				Name:     "requirement-groups",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&requirementGroupsCreate,
					&requirementGroupsRetrieve,
					&requirementGroupsUpdate,
					&requirementGroupsList,
					&requirementGroupsDelete,
					&requirementGroupsSubmitForApproval,
				},
			},
			{
				Name:     "requirement-types",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&requirementTypesRetrieve,
					&requirementTypesList,
				},
			},
			{
				Name:     "requirements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&requirementsRetrieve,
					&requirementsList,
				},
			},
			{
				Name:     "room-compositions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&roomCompositionsCreate,
					&roomCompositionsRetrieve,
					&roomCompositionsList,
					&roomCompositionsDelete,
				},
			},
			{
				Name:     "room-participants",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&roomParticipantsRetrieve,
					&roomParticipantsList,
				},
			},
			{
				Name:     "room-recordings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&roomRecordingsRetrieve,
					&roomRecordingsList,
					&roomRecordingsDelete,
					&roomRecordingsDeleteBulk,
				},
			},
			{
				Name:     "rooms",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&roomsCreate,
					&roomsRetrieve,
					&roomsUpdate,
					&roomsList,
					&roomsDelete,
				},
			},
			{
				Name:     "rooms:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&roomsActionsGenerateJoinClientToken,
					&roomsActionsRefreshClientToken,
				},
			},
			{
				Name:     "rooms:sessions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&roomsSessionsRetrieve,
					&roomsSessionsList0,
					&roomsSessionsList1,
					&roomsSessionsRetrieveParticipants,
				},
			},
			{
				Name:     "rooms:sessions:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&roomsSessionsActionsEnd,
					&roomsSessionsActionsKick,
					&roomsSessionsActionsMute,
					&roomsSessionsActionsUnmute,
				},
			},
			{
				Name:     "seti",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&setiRetrieveBlackBoxTestResults,
				},
			},
			{
				Name:     "short-codes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&shortCodesRetrieve,
					&shortCodesUpdate,
					&shortCodesList,
				},
			},
			{
				Name:     "sim-card-data-usage-notifications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simCardDataUsageNotificationsCreate,
					&simCardDataUsageNotificationsRetrieve,
					&simCardDataUsageNotificationsUpdate,
					&simCardDataUsageNotificationsList,
					&simCardDataUsageNotificationsDelete,
				},
			},
			{
				Name:     "sim-card-groups",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simCardGroupsCreate,
					&simCardGroupsRetrieve,
					&simCardGroupsUpdate,
					&simCardGroupsList,
					&simCardGroupsDelete,
				},
			},
			{
				Name:     "sim-card-groups:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simCardGroupsActionsRetrieve,
					&simCardGroupsActionsList,
					&simCardGroupsActionsRemovePrivateWirelessGateway,
					&simCardGroupsActionsRemoveWirelessBlocklist,
					&simCardGroupsActionsSetPrivateWirelessGateway,
					&simCardGroupsActionsSetWirelessBlocklist,
				},
			},
			{
				Name:     "sim-card-order-preview",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simCardOrderPreviewPreview,
				},
			},
			{
				Name:     "sim-card-orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simCardOrdersCreate,
					&simCardOrdersRetrieve,
					&simCardOrdersList,
				},
			},
			{
				Name:     "sim-cards",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simCardsRetrieve,
					&simCardsUpdate,
					&simCardsList,
					&simCardsDelete,
					&simCardsGetActivationCode,
					&simCardsGetDeviceDetails,
					&simCardsGetPublicIP,
					&simCardsListWirelessConnectivityLogs,
				},
			},
			{
				Name:     "sim-cards:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simCardsActionsRetrieve,
					&simCardsActionsList,
					&simCardsActionsBulkSetPublicIPs,
					&simCardsActionsDisable,
					&simCardsActionsEnable,
					&simCardsActionsRemovePublicIP,
					&simCardsActionsSetPublicIP,
					&simCardsActionsSetStandby,
					&simCardsActionsValidateRegistrationCodes,
				},
			},
			{
				Name:     "siprec-connectors",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&siprecConnectorsCreate,
					&siprecConnectorsRetrieve,
					&siprecConnectorsUpdate,
					&siprecConnectorsDelete,
				},
			},
			{
				Name:     "storage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&storageListMigrationSourceCoverage,
				},
			},
			{
				Name:     "storage:buckets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&storageBucketsCreatePresignedURL,
				},
			},
			{
				Name:     "storage:buckets:ssl-certificate",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&storageBucketsSslCertificateCreate,
					&storageBucketsSslCertificateRetrieve,
					&storageBucketsSslCertificateDelete,
				},
			},
			{
				Name:     "storage:buckets:usage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&storageBucketsUsageGetAPIUsage,
					&storageBucketsUsageGetBucketUsage,
				},
			},
			{
				Name:     "storage:migration-sources",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&storageMigrationSourcesCreate,
					&storageMigrationSourcesRetrieve,
					&storageMigrationSourcesList,
					&storageMigrationSourcesDelete,
				},
			},
			{
				Name:     "storage:migrations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&storageMigrationsCreate,
					&storageMigrationsRetrieve,
					&storageMigrationsList,
				},
			},
			{
				Name:     "storage:migrations:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&storageMigrationsActionsStop,
				},
			},
			{
				Name:     "sub-number-orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&subNumberOrdersRetrieve,
					&subNumberOrdersUpdate,
					&subNumberOrdersList,
					&subNumberOrdersCancel,
					&subNumberOrdersUpdateRequirementGroup,
				},
			},
			{
				Name:     "sub-number-orders-report",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&subNumberOrdersReportCreate,
					&subNumberOrdersReportRetrieve,
				},
			},
			{
				Name:     "telephony-credentials",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&telephonyCredentialsCreate,
					&telephonyCredentialsRetrieve,
					&telephonyCredentialsUpdate,
					&telephonyCredentialsList,
					&telephonyCredentialsDelete,
				},
			},
			{
				Name:     "texml",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlSecrets,
				},
			},
			{
				Name:     "texml:accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsRetrieveRecordingsJson,
					&texmlAccountsRetrieveTranscriptionsJson,
				},
			},
			{
				Name:     "texml:accounts:calls",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsCallsRetrieve,
					&texmlAccountsCallsUpdate,
					&texmlAccountsCallsCalls,
					&texmlAccountsCallsRetrieveCalls,
					&texmlAccountsCallsSiprecJson,
					&texmlAccountsCallsStreamsJson,
				},
			},
			{
				Name:     "texml:accounts:calls:recordings-json",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsCallsRecordingsJsonRecordingsJson,
					&texmlAccountsCallsRecordingsJsonRetrieveRecordingsJson,
				},
			},
			{
				Name:     "texml:accounts:calls:recordings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsCallsRecordingsRecordingSidJson,
				},
			},
			{
				Name:     "texml:accounts:calls:siprec",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsCallsSiprecSiprecSidJson,
				},
			},
			{
				Name:     "texml:accounts:calls:streams",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsCallsStreamsStreamingSidJson,
				},
			},
			{
				Name:     "texml:accounts:conferences",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsConferencesRetrieve,
					&texmlAccountsConferencesUpdate,
					&texmlAccountsConferencesRetrieveConferences,
					&texmlAccountsConferencesRetrieveRecordings,
					&texmlAccountsConferencesRetrieveRecordingsJson,
				},
			},
			{
				Name:     "texml:accounts:conferences:participants",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsConferencesParticipantsRetrieve,
					&texmlAccountsConferencesParticipantsUpdate,
					&texmlAccountsConferencesParticipantsDelete,
					&texmlAccountsConferencesParticipantsParticipants,
					&texmlAccountsConferencesParticipantsRetrieveParticipants,
				},
			},
			{
				Name:     "texml:accounts:recordings:json",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsRecordingsJsonDeleteRecordingSidJson,
					&texmlAccountsRecordingsJsonRetrieveRecordingSidJson,
				},
			},
			{
				Name:     "texml:accounts:transcriptions:json",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsTranscriptionsJsonDeleteRecordingTranscriptionSidJson,
					&texmlAccountsTranscriptionsJsonRetrieveRecordingTranscriptionSidJson,
				},
			},
			{
				Name:     "texml:accounts:queues",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlAccountsQueuesCreate,
					&texmlAccountsQueuesRetrieve,
					&texmlAccountsQueuesUpdate,
					&texmlAccountsQueuesList,
					&texmlAccountsQueuesDelete,
				},
			},
			{
				Name:     "texml-applications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&texmlApplicationsCreate,
					&texmlApplicationsRetrieve,
					&texmlApplicationsUpdate,
					&texmlApplicationsList,
					&texmlApplicationsDelete,
				},
			},
			{
				Name:     "text-to-speech",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&textToSpeechListVoices,
				},
			},
			{
				Name:     "usage-reports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usageReportsList,
					&usageReportsGetOptions,
				},
			},
			{
				Name:     "user-addresses",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&userAddressesCreate,
					&userAddressesRetrieve,
					&userAddressesList,
				},
			},
			{
				Name:     "user-tags",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&userTagsList,
				},
			},
			{
				Name:     "verifications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&verificationsRetrieve,
					&verificationsTriggerCall,
					&verificationsTriggerFlashcall,
					&verificationsTriggerSMS,
				},
			},
			{
				Name:     "verifications:by-phone-number",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&verificationsByPhoneNumberList,
				},
			},
			{
				Name:     "verifications:by-phone-number:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&verificationsByPhoneNumberActionsVerify,
				},
			},
			{
				Name:     "verifications:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&verificationsActionsVerify,
				},
			},
			{
				Name:     "verified-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&verifiedNumbersCreate,
					&verifiedNumbersRetrieve,
					&verifiedNumbersList,
					&verifiedNumbersDelete,
				},
			},
			{
				Name:     "verified-numbers:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&verifiedNumbersActionsSubmitVerificationCode,
				},
			},
			{
				Name:     "verify-profiles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&verifyProfilesCreate,
					&verifyProfilesRetrieve,
					&verifyProfilesUpdate,
					&verifyProfilesList,
					&verifyProfilesDelete,
					&verifyProfilesCreateTemplate,
					&verifyProfilesRetrieveTemplates,
					&verifyProfilesUpdateTemplate,
				},
			},
			{
				Name:     "virtual-cross-connects",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&virtualCrossConnectsCreate,
					&virtualCrossConnectsRetrieve,
					&virtualCrossConnectsUpdate,
					&virtualCrossConnectsList,
					&virtualCrossConnectsDelete,
				},
			},
			{
				Name:     "virtual-cross-connects-coverage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&virtualCrossConnectsCoverageList,
				},
			},
			{
				Name:     "webhook-deliveries",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&webhookDeliveriesRetrieve,
					&webhookDeliveriesList,
				},
			},
			{
				Name:     "wireguard-interfaces",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wireguardInterfacesCreate,
					&wireguardInterfacesRetrieve,
					&wireguardInterfacesList,
					&wireguardInterfacesDelete,
				},
			},
			{
				Name:     "wireguard-peers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wireguardPeersCreate,
					&wireguardPeersRetrieve,
					&wireguardPeersUpdate,
					&wireguardPeersList,
					&wireguardPeersDelete,
				},
			},
			{
				Name:     "wireless",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wirelessRetrieveRegions,
				},
			},
			{
				Name:     "wireless:detail-records-reports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wirelessDetailRecordsReportsCreate,
					&wirelessDetailRecordsReportsRetrieve,
					&wirelessDetailRecordsReportsList,
					&wirelessDetailRecordsReportsDelete,
				},
			},
			{
				Name:     "wireless-blocklist-values",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wirelessBlocklistValuesList,
				},
			},
			{
				Name:     "wireless-blocklists",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wirelessBlocklistsCreate,
					&wirelessBlocklistsRetrieve,
					&wirelessBlocklistsUpdate,
					&wirelessBlocklistsList,
					&wirelessBlocklistsDelete,
				},
			},
			{
				Name:     "well-known",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wellKnownRetrieveAuthorizationServerMetadata,
					&wellKnownRetrieveProtectedResourceMetadata,
				},
			},
			{
				Name:     "inexplicit-number-orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inexplicitNumberOrdersCreate,
					&inexplicitNumberOrdersRetrieve,
					&inexplicitNumberOrdersList,
				},
			},
			{
				Name:     "mobile-phone-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mobilePhoneNumbersRetrieve,
					&mobilePhoneNumbersUpdate,
					&mobilePhoneNumbersList,
				},
			},
			{
				Name:     "mobile-phone-numbers:messaging",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mobilePhoneNumbersMessagingRetrieve,
					&mobilePhoneNumbersMessagingList,
				},
			},
			{
				Name:     "mobile-voice-connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mobileVoiceConnectionsCreate,
					&mobileVoiceConnectionsRetrieve,
					&mobileVoiceConnectionsUpdate,
					&mobileVoiceConnectionsList,
					&mobileVoiceConnectionsDelete,
				},
			},
			{
				Name:     "messaging-10dlc",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcGetEnum,
				},
			},
			{
				Name:     "messaging-10dlc:brand",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcBrandCreate,
					&messaging10dlcBrandRetrieve,
					&messaging10dlcBrandUpdate,
					&messaging10dlcBrandList,
					&messaging10dlcBrandDelete,
					&messaging10dlcBrandGetFeedback,
					&messaging10dlcBrandGetSMSOtpByReference,
					&messaging10dlcBrandResend2faEmail,
					&messaging10dlcBrandRetrieveSMSOtpStatus,
					&messaging10dlcBrandRevet,
					&messaging10dlcBrandTriggerSMSOtp,
					&messaging10dlcBrandVerifySMSOtp,
				},
			},
			{
				Name:     "messaging-10dlc:brand:external-vetting",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcBrandExternalVettingList,
					&messaging10dlcBrandExternalVettingImports,
					&messaging10dlcBrandExternalVettingOrder,
				},
			},
			{
				Name:     "messaging-10dlc:campaign",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcCampaignRetrieve,
					&messaging10dlcCampaignUpdate,
					&messaging10dlcCampaignList,
					&messaging10dlcCampaignAcceptSharing,
					&messaging10dlcCampaignDeactivate,
					&messaging10dlcCampaignGetMnoMetadata,
					&messaging10dlcCampaignGetOperationStatus,
					&messaging10dlcCampaignGetSharingStatus,
					&messaging10dlcCampaignSubmitAppeal,
				},
			},
			{
				Name:     "messaging-10dlc:campaign:usecase",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcCampaignUsecaseGetCost,
				},
			},
			{
				Name:     "messaging-10dlc:campaign:osr",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcCampaignOsrGetAttributes,
				},
			},
			{
				Name:     "messaging-10dlc:campaign-builder",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcCampaignBuilderSubmit,
				},
			},
			{
				Name:     "messaging-10dlc:campaign-builder:brand",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcCampaignBuilderBrandQualifyByUsecase,
				},
			},
			{
				Name:     "messaging-10dlc:partner-campaigns",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcPartnerCampaignsRetrieve,
					&messaging10dlcPartnerCampaignsUpdate,
					&messaging10dlcPartnerCampaignsList,
					&messaging10dlcPartnerCampaignsListSharedByMe,
					&messaging10dlcPartnerCampaignsRetrieveSharingStatus,
				},
			},
			{
				Name:     "messaging-10dlc:phone-number-campaigns",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcPhoneNumberCampaignsCreate,
					&messaging10dlcPhoneNumberCampaignsRetrieve,
					&messaging10dlcPhoneNumberCampaignsUpdate,
					&messaging10dlcPhoneNumberCampaignsList,
					&messaging10dlcPhoneNumberCampaignsDelete,
				},
			},
			{
				Name:     "messaging-10dlc:phone-number-assignment-by-profile",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messaging10dlcPhoneNumberAssignmentByProfileAssign,
					&messaging10dlcPhoneNumberAssignmentByProfileListPhoneNumberStatus,
					&messaging10dlcPhoneNumberAssignmentByProfileRetrievePhoneNumberStatus,
					&messaging10dlcPhoneNumberAssignmentByProfileRetrieveStatus,
				},
			},
			{
				Name:     "speech-to-text",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&speechToTextTranscribe,
				},
			},
			{
				Name:     "organizations:users",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&organizationsUsersRetrieve,
					&organizationsUsersList,
					&organizationsUsersGetGroupsReport,
				},
			},
			{
				Name:     "organizations:users:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&organizationsUsersActionsRemove,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "telnyx @manpages [-o telnyx.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "telnyx.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "telnyx.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
