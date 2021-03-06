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

// MicrosoftGraphKeyCredential struct for MicrosoftGraphKeyCredential
type MicrosoftGraphKeyCredential struct {
	CustomKeyIdentifier string    `json:"customKeyIdentifier,omitempty"`
	DisplayName         string    `json:"displayName,omitempty"`
	EndDateTime         time.Time `json:"endDateTime,omitempty"`
	Key                 string    `json:"key,omitempty"`
	KeyId               string    `json:"keyId,omitempty"`
	StartDateTime       time.Time `json:"startDateTime,omitempty"`
	Type                string    `json:"type,omitempty"`
	Usage               string    `json:"usage,omitempty"`
}
