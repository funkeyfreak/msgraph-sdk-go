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

// MicrosoftGraphUserActivity struct for MicrosoftGraphUserActivity
type MicrosoftGraphUserActivity struct {
	Id                   string                              `json:"id,omitempty"`
	ActivationUrl        string                              `json:"activationUrl,omitempty"`
	ActivitySourceHost   string                              `json:"activitySourceHost,omitempty"`
	AppActivityId        string                              `json:"appActivityId,omitempty"`
	AppDisplayName       string                              `json:"appDisplayName,omitempty"`
	ContentInfo          interface{}                         `json:"contentInfo,omitempty"`
	ContentUrl           string                              `json:"contentUrl,omitempty"`
	CreatedDateTime      time.Time                           `json:"createdDateTime,omitempty"`
	ExpirationDateTime   time.Time                           `json:"expirationDateTime,omitempty"`
	FallbackUrl          string                              `json:"fallbackUrl,omitempty"`
	LastModifiedDateTime time.Time                           `json:"lastModifiedDateTime,omitempty"`
	Status               interface{}                         `json:"status,omitempty"`
	UserTimezone         string                              `json:"userTimezone,omitempty"`
	VisualElements       MicrosoftGraphVisualInfo            `json:"visualElements,omitempty"`
	HistoryItems         []MicrosoftGraphActivityHistoryItem `json:"historyItems,omitempty"`
}
