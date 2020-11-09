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

// VppToken You purchase multiple licenses for iOS apps through the Apple Volume Purchase Program for Business or Education. This involves setting up an Apple VPP account from the Apple website and uploading the Apple VPP Business or Education token to Intune. You can then synchronize your volume purchase information with Intune and track your volume-purchased app use. You can upload multiple Apple VPP Business or Education tokens.
type VppToken struct {
	// The apple Id associated with the given Apple Volume Purchase Program Token.
	AppleId string `json:"appleId,omitempty"`
	// Whether or not apps for the VPP token will be automatically updated.
	AutomaticallyUpdateApps bool `json:"automaticallyUpdateApps,omitempty"`
	// Whether or not apps for the VPP token will be automatically updated.
	CountryOrRegion string `json:"countryOrRegion,omitempty"`
	// The expiration date time of the Apple Volume Purchase Program Token.
	ExpirationDateTime time.Time `json:"expirationDateTime,omitempty"`
	// Last modification date time associated with the Apple Volume Purchase Program Token.
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
	LastSyncDateTime time.Time `json:"lastSyncDateTime,omitempty"`
	// Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: `none`, `inProgress`, `completed`, `failed`.
	LastSyncStatus interface{} `json:"lastSyncStatus,omitempty"`
	// The organization associated with the Apple Volume Purchase Program Token
	OrganizationName string `json:"organizationName,omitempty"`
	// Current state of the Apple Volume Purchase Program Token. Possible values are: `unknown`, `valid`, `expired`, `invalid`, `assignedToExternalMDM`.
	State interface{} `json:"state,omitempty"`
	// The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
	Token string `json:"token,omitempty"`
	// The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: `business`, `education`.
	VppTokenAccountType interface{} `json:"vppTokenAccountType,omitempty"`
}