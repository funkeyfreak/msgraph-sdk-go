/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphTimeSlot struct for MicrosoftGraphTimeSlot
type MicrosoftGraphTimeSlot struct {
	End   MicrosoftGraphDateTimeTimeZone `json:"end,omitempty"`
	Start MicrosoftGraphDateTimeTimeZone `json:"start,omitempty"`
}
