/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// TimeOff struct for TimeOff
type TimeOff struct {
	DraftTimeOff  interface{} `json:"draftTimeOff,omitempty"`
	SharedTimeOff interface{} `json:"sharedTimeOff,omitempty"`
	UserId        string      `json:"userId,omitempty"`
}