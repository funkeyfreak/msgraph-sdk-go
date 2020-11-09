# \DeviceManagementDetectedAppApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDetectedApps**](DeviceManagementDetectedAppApi.md#DeviceManagementCreateDetectedApps) | **Post** /deviceManagement/detectedApps | Create new navigation property to detectedApps for deviceManagement
[**DeviceManagementDetectedAppsGetManagedDevices**](DeviceManagementDetectedAppApi.md#DeviceManagementDetectedAppsGetManagedDevices) | **Get** /deviceManagement/detectedApps/{detectedApp-id}/managedDevices/{managedDevice-id} | Get managedDevices from deviceManagement
[**DeviceManagementDetectedAppsListManagedDevices**](DeviceManagementDetectedAppApi.md#DeviceManagementDetectedAppsListManagedDevices) | **Get** /deviceManagement/detectedApps/{detectedApp-id}/managedDevices | Get managedDevices from deviceManagement
[**DeviceManagementGetDetectedApps**](DeviceManagementDetectedAppApi.md#DeviceManagementGetDetectedApps) | **Get** /deviceManagement/detectedApps/{detectedApp-id} | Get detectedApps from deviceManagement
[**DeviceManagementListDetectedApps**](DeviceManagementDetectedAppApi.md#DeviceManagementListDetectedApps) | **Get** /deviceManagement/detectedApps | Get detectedApps from deviceManagement
[**DeviceManagementUpdateDetectedApps**](DeviceManagementDetectedAppApi.md#DeviceManagementUpdateDetectedApps) | **Patch** /deviceManagement/detectedApps/{detectedApp-id} | Update the navigation property detectedApps in deviceManagement



## DeviceManagementCreateDetectedApps

> MicrosoftGraphDetectedApp DeviceManagementCreateDetectedApps(ctx, microsoftGraphDetectedApp)

Create new navigation property to detectedApps for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDetectedApp** | [**MicrosoftGraphDetectedApp**](MicrosoftGraphDetectedApp.md)| New navigation property | 

### Return type

[**MicrosoftGraphDetectedApp**](microsoft.graph.detectedApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDetectedAppsGetManagedDevices

> MicrosoftGraphManagedDevice DeviceManagementDetectedAppsGetManagedDevices(ctx, detectedAppId, managedDeviceId, optional)

Get managedDevices from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**detectedAppId** | **string**| key: detectedApp-id of detectedApp | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***DeviceManagementDetectedAppsGetManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDetectedAppsGetManagedDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDetectedAppsListManagedDevices

> CollectionOfManagedDevice DeviceManagementDetectedAppsListManagedDevices(ctx, detectedAppId, optional)

Get managedDevices from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**detectedAppId** | **string**| key: detectedApp-id of detectedApp | 
 **optional** | ***DeviceManagementDetectedAppsListManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDetectedAppsListManagedDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**CollectionOfManagedDevice**](Collection_of_managedDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetDetectedApps

> MicrosoftGraphDetectedApp DeviceManagementGetDetectedApps(ctx, detectedAppId, optional)

Get detectedApps from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**detectedAppId** | **string**| key: detectedApp-id of detectedApp | 
 **optional** | ***DeviceManagementGetDetectedAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDetectedAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDetectedApp**](microsoft.graph.detectedApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDetectedApps

> CollectionOfDetectedApp DeviceManagementListDetectedApps(ctx, optional)

Get detectedApps from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListDetectedAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListDetectedAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**CollectionOfDetectedApp**](Collection_of_detectedApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDetectedApps

> DeviceManagementUpdateDetectedApps(ctx, detectedAppId, microsoftGraphDetectedApp)

Update the navigation property detectedApps in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**detectedAppId** | **string**| key: detectedApp-id of detectedApp | 
**microsoftGraphDetectedApp** | [**MicrosoftGraphDetectedApp**](MicrosoftGraphDetectedApp.md)| New navigation property values | 

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

