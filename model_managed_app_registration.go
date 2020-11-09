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

// ManagedAppRegistration The ManagedAppEntity is the base entity type for all other entity types under app management workflow.
type ManagedAppRegistration struct {
	// The app package Identifier
	AppIdentifier interface{} `json:"appIdentifier,omitempty"`
	// App version
	ApplicationVersion string `json:"applicationVersion,omitempty"`
	// Date and time of creation
	CreatedDateTime time.Time `json:"createdDateTime,omitempty"`
	// Host device name
	DeviceName string `json:"deviceName,omitempty"`
	// App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
	DeviceTag string `json:"deviceTag,omitempty"`
	// Host device type
	DeviceType string `json:"deviceType,omitempty"`
	// Zero or more reasons an app registration is flagged. E.g. app running on rooted device
	FlaggedReasons []interface{} `json:"flaggedReasons,omitempty"`
	// Date and time of last the app synced with management service.
	LastSyncDateTime time.Time `json:"lastSyncDateTime,omitempty"`
	// App management SDK version
	ManagementSdkVersion string `json:"managementSdkVersion,omitempty"`
	// Operating System version
	PlatformVersion string `json:"platformVersion,omitempty"`
	// The user Id to who this app registration belongs.
	UserId string `json:"userId,omitempty"`
	// Version of the entity.
	Version          string                              `json:"version,omitempty"`
	AppliedPolicies  []MicrosoftGraphManagedAppPolicy    `json:"appliedPolicies,omitempty"`
	IntendedPolicies []MicrosoftGraphManagedAppPolicy    `json:"intendedPolicies,omitempty"`
	Operations       []MicrosoftGraphManagedAppOperation `json:"operations,omitempty"`
}
