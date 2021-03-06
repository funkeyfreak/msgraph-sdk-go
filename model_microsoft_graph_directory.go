/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphDirectory struct for MicrosoftGraphDirectory
type MicrosoftGraphDirectory struct {
	Id                  string                             `json:"id,omitempty"`
	AdministrativeUnits []MicrosoftGraphAdministrativeUnit `json:"administrativeUnits,omitempty"`
	DeletedItems        []MicrosoftGraphDirectoryObject    `json:"deletedItems,omitempty"`
}
