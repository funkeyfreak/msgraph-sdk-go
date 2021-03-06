/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphSharedInsight struct for MicrosoftGraphSharedInsight
type MicrosoftGraphSharedInsight struct {
	Id                    string        `json:"id,omitempty"`
	LastShared            interface{}   `json:"lastShared,omitempty"`
	ResourceReference     interface{}   `json:"resourceReference,omitempty"`
	ResourceVisualization interface{}   `json:"resourceVisualization,omitempty"`
	SharingHistory        []interface{} `json:"sharingHistory,omitempty"`
	LastSharedMethod      interface{}   `json:"lastSharedMethod,omitempty"`
	Resource              interface{}   `json:"resource,omitempty"`
}
