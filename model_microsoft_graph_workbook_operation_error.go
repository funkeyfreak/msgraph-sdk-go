/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWorkbookOperationError struct for MicrosoftGraphWorkbookOperationError
type MicrosoftGraphWorkbookOperationError struct {
	Code       string      `json:"code,omitempty"`
	InnerError interface{} `json:"innerError,omitempty"`
	Message    string      `json:"message,omitempty"`
}