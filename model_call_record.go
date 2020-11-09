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

// CallRecord struct for CallRecord
type CallRecord struct {
	EndDateTime          time.Time                          `json:"endDateTime,omitempty"`
	JoinWebUrl           string                             `json:"joinWebUrl,omitempty"`
	LastModifiedDateTime time.Time                          `json:"lastModifiedDateTime,omitempty"`
	Modalities           []interface{}                      `json:"modalities,omitempty"`
	Organizer            interface{}                        `json:"organizer,omitempty"`
	Participants         []interface{}                      `json:"participants,omitempty"`
	StartDateTime        time.Time                          `json:"startDateTime,omitempty"`
	Type                 interface{}                        `json:"type,omitempty"`
	Version              int64                              `json:"version,omitempty"`
	Sessions             []MicrosoftGraphCallRecordsSession `json:"sessions,omitempty"`
}
