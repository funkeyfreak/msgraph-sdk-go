# \DevicesDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesDeviceCreateDevice**](DevicesDeviceApi.md#DevicesDeviceCreateDevice) | **Post** /devices | Add new entity to devices
[**DevicesDeviceDeleteDevice**](DevicesDeviceApi.md#DevicesDeviceDeleteDevice) | **Delete** /devices/{device-id} | Delete entity from devices
[**DevicesDeviceGetDevice**](DevicesDeviceApi.md#DevicesDeviceGetDevice) | **Get** /devices/{device-id} | Get entity from devices by key
[**DevicesDeviceListDevice**](DevicesDeviceApi.md#DevicesDeviceListDevice) | **Get** /devices | Get entities from devices
[**DevicesDeviceUpdateDevice**](DevicesDeviceApi.md#DevicesDeviceUpdateDevice) | **Patch** /devices/{device-id} | Update entity in devices



## DevicesDeviceCreateDevice

> MicrosoftGraphDevice DevicesDeviceCreateDevice(ctx, microsoftGraphDevice)

Add new entity to devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDevice** | [**MicrosoftGraphDevice**](MicrosoftGraphDevice.md)| New entity | 

### Return type

[**MicrosoftGraphDevice**](microsoft.graph.device.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceDeleteDevice

> DevicesDeviceDeleteDevice(ctx, deviceId, optional)

Delete entity from devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| key: device-id of device | 
 **optional** | ***DevicesDeviceDeleteDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesDeviceDeleteDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| ETag | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceGetDevice

> MicrosoftGraphDevice DevicesDeviceGetDevice(ctx, deviceId, optional)

Get entity from devices by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| key: device-id of device | 
 **optional** | ***DevicesDeviceGetDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesDeviceGetDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphDevice**](microsoft.graph.device.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceListDevice

> CollectionOfDevice DevicesDeviceListDevice(ctx, optional)

Get entities from devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesDeviceListDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesDeviceListDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfDevice**](Collection_of_device.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceUpdateDevice

> DevicesDeviceUpdateDevice(ctx, deviceId, microsoftGraphDevice)

Update entity in devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| key: device-id of device | 
**microsoftGraphDevice** | [**MicrosoftGraphDevice**](MicrosoftGraphDevice.md)| New property values | 

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

