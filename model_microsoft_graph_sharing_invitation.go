/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphSharingInvitation struct for MicrosoftGraphSharingInvitation
type MicrosoftGraphSharingInvitation struct {
	Email          string      `json:"email,omitempty"`
	InvitedBy      interface{} `json:"invitedBy,omitempty"`
	RedeemedBy     string      `json:"redeemedBy,omitempty"`
	SignInRequired bool        `json:"signInRequired,omitempty"`
}