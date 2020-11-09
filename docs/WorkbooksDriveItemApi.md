# \WorkbooksDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksCreateChildren**](WorkbooksDriveItemApi.md#WorkbooksCreateChildren) | **Post** /workbooks/{driveItem-id}/children | Create new navigation property to children for workbooks
[**WorkbooksDriveItemCreateDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemCreateDriveItem) | **Post** /workbooks | Add new entity to workbooks
[**WorkbooksDriveItemDeleteDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemDeleteDriveItem) | **Delete** /workbooks/{driveItem-id} | Delete entity from workbooks
[**WorkbooksDriveItemGetDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemGetDriveItem) | **Get** /workbooks/{driveItem-id} | Get entity from workbooks by key
[**WorkbooksDriveItemListDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemListDriveItem) | **Get** /workbooks | Get entities from workbooks
[**WorkbooksDriveItemUpdateDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemUpdateDriveItem) | **Patch** /workbooks/{driveItem-id} | Update entity in workbooks
[**WorkbooksGetChildren**](WorkbooksDriveItemApi.md#WorkbooksGetChildren) | **Get** /workbooks/{driveItem-id}/children/{driveItem-id1} | Get children from workbooks
[**WorkbooksListChildren**](WorkbooksDriveItemApi.md#WorkbooksListChildren) | **Get** /workbooks/{driveItem-id}/children | Get children from workbooks
[**WorkbooksUpdateChildren**](WorkbooksDriveItemApi.md#WorkbooksUpdateChildren) | **Patch** /workbooks/{driveItem-id}/children/{driveItem-id1} | Update the navigation property children in workbooks



## WorkbooksCreateChildren

> MicrosoftGraphDriveItem WorkbooksCreateChildren(ctx, driveItemId, microsoftGraphDriveItem)

Create new navigation property to children for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDriveItemCreateDriveItem

> MicrosoftGraphDriveItem WorkbooksDriveItemCreateDriveItem(ctx, microsoftGraphDriveItem)

Add new entity to workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New entity | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDriveItemDeleteDriveItem

> WorkbooksDriveItemDeleteDriveItem(ctx, driveItemId, optional)

Delete entity from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksDriveItemDeleteDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksDriveItemDeleteDriveItemOpts struct


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


## WorkbooksDriveItemGetDriveItem

> MicrosoftGraphDriveItem WorkbooksDriveItemGetDriveItem(ctx, driveItemId, optional)

Get entity from workbooks by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksDriveItemGetDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksDriveItemGetDriveItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDriveItemListDriveItem

> CollectionOfDriveItem WorkbooksDriveItemListDriveItem(ctx, optional)

Get entities from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkbooksDriveItemListDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksDriveItemListDriveItemOpts struct


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

[**CollectionOfDriveItem**](Collection_of_driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDriveItemUpdateDriveItem

> WorkbooksDriveItemUpdateDriveItem(ctx, driveItemId, microsoftGraphDriveItem)

Update entity in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New property values | 

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


## WorkbooksGetChildren

> MicrosoftGraphDriveItem WorkbooksGetChildren(ctx, driveItemId, driveItemId1, optional)

Get children from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**driveItemId1** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksGetChildrenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksGetChildrenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListChildren

> CollectionOfDriveItem WorkbooksListChildren(ctx, driveItemId, optional)

Get children from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksListChildrenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListChildrenOpts struct


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

[**CollectionOfDriveItem**](Collection_of_driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksUpdateChildren

> WorkbooksUpdateChildren(ctx, driveItemId, driveItemId1, microsoftGraphDriveItem)

Update the navigation property children in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**driveItemId1** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property values | 

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

