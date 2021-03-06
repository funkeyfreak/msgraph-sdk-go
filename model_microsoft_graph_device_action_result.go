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

// MicrosoftGraphDeviceActionResult Device action result
type MicrosoftGraphDeviceActionResult struct {
	// Action name
	ActionName string `json:"actionName,omitempty"`
	// State of the action
	ActionState interface{} `json:"actionState,omitempty"`
	// Time the action state was last updated
	LastUpdatedDateTime time.Time `json:"lastUpdatedDateTime,omitempty"`
	// Time the action was initiated
	StartDateTime time.Time `json:"startDateTime,omitempty"`
}
