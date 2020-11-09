/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphSensitivity the model 'MicrosoftGraphSensitivity'
type MicrosoftGraphSensitivity string

// List of microsoft.graph.sensitivity
const (
	MICROSOFTGRAPHSENSITIVITY_NORMAL       MicrosoftGraphSensitivity = "normal"
	MICROSOFTGRAPHSENSITIVITY_PERSONAL     MicrosoftGraphSensitivity = "personal"
	MICROSOFTGRAPHSENSITIVITY_PRIVATE      MicrosoftGraphSensitivity = "private"
	MICROSOFTGRAPHSENSITIVITY_CONFIDENTIAL MicrosoftGraphSensitivity = "confidential"
)