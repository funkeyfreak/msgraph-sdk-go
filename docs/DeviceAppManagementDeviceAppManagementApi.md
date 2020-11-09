# \DeviceAppManagementDeviceAppManagementApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementDeviceAppManagementGetDeviceAppManagement**](DeviceAppManagementDeviceAppManagementApi.md#DeviceAppManagementDeviceAppManagementGetDeviceAppManagement) | **Get** /deviceAppManagement | Get deviceAppManagement
[**DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement**](DeviceAppManagementDeviceAppManagementApi.md#DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement) | **Patch** /deviceAppManagement | Update deviceAppManagement



## DeviceAppManagementDeviceAppManagementGetDeviceAppManagement

> MicrosoftGraphDeviceAppManagement DeviceAppManagementDeviceAppManagementGetDeviceAppManagement(ctx, optional)

Get deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementDeviceAppManagementGetDeviceAppManagementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementDeviceAppManagementGetDeviceAppManagementOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceAppManagement**](microsoft.graph.deviceAppManagement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement

> DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement(ctx, microsoftGraphDeviceAppManagement)

Update deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceAppManagement** | [**MicrosoftGraphDeviceAppManagement**](MicrosoftGraphDeviceAppManagement.md)| New property values | 

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

