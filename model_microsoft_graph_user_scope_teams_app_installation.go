/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphUserScopeTeamsAppInstallation struct for MicrosoftGraphUserScopeTeamsAppInstallation
type MicrosoftGraphUserScopeTeamsAppInstallation struct {
	Id                 string      `json:"id,omitempty"`
	TeamsApp           interface{} `json:"teamsApp,omitempty"`
	TeamsAppDefinition interface{} `json:"teamsAppDefinition,omitempty"`
	Chat               interface{} `json:"chat,omitempty"`
}