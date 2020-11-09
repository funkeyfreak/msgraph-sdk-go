# \DeviceManagementDeviceCategoryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceCategories**](DeviceManagementDeviceCategoryApi.md#DeviceManagementCreateDeviceCategories) | **Post** /deviceManagement/deviceCategories | Create new navigation property to deviceCategories for deviceManagement
[**DeviceManagementGetDeviceCategories**](DeviceManagementDeviceCategoryApi.md#DeviceManagementGetDeviceCategories) | **Get** /deviceManagement/deviceCategories/{deviceCategory-id} | Get deviceCategories from deviceManagement
[**DeviceManagementListDeviceCategories**](DeviceManagementDeviceCategoryApi.md#DeviceManagementListDeviceCategories) | **Get** /deviceManagement/deviceCategories | Get deviceCategories from deviceManagement
[**DeviceManagementUpdateDeviceCategories**](DeviceManagementDeviceCategoryApi.md#DeviceManagementUpdateDeviceCategories) | **Patch** /deviceManagement/deviceCategories/{deviceCategory-id} | Update the navigation property deviceCategories in deviceManagement



## DeviceManagementCreateDeviceCategories

> MicrosoftGraphDeviceCategory DeviceManagementCreateDeviceCategories(ctx, microsoftGraphDeviceCategory)

Create new navigation property to deviceCategories for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceCategory** | [**MicrosoftGraphDeviceCategory**](MicrosoftGraphDeviceCategory.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceCategory**](microsoft.graph.deviceCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetDeviceCategories

> MicrosoftGraphDeviceCategory DeviceManagementGetDeviceCategories(ctx, deviceCategoryId, optional)

Get deviceCategories from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCategoryId** | **string**| key: deviceCategory-id of deviceCategory | 
 **optional** | ***DeviceManagementGetDeviceCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDeviceCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCategory**](microsoft.graph.deviceCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceCategories

> CollectionOfDeviceCategory DeviceManagementListDeviceCategories(ctx, optional)

Get deviceCategories from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListDeviceCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListDeviceCategoriesOpts struct


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

[**CollectionOfDeviceCategory**](Collection_of_deviceCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceCategories

> DeviceManagementUpdateDeviceCategories(ctx, deviceCategoryId, microsoftGraphDeviceCategory)

Update the navigation property deviceCategories in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCategoryId** | **string**| key: deviceCategory-id of deviceCategory | 
**microsoftGraphDeviceCategory** | [**MicrosoftGraphDeviceCategory**](MicrosoftGraphDeviceCategory.md)| New navigation property values | 

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

