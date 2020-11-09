/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphObjectIdentity struct for MicrosoftGraphObjectIdentity
type MicrosoftGraphObjectIdentity struct {
	Issuer           string `json:"issuer,omitempty"`
	IssuerAssignedId string `json:"issuerAssignedId,omitempty"`
	SignInType       string `json:"signInType,omitempty"`
}
