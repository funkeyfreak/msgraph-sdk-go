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

// MicrosoftGraphNotificationMessageTemplate struct for MicrosoftGraphNotificationMessageTemplate
type MicrosoftGraphNotificationMessageTemplate struct {
	Id string `json:"id,omitempty"`
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
