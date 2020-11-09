/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphVppTokenSyncStatus the model 'MicrosoftGraphVppTokenSyncStatus'
type MicrosoftGraphVppTokenSyncStatus string

// List of microsoft.graph.vppTokenSyncStatus
const (
	MICROSOFTGRAPHVPPTOKENSYNCSTATUS_NONE        MicrosoftGraphVppTokenSyncStatus = "none"
	MICROSOFTGRAPHVPPTOKENSYNCSTATUS_IN_PROGRESS MicrosoftGraphVppTokenSyncStatus = "inProgress"
	MICROSOFTGRAPHVPPTOKENSYNCSTATUS_COMPLETED   MicrosoftGraphVppTokenSyncStatus = "completed"
	MICROSOFTGRAPHVPPTOKENSYNCSTATUS_FAILED      MicrosoftGraphVppTokenSyncStatus = "failed"
)