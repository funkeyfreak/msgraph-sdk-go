/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// InlineObject629 struct for InlineObject629
type InlineObject629 struct {
	Schedules                []string    `json:"Schedules,omitempty"`
	EndTime                  interface{} `json:"EndTime,omitempty"`
	StartTime                interface{} `json:"StartTime,omitempty"`
	AvailabilityViewInterval int32       `json:"AvailabilityViewInterval,omitempty"`
}
