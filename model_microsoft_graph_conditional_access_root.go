/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphConditionalAccessRoot struct for MicrosoftGraphConditionalAccessRoot
type MicrosoftGraphConditionalAccessRoot struct {
	Id             string                                  `json:"id,omitempty"`
	NamedLocations []MicrosoftGraphNamedLocation           `json:"namedLocations,omitempty"`
	Policies       []MicrosoftGraphConditionalAccessPolicy `json:"policies,omitempty"`
}
