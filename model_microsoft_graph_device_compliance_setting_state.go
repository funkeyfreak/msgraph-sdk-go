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

// MicrosoftGraphDeviceComplianceSettingState struct for MicrosoftGraphDeviceComplianceSettingState
type MicrosoftGraphDeviceComplianceSettingState struct {
	Id string `json:"id,omitempty"`
	// The DateTime when device compliance grace period expires
	ComplianceGracePeriodExpirationDateTime time.Time `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	// The Device Id that is being reported
	DeviceId string `json:"deviceId,omitempty"`
	// The device model that is being reported
	DeviceModel string `json:"deviceModel,omitempty"`
	// The Device Name that is being reported
	DeviceName string `json:"deviceName,omitempty"`
	// The setting class name and property name.
	Setting string `json:"setting,omitempty"`
	// The Setting Name that is being reported
	SettingName string `json:"settingName,omitempty"`
	// The compliance state of the setting
	State interface{} `json:"state,omitempty"`
	// The User email address that is being reported
	UserEmail string `json:"userEmail,omitempty"`
	// The user Id that is being reported
	UserId string `json:"userId,omitempty"`
	// The User Name that is being reported
	UserName string `json:"userName,omitempty"`
	// The User PrincipalName that is being reported
	UserPrincipalName string `json:"userPrincipalName,omitempty"`
}
