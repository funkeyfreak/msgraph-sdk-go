/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// ConditionalAccessRoot struct for ConditionalAccessRoot
type ConditionalAccessRoot struct {
	NamedLocations []MicrosoftGraphNamedLocation           `json:"namedLocations,omitempty"`
	Policies       []MicrosoftGraphConditionalAccessPolicy `json:"policies,omitempty"`
}
