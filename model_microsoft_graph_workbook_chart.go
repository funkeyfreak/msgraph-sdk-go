/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWorkbookChart struct for MicrosoftGraphWorkbookChart
type MicrosoftGraphWorkbookChart struct {
	Id         string                              `json:"id,omitempty"`
	Height     interface{}                         `json:"height,omitempty"`
	Left       interface{}                         `json:"left,omitempty"`
	Name       string                              `json:"name,omitempty"`
	Top        interface{}                         `json:"top,omitempty"`
	Width      interface{}                         `json:"width,omitempty"`
	Axes       interface{}                         `json:"axes,omitempty"`
	DataLabels interface{}                         `json:"dataLabels,omitempty"`
	Format     interface{}                         `json:"format,omitempty"`
	Legend     interface{}                         `json:"legend,omitempty"`
	Series     []MicrosoftGraphWorkbookChartSeries `json:"series,omitempty"`
	Title      interface{}                         `json:"title,omitempty"`
	Worksheet  interface{}                         `json:"worksheet,omitempty"`
}
