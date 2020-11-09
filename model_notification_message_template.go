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

// NotificationMessageTemplate Notification messages are messages that are sent to end users who are determined to be not-compliant with the compliance policies defined by the administrator. Administrators choose notifications and configure them in the Intune Admin Console using the compliance policy creation page under the “Actions for non-compliance” section. Use the notificationMessageTemplate object to create your own custom notifications for administrators to choose while configuring actions for non-compliance.
type NotificationMessageTemplate struct {
	// The Message Template Branding Options. Branding is defined in the Intune Admin Console.
	BrandingOptions interface{} `json:"brandingOptions,omitempty"`
	// The default locale to fallback onto when the requested locale is not available.
	DefaultLocale string `json:"defaultLocale,omitempty"`
	// Display name for the Notification Message Template.
	DisplayName string `json:"displayName,omitempty"`
	// DateTime the object was last modified.
	LastModifiedDateTime          time.Time                                    `json:"lastModifiedDateTime,omitempty"`
	LocalizedNotificationMessages []MicrosoftGraphLocalizedNotificationMessage `json:"localizedNotificationMessages,omitempty"`
}
