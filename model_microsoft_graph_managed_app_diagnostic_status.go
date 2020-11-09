/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphManagedAppDiagnosticStatus Represents diagnostics status.
type MicrosoftGraphManagedAppDiagnosticStatus struct {
	// Instruction on how to mitigate a failed validation
	MitigationInstruction string `json:"mitigationInstruction,omitempty"`
	// The state of the operation
	State string `json:"state,omitempty"`
	// The validation friendly name
	ValidationName string `json:"validationName,omitempty"`
}
