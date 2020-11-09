/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphInviteParticipantsOperation struct for MicrosoftGraphInviteParticipantsOperation
type MicrosoftGraphInviteParticipantsOperation struct {
	Id            string                                    `json:"id,omitempty"`
	ClientContext string                                    `json:"clientContext,omitempty"`
	ResultInfo    interface{}                               `json:"resultInfo,omitempty"`
	Status        interface{}                               `json:"status,omitempty"`
	Participants  []MicrosoftGraphInvitationParticipantInfo `json:"participants,omitempty"`
}
