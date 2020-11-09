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

// MicrosoftGraphOperation struct for MicrosoftGraphOperation
type MicrosoftGraphOperation struct {
	Id                 string      `json:"id,omitempty"`
	CreatedDateTime    time.Time   `json:"createdDateTime,omitempty"`
	LastActionDateTime time.Time   `json:"lastActionDateTime,omitempty"`
	Status             interface{} `json:"status,omitempty"`
}
