/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphAutomaticRepliesMailTips struct for MicrosoftGraphAutomaticRepliesMailTips
type MicrosoftGraphAutomaticRepliesMailTips struct {
	Message            string      `json:"message,omitempty"`
	MessageLanguage    interface{} `json:"messageLanguage,omitempty"`
	ScheduledEndTime   interface{} `json:"scheduledEndTime,omitempty"`
	ScheduledStartTime interface{} `json:"scheduledStartTime,omitempty"`
}