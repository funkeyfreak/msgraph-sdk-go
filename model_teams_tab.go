/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// TeamsTab struct for TeamsTab
type TeamsTab struct {
	Configuration interface{} `json:"configuration,omitempty"`
	DisplayName   string      `json:"displayName,omitempty"`
	WebUrl        string      `json:"webUrl,omitempty"`
	TeamsApp      interface{} `json:"teamsApp,omitempty"`
}