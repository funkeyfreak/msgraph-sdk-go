/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphPreAuthorizedApplication struct for MicrosoftGraphPreAuthorizedApplication
type MicrosoftGraphPreAuthorizedApplication struct {
	AppId                  string   `json:"appId,omitempty"`
	DelegatedPermissionIds []string `json:"delegatedPermissionIds,omitempty"`
}
