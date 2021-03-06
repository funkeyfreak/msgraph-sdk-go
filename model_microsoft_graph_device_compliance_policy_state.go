/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphDeviceCompliancePolicyState struct for MicrosoftGraphDeviceCompliancePolicyState
type MicrosoftGraphDeviceCompliancePolicyState struct {
	Id string `json:"id,omitempty"`
	// The name of the policy for this policyBase
	DisplayName string `json:"displayName,omitempty"`
	// Platform type that the policy applies to
	PlatformType interface{} `json:"platformType,omitempty"`
	// Count of how many setting a policy holds
	SettingCount  int32         `json:"settingCount,omitempty"`
	SettingStates []interface{} `json:"settingStates,omitempty"`
	// The compliance state of the policy
	State interface{} `json:"state,omitempty"`
	// The version of the policy
	Version int32 `json:"version,omitempty"`
}
