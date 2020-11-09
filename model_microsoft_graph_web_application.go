/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphWebApplication struct for MicrosoftGraphWebApplication
type MicrosoftGraphWebApplication struct {
	HomePageUrl           string      `json:"homePageUrl,omitempty"`
	ImplicitGrantSettings interface{} `json:"implicitGrantSettings,omitempty"`
	LogoutUrl             string      `json:"logoutUrl,omitempty"`
	RedirectUris          []string    `json:"redirectUris,omitempty"`
}
