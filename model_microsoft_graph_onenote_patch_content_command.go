/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphOnenotePatchContentCommand struct for MicrosoftGraphOnenotePatchContentCommand
type MicrosoftGraphOnenotePatchContentCommand struct {
	Action   interface{} `json:"action,omitempty"`
	Content  string      `json:"content,omitempty"`
	Position interface{} `json:"position,omitempty"`
	Target   string      `json:"target,omitempty"`
}