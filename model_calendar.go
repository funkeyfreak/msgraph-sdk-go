/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// Calendar struct for Calendar
type Calendar struct {
	AllowedOnlineMeetingProviders []interface{}                                     `json:"allowedOnlineMeetingProviders,omitempty"`
	CanEdit                       bool                                              `json:"canEdit,omitempty"`
	CanShare                      bool                                              `json:"canShare,omitempty"`
	CanViewPrivateItems           bool                                              `json:"canViewPrivateItems,omitempty"`
	ChangeKey                     string                                            `json:"changeKey,omitempty"`
	Color                         interface{}                                       `json:"color,omitempty"`
	DefaultOnlineMeetingProvider  interface{}                                       `json:"defaultOnlineMeetingProvider,omitempty"`
	HexColor                      string                                            `json:"hexColor,omitempty"`
	IsDefaultCalendar             bool                                              `json:"isDefaultCalendar,omitempty"`
	IsRemovable                   bool                                              `json:"isRemovable,omitempty"`
	IsTallyingResponses           bool                                              `json:"isTallyingResponses,omitempty"`
	Name                          string                                            `json:"name,omitempty"`
	Owner                         interface{}                                       `json:"owner,omitempty"`
	CalendarPermissions           []MicrosoftGraphCalendarPermission                `json:"calendarPermissions,omitempty"`
	CalendarView                  []MicrosoftGraphEvent                             `json:"calendarView,omitempty"`
	Events                        []MicrosoftGraphEvent                             `json:"events,omitempty"`
	MultiValueExtendedProperties  []MicrosoftGraphMultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	SingleValueExtendedProperties []MicrosoftGraphSingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}
