# MicrosoftGraphManagedAppRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**AppIdentifier** | [**interface{}**](.md) | The app package Identifier | [optional] 
**ApplicationVersion** | **string** | App version | [optional] 
**CreatedDateTime** | [**time.Time**](time.Time.md) | Date and time of creation | [optional] 
**DeviceName** | **string** | Host device name | [optional] 
**DeviceTag** | **string** | App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions. | [optional] 
**DeviceType** | **string** | Host device type | [optional] 
**FlaggedReasons** | **[]interface{}** | Zero or more reasons an app registration is flagged. E.g. app running on rooted device | [optional] 
**LastSyncDateTime** | [**time.Time**](time.Time.md) | Date and time of last the app synced with management service. | [optional] 
**ManagementSdkVersion** | **string** | App management SDK version | [optional] 
**PlatformVersion** | **string** | Operating System version | [optional] 
**UserId** | **string** | The user Id to who this app registration belongs. | [optional] 
**Version** | **string** | Version of the entity. | [optional] 
**AppliedPolicies** | [**[]MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md) |  | [optional] 
**IntendedPolicies** | [**[]MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md) |  | [optional] 
**Operations** | [**[]MicrosoftGraphManagedAppOperation**](microsoft.graph.managedAppOperation.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


