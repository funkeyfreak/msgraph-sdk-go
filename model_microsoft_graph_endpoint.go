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

// MicrosoftGraphEndpoint struct for MicrosoftGraphEndpoint
type MicrosoftGraphEndpoint struct {
	Id                 string    `json:"id,omitempty"`
	DeletedDateTime    time.Time `json:"deletedDateTime,omitempty"`
	Capability         string    `json:"capability,omitempty"`
	ProviderId         string    `json:"providerId,omitempty"`
	ProviderName       string    `json:"providerName,omitempty"`
	ProviderResourceId string    `json:"providerResourceId,omitempty"`
	Uri                string    `json:"uri,omitempty"`
}
