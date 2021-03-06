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

// MicrosoftGraphExtensionProperty struct for MicrosoftGraphExtensionProperty
type MicrosoftGraphExtensionProperty struct {
	Id                     string    `json:"id,omitempty"`
	DeletedDateTime        time.Time `json:"deletedDateTime,omitempty"`
	AppDisplayName         string    `json:"appDisplayName,omitempty"`
	DataType               string    `json:"dataType,omitempty"`
	IsSyncedFromOnPremises bool      `json:"isSyncedFromOnPremises,omitempty"`
	Name                   string    `json:"name,omitempty"`
	TargetObjects          []string  `json:"targetObjects,omitempty"`
}
