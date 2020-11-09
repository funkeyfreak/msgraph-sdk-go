/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// IosManagedAppProtection Policy used to configure detailed management settings targeted to specific security groups and for a specified set of apps on an iOS device
type IosManagedAppProtection struct {
	// Type of encryption which should be used for data in a managed app.
	AppDataEncryptionType interface{} `json:"appDataEncryptionType,omitempty"`
	// A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
	CustomBrowserProtocol string `json:"customBrowserProtocol,omitempty"`
	// Count of apps to which the current policy is deployed.
	DeployedAppCount int32 `json:"deployedAppCount,omitempty"`
	// Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
	FaceIdBlocked bool `json:"faceIdBlocked,omitempty"`
	// Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredSdkVersion string                           `json:"minimumRequiredSdkVersion,omitempty"`
	Apps                      []MicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
	DeploymentSummary         interface{}                      `json:"deploymentSummary,omitempty"`
}
