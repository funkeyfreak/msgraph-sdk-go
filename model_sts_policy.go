/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// StsPolicy Represents an Azure Active Directory object. The directoryObject type is the base type for many other directory entity types.
type StsPolicy struct {
	Definition            []string                        `json:"definition,omitempty"`
	IsOrganizationDefault bool                            `json:"isOrganizationDefault,omitempty"`
	AppliesTo             []MicrosoftGraphDirectoryObject `json:"appliesTo,omitempty"`
}
