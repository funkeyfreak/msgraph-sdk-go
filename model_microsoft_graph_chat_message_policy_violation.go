/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphChatMessagePolicyViolation struct for MicrosoftGraphChatMessagePolicyViolation
type MicrosoftGraphChatMessagePolicyViolation struct {
	DlpAction         interface{} `json:"dlpAction,omitempty"`
	JustificationText string      `json:"justificationText,omitempty"`
	PolicyTip         interface{} `json:"policyTip,omitempty"`
	UserAction        interface{} `json:"userAction,omitempty"`
	VerdictDetails    interface{} `json:"verdictDetails,omitempty"`
}
