/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// InlineObject416 struct for InlineObject416
type InlineObject416 struct {
	Attendees                 []interface{} `json:"attendees,omitempty"`
	LocationConstraint        interface{}   `json:"locationConstraint,omitempty"`
	TimeConstraint            interface{}   `json:"timeConstraint,omitempty"`
	MeetingDuration           string        `json:"meetingDuration,omitempty"`
	MaxCandidates             int32         `json:"maxCandidates,omitempty"`
	IsOrganizerOptional       bool          `json:"isOrganizerOptional,omitempty"`
	ReturnSuggestionReasons   bool          `json:"returnSuggestionReasons,omitempty"`
	MinimumAttendeePercentage interface{}   `json:"minimumAttendeePercentage,omitempty"`
}