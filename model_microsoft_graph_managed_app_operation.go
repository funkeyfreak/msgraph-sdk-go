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

// MicrosoftGraphManagedAppOperation struct for MicrosoftGraphManagedAppOperation
type MicrosoftGraphManagedAppOperation struct {
	Id string `json:"id,omitempty"`
	// The operation name.
	DisplayName string `json:"displayName,omitempty"`
	// The last time the app operation was modified.
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The current state of the operation
	State string `json:"state,omitempty"`
	// Version of the entity.
	Version string `json:"version,omitempty"`
}