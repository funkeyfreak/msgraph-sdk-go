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

// Attachment struct for Attachment
type Attachment struct {
	ContentType          string    `json:"contentType,omitempty"`
	IsInline             bool      `json:"isInline,omitempty"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	Name                 string    `json:"name,omitempty"`
	Size                 int32     `json:"size,omitempty"`
}
