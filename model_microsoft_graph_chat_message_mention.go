/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphChatMessageMention struct for MicrosoftGraphChatMessageMention
type MicrosoftGraphChatMessageMention struct {
	Id          int32       `json:"id,omitempty"`
	Mentioned   interface{} `json:"mentioned,omitempty"`
	MentionText string      `json:"mentionText,omitempty"`
}
