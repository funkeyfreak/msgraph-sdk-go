/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphPublicInnerError struct for MicrosoftGraphPublicInnerError
type MicrosoftGraphPublicInnerError struct {
	Code    string        `json:"code,omitempty"`
	Details []interface{} `json:"details,omitempty"`
	Message string        `json:"message,omitempty"`
	Target  string        `json:"target,omitempty"`
}
