/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphParticipant struct for MicrosoftGraphParticipant
type MicrosoftGraphParticipant struct {
	Id            string                        `json:"id,omitempty"`
	Info          MicrosoftGraphParticipantInfo `json:"info,omitempty"`
	IsInLobby     bool                          `json:"isInLobby,omitempty"`
	IsMuted       bool                          `json:"isMuted,omitempty"`
	MediaStreams  []interface{}                 `json:"mediaStreams,omitempty"`
	RecordingInfo interface{}                   `json:"recordingInfo,omitempty"`
}
