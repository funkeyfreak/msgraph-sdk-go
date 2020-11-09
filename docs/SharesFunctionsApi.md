# \SharesFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesListItemGetActivitiesByInterval53ee**](SharesFunctionsApi.md#SharesListItemGetActivitiesByInterval53ee) | **Get** /shares/{sharedDriveItem-id}/listItem/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;{startDateTime},endDateTime&#x3D;{endDateTime},interval&#x3D;{interval}) | Invoke function getActivitiesByInterval
[**SharesListItemGetActivitiesByInterval96b0**](SharesFunctionsApi.md#SharesListItemGetActivitiesByInterval96b0) | **Get** /shares/{sharedDriveItem-id}/listItem/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**SharesListItemsGetActivitiesByInterval53ee**](SharesFunctionsApi.md#SharesListItemsGetActivitiesByInterval53ee) | **Get** /shares/{sharedDriveItem-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;{startDateTime},endDateTime&#x3D;{endDateTime},interval&#x3D;{interval}) | Invoke function getActivitiesByInterval
[**SharesListItemsGetActivitiesByInterval96b0**](SharesFunctionsApi.md#SharesListItemsGetActivitiesByInterval96b0) | **Get** /shares/{sharedDriveItem-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval



## SharesListItemGetActivitiesByInterval53ee

> []interface{} SharesListItemGetActivitiesByInterval53ee(ctx, sharedDriveItemId, startDateTime, endDateTime, interval)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**startDateTime** | **string**|  | 
**endDateTime** | **string**|  | 
**interval** | **string**|  | 

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemGetActivitiesByInterval96b0

> []interface{} SharesListItemGetActivitiesByInterval96b0(ctx, sharedDriveItemId)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemsGetActivitiesByInterval53ee

> []interface{} SharesListItemsGetActivitiesByInterval53ee(ctx, sharedDriveItemId, listItemId, startDateTime, endDateTime, interval)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
**startDateTime** | **string**|  | 
**endDateTime** | **string**|  | 
**interval** | **string**|  | 

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemsGetActivitiesByInterval96b0

> []interface{} SharesListItemsGetActivitiesByInterval96b0(ctx, sharedDriveItemId, listItemId)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

