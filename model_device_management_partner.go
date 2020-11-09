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

// DeviceManagementPartner Entity which represents a connection to device management partner.
type DeviceManagementPartner struct {
	// Partner display name
	DisplayName string `json:"displayName,omitempty"`
	// Whether device management partner is configured or not
	IsConfigured bool `json:"isConfigured,omitempty"`
	// Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
	LastHeartbeatDateTime time.Time `json:"lastHeartbeatDateTime,omitempty"`
	// Partner App type
	PartnerAppType interface{} `json:"partnerAppType,omitempty"`
	// Partner state of this tenant
	PartnerState interface{} `json:"partnerState,omitempty"`
	// Partner Single tenant App id
	SingleTenantAppId string `json:"singleTenantAppId,omitempty"`
	// DateTime in UTC when PartnerDevices will be marked as NonCompliant
	WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime time.Time `json:"whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime,omitempty"`
	// DateTime in UTC when PartnerDevices will be removed
	WhenPartnerDevicesWillBeRemovedDateTime time.Time `json:"whenPartnerDevicesWillBeRemovedDateTime,omitempty"`
}
