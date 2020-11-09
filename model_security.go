/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// Security struct for Security
type Security struct {
	Alerts                     []MicrosoftGraphAlert                     `json:"alerts,omitempty"`
	SecureScoreControlProfiles []MicrosoftGraphSecureScoreControlProfile `json:"secureScoreControlProfiles,omitempty"`
	SecureScores               []MicrosoftGraphSecureScore               `json:"secureScores,omitempty"`
}
