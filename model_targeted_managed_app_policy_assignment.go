/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// TargetedManagedAppPolicyAssignment The type for deployment of groups or apps.
type TargetedManagedAppPolicyAssignment struct {
	// Identifier for deployment to a group or app
	Target interface{} `json:"target,omitempty"`
}