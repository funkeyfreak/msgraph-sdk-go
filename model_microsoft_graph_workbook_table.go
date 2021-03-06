/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWorkbookTable struct for MicrosoftGraphWorkbookTable
type MicrosoftGraphWorkbookTable struct {
	Id                   string                              `json:"id,omitempty"`
	HighlightFirstColumn bool                                `json:"highlightFirstColumn,omitempty"`
	HighlightLastColumn  bool                                `json:"highlightLastColumn,omitempty"`
	LegacyId             string                              `json:"legacyId,omitempty"`
	Name                 string                              `json:"name,omitempty"`
	ShowBandedColumns    bool                                `json:"showBandedColumns,omitempty"`
	ShowBandedRows       bool                                `json:"showBandedRows,omitempty"`
	ShowFilterButton     bool                                `json:"showFilterButton,omitempty"`
	ShowHeaders          bool                                `json:"showHeaders,omitempty"`
	ShowTotals           bool                                `json:"showTotals,omitempty"`
	Style                string                              `json:"style,omitempty"`
	Columns              []MicrosoftGraphWorkbookTableColumn `json:"columns,omitempty"`
	Rows                 []MicrosoftGraphWorkbookTableRow    `json:"rows,omitempty"`
	Sort                 interface{}                         `json:"sort,omitempty"`
	Worksheet            interface{}                         `json:"worksheet,omitempty"`
}
