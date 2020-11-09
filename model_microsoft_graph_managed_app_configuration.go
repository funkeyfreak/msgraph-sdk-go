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

// MicrosoftGraphManagedAppConfiguration struct for MicrosoftGraphManagedAppConfiguration
type MicrosoftGraphManagedAppConfiguration struct {
	Id string `json:"id,omitempty"`
	// The date and time the policy was created.
	CreatedDateTime time.Time `json:"createdDateTime,omitempty"`
	// The policy's description.
	Description string `json:"description,omitempty"`
	// Policy display name.
	DisplayName string `json:"displayName,omitempty"`
	// Last time the policy was modified.
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// Version of the entity.
	Version string `json:"version,omitempty"`
	// A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
	CustomSettings []MicrosoftGraphKeyValuePair `json:"customSettings,omitempty"`
}