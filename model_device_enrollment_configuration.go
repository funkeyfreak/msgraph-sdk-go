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

// DeviceEnrollmentConfiguration The Base Class of Device Enrollment Configuration
type DeviceEnrollmentConfiguration struct {
	// Created date time in UTC of the device enrollment configuration
	CreatedDateTime time.Time `json:"createdDateTime,omitempty"`
	// The description of the device enrollment configuration
	Description string `json:"description,omitempty"`
	// The display name of the device enrollment configuration
	DisplayName string `json:"displayName,omitempty"`
	// Last modified date time in UTC of the device enrollment configuration
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject only to the configuration with the lowest priority value.
	Priority int32 `json:"priority,omitempty"`
	// The version of the device enrollment configuration
	Version     int32                                             `json:"version,omitempty"`
	Assignments []MicrosoftGraphEnrollmentConfigurationAssignment `json:"assignments,omitempty"`
}
