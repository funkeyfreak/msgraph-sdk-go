/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphPermissionGrantConditionSet struct for MicrosoftGraphPermissionGrantConditionSet
type MicrosoftGraphPermissionGrantConditionSet struct {
	Id                                          string      `json:"id,omitempty"`
	ClientApplicationIds                        []string    `json:"clientApplicationIds,omitempty"`
	ClientApplicationPublisherIds               []string    `json:"clientApplicationPublisherIds,omitempty"`
	ClientApplicationsFromVerifiedPublisherOnly bool        `json:"clientApplicationsFromVerifiedPublisherOnly,omitempty"`
	ClientApplicationTenantIds                  []string    `json:"clientApplicationTenantIds,omitempty"`
	PermissionClassification                    string      `json:"permissionClassification,omitempty"`
	Permissions                                 []string    `json:"permissions,omitempty"`
	PermissionType                              interface{} `json:"permissionType,omitempty"`
	ResourceApplication                         string      `json:"resourceApplication,omitempty"`
}
