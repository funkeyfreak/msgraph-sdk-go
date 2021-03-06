/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// CalendarPermission struct for CalendarPermission
type CalendarPermission struct {
	AllowedRoles         []interface{} `json:"allowedRoles,omitempty"`
	EmailAddress         interface{}   `json:"emailAddress,omitempty"`
	IsInsideOrganization bool          `json:"isInsideOrganization,omitempty"`
	IsRemovable          bool          `json:"isRemovable,omitempty"`
	Role                 interface{}   `json:"role,omitempty"`
}
