# \SubscribedSkusSubscribedSkuApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribedSkusSubscribedSkuCreateSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuCreateSubscribedSku) | **Post** /subscribedSkus | Add new entity to subscribedSkus
[**SubscribedSkusSubscribedSkuDeleteSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuDeleteSubscribedSku) | **Delete** /subscribedSkus/{subscribedSku-id} | Delete entity from subscribedSkus
[**SubscribedSkusSubscribedSkuGetSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuGetSubscribedSku) | **Get** /subscribedSkus/{subscribedSku-id} | Get entity from subscribedSkus by key
[**SubscribedSkusSubscribedSkuListSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuListSubscribedSku) | **Get** /subscribedSkus | Get entities from subscribedSkus
[**SubscribedSkusSubscribedSkuUpdateSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuUpdateSubscribedSku) | **Patch** /subscribedSkus/{subscribedSku-id} | Update entity in subscribedSkus



## SubscribedSkusSubscribedSkuCreateSubscribedSku

> MicrosoftGraphSubscribedSku SubscribedSkusSubscribedSkuCreateSubscribedSku(ctx, microsoftGraphSubscribedSku)

Add new entity to subscribedSkus

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSubscribedSku** | [**MicrosoftGraphSubscribedSku**](MicrosoftGraphSubscribedSku.md)| New entity | 

### Return type

[**MicrosoftGraphSubscribedSku**](microsoft.graph.subscribedSku.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribedSkusSubscribedSkuDeleteSubscribedSku

> SubscribedSkusSubscribedSkuDeleteSubscribedSku(ctx, subscribedSkuId, optional)

Delete entity from subscribedSkus

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscribedSkuId** | **string**| key: subscribedSku-id of subscribedSku | 
 **optional** | ***SubscribedSkusSubscribedSkuDeleteSubscribedSkuOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscribedSkusSubscribedSkuDeleteSubscribedSkuOpts struct


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


## SubscribedSkusSubscribedSkuGetSubscribedSku

> MicrosoftGraphSubscribedSku SubscribedSkusSubscribedSkuGetSubscribedSku(ctx, subscribedSkuId, optional)

Get entity from subscribedSkus by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscribedSkuId** | **string**| key: subscribedSku-id of subscribedSku | 
 **optional** | ***SubscribedSkusSubscribedSkuGetSubscribedSkuOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscribedSkusSubscribedSkuGetSubscribedSkuOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphSubscribedSku**](microsoft.graph.subscribedSku.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribedSkusSubscribedSkuListSubscribedSku

> CollectionOfSubscribedSku SubscribedSkusSubscribedSkuListSubscribedSku(ctx, optional)

Get entities from subscribedSkus

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscribedSkusSubscribedSkuListSubscribedSkuOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscribedSkusSubscribedSkuListSubscribedSkuOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfSubscribedSku**](Collection_of_subscribedSku.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribedSkusSubscribedSkuUpdateSubscribedSku

> SubscribedSkusSubscribedSkuUpdateSubscribedSku(ctx, subscribedSkuId, microsoftGraphSubscribedSku)

Update entity in subscribedSkus

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscribedSkuId** | **string**| key: subscribedSku-id of subscribedSku | 
**microsoftGraphSubscribedSku** | [**MicrosoftGraphSubscribedSku**](MicrosoftGraphSubscribedSku.md)| New property values | 

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

