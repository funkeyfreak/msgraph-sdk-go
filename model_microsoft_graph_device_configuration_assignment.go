/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphDeviceConfigurationAssignment struct for MicrosoftGraphDeviceConfigurationAssignment
type MicrosoftGraphDeviceConfigurationAssignment struct {
	Id string `json:"id,omitempty"`
	// The assignment target for the device configuration.
	Target interface{} `json:"target,omitempty"`
}
