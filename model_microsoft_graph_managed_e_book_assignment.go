/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphManagedEBookAssignment struct for MicrosoftGraphManagedEBookAssignment
type MicrosoftGraphManagedEBookAssignment struct {
	Id string `json:"id,omitempty"`
	// The install intent for eBook.
	InstallIntent interface{} `json:"installIntent,omitempty"`
	// The assignment target for eBook.
	Target interface{} `json:"target,omitempty"`
}