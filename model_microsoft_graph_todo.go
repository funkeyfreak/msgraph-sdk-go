/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphTodo struct for MicrosoftGraphTodo
type MicrosoftGraphTodo struct {
	Id    string                       `json:"id,omitempty"`
	Lists []MicrosoftGraphTodoTaskList `json:"lists,omitempty"`
}