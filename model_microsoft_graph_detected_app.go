/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphDetectedApp struct for MicrosoftGraphDetectedApp
type MicrosoftGraphDetectedApp struct {
	Id string `json:"id,omitempty"`
	// The number of devices that have installed this application
	DeviceCount int32 `json:"deviceCount,omitempty"`
	// Name of the discovered application. Read-only
	DisplayName string `json:"displayName,omitempty"`
	// Discovered application size in bytes. Read-only
	SizeInByte int64 `json:"sizeInByte,omitempty"`
	// Version of the discovered application. Read-only
	Version        string                        `json:"version,omitempty"`
	ManagedDevices []MicrosoftGraphManagedDevice `json:"managedDevices,omitempty"`
}