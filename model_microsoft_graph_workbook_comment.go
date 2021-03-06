/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWorkbookComment struct for MicrosoftGraphWorkbookComment
type MicrosoftGraphWorkbookComment struct {
	Id          string                               `json:"id,omitempty"`
	Content     string                               `json:"content,omitempty"`
	ContentType string                               `json:"contentType,omitempty"`
	Replies     []MicrosoftGraphWorkbookCommentReply `json:"replies,omitempty"`
}
