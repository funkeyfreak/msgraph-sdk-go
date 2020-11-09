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

// MicrosoftGraphSubscription struct for MicrosoftGraphSubscription
type MicrosoftGraphSubscription struct {
	Id                        string    `json:"id,omitempty"`
	ApplicationId             string    `json:"applicationId,omitempty"`
	ChangeType                string    `json:"changeType,omitempty"`
	ClientState               string    `json:"clientState,omitempty"`
	CreatorId                 string    `json:"creatorId,omitempty"`
	EncryptionCertificate     string    `json:"encryptionCertificate,omitempty"`
	EncryptionCertificateId   string    `json:"encryptionCertificateId,omitempty"`
	ExpirationDateTime        time.Time `json:"expirationDateTime,omitempty"`
	IncludeResourceData       bool      `json:"includeResourceData,omitempty"`
	LatestSupportedTlsVersion string    `json:"latestSupportedTlsVersion,omitempty"`
	LifecycleNotificationUrl  string    `json:"lifecycleNotificationUrl,omitempty"`
	NotificationUrl           string    `json:"notificationUrl,omitempty"`
	Resource                  string    `json:"resource,omitempty"`
}
