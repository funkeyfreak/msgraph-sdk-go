/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphConditionalAccessApplications struct for MicrosoftGraphConditionalAccessApplications
type MicrosoftGraphConditionalAccessApplications struct {
	ExcludeApplications []string `json:"excludeApplications,omitempty"`
	IncludeApplications []string `json:"includeApplications,omitempty"`
	IncludeUserActions  []string `json:"includeUserActions,omitempty"`
}
