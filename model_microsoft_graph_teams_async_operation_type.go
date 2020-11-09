/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphTeamsAsyncOperationType the model 'MicrosoftGraphTeamsAsyncOperationType'
type MicrosoftGraphTeamsAsyncOperationType string

// List of microsoft.graph.teamsAsyncOperationType
const (
	MICROSOFTGRAPHTEAMSASYNCOPERATIONTYPE_INVALID              MicrosoftGraphTeamsAsyncOperationType = "invalid"
	MICROSOFTGRAPHTEAMSASYNCOPERATIONTYPE_CLONE_TEAM           MicrosoftGraphTeamsAsyncOperationType = "cloneTeam"
	MICROSOFTGRAPHTEAMSASYNCOPERATIONTYPE_ARCHIVE_TEAM         MicrosoftGraphTeamsAsyncOperationType = "archiveTeam"
	MICROSOFTGRAPHTEAMSASYNCOPERATIONTYPE_UNARCHIVE_TEAM       MicrosoftGraphTeamsAsyncOperationType = "unarchiveTeam"
	MICROSOFTGRAPHTEAMSASYNCOPERATIONTYPE_CREATE_TEAM          MicrosoftGraphTeamsAsyncOperationType = "createTeam"
	MICROSOFTGRAPHTEAMSASYNCOPERATIONTYPE_UNKNOWN_FUTURE_VALUE MicrosoftGraphTeamsAsyncOperationType = "unknownFutureValue"
)
