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

// MicrosoftGraphListItem struct for MicrosoftGraphListItem
type MicrosoftGraphListItem struct {
	Id                   string                          `json:"id,omitempty"`
	CreatedBy            interface{}                     `json:"createdBy,omitempty"`
	CreatedDateTime      time.Time                       `json:"createdDateTime,omitempty"`
	Description          string                          `json:"description,omitempty"`
	ETag                 string                          `json:"eTag,omitempty"`
	LastModifiedBy       interface{}                     `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime time.Time                       `json:"lastModifiedDateTime,omitempty"`
	Name                 string                          `json:"name,omitempty"`
	ParentReference      interface{}                     `json:"parentReference,omitempty"`
	WebUrl               string                          `json:"webUrl,omitempty"`
	CreatedByUser        interface{}                     `json:"createdByUser,omitempty"`
	LastModifiedByUser   interface{}                     `json:"lastModifiedByUser,omitempty"`
	ContentType          interface{}                     `json:"contentType,omitempty"`
	SharepointIds        interface{}                     `json:"sharepointIds,omitempty"`
	Analytics            interface{}                     `json:"analytics,omitempty"`
	DriveItem            interface{}                     `json:"driveItem,omitempty"`
	Fields               interface{}                     `json:"fields,omitempty"`
	Versions             []MicrosoftGraphListItemVersion `json:"versions,omitempty"`
}
