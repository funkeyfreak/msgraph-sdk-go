/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// DeviceCompliancePolicySettingStateSummary Device Compilance Policy Setting State summary across the account.
type DeviceCompliancePolicySettingStateSummary struct {
	// Number of compliant devices
	CompliantDeviceCount int32 `json:"compliantDeviceCount,omitempty"`
	// Number of conflict devices
	ConflictDeviceCount int32 `json:"conflictDeviceCount,omitempty"`
	// Number of error devices
	ErrorDeviceCount int32 `json:"errorDeviceCount,omitempty"`
	// Number of NonCompliant devices
	NonCompliantDeviceCount int32 `json:"nonCompliantDeviceCount,omitempty"`
	// Number of not applicable devices
	NotApplicableDeviceCount int32 `json:"notApplicableDeviceCount,omitempty"`
	// Setting platform
	PlatformType interface{} `json:"platformType,omitempty"`
	// Number of remediated devices
	RemediatedDeviceCount int32 `json:"remediatedDeviceCount,omitempty"`
	// The setting class name and property name.
	Setting string `json:"setting,omitempty"`
	// Name of the setting.
	SettingName string `json:"settingName,omitempty"`
	// Number of unknown devices
	UnknownDeviceCount            int32                                        `json:"unknownDeviceCount,omitempty"`
	DeviceComplianceSettingStates []MicrosoftGraphDeviceComplianceSettingState `json:"deviceComplianceSettingStates,omitempty"`
}