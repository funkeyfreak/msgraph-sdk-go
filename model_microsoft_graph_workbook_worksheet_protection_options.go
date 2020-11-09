/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWorkbookWorksheetProtectionOptions struct for MicrosoftGraphWorkbookWorksheetProtectionOptions
type MicrosoftGraphWorkbookWorksheetProtectionOptions struct {
	AllowAutoFilter       bool `json:"allowAutoFilter,omitempty"`
	AllowDeleteColumns    bool `json:"allowDeleteColumns,omitempty"`
	AllowDeleteRows       bool `json:"allowDeleteRows,omitempty"`
	AllowFormatCells      bool `json:"allowFormatCells,omitempty"`
	AllowFormatColumns    bool `json:"allowFormatColumns,omitempty"`
	AllowFormatRows       bool `json:"allowFormatRows,omitempty"`
	AllowInsertColumns    bool `json:"allowInsertColumns,omitempty"`
	AllowInsertHyperlinks bool `json:"allowInsertHyperlinks,omitempty"`
	AllowInsertRows       bool `json:"allowInsertRows,omitempty"`
	AllowPivotTables      bool `json:"allowPivotTables,omitempty"`
	AllowSort             bool `json:"allowSort,omitempty"`
}
