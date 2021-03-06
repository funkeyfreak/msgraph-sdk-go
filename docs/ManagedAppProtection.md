# ManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDataStorageLocations** | **[]interface{}** | Data storage locations where a user may store managed data. | [optional] 
**AllowedInboundDataTransferSources** | [**interface{}**](.md) | Sources from which data is allowed to be transferred. | [optional] 
**AllowedOutboundClipboardSharingLevel** | [**interface{}**](.md) | The level to which the clipboard may be shared between apps on the managed device. | [optional] 
**AllowedOutboundDataTransferDestinations** | [**interface{}**](.md) | Destinations to which data is allowed to be transferred. | [optional] 
**ContactSyncBlocked** | **bool** | Indicates whether contacts can be synced to the user&#39;s device. | [optional] 
**DataBackupBlocked** | **bool** | Indicates whether the backup of a managed app&#39;s data is blocked. | [optional] 
**DeviceComplianceRequired** | **bool** | Indicates whether device compliance is required. | [optional] 
**DisableAppPinIfDevicePinIsSet** | **bool** | Indicates whether use of the app pin is required if the device pin is set. | [optional] 
**FingerprintBlocked** | **bool** | Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True. | [optional] 
**ManagedBrowser** | [**interface{}**](.md) | Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. | [optional] 
**ManagedBrowserToOpenLinksRequired** | **bool** | Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android) | [optional] 
**MaximumPinRetries** | **int32** | Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped. | [optional] 
**MinimumPinLength** | **int32** | Minimum pin length required for an app-level pin if PinRequired is set to True | [optional] 
**MinimumRequiredAppVersion** | **string** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**MinimumRequiredOsVersion** | **string** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**MinimumWarningAppVersion** | **string** | Versions less than the specified version will result in warning message on the managed app. | [optional] 
**MinimumWarningOsVersion** | **string** | Versions less than the specified version will result in warning message on the managed app from accessing company data. | [optional] 
**OrganizationalCredentialsRequired** | **bool** | Indicates whether organizational credentials are required for app use. | [optional] 
**PeriodBeforePinReset** | **string** | TimePeriod before the all-level pin must be reset if PinRequired is set to True. | [optional] 
**PeriodOfflineBeforeAccessCheck** | **string** | The period after which access is checked when the device is not connected to the internet. | [optional] 
**PeriodOfflineBeforeWipeIsEnforced** | **string** | The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped. | [optional] 
**PeriodOnlineBeforeAccessCheck** | **string** | The period after which access is checked when the device is connected to the internet. | [optional] 
**PinCharacterSet** | [**interface{}**](.md) | Character set which may be used for an app-level pin if PinRequired is set to True. | [optional] 
**PinRequired** | **bool** | Indicates whether an app-level pin is required. | [optional] 
**PrintBlocked** | **bool** | Indicates whether printing is allowed from managed apps. | [optional] 
**SaveAsBlocked** | **bool** | Indicates whether users may use the \&quot;Save As\&quot; menu item to save a copy of protected files. | [optional] 
**SimplePinBlocked** | **bool** | Indicates whether simplePin is blocked. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


