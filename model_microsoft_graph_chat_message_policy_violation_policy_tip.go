/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphChatMessagePolicyViolationPolicyTip struct for MicrosoftGraphChatMessagePolicyViolationPolicyTip
type MicrosoftGraphChatMessagePolicyViolationPolicyTip struct {
	ComplianceUrl                string   `json:"complianceUrl,omitempty"`
	GeneralText                  string   `json:"generalText,omitempty"`
	MatchedConditionDescriptions []string `json:"matchedConditionDescriptions,omitempty"`
}
