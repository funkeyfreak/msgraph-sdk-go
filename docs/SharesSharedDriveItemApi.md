# \SharesSharedDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesSharedDriveItemCreateSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemCreateSharedDriveItem) | **Post** /shares | Add new entity to shares
[**SharesSharedDriveItemDeleteSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemDeleteSharedDriveItem) | **Delete** /shares/{sharedDriveItem-id} | Delete entity from shares
[**SharesSharedDriveItemGetSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemGetSharedDriveItem) | **Get** /shares/{sharedDriveItem-id} | Get entity from shares by key
[**SharesSharedDriveItemListSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemListSharedDriveItem) | **Get** /shares | Get entities from shares
[**SharesSharedDriveItemUpdateSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemUpdateSharedDriveItem) | **Patch** /shares/{sharedDriveItem-id} | Update entity in shares



## SharesSharedDriveItemCreateSharedDriveItem

> MicrosoftGraphSharedDriveItem SharesSharedDriveItemCreateSharedDriveItem(ctx, microsoftGraphSharedDriveItem)

Add new entity to shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSharedDriveItem** | [**MicrosoftGraphSharedDriveItem**](MicrosoftGraphSharedDriveItem.md)| New entity | 

### Return type

[**MicrosoftGraphSharedDriveItem**](microsoft.graph.sharedDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemDeleteSharedDriveItem

> SharesSharedDriveItemDeleteSharedDriveItem(ctx, sharedDriveItemId, optional)

Delete entity from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesSharedDriveItemDeleteSharedDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesSharedDriveItemDeleteSharedDriveItemOpts struct


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


## SharesSharedDriveItemGetSharedDriveItem

> MicrosoftGraphSharedDriveItem SharesSharedDriveItemGetSharedDriveItem(ctx, sharedDriveItemId, optional)

Get entity from shares by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesSharedDriveItemGetSharedDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesSharedDriveItemGetSharedDriveItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSharedDriveItem**](microsoft.graph.sharedDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemListSharedDriveItem

> CollectionOfSharedDriveItem SharesSharedDriveItemListSharedDriveItem(ctx, optional)

Get entities from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharesSharedDriveItemListSharedDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesSharedDriveItemListSharedDriveItemOpts struct


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

[**CollectionOfSharedDriveItem**](Collection_of_sharedDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemUpdateSharedDriveItem

> SharesSharedDriveItemUpdateSharedDriveItem(ctx, sharedDriveItemId, microsoftGraphSharedDriveItem)

Update entity in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**microsoftGraphSharedDriveItem** | [**MicrosoftGraphSharedDriveItem**](MicrosoftGraphSharedDriveItem.md)| New property values | 

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

