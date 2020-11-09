/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	CallbackUri        string                 `json:"callbackUri,omitempty"`
	MediaConfig        map[string]interface{} `json:"mediaConfig,omitempty"`
	AcceptedModalities []interface{}          `json:"acceptedModalities,omitempty"`
}
