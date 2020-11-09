/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphUserSettings struct for MicrosoftGraphUserSettings
type MicrosoftGraphUserSettings struct {
	Id                                                   string      `json:"id,omitempty"`
	ContributionToContentDiscoveryAsOrganizationDisabled bool        `json:"contributionToContentDiscoveryAsOrganizationDisabled,omitempty"`
	ContributionToContentDiscoveryDisabled               bool        `json:"contributionToContentDiscoveryDisabled,omitempty"`
	ShiftPreferences                                     interface{} `json:"shiftPreferences,omitempty"`
}
