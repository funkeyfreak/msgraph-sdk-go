/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// WindowsInformationProtectionAppLearningSummary Windows Information Protection AppLearning Summary entity.
type WindowsInformationProtectionAppLearningSummary struct {
	// Application Name
	ApplicationName string `json:"applicationName,omitempty"`
	// Application Type
	ApplicationType interface{} `json:"applicationType,omitempty"`
	// Device Count
	DeviceCount int32 `json:"deviceCount,omitempty"`
}
