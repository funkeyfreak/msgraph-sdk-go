/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphEnrollmentConfigurationAssignment struct for MicrosoftGraphEnrollmentConfigurationAssignment
type MicrosoftGraphEnrollmentConfigurationAssignment struct {
	Id string `json:"id,omitempty"`
	// Represents an assignment to managed devices in the tenant
	Target interface{} `json:"target,omitempty"`
}
