# DefaultManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDataEncryptionType** | [**interface{}**](.md) | Type of encryption which should be used for data in a managed app. (iOS Only) | [optional] 
**CustomSettings** | [**[]MicrosoftGraphKeyValuePair**](microsoft.graph.keyValuePair.md) | A set of string key and string value pairs to be sent to the affected users, unalterned by this service | [optional] 
**DeployedAppCount** | **int32** | Count of apps to which the current policy is deployed. | [optional] 
**DisableAppEncryptionIfDeviceEncryptionIsEnabled** | **bool** | When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only) | [optional] 
**EncryptAppData** | **bool** | Indicates whether managed-app data should be encrypted. (Android only) | [optional] 
**FaceIdBlocked** | **bool** | Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only) | [optional] 
**MinimumRequiredPatchVersion** | **string** | Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only) | [optional] 
**MinimumRequiredSdkVersion** | **string** | Versions less than the specified version will block the managed app from accessing company data. (iOS Only) | [optional] 
**MinimumWarningPatchVersion** | **string** | Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only) | [optional] 
**ScreenCaptureBlocked** | **bool** | Indicates whether screen capture is blocked. (Android only) | [optional] 
**Apps** | [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**DeploymentSummary** | [**interface{}**](.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


