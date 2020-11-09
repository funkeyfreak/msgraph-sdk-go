# AndroidManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomBrowserDisplayName** | **string** | Friendly name of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. | [optional] 
**CustomBrowserPackageId** | **string** | Unique identifier of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. | [optional] 
**DeployedAppCount** | **int32** | Count of apps to which the current policy is deployed. | [optional] 
**DisableAppEncryptionIfDeviceEncryptionIsEnabled** | **bool** | When this setting is enabled, app level encryption is disabled if device level encryption is enabled | [optional] 
**EncryptAppData** | **bool** | Indicates whether application data for managed apps should be encrypted | [optional] 
**MinimumRequiredPatchVersion** | **string** | Define the oldest required Android security patch level a user can have to gain secure access to the app. | [optional] 
**MinimumWarningPatchVersion** | **string** | Define the oldest recommended Android security patch level a user can have for secure access to the app. | [optional] 
**ScreenCaptureBlocked** | **bool** | Indicates whether a managed user can take screen captures of managed apps | [optional] 
**Apps** | [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**DeploymentSummary** | [**interface{}**](.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


