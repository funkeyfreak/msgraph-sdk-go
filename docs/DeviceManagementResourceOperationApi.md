# \DeviceManagementResourceOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateResourceOperations**](DeviceManagementResourceOperationApi.md#DeviceManagementCreateResourceOperations) | **Post** /deviceManagement/resourceOperations | Create new navigation property to resourceOperations for deviceManagement
[**DeviceManagementGetResourceOperations**](DeviceManagementResourceOperationApi.md#DeviceManagementGetResourceOperations) | **Get** /deviceManagement/resourceOperations/{resourceOperation-id} | Get resourceOperations from deviceManagement
[**DeviceManagementListResourceOperations**](DeviceManagementResourceOperationApi.md#DeviceManagementListResourceOperations) | **Get** /deviceManagement/resourceOperations | Get resourceOperations from deviceManagement
[**DeviceManagementUpdateResourceOperations**](DeviceManagementResourceOperationApi.md#DeviceManagementUpdateResourceOperations) | **Patch** /deviceManagement/resourceOperations/{resourceOperation-id} | Update the navigation property resourceOperations in deviceManagement



## DeviceManagementCreateResourceOperations

> MicrosoftGraphResourceOperation DeviceManagementCreateResourceOperations(ctx, microsoftGraphResourceOperation)

Create new navigation property to resourceOperations for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphResourceOperation** | [**MicrosoftGraphResourceOperation**](MicrosoftGraphResourceOperation.md)| New navigation property | 

### Return type

[**MicrosoftGraphResourceOperation**](microsoft.graph.resourceOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetResourceOperations

> MicrosoftGraphResourceOperation DeviceManagementGetResourceOperations(ctx, resourceOperationId, optional)

Get resourceOperations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceOperationId** | **string**| key: resourceOperation-id of resourceOperation | 
 **optional** | ***DeviceManagementGetResourceOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetResourceOperationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphResourceOperation**](microsoft.graph.resourceOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListResourceOperations

> CollectionOfResourceOperation DeviceManagementListResourceOperations(ctx, optional)

Get resourceOperations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListResourceOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListResourceOperationsOpts struct


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

[**CollectionOfResourceOperation**](Collection_of_resourceOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateResourceOperations

> DeviceManagementUpdateResourceOperations(ctx, resourceOperationId, microsoftGraphResourceOperation)

Update the navigation property resourceOperations in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceOperationId** | **string**| key: resourceOperation-id of resourceOperation | 
**microsoftGraphResourceOperation** | [**MicrosoftGraphResourceOperation**](MicrosoftGraphResourceOperation.md)| New navigation property values | 

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

