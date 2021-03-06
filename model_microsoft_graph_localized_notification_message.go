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

// MicrosoftGraphLocalizedNotificationMessage struct for MicrosoftGraphLocalizedNotificationMessage
type MicrosoftGraphLocalizedNotificationMessage struct {
	Id string `json:"id,omitempty"`
	// Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
	IsDefault bool `json:"isDefault,omitempty"`
	// DateTime the object was last modified.
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The Locale for which this message is destined.
	Locale string `json:"locale,omitempty"`
	// The Message Template content.
	MessageTemplate string `json:"messageTemplate,omitempty"`
	// The Message Template Subject.
	Subject string `json:"subject,omitempty"`
}
