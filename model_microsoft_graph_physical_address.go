/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphPhysicalAddress struct for MicrosoftGraphPhysicalAddress
type MicrosoftGraphPhysicalAddress struct {
	City            string `json:"city,omitempty"`
	CountryOrRegion string `json:"countryOrRegion,omitempty"`
	PostalCode      string `json:"postalCode,omitempty"`
	State           string `json:"state,omitempty"`
	Street          string `json:"street,omitempty"`
}
