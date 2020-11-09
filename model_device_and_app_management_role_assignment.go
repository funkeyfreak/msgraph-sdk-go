/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// DeviceAndAppManagementRoleAssignment The Role Assignment resource. Role assignments tie together a role definition with members and scopes. There can be one or more role assignments per role. This applies to custom and built-in roles.
type DeviceAndAppManagementRoleAssignment struct {
	// The list of ids of role member security groups. These are IDs from Azure Active Directory.
	Members []string `json:"members,omitempty"`
}
