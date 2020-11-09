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

// MicrosoftGraphSectionGroup struct for MicrosoftGraphSectionGroup
type MicrosoftGraphSectionGroup struct {
	Id                   string                         `json:"id,omitempty"`
	Self                 string                         `json:"self,omitempty"`
	CreatedDateTime      time.Time                      `json:"createdDateTime,omitempty"`
	CreatedBy            interface{}                    `json:"createdBy,omitempty"`
	DisplayName          string                         `json:"displayName,omitempty"`
	LastModifiedBy       interface{}                    `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime time.Time                      `json:"lastModifiedDateTime,omitempty"`
	SectionGroupsUrl     string                         `json:"sectionGroupsUrl,omitempty"`
	SectionsUrl          string                         `json:"sectionsUrl,omitempty"`
	ParentNotebook       interface{}                    `json:"parentNotebook,omitempty"`
	ParentSectionGroup   interface{}                    `json:"parentSectionGroup,omitempty"`
	SectionGroups        []MicrosoftGraphSectionGroup   `json:"sectionGroups,omitempty"`
	Sections             []MicrosoftGraphOnenoteSection `json:"sections,omitempty"`
}
