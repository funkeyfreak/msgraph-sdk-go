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

// MicrosoftGraphPost struct for MicrosoftGraphPost
type MicrosoftGraphPost struct {
	Id                            string                                            `json:"id,omitempty"`
	Categories                    []string                                          `json:"categories,omitempty"`
	ChangeKey                     string                                            `json:"changeKey,omitempty"`
	CreatedDateTime               time.Time                                         `json:"createdDateTime,omitempty"`
	LastModifiedDateTime          time.Time                                         `json:"lastModifiedDateTime,omitempty"`
	Body                          interface{}                                       `json:"body,omitempty"`
	ConversationId                string                                            `json:"conversationId,omitempty"`
	ConversationThreadId          string                                            `json:"conversationThreadId,omitempty"`
	From                          MicrosoftGraphRecipient                           `json:"from,omitempty"`
	HasAttachments                bool                                              `json:"hasAttachments,omitempty"`
	NewParticipants               []MicrosoftGraphRecipient                         `json:"newParticipants,omitempty"`
	ReceivedDateTime              time.Time                                         `json:"receivedDateTime,omitempty"`
	Sender                        interface{}                                       `json:"sender,omitempty"`
	Attachments                   []MicrosoftGraphAttachment                        `json:"attachments,omitempty"`
	Extensions                    []MicrosoftGraphExtension                         `json:"extensions,omitempty"`
	InReplyTo                     interface{}                                       `json:"inReplyTo,omitempty"`
	MultiValueExtendedProperties  []MicrosoftGraphMultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	SingleValueExtendedProperties []MicrosoftGraphSingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}
