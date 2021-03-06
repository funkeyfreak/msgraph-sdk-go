/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// ResourceOperation Describes the resourceOperation resource (entity) of the Microsoft Graph API (REST), which supports Intune workflows related to role-based access control (RBAC).
type ResourceOperation struct {
	// Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
	ActionName string `json:"actionName,omitempty"`
	// Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
	Description string `json:"description,omitempty"`
	// Name of the Resource this operation is performed on.
	ResourceName string `json:"resourceName,omitempty"`
}
