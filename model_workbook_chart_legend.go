/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// WorkbookChartLegend struct for WorkbookChartLegend
type WorkbookChartLegend struct {
	Overlay  bool        `json:"overlay,omitempty"`
	Position string      `json:"position,omitempty"`
	Visible  bool        `json:"visible,omitempty"`
	Format   interface{} `json:"format,omitempty"`
}
