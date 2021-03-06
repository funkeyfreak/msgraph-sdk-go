/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// Invitation struct for Invitation
type Invitation struct {
	InvitedUserDisplayName  string      `json:"invitedUserDisplayName,omitempty"`
	InvitedUserEmailAddress string      `json:"invitedUserEmailAddress,omitempty"`
	InvitedUserMessageInfo  interface{} `json:"invitedUserMessageInfo,omitempty"`
	InvitedUserType         string      `json:"invitedUserType,omitempty"`
	InviteRedeemUrl         string      `json:"inviteRedeemUrl,omitempty"`
	InviteRedirectUrl       string      `json:"inviteRedirectUrl,omitempty"`
	SendInvitationMessage   bool        `json:"sendInvitationMessage,omitempty"`
	Status                  string      `json:"status,omitempty"`
	InvitedUser             interface{} `json:"invitedUser,omitempty"`
}
