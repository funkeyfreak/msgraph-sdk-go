/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphMessageRule struct for MicrosoftGraphMessageRule
type MicrosoftGraphMessageRule struct {
	Id          string      `json:"id,omitempty"`
	Actions     interface{} `json:"actions,omitempty"`
	Conditions  interface{} `json:"conditions,omitempty"`
	DisplayName string      `json:"displayName,omitempty"`
	Exceptions  interface{} `json:"exceptions,omitempty"`
	HasError    bool        `json:"hasError,omitempty"`
	IsEnabled   bool        `json:"isEnabled,omitempty"`
	IsReadOnly  bool        `json:"isReadOnly,omitempty"`
	Sequence    int32       `json:"sequence,omitempty"`
}
