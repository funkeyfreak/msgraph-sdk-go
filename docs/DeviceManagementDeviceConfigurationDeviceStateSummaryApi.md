# \DeviceManagementDeviceConfigurationDeviceStateSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementGetDeviceConfigurationDeviceStateSummaries**](DeviceManagementDeviceConfigurationDeviceStateSummaryApi.md#DeviceManagementGetDeviceConfigurationDeviceStateSummaries) | **Get** /deviceManagement/deviceConfigurationDeviceStateSummaries | Get deviceConfigurationDeviceStateSummaries from deviceManagement
[**DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries**](DeviceManagementDeviceConfigurationDeviceStateSummaryApi.md#DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries) | **Patch** /deviceManagement/deviceConfigurationDeviceStateSummaries | Update the navigation property deviceConfigurationDeviceStateSummaries in deviceManagement



## DeviceManagementGetDeviceConfigurationDeviceStateSummaries

> MicrosoftGraphDeviceConfigurationDeviceStateSummary DeviceManagementGetDeviceConfigurationDeviceStateSummaries(ctx, optional)

Get deviceConfigurationDeviceStateSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementGetDeviceConfigurationDeviceStateSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDeviceConfigurationDeviceStateSummariesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationDeviceStateSummary**](microsoft.graph.deviceConfigurationDeviceStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries

> DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries(ctx, microsoftGraphDeviceConfigurationDeviceStateSummary)

Update the navigation property deviceConfigurationDeviceStateSummaries in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceConfigurationDeviceStateSummary** | [**MicrosoftGraphDeviceConfigurationDeviceStateSummary**](MicrosoftGraphDeviceConfigurationDeviceStateSummary.md)| New navigation property values | 

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

