/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphSchemaExtension struct for MicrosoftGraphSchemaExtension
type MicrosoftGraphSchemaExtension struct {
	Id          string                                  `json:"id,omitempty"`
	Description string                                  `json:"description,omitempty"`
	Owner       string                                  `json:"owner,omitempty"`
	Properties  []MicrosoftGraphExtensionSchemaProperty `json:"properties,omitempty"`
	Status      string                                  `json:"status,omitempty"`
	TargetTypes []string                                `json:"targetTypes,omitempty"`
}
