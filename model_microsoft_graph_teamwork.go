/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphTeamwork struct for MicrosoftGraphTeamwork
type MicrosoftGraphTeamwork struct {
	Id                    string                               `json:"id,omitempty"`
	WorkforceIntegrations []MicrosoftGraphWorkforceIntegration `json:"workforceIntegrations,omitempty"`
}
