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

// MicrosoftGraphResponseStatus struct for MicrosoftGraphResponseStatus
type MicrosoftGraphResponseStatus struct {
	Response interface{} `json:"response,omitempty"`
	Time     time.Time   `json:"time,omitempty"`
}