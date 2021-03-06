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

// MicrosoftGraphShiftActivity struct for MicrosoftGraphShiftActivity
type MicrosoftGraphShiftActivity struct {
	Code          string      `json:"code,omitempty"`
	DisplayName   string      `json:"displayName,omitempty"`
	EndDateTime   time.Time   `json:"endDateTime,omitempty"`
	IsPaid        bool        `json:"isPaid,omitempty"`
	StartDateTime time.Time   `json:"startDateTime,omitempty"`
	Theme         interface{} `json:"theme,omitempty"`
}
