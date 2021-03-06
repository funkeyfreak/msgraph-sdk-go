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

// ApplePushNotificationCertificate Apple push notification certificate.
type ApplePushNotificationCertificate struct {
	// Apple Id of the account used to create the MDM push certificate.
	AppleIdentifier string `json:"appleIdentifier,omitempty"`
	Certificate     string `json:"certificate,omitempty"`
	// The expiration date and time for Apple push notification certificate.
	ExpirationDateTime time.Time `json:"expirationDateTime,omitempty"`
	// Last modified date and time for Apple push notification certificate.
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// Topic Id.
	TopicIdentifier string `json:"topicIdentifier,omitempty"`
}
