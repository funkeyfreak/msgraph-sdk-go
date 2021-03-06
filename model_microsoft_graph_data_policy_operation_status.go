/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphDataPolicyOperationStatus the model 'MicrosoftGraphDataPolicyOperationStatus'
type MicrosoftGraphDataPolicyOperationStatus string

// List of microsoft.graph.dataPolicyOperationStatus
const (
	MICROSOFTGRAPHDATAPOLICYOPERATIONSTATUS_NOT_STARTED          MicrosoftGraphDataPolicyOperationStatus = "notStarted"
	MICROSOFTGRAPHDATAPOLICYOPERATIONSTATUS_RUNNING              MicrosoftGraphDataPolicyOperationStatus = "running"
	MICROSOFTGRAPHDATAPOLICYOPERATIONSTATUS_COMPLETE             MicrosoftGraphDataPolicyOperationStatus = "complete"
	MICROSOFTGRAPHDATAPOLICYOPERATIONSTATUS_FAILED               MicrosoftGraphDataPolicyOperationStatus = "failed"
	MICROSOFTGRAPHDATAPOLICYOPERATIONSTATUS_UNKNOWN_FUTURE_VALUE MicrosoftGraphDataPolicyOperationStatus = "unknownFutureValue"
)
