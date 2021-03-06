/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphCalendarPermission struct for MicrosoftGraphCalendarPermission
type MicrosoftGraphCalendarPermission struct {
	Id                   string        `json:"id,omitempty"`
	AllowedRoles         []interface{} `json:"allowedRoles,omitempty"`
	EmailAddress         interface{}   `json:"emailAddress,omitempty"`
	IsInsideOrganization bool          `json:"isInsideOrganization,omitempty"`
	IsRemovable          bool          `json:"isRemovable,omitempty"`
	Role                 interface{}   `json:"role,omitempty"`
}
