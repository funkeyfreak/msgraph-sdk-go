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

// DeviceManagementTroubleshootingEvent Event representing an general failure.
type DeviceManagementTroubleshootingEvent struct {
	// Id used for tracing the failure in the service.
	CorrelationId string `json:"correlationId,omitempty"`
	// Time when the event occurred .
	EventDateTime time.Time `json:"eventDateTime,omitempty"`
}
