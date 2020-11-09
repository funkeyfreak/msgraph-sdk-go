/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// List struct for List
type List struct {
	DisplayName   string                           `json:"displayName,omitempty"`
	List          interface{}                      `json:"list,omitempty"`
	SharepointIds interface{}                      `json:"sharepointIds,omitempty"`
	System        interface{}                      `json:"system,omitempty"`
	Columns       []MicrosoftGraphColumnDefinition `json:"columns,omitempty"`
	ContentTypes  []MicrosoftGraphContentType      `json:"contentTypes,omitempty"`
	Drive         interface{}                      `json:"drive,omitempty"`
	Items         []MicrosoftGraphListItem         `json:"items,omitempty"`
	Subscriptions []MicrosoftGraphSubscription     `json:"subscriptions,omitempty"`
}