/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphResourceVisualization struct for MicrosoftGraphResourceVisualization
type MicrosoftGraphResourceVisualization struct {
	ContainerDisplayName string `json:"containerDisplayName,omitempty"`
	ContainerType        string `json:"containerType,omitempty"`
	ContainerWebUrl      string `json:"containerWebUrl,omitempty"`
	MediaType            string `json:"mediaType,omitempty"`
	PreviewImageUrl      string `json:"previewImageUrl,omitempty"`
	PreviewText          string `json:"previewText,omitempty"`
	Title                string `json:"title,omitempty"`
	Type                 string `json:"type,omitempty"`
}
