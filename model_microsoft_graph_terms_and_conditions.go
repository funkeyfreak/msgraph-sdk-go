/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

import (
	"time"
)

// MicrosoftGraphTermsAndConditions struct for MicrosoftGraphTermsAndConditions
type MicrosoftGraphTermsAndConditions struct {
	Id string `json:"id,omitempty"`
	// Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&C policy. This is shown to the user on prompts to accept the T&C policy.
	AcceptanceStatement string `json:"acceptanceStatement,omitempty"`
	// Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&C policy.
	BodyText string `json:"bodyText,omitempty"`
	// DateTime the object was created.
	CreatedDateTime time.Time `json:"createdDateTime,omitempty"`
	// Administrator-supplied description of the T&C policy.
	Description string `json:"description,omitempty"`
	// Administrator-supplied name for the T&C policy.
	DisplayName string `json:"displayName,omitempty"`
	// DateTime the object was last modified.
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&C policy.
	Title string `json:"title,omitempty"`
	// Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&C policy.
	Version            int32                                              `json:"version,omitempty"`
	AcceptanceStatuses []MicrosoftGraphTermsAndConditionsAcceptanceStatus `json:"acceptanceStatuses,omitempty"`
	Assignments        []MicrosoftGraphTermsAndConditionsAssignment       `json:"assignments,omitempty"`
}
