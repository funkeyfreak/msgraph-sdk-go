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

// MicrosoftGraphNamedLocation struct for MicrosoftGraphNamedLocation
type MicrosoftGraphNamedLocation struct {
	Id               string    `json:"id,omitempty"`
	CreatedDateTime  time.Time `json:"createdDateTime,omitempty"`
	DisplayName      string    `json:"displayName,omitempty"`
	ModifiedDateTime time.Time `json:"modifiedDateTime,omitempty"`
}
