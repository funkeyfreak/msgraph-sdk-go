/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// WorkforceIntegration struct for WorkforceIntegration
type WorkforceIntegration struct {
	ApiVersion        int32       `json:"apiVersion,omitempty"`
	DisplayName       string      `json:"displayName,omitempty"`
	Encryption        interface{} `json:"encryption,omitempty"`
	IsActive          bool        `json:"isActive,omitempty"`
	SupportedEntities interface{} `json:"supportedEntities,omitempty"`
	Url               string      `json:"url,omitempty"`
}