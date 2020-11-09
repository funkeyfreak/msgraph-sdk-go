/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

import (
	"time"
)

// Conversation struct for Conversation
type Conversation struct {
	HasAttachments        bool                               `json:"hasAttachments,omitempty"`
	LastDeliveredDateTime time.Time                          `json:"lastDeliveredDateTime,omitempty"`
	Preview               string                             `json:"preview,omitempty"`
	Topic                 string                             `json:"topic,omitempty"`
	UniqueSenders         []string                           `json:"uniqueSenders,omitempty"`
	Threads               []MicrosoftGraphConversationThread `json:"threads,omitempty"`
}