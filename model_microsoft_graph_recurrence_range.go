/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphRecurrenceRange struct for MicrosoftGraphRecurrenceRange
type MicrosoftGraphRecurrenceRange struct {
	EndDate             string      `json:"endDate,omitempty"`
	NumberOfOccurrences int32       `json:"numberOfOccurrences,omitempty"`
	RecurrenceTimeZone  string      `json:"recurrenceTimeZone,omitempty"`
	StartDate           string      `json:"startDate,omitempty"`
	Type                interface{} `json:"type,omitempty"`
}