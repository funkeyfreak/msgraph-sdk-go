/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphConversationMember struct for MicrosoftGraphConversationMember
type MicrosoftGraphConversationMember struct {
	Id          string   `json:"id,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
	Roles       []string `json:"roles,omitempty"`
}
