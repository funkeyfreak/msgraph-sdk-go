/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphDeviceCompliancePolicySettingState Device Compilance Policy Setting State for a given device.
type MicrosoftGraphDeviceCompliancePolicySettingState struct {
	// Current value of setting on device
	CurrentValue string `json:"currentValue,omitempty"`
	// Error code for the setting
	ErrorCode int64 `json:"errorCode,omitempty"`
	// Error description
	ErrorDescription string `json:"errorDescription,omitempty"`
	// Name of setting instance that is being reported.
	InstanceDisplayName string `json:"instanceDisplayName,omitempty"`
	// The setting that is being reported
	Setting string `json:"setting,omitempty"`
	// Localized/user friendly setting name that is being reported
	SettingName string `json:"settingName,omitempty"`
	// Contributing policies
	Sources []interface{} `json:"sources,omitempty"`
	// The compliance state of the setting
	State interface{} `json:"state,omitempty"`
	// UserEmail
	UserEmail string `json:"userEmail,omitempty"`
	// UserId
	UserId string `json:"userId,omitempty"`
	// UserName
	UserName string `json:"userName,omitempty"`
	// UserPrincipalName.
	UserPrincipalName string `json:"userPrincipalName,omitempty"`
}
