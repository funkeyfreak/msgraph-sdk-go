# MicrosoftGraphDeviceEnrollmentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**CreatedDateTime** | [**time.Time**](time.Time.md) | Created date time in UTC of the device enrollment configuration | [optional] 
**Description** | **string** | The description of the device enrollment configuration | [optional] 
**DisplayName** | **string** | The display name of the device enrollment configuration | [optional] 
**LastModifiedDateTime** | [**time.Time**](time.Time.md) | Last modified date time in UTC of the device enrollment configuration | [optional] 
**Priority** | **int32** | Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject only to the configuration with the lowest priority value. | [optional] 
**Version** | **int32** | The version of the device enrollment configuration | [optional] 
**Assignments** | [**[]MicrosoftGraphEnrollmentConfigurationAssignment**](microsoft.graph.enrollmentConfigurationAssignment.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


