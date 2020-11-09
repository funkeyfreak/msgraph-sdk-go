/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphTeamMessagingSettings struct for MicrosoftGraphTeamMessagingSettings
type MicrosoftGraphTeamMessagingSettings struct {
	AllowChannelMentions     bool `json:"allowChannelMentions,omitempty"`
	AllowOwnerDeleteMessages bool `json:"allowOwnerDeleteMessages,omitempty"`
	AllowTeamMentions        bool `json:"allowTeamMentions,omitempty"`
	AllowUserDeleteMessages  bool `json:"allowUserDeleteMessages,omitempty"`
	AllowUserEditMessages    bool `json:"allowUserEditMessages,omitempty"`
}
