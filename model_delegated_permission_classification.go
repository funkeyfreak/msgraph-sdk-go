/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// DelegatedPermissionClassification struct for DelegatedPermissionClassification
type DelegatedPermissionClassification struct {
	Classification interface{} `json:"classification,omitempty"`
	PermissionId   string      `json:"permissionId,omitempty"`
	PermissionName string      `json:"permissionName,omitempty"`
}
