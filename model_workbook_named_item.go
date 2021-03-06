/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// WorkbookNamedItem struct for WorkbookNamedItem
type WorkbookNamedItem struct {
	Comment   string      `json:"comment,omitempty"`
	Name      string      `json:"name,omitempty"`
	Scope     string      `json:"scope,omitempty"`
	Type      string      `json:"type,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Visible   bool        `json:"visible,omitempty"`
	Worksheet interface{} `json:"worksheet,omitempty"`
}
