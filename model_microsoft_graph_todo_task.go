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

// MicrosoftGraphTodoTask struct for MicrosoftGraphTodoTask
type MicrosoftGraphTodoTask struct {
	Id                       string                         `json:"id,omitempty"`
	Body                     interface{}                    `json:"body,omitempty"`
	BodyLastModifiedDateTime time.Time                      `json:"bodyLastModifiedDateTime,omitempty"`
	CompletedDateTime        interface{}                    `json:"completedDateTime,omitempty"`
	CreatedDateTime          time.Time                      `json:"createdDateTime,omitempty"`
	DueDateTime              interface{}                    `json:"dueDateTime,omitempty"`
	Importance               interface{}                    `json:"importance,omitempty"`
	IsReminderOn             bool                           `json:"isReminderOn,omitempty"`
	LastModifiedDateTime     time.Time                      `json:"lastModifiedDateTime,omitempty"`
	Recurrence               interface{}                    `json:"recurrence,omitempty"`
	ReminderDateTime         interface{}                    `json:"reminderDateTime,omitempty"`
	Status                   interface{}                    `json:"status,omitempty"`
	Title                    string                         `json:"title,omitempty"`
	Extensions               []MicrosoftGraphExtension      `json:"extensions,omitempty"`
	LinkedResources          []MicrosoftGraphLinkedResource `json:"linkedResources,omitempty"`
}
