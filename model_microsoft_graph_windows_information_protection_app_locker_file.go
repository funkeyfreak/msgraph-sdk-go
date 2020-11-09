/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWindowsInformationProtectionAppLockerFile struct for MicrosoftGraphWindowsInformationProtectionAppLockerFile
type MicrosoftGraphWindowsInformationProtectionAppLockerFile struct {
	Id string `json:"id,omitempty"`
	// The friendly name
	DisplayName string `json:"displayName,omitempty"`
	// File as a byte array
	File string `json:"file,omitempty"`
	// SHA256 hash of the file
	FileHash string `json:"fileHash,omitempty"`
	// Version of the entity.
	Version string `json:"version,omitempty"`
}
