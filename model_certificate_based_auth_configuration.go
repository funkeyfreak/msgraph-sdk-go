/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// CertificateBasedAuthConfiguration struct for CertificateBasedAuthConfiguration
type CertificateBasedAuthConfiguration struct {
	CertificateAuthorities []MicrosoftGraphCertificateAuthority `json:"certificateAuthorities,omitempty"`
}