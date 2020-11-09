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

// MicrosoftGraphStsPolicy struct for MicrosoftGraphStsPolicy
type MicrosoftGraphStsPolicy struct {
	Id                    string                          `json:"id,omitempty"`
	DeletedDateTime       time.Time                       `json:"deletedDateTime,omitempty"`
	Description           string                          `json:"description,omitempty"`
	DisplayName           string                          `json:"displayName,omitempty"`
	Definition            []string                        `json:"definition,omitempty"`
	IsOrganizationDefault bool                            `json:"isOrganizationDefault,omitempty"`
	AppliesTo             []MicrosoftGraphDirectoryObject `json:"appliesTo,omitempty"`
}
