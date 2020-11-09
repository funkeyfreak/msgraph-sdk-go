/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWorkbookOperationStatus the model 'MicrosoftGraphWorkbookOperationStatus'
type MicrosoftGraphWorkbookOperationStatus string

// List of microsoft.graph.workbookOperationStatus
const (
	MICROSOFTGRAPHWORKBOOKOPERATIONSTATUS_NOT_STARTED MicrosoftGraphWorkbookOperationStatus = "notStarted"
	MICROSOFTGRAPHWORKBOOKOPERATIONSTATUS_RUNNING     MicrosoftGraphWorkbookOperationStatus = "running"
	MICROSOFTGRAPHWORKBOOKOPERATIONSTATUS_SUCCEEDED   MicrosoftGraphWorkbookOperationStatus = "succeeded"
	MICROSOFTGRAPHWORKBOOKOPERATIONSTATUS_FAILED      MicrosoftGraphWorkbookOperationStatus = "failed"
)
