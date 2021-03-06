/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// ExtensionProperty Represents an Azure Active Directory object. The directoryObject type is the base type for many other directory entity types.
type ExtensionProperty struct {
	AppDisplayName         string   `json:"appDisplayName,omitempty"`
	DataType               string   `json:"dataType,omitempty"`
	IsSyncedFromOnPremises bool     `json:"isSyncedFromOnPremises,omitempty"`
	Name                   string   `json:"name,omitempty"`
	TargetObjects          []string `json:"targetObjects,omitempty"`
}
