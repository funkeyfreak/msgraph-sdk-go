/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphFile struct for MicrosoftGraphFile
type MicrosoftGraphFile struct {
	Hashes             interface{} `json:"hashes,omitempty"`
	MimeType           string      `json:"mimeType,omitempty"`
	ProcessingMetadata bool        `json:"processingMetadata,omitempty"`
}
