/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// ItemAnalytics struct for ItemAnalytics
type ItemAnalytics struct {
	AllTime           interface{}                      `json:"allTime,omitempty"`
	ItemActivityStats []MicrosoftGraphItemActivityStat `json:"itemActivityStats,omitempty"`
	LastSevenDays     interface{}                      `json:"lastSevenDays,omitempty"`
}
