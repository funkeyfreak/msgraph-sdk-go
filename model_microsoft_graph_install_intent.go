/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphInstallIntent the model 'MicrosoftGraphInstallIntent'
type MicrosoftGraphInstallIntent string

// List of microsoft.graph.installIntent
const (
	MICROSOFTGRAPHINSTALLINTENT_AVAILABLE                    MicrosoftGraphInstallIntent = "available"
	MICROSOFTGRAPHINSTALLINTENT_REQUIRED                     MicrosoftGraphInstallIntent = "required"
	MICROSOFTGRAPHINSTALLINTENT_UNINSTALL                    MicrosoftGraphInstallIntent = "uninstall"
	MICROSOFTGRAPHINSTALLINTENT_AVAILABLE_WITHOUT_ENROLLMENT MicrosoftGraphInstallIntent = "availableWithoutEnrollment"
)
