/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWorkbook struct for MicrosoftGraphWorkbook
type MicrosoftGraphWorkbook struct {
	Id          string                            `json:"id,omitempty"`
	Application interface{}                       `json:"application,omitempty"`
	Comments    []MicrosoftGraphWorkbookComment   `json:"comments,omitempty"`
	Functions   interface{}                       `json:"functions,omitempty"`
	Names       []MicrosoftGraphWorkbookNamedItem `json:"names,omitempty"`
	Operations  []MicrosoftGraphWorkbookOperation `json:"operations,omitempty"`
	Tables      []MicrosoftGraphWorkbookTable     `json:"tables,omitempty"`
	Worksheets  []MicrosoftGraphWorkbookWorksheet `json:"worksheets,omitempty"`
}