/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// DeviceCompliancePolicyDeviceStateSummary struct for DeviceCompliancePolicyDeviceStateSummary
type DeviceCompliancePolicyDeviceStateSummary struct {
	// Number of compliant devices
	CompliantDeviceCount int32 `json:"compliantDeviceCount,omitempty"`
	// Number of devices that have compliance managed by System Center Configuration Manager
	ConfigManagerCount int32 `json:"configManagerCount,omitempty"`
	// Number of conflict devices
	ConflictDeviceCount int32 `json:"conflictDeviceCount,omitempty"`
	// Number of error devices
	ErrorDeviceCount int32 `json:"errorDeviceCount,omitempty"`
	// Number of devices that are in grace period
	InGracePeriodCount int32 `json:"inGracePeriodCount,omitempty"`
	// Number of NonCompliant devices
	NonCompliantDeviceCount int32 `json:"nonCompliantDeviceCount,omitempty"`
	// Number of not applicable devices
	NotApplicableDeviceCount int32 `json:"notApplicableDeviceCount,omitempty"`
	// Number of remediated devices
	RemediatedDeviceCount int32 `json:"remediatedDeviceCount,omitempty"`
	// Number of unknown devices
	UnknownDeviceCount int32 `json:"unknownDeviceCount,omitempty"`
}
