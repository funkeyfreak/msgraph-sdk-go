# \DirectoryObjectsDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryObjectsDirectoryObjectCreateDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectCreateDirectoryObject) | **Post** /directoryObjects | Add new entity to directoryObjects
[**DirectoryObjectsDirectoryObjectDeleteDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectDeleteDirectoryObject) | **Delete** /directoryObjects/{directoryObject-id} | Delete entity from directoryObjects
[**DirectoryObjectsDirectoryObjectGetDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectGetDirectoryObject) | **Get** /directoryObjects/{directoryObject-id} | Get entity from directoryObjects by key
[**DirectoryObjectsDirectoryObjectListDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectListDirectoryObject) | **Get** /directoryObjects | Get entities from directoryObjects
[**DirectoryObjectsDirectoryObjectUpdateDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectUpdateDirectoryObject) | **Patch** /directoryObjects/{directoryObject-id} | Update entity in directoryObjects



## DirectoryObjectsDirectoryObjectCreateDirectoryObject

> MicrosoftGraphDirectoryObject DirectoryObjectsDirectoryObjectCreateDirectoryObject(ctx, microsoftGraphDirectoryObject)

Add new entity to directoryObjects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)| New entity | 

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


## DirectoryObjectsDirectoryObjectDeleteDirectoryObject

> DirectoryObjectsDirectoryObjectDeleteDirectoryObject(ctx, directoryObjectId, optional)

Delete entity from directoryObjects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***DirectoryObjectsDirectoryObjectDeleteDirectoryObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryObjectsDirectoryObjectDeleteDirectoryObjectOpts struct


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


## DirectoryObjectsDirectoryObjectGetDirectoryObject

> MicrosoftGraphDirectoryObject DirectoryObjectsDirectoryObjectGetDirectoryObject(ctx, directoryObjectId, optional)

Get entity from directoryObjects by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***DirectoryObjectsDirectoryObjectGetDirectoryObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryObjectsDirectoryObjectGetDirectoryObjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## DirectoryObjectsDirectoryObjectListDirectoryObject

> CollectionOfDirectoryObject DirectoryObjectsDirectoryObjectListDirectoryObject(ctx, optional)

Get entities from directoryObjects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryObjectsDirectoryObjectListDirectoryObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryObjectsDirectoryObjectListDirectoryObjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## DirectoryObjectsDirectoryObjectUpdateDirectoryObject

> DirectoryObjectsDirectoryObjectUpdateDirectoryObject(ctx, directoryObjectId, microsoftGraphDirectoryObject)

Update entity in directoryObjects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
**microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)| New property values | 

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

