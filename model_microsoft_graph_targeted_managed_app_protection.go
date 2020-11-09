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

// MicrosoftGraphTargetedManagedAppProtection struct for MicrosoftGraphTargetedManagedAppProtection
type MicrosoftGraphTargetedManagedAppProtection struct {
	Id string `json:"id,omitempty"`
	// The date and time the policy was created.
	CreatedDateTime time.Time `json:"createdDateTime,omitempty"`
	// The policy's description.
	Description string `json:"description,omitempty"`
	// Policy display name.
	DisplayName string `json:"displayName,omitempty"`
	// Last time the policy was modified.
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// Version of the entity.
	Version string `json:"version,omitempty"`
	// Data storage locations where a user may store managed data.
	AllowedDataStorageLocations []interface{} `json:"allowedDataStorageLocations,omitempty"`
	// Sources from which data is allowed to be transferred.
	AllowedInboundDataTransferSources interface{} `json:"allowedInboundDataTransferSources,omitempty"`
	// The level to which the clipboard may be shared between apps on the managed device.
	AllowedOutboundClipboardSharingLevel interface{} `json:"allowedOutboundClipboardSharingLevel,omitempty"`
	// Destinations to which data is allowed to be transferred.
	AllowedOutboundDataTransferDestinations interface{} `json:"allowedOutboundDataTransferDestinations,omitempty"`
	// Indicates whether contacts can be synced to the user's device.
	ContactSyncBlocked bool `json:"contactSyncBlocked,omitempty"`
	// Indicates whether the backup of a managed app's data is blocked.
	DataBackupBlocked bool `json:"dataBackupBlocked,omitempty"`
	// Indicates whether device compliance is required.
	DeviceComplianceRequired bool `json:"deviceComplianceRequired,omitempty"`
	// Indicates whether use of the app pin is required if the device pin is set.
	DisableAppPinIfDevicePinIsSet bool `json:"disableAppPinIfDevicePinIsSet,omitempty"`
	// Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
	FingerprintBlocked bool `json:"fingerprintBlocked,omitempty"`
	// Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
	ManagedBrowser interface{} `json:"managedBrowser,omitempty"`
	// Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
	ManagedBrowserToOpenLinksRequired bool `json:"managedBrowserToOpenLinksRequired,omitempty"`
	// Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
	MaximumPinRetries int32 `json:"maximumPinRetries,omitempty"`
	// Minimum pin length required for an app-level pin if PinRequired is set to True
	MinimumPinLength int32 `json:"minimumPinLength,omitempty"`
	// Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredAppVersion string `json:"minimumRequiredAppVersion,omitempty"`
	// Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredOsVersion string `json:"minimumRequiredOsVersion,omitempty"`
	// Versions less than the specified version will result in warning message on the managed app.
	MinimumWarningAppVersion string `json:"minimumWarningAppVersion,omitempty"`
	// Versions less than the specified version will result in warning message on the managed app from accessing company data.
	MinimumWarningOsVersion string `json:"minimumWarningOsVersion,omitempty"`
	// Indicates whether organizational credentials are required for app use.
	OrganizationalCredentialsRequired bool `json:"organizationalCredentialsRequired,omitempty"`
	// TimePeriod before the all-level pin must be reset if PinRequired is set to True.
	PeriodBeforePinReset string `json:"periodBeforePinReset,omitempty"`
	// The period after which access is checked when the device is not connected to the internet.
	PeriodOfflineBeforeAccessCheck string `json:"periodOfflineBeforeAccessCheck,omitempty"`
	// The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
	PeriodOfflineBeforeWipeIsEnforced string `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`
	// The period after which access is checked when the device is connected to the internet.
	PeriodOnlineBeforeAccessCheck string `json:"periodOnlineBeforeAccessCheck,omitempty"`
	// Character set which may be used for an app-level pin if PinRequired is set to True.
	PinCharacterSet interface{} `json:"pinCharacterSet,omitempty"`
	// Indicates whether an app-level pin is required.
	PinRequired bool `json:"pinRequired,omitempty"`
	// Indicates whether printing is allowed from managed apps.
	PrintBlocked bool `json:"printBlocked,omitempty"`
	// Indicates whether users may use the \"Save As\" menu item to save a copy of protected files.
	SaveAsBlocked bool `json:"saveAsBlocked,omitempty"`
	// Indicates whether simplePin is blocked.
	SimplePinBlocked bool `json:"simplePinBlocked,omitempty"`
	// Indicates if the policy is deployed to any inclusion groups or not.
	IsAssigned  bool                                               `json:"isAssigned,omitempty"`
	Assignments []MicrosoftGraphTargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}
