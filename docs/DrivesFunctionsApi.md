# \DrivesFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesListItemsGetActivitiesByInterval53ee**](DrivesFunctionsApi.md#DrivesListItemsGetActivitiesByInterval53ee) | **Get** /drives/{drive-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;{startDateTime},endDateTime&#x3D;{endDateTime},interval&#x3D;{interval}) | Invoke function getActivitiesByInterval
[**DrivesListItemsGetActivitiesByInterval96b0**](DrivesFunctionsApi.md#DrivesListItemsGetActivitiesByInterval96b0) | **Get** /drives/{drive-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**DrivesRecent**](DrivesFunctionsApi.md#DrivesRecent) | **Get** /drives/{drive-id}/microsoft.graph.recent() | Invoke function recent
[**DrivesSearch**](DrivesFunctionsApi.md#DrivesSearch) | **Get** /drives/{drive-id}/microsoft.graph.search(q&#x3D;{q}) | Invoke function search
[**DrivesSharedWithMe**](DrivesFunctionsApi.md#DrivesSharedWithMe) | **Get** /drives/{drive-id}/microsoft.graph.sharedWithMe() | Invoke function sharedWithMe



## DrivesListItemsGetActivitiesByInterval53ee

> []interface{} DrivesListItemsGetActivitiesByInterval53ee(ctx, driveId, listItemId, startDateTime, endDateTime, interval)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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


## DrivesListItemsGetActivitiesByInterval96b0

> []interface{} DrivesListItemsGetActivitiesByInterval96b0(ctx, driveId, listItemId)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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


## DrivesRecent

> []interface{} DrivesRecent(ctx, driveId)

Invoke function recent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 

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


## DrivesSearch

> []interface{} DrivesSearch(ctx, driveId, q)

Invoke function search

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**q** | **string**|  | 

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


## DrivesSharedWithMe

> []interface{} DrivesSharedWithMe(ctx, driveId)

Invoke function sharedWithMe

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 

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

