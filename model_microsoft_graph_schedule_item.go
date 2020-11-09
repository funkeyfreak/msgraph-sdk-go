/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphScheduleItem struct for MicrosoftGraphScheduleItem
type MicrosoftGraphScheduleItem struct {
	End       interface{} `json:"end,omitempty"`
	IsPrivate bool        `json:"isPrivate,omitempty"`
	Location  string      `json:"location,omitempty"`
	Start     interface{} `json:"start,omitempty"`
	Status    interface{} `json:"status,omitempty"`
	Subject   string      `json:"subject,omitempty"`
}
