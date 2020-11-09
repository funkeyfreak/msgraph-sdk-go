# \DeviceManagementDeviceManagementApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceManagementGetDeviceManagement**](DeviceManagementDeviceManagementApi.md#DeviceManagementDeviceManagementGetDeviceManagement) | **Get** /deviceManagement | Get deviceManagement
[**DeviceManagementDeviceManagementUpdateDeviceManagement**](DeviceManagementDeviceManagementApi.md#DeviceManagementDeviceManagementUpdateDeviceManagement) | **Patch** /deviceManagement | Update deviceManagement



## DeviceManagementDeviceManagementGetDeviceManagement

> MicrosoftGraphDeviceManagement DeviceManagementDeviceManagementGetDeviceManagement(ctx, optional)

Get deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementDeviceManagementGetDeviceManagementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceManagementGetDeviceManagementOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagement**](microsoft.graph.deviceManagement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceManagementUpdateDeviceManagement

> DeviceManagementDeviceManagementUpdateDeviceManagement(ctx, microsoftGraphDeviceManagement)

Update deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceManagement** | [**MicrosoftGraphDeviceManagement**](MicrosoftGraphDeviceManagement.md)| New property values | 

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

