# IosManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDataEncryptionType** | [**interface{}**](.md) | Type of encryption which should be used for data in a managed app. | [optional] 
**CustomBrowserProtocol** | **string** | A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. | [optional] 
**DeployedAppCount** | **int32** | Count of apps to which the current policy is deployed. | [optional] 
**FaceIdBlocked** | **bool** | Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. | [optional] 
**MinimumRequiredSdkVersion** | **string** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**Apps** | [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**DeploymentSummary** | [**interface{}**](.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


