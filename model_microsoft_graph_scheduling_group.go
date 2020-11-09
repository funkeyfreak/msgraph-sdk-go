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

// MicrosoftGraphSchedulingGroup struct for MicrosoftGraphSchedulingGroup
type MicrosoftGraphSchedulingGroup struct {
	Id                   string      `json:"id,omitempty"`
	CreatedDateTime      time.Time   `json:"createdDateTime,omitempty"`
	LastModifiedBy       interface{} `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime time.Time   `json:"lastModifiedDateTime,omitempty"`
	DisplayName          string      `json:"displayName,omitempty"`
	IsActive             bool        `json:"isActive,omitempty"`
	UserIds              []string    `json:"userIds,omitempty"`
}