/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// Notebook struct for Notebook
type Notebook struct {
	IsDefault        bool                           `json:"isDefault,omitempty"`
	IsShared         bool                           `json:"isShared,omitempty"`
	Links            interface{}                    `json:"links,omitempty"`
	SectionGroupsUrl string                         `json:"sectionGroupsUrl,omitempty"`
	SectionsUrl      string                         `json:"sectionsUrl,omitempty"`
	UserRole         interface{}                    `json:"userRole,omitempty"`
	SectionGroups    []MicrosoftGraphSectionGroup   `json:"sectionGroups,omitempty"`
	Sections         []MicrosoftGraphOnenoteSection `json:"sections,omitempty"`
}
