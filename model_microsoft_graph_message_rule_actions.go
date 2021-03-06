/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphMessageRuleActions struct for MicrosoftGraphMessageRuleActions
type MicrosoftGraphMessageRuleActions struct {
	AssignCategories      []string      `json:"assignCategories,omitempty"`
	CopyToFolder          string        `json:"copyToFolder,omitempty"`
	Delete                bool          `json:"delete,omitempty"`
	ForwardAsAttachmentTo []interface{} `json:"forwardAsAttachmentTo,omitempty"`
	ForwardTo             []interface{} `json:"forwardTo,omitempty"`
	MarkAsRead            bool          `json:"markAsRead,omitempty"`
	MarkImportance        interface{}   `json:"markImportance,omitempty"`
	MoveToFolder          string        `json:"moveToFolder,omitempty"`
	PermanentDelete       bool          `json:"permanentDelete,omitempty"`
	RedirectTo            []interface{} `json:"redirectTo,omitempty"`
	StopProcessingRules   bool          `json:"stopProcessingRules,omitempty"`
}
