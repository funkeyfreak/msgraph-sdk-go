/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// EducationSchool struct for EducationSchool
type EducationSchool struct {
	Address             interface{}                    `json:"address,omitempty"`
	CreatedBy           interface{}                    `json:"createdBy,omitempty"`
	ExternalId          string                         `json:"externalId,omitempty"`
	ExternalPrincipalId string                         `json:"externalPrincipalId,omitempty"`
	Fax                 string                         `json:"fax,omitempty"`
	HighestGrade        string                         `json:"highestGrade,omitempty"`
	LowestGrade         string                         `json:"lowestGrade,omitempty"`
	Phone               string                         `json:"phone,omitempty"`
	PrincipalEmail      string                         `json:"principalEmail,omitempty"`
	PrincipalName       string                         `json:"principalName,omitempty"`
	SchoolNumber        string                         `json:"schoolNumber,omitempty"`
	Classes             []MicrosoftGraphEducationClass `json:"classes,omitempty"`
	Users               []MicrosoftGraphEducationUser  `json:"users,omitempty"`
}
