/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphConditionalAccessGrantControls struct for MicrosoftGraphConditionalAccessGrantControls
type MicrosoftGraphConditionalAccessGrantControls struct {
	BuiltInControls             []interface{} `json:"builtInControls,omitempty"`
	CustomAuthenticationFactors []string      `json:"customAuthenticationFactors,omitempty"`
	Operator                    string        `json:"operator,omitempty"`
	TermsOfUse                  []string      `json:"termsOfUse,omitempty"`
}