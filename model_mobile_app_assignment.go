/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MobileAppAssignment A class containing the properties used for Group Assignment of a Mobile App.
type MobileAppAssignment struct {
	// The install intent defined by the admin.
	Intent interface{} `json:"intent,omitempty"`
	// The settings for target assignment defined by the admin.
	Settings interface{} `json:"settings,omitempty"`
	// The target group assignment defined by the admin.
	Target interface{} `json:"target,omitempty"`
}
