/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphMediaStream struct for MicrosoftGraphMediaStream
type MicrosoftGraphMediaStream struct {
	Direction   interface{} `json:"direction,omitempty"`
	Label       string      `json:"label,omitempty"`
	MediaType   interface{} `json:"mediaType,omitempty"`
	ServerMuted bool        `json:"serverMuted,omitempty"`
	SourceId    string      `json:"sourceId,omitempty"`
}