/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWorkbookTableSort struct for MicrosoftGraphWorkbookTableSort
type MicrosoftGraphWorkbookTableSort struct {
	Id        string        `json:"id,omitempty"`
	Fields    []interface{} `json:"fields,omitempty"`
	MatchCase bool          `json:"matchCase,omitempty"`
	Method    string        `json:"method,omitempty"`
}