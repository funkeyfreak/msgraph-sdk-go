/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphNumberColumn struct for MicrosoftGraphNumberColumn
type MicrosoftGraphNumberColumn struct {
	DecimalPlaces string      `json:"decimalPlaces,omitempty"`
	DisplayAs     string      `json:"displayAs,omitempty"`
	Maximum       interface{} `json:"maximum,omitempty"`
	Minimum       interface{} `json:"minimum,omitempty"`
}
