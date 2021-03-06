/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphHashes struct for MicrosoftGraphHashes
type MicrosoftGraphHashes struct {
	Crc32Hash    string `json:"crc32Hash,omitempty"`
	QuickXorHash string `json:"quickXorHash,omitempty"`
	Sha1Hash     string `json:"sha1Hash,omitempty"`
	Sha256Hash   string `json:"sha256Hash,omitempty"`
}
