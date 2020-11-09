/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// Onenote struct for Onenote
type Onenote struct {
	Notebooks     []MicrosoftGraphNotebook         `json:"notebooks,omitempty"`
	Operations    []MicrosoftGraphOnenoteOperation `json:"operations,omitempty"`
	Pages         []MicrosoftGraphOnenotePage      `json:"pages,omitempty"`
	Resources     []MicrosoftGraphOnenoteResource  `json:"resources,omitempty"`
	SectionGroups []MicrosoftGraphSectionGroup     `json:"sectionGroups,omitempty"`
	Sections      []MicrosoftGraphOnenoteSection   `json:"sections,omitempty"`
}
