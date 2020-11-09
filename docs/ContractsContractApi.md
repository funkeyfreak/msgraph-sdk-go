# \ContractsContractApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContractsContractCreateContract**](ContractsContractApi.md#ContractsContractCreateContract) | **Post** /contracts | Add new entity to contracts
[**ContractsContractDeleteContract**](ContractsContractApi.md#ContractsContractDeleteContract) | **Delete** /contracts/{contract-id} | Delete entity from contracts
[**ContractsContractGetContract**](ContractsContractApi.md#ContractsContractGetContract) | **Get** /contracts/{contract-id} | Get entity from contracts by key
[**ContractsContractListContract**](ContractsContractApi.md#ContractsContractListContract) | **Get** /contracts | Get entities from contracts
[**ContractsContractUpdateContract**](ContractsContractApi.md#ContractsContractUpdateContract) | **Patch** /contracts/{contract-id} | Update entity in contracts



## ContractsContractCreateContract

> MicrosoftGraphContract ContractsContractCreateContract(ctx, microsoftGraphContract)

Add new entity to contracts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphContract** | [**MicrosoftGraphContract**](MicrosoftGraphContract.md)| New entity | 

### Return type

[**MicrosoftGraphContract**](microsoft.graph.contract.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractDeleteContract

> ContractsContractDeleteContract(ctx, contractId, optional)

Delete entity from contracts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string**| key: contract-id of contract | 
 **optional** | ***ContractsContractDeleteContractOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContractsContractDeleteContractOpts struct


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


## ContractsContractGetContract

> MicrosoftGraphContract ContractsContractGetContract(ctx, contractId, optional)

Get entity from contracts by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string**| key: contract-id of contract | 
 **optional** | ***ContractsContractGetContractOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContractsContractGetContractOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphContract**](microsoft.graph.contract.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractListContract

> CollectionOfContract ContractsContractListContract(ctx, optional)

Get entities from contracts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContractsContractListContractOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContractsContractListContractOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfContract**](Collection_of_contract.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractUpdateContract

> ContractsContractUpdateContract(ctx, contractId, microsoftGraphContract)

Update entity in contracts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string**| key: contract-id of contract | 
**microsoftGraphContract** | [**MicrosoftGraphContract**](MicrosoftGraphContract.md)| New property values | 

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

