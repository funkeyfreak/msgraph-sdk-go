/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// DeviceConfigurationAssignment The device configuration assignment entity assigns an AAD group to a specific device configuration.
type DeviceConfigurationAssignment struct {
	// The assignment target for the device configuration.
	Target interface{} `json:"target,omitempty"`
}
