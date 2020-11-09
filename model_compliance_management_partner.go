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

// ComplianceManagementPartner Compliance management partner for all platforms
type ComplianceManagementPartner struct {
	// User groups which enroll Android devices through partner.
	AndroidEnrollmentAssignments []interface{} `json:"androidEnrollmentAssignments,omitempty"`
	// Partner onboarded for Android devices.
	AndroidOnboarded bool `json:"androidOnboarded,omitempty"`
	// Partner display name
	DisplayName string `json:"displayName,omitempty"`
	// User groups which enroll ios devices through partner.
	IosEnrollmentAssignments []interface{} `json:"iosEnrollmentAssignments,omitempty"`
	// Partner onboarded for ios devices.
	IosOnboarded bool `json:"iosOnboarded,omitempty"`
	// Timestamp of last heartbeat after admin onboarded to the compliance management partner
	LastHeartbeatDateTime time.Time `json:"lastHeartbeatDateTime,omitempty"`
	// User groups which enroll Mac devices through partner.
	MacOsEnrollmentAssignments []interface{} `json:"macOsEnrollmentAssignments,omitempty"`
	// Partner onboarded for Mac devices.
	MacOsOnboarded bool `json:"macOsOnboarded,omitempty"`
	// Partner state of this tenant
	PartnerState interface{} `json:"partnerState,omitempty"`
}
