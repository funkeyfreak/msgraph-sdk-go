# \DirectoryDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryCreateDeletedItems**](DirectoryDirectoryObjectApi.md#DirectoryCreateDeletedItems) | **Post** /directory/deletedItems | Create new navigation property to deletedItems for directory
[**DirectoryGetDeletedItems**](DirectoryDirectoryObjectApi.md#DirectoryGetDeletedItems) | **Get** /directory/deletedItems/{directoryObject-id} | Get deletedItems from directory
[**DirectoryListDeletedItems**](DirectoryDirectoryObjectApi.md#DirectoryListDeletedItems) | **Get** /directory/deletedItems | Get deletedItems from directory
[**DirectoryUpdateDeletedItems**](DirectoryDirectoryObjectApi.md#DirectoryUpdateDeletedItems) | **Patch** /directory/deletedItems/{directoryObject-id} | Update the navigation property deletedItems in directory



## DirectoryCreateDeletedItems

> MicrosoftGraphDirectoryObject DirectoryCreateDeletedItems(ctx, microsoftGraphDirectoryObject)

Create new navigation property to deletedItems for directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)| New navigation property | 

### Return type

[**MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryGetDeletedItems

> MicrosoftGraphDirectoryObject DirectoryGetDeletedItems(ctx, directoryObjectId, optional)

Get deletedItems from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***DirectoryGetDeletedItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryGetDeletedItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryListDeletedItems

> CollectionOfDirectoryObject DirectoryListDeletedItems(ctx, optional)

Get deletedItems from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryListDeletedItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryListDeletedItemsOpts struct


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

[**CollectionOfDirectoryObject**](Collection_of_directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryUpdateDeletedItems

> DirectoryUpdateDeletedItems(ctx, directoryObjectId, microsoftGraphDirectoryObject)

Update the navigation property deletedItems in directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
**microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)| New navigation property values | 

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

