/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// InlineObject656 struct for InlineObject656
type InlineObject656 struct {
	Schedules                []string    `json:"Schedules,omitempty"`
	EndTime                  interface{} `json:"EndTime,omitempty"`
	StartTime                interface{} `json:"StartTime,omitempty"`
	AvailabilityViewInterval int32       `json:"AvailabilityViewInterval,omitempty"`
}
