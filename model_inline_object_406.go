/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// InlineObject406 struct for InlineObject406
type InlineObject406 struct {
	ToRecipients []interface{} `json:"ToRecipients,omitempty"`
	Message      interface{}   `json:"Message,omitempty"`
	Comment      string        `json:"Comment,omitempty"`
}