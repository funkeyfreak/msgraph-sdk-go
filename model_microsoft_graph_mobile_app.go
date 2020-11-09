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

// MicrosoftGraphMobileApp struct for MicrosoftGraphMobileApp
type MicrosoftGraphMobileApp struct {
	Id string `json:"id,omitempty"`
	// The date and time the app was created.
	CreatedDateTime time.Time `json:"createdDateTime,omitempty"`
	// The description of the app.
	Description string `json:"description,omitempty"`
	// The developer of the app.
	Developer string `json:"developer,omitempty"`
	// The admin provided or imported title of the app.
	DisplayName string `json:"displayName,omitempty"`
	// The more information Url.
	InformationUrl string `json:"informationUrl,omitempty"`
	// The value indicating whether the app is marked as featured by the admin.
	IsFeatured bool `json:"isFeatured,omitempty"`
	// The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon interface{} `json:"largeIcon,omitempty"`
	// The date and time the app was last modified.
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// Notes for the app.
	Notes string `json:"notes,omitempty"`
	// The owner of the app.
	Owner string `json:"owner,omitempty"`
	// The privacy statement Url.
	PrivacyInformationUrl string `json:"privacyInformationUrl,omitempty"`
	// The publisher of the app.
	Publisher string `json:"publisher,omitempty"`
	// The publishing state for the app. The app cannot be assigned unless the app is published.
	PublishingState interface{}                         `json:"publishingState,omitempty"`
	Assignments     []MicrosoftGraphMobileAppAssignment `json:"assignments,omitempty"`
	Categories      []MicrosoftGraphMobileAppCategory   `json:"categories,omitempty"`
}
