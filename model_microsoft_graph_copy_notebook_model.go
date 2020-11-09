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

// MicrosoftGraphCopyNotebookModel struct for MicrosoftGraphCopyNotebookModel
type MicrosoftGraphCopyNotebookModel struct {
	CreatedBy              string      `json:"createdBy,omitempty"`
	CreatedByIdentity      interface{} `json:"createdByIdentity,omitempty"`
	CreatedTime            time.Time   `json:"createdTime,omitempty"`
	Id                     string      `json:"id,omitempty"`
	IsDefault              bool        `json:"isDefault,omitempty"`
	IsShared               bool        `json:"isShared,omitempty"`
	LastModifiedBy         string      `json:"lastModifiedBy,omitempty"`
	LastModifiedByIdentity interface{} `json:"lastModifiedByIdentity,omitempty"`
	LastModifiedTime       time.Time   `json:"lastModifiedTime,omitempty"`
	Links                  interface{} `json:"links,omitempty"`
	Name                   string      `json:"name,omitempty"`
	SectionGroupsUrl       string      `json:"sectionGroupsUrl,omitempty"`
	SectionsUrl            string      `json:"sectionsUrl,omitempty"`
	Self                   string      `json:"self,omitempty"`
	UserRole               interface{} `json:"userRole,omitempty"`
}