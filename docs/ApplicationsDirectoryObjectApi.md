# \ApplicationsDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsGetCreatedOnBehalfOf**](ApplicationsDirectoryObjectApi.md#ApplicationsGetCreatedOnBehalfOf) | **Get** /applications/{application-id}/createdOnBehalfOf | Get createdOnBehalfOf from applications
[**ApplicationsGetOwners**](ApplicationsDirectoryObjectApi.md#ApplicationsGetOwners) | **Get** /applications/{application-id}/owners/{directoryObject-id} | Get owners from applications
[**ApplicationsListOwners**](ApplicationsDirectoryObjectApi.md#ApplicationsListOwners) | **Get** /applications/{application-id}/owners | Get owners from applications



## ApplicationsGetCreatedOnBehalfOf

> MicrosoftGraphDirectoryObject ApplicationsGetCreatedOnBehalfOf(ctx, applicationId, optional)

Get createdOnBehalfOf from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
 **optional** | ***ApplicationsGetCreatedOnBehalfOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsGetCreatedOnBehalfOfOpts struct


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


## ApplicationsGetOwners

> MicrosoftGraphDirectoryObject ApplicationsGetOwners(ctx, applicationId, directoryObjectId, optional)

Get owners from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***ApplicationsGetOwnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsGetOwnersOpts struct


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


## ApplicationsListOwners

> CollectionOfDirectoryObject ApplicationsListOwners(ctx, applicationId, optional)

Get owners from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
 **optional** | ***ApplicationsListOwnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsListOwnersOpts struct


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

