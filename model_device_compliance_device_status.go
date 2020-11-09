/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

import (
	"time"
)

// DeviceComplianceDeviceStatus struct for DeviceComplianceDeviceStatus
type DeviceComplianceDeviceStatus struct {
	// The DateTime when device compliance grace period expires
	ComplianceGracePeriodExpirationDateTime time.Time `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	// Device name of the DevicePolicyStatus.
	DeviceDisplayName string `json:"deviceDisplayName,omitempty"`
	// The device model that is being reported
	DeviceModel string `json:"deviceModel,omitempty"`
	// Last modified date time of the policy report.
	LastReportedDateTime time.Time `json:"lastReportedDateTime,omitempty"`
	// Compliance status of the policy report.
	Status interface{} `json:"status,omitempty"`
	// The User Name that is being reported
	UserName string `json:"userName,omitempty"`
	// UserPrincipalName.
	UserPrincipalName string `json:"userPrincipalName,omitempty"`
}
