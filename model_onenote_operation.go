/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// OnenoteOperation struct for OnenoteOperation
type OnenoteOperation struct {
	Error            interface{} `json:"error,omitempty"`
	PercentComplete  string      `json:"percentComplete,omitempty"`
	ResourceId       string      `json:"resourceId,omitempty"`
	ResourceLocation string      `json:"resourceLocation,omitempty"`
}
