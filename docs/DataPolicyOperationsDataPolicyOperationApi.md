# \DataPolicyOperationsDataPolicyOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation) | **Post** /dataPolicyOperations | Add new entity to dataPolicyOperations
[**DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation) | **Delete** /dataPolicyOperations/{dataPolicyOperation-id} | Delete entity from dataPolicyOperations
[**DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation) | **Get** /dataPolicyOperations/{dataPolicyOperation-id} | Get entity from dataPolicyOperations by key
[**DataPolicyOperationsDataPolicyOperationListDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationListDataPolicyOperation) | **Get** /dataPolicyOperations | Get entities from dataPolicyOperations
[**DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation) | **Patch** /dataPolicyOperations/{dataPolicyOperation-id} | Update entity in dataPolicyOperations



## DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation

> MicrosoftGraphDataPolicyOperation DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation(ctx, microsoftGraphDataPolicyOperation)

Add new entity to dataPolicyOperations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDataPolicyOperation** | [**MicrosoftGraphDataPolicyOperation**](MicrosoftGraphDataPolicyOperation.md)| New entity | 

### Return type

[**MicrosoftGraphDataPolicyOperation**](microsoft.graph.dataPolicyOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation

> DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation(ctx, dataPolicyOperationId, optional)

Delete entity from dataPolicyOperations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataPolicyOperationId** | **string**| key: dataPolicyOperation-id of dataPolicyOperation | 
 **optional** | ***DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperationOpts struct


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


## DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation

> MicrosoftGraphDataPolicyOperation DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation(ctx, dataPolicyOperationId, optional)

Get entity from dataPolicyOperations by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataPolicyOperationId** | **string**| key: dataPolicyOperation-id of dataPolicyOperation | 
 **optional** | ***DataPolicyOperationsDataPolicyOperationGetDataPolicyOperationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DataPolicyOperationsDataPolicyOperationGetDataPolicyOperationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDataPolicyOperation**](microsoft.graph.dataPolicyOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataPolicyOperationsDataPolicyOperationListDataPolicyOperation

> CollectionOfDataPolicyOperation DataPolicyOperationsDataPolicyOperationListDataPolicyOperation(ctx, optional)

Get entities from dataPolicyOperations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataPolicyOperationsDataPolicyOperationListDataPolicyOperationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DataPolicyOperationsDataPolicyOperationListDataPolicyOperationOpts struct


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

[**CollectionOfDataPolicyOperation**](Collection_of_dataPolicyOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation

> DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation(ctx, dataPolicyOperationId, microsoftGraphDataPolicyOperation)

Update entity in dataPolicyOperations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataPolicyOperationId** | **string**| key: dataPolicyOperation-id of dataPolicyOperation | 
**microsoftGraphDataPolicyOperation** | [**MicrosoftGraphDataPolicyOperation**](MicrosoftGraphDataPolicyOperation.md)| New property values | 

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

