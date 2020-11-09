/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphDeviceManagementSettings struct for MicrosoftGraphDeviceManagementSettings
type MicrosoftGraphDeviceManagementSettings struct {
	// The number of days a device is allowed to go without checking in to remain compliant.
	DeviceComplianceCheckinThresholdDays int32 `json:"deviceComplianceCheckinThresholdDays,omitempty"`
	// Is feature enabled or not for scheduled action for rule.
	IsScheduledActionEnabled bool `json:"isScheduledActionEnabled,omitempty"`
	// Device should be noncompliant when there is no compliance policy targeted when this is true
	SecureByDefault bool `json:"secureByDefault,omitempty"`
}
