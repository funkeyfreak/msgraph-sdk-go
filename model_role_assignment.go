/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// RoleAssignment The Role Assignment resource. Role assignments tie together a role definition with members and scopes. There can be one or more role assignments per role. This applies to custom and built-in roles.
type RoleAssignment struct {
	// Description of the Role Assignment.
	Description string `json:"description,omitempty"`
	// The display or friendly name of the role Assignment.
	DisplayName string `json:"displayName,omitempty"`
	// List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
	ResourceScopes []string    `json:"resourceScopes,omitempty"`
	RoleDefinition interface{} `json:"roleDefinition,omitempty"`
}
