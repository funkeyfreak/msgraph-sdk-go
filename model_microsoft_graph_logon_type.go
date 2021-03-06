/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphLogonType the model 'MicrosoftGraphLogonType'
type MicrosoftGraphLogonType string

// List of microsoft.graph.logonType
const (
	MICROSOFTGRAPHLOGONTYPE_UNKNOWN              MicrosoftGraphLogonType = "unknown"
	MICROSOFTGRAPHLOGONTYPE_INTERACTIVE          MicrosoftGraphLogonType = "interactive"
	MICROSOFTGRAPHLOGONTYPE_REMOTE_INTERACTIVE   MicrosoftGraphLogonType = "remoteInteractive"
	MICROSOFTGRAPHLOGONTYPE_NETWORK              MicrosoftGraphLogonType = "network"
	MICROSOFTGRAPHLOGONTYPE_BATCH                MicrosoftGraphLogonType = "batch"
	MICROSOFTGRAPHLOGONTYPE_SERVICE              MicrosoftGraphLogonType = "service"
	MICROSOFTGRAPHLOGONTYPE_UNKNOWN_FUTURE_VALUE MicrosoftGraphLogonType = "unknownFutureValue"
)
