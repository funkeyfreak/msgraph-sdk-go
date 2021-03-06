/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphFileSecurityState struct for MicrosoftGraphFileSecurityState
type MicrosoftGraphFileSecurityState struct {
	FileHash  interface{} `json:"fileHash,omitempty"`
	Name      string      `json:"name,omitempty"`
	Path      string      `json:"path,omitempty"`
	RiskScore string      `json:"riskScore,omitempty"`
}
