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

// MicrosoftGraphPendingContentUpdate struct for MicrosoftGraphPendingContentUpdate
type MicrosoftGraphPendingContentUpdate struct {
	QueuedDateTime time.Time `json:"queuedDateTime,omitempty"`
}
