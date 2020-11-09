# \DeviceManagementOnPremisesConditionalAccessSettingsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementGetConditionalAccessSettings**](DeviceManagementOnPremisesConditionalAccessSettingsApi.md#DeviceManagementGetConditionalAccessSettings) | **Get** /deviceManagement/conditionalAccessSettings | Get conditionalAccessSettings from deviceManagement
[**DeviceManagementUpdateConditionalAccessSettings**](DeviceManagementOnPremisesConditionalAccessSettingsApi.md#DeviceManagementUpdateConditionalAccessSettings) | **Patch** /deviceManagement/conditionalAccessSettings | Update the navigation property conditionalAccessSettings in deviceManagement



## DeviceManagementGetConditionalAccessSettings

> MicrosoftGraphOnPremisesConditionalAccessSettings DeviceManagementGetConditionalAccessSettings(ctx, optional)

Get conditionalAccessSettings from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementGetConditionalAccessSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetConditionalAccessSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnPremisesConditionalAccessSettings**](microsoft.graph.onPremisesConditionalAccessSettings.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateConditionalAccessSettings

> DeviceManagementUpdateConditionalAccessSettings(ctx, microsoftGraphOnPremisesConditionalAccessSettings)

Update the navigation property conditionalAccessSettings in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOnPremisesConditionalAccessSettings** | [**MicrosoftGraphOnPremisesConditionalAccessSettings**](MicrosoftGraphOnPremisesConditionalAccessSettings.md)| New navigation property values | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

