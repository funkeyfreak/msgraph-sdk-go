/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphOptionalClaims struct for MicrosoftGraphOptionalClaims
type MicrosoftGraphOptionalClaims struct {
	AccessToken []interface{} `json:"accessToken,omitempty"`
	IdToken     []interface{} `json:"idToken,omitempty"`
	Saml2Token  []interface{} `json:"saml2Token,omitempty"`
}
