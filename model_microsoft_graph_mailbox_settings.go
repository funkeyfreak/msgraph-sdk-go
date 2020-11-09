/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphMailboxSettings struct for MicrosoftGraphMailboxSettings
type MicrosoftGraphMailboxSettings struct {
	ArchiveFolder                         string      `json:"archiveFolder,omitempty"`
	AutomaticRepliesSetting               interface{} `json:"automaticRepliesSetting,omitempty"`
	DateFormat                            string      `json:"dateFormat,omitempty"`
	DelegateMeetingMessageDeliveryOptions interface{} `json:"delegateMeetingMessageDeliveryOptions,omitempty"`
	Language                              interface{} `json:"language,omitempty"`
	TimeFormat                            string      `json:"timeFormat,omitempty"`
	TimeZone                              string      `json:"timeZone,omitempty"`
	WorkingHours                          interface{} `json:"workingHours,omitempty"`
}