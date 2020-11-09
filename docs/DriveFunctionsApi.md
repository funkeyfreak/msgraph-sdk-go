# \DriveFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveListItemsGetActivitiesByInterval53ee**](DriveFunctionsApi.md#DriveListItemsGetActivitiesByInterval53ee) | **Get** /drive/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;{startDateTime},endDateTime&#x3D;{endDateTime},interval&#x3D;{interval}) | Invoke function getActivitiesByInterval
[**DriveListItemsGetActivitiesByInterval96b0**](DriveFunctionsApi.md#DriveListItemsGetActivitiesByInterval96b0) | **Get** /drive/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**DriveRecent**](DriveFunctionsApi.md#DriveRecent) | **Get** /drive/microsoft.graph.recent() | Invoke function recent
[**DriveSearch**](DriveFunctionsApi.md#DriveSearch) | **Get** /drive/microsoft.graph.search(q&#x3D;{q}) | Invoke function search
[**DriveSharedWithMe**](DriveFunctionsApi.md#DriveSharedWithMe) | **Get** /drive/microsoft.graph.sharedWithMe() | Invoke function sharedWithMe



## DriveListItemsGetActivitiesByInterval53ee

> []interface{} DriveListItemsGetActivitiesByInterval53ee(ctx, listItemId, startDateTime, endDateTime, interval)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## DriveListItemsGetActivitiesByInterval96b0

> []interface{} DriveListItemsGetActivitiesByInterval96b0(ctx, listItemId)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## DriveRecent

> []interface{} DriveRecent(ctx, )

Invoke function recent

### Required Parameters

This endpoint does not need any parameter.

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


## DriveSearch

> []interface{} DriveSearch(ctx, q)

Invoke function search

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## DriveSharedWithMe

> []interface{} DriveSharedWithMe(ctx, )

Invoke function sharedWithMe

### Required Parameters

This endpoint does not need any parameter.

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

