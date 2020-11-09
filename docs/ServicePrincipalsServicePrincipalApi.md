# \ServicePrincipalsServicePrincipalApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsServicePrincipalCreateServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalCreateServicePrincipal) | **Post** /servicePrincipals | Add new entity to servicePrincipals
[**ServicePrincipalsServicePrincipalDeleteServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalDeleteServicePrincipal) | **Delete** /servicePrincipals/{servicePrincipal-id} | Delete entity from servicePrincipals
[**ServicePrincipalsServicePrincipalGetServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalGetServicePrincipal) | **Get** /servicePrincipals/{servicePrincipal-id} | Get entity from servicePrincipals by key
[**ServicePrincipalsServicePrincipalListServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalListServicePrincipal) | **Get** /servicePrincipals | Get entities from servicePrincipals
[**ServicePrincipalsServicePrincipalUpdateServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalUpdateServicePrincipal) | **Patch** /servicePrincipals/{servicePrincipal-id} | Update entity in servicePrincipals



## ServicePrincipalsServicePrincipalCreateServicePrincipal

> MicrosoftGraphServicePrincipal ServicePrincipalsServicePrincipalCreateServicePrincipal(ctx, microsoftGraphServicePrincipal)

Add new entity to servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphServicePrincipal** | [**MicrosoftGraphServicePrincipal**](MicrosoftGraphServicePrincipal.md)| New entity | 

### Return type

[**MicrosoftGraphServicePrincipal**](microsoft.graph.servicePrincipal.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsServicePrincipalDeleteServicePrincipal

> ServicePrincipalsServicePrincipalDeleteServicePrincipal(ctx, servicePrincipalId, optional)

Delete entity from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsServicePrincipalDeleteServicePrincipalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsServicePrincipalDeleteServicePrincipalOpts struct


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


## ServicePrincipalsServicePrincipalGetServicePrincipal

> MicrosoftGraphServicePrincipal ServicePrincipalsServicePrincipalGetServicePrincipal(ctx, servicePrincipalId, optional)

Get entity from servicePrincipals by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsServicePrincipalGetServicePrincipalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsServicePrincipalGetServicePrincipalOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphServicePrincipal**](microsoft.graph.servicePrincipal.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsServicePrincipalListServicePrincipal

> CollectionOfServicePrincipal ServicePrincipalsServicePrincipalListServicePrincipal(ctx, optional)

Get entities from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicePrincipalsServicePrincipalListServicePrincipalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsServicePrincipalListServicePrincipalOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfServicePrincipal**](Collection_of_servicePrincipal.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsServicePrincipalUpdateServicePrincipal

> ServicePrincipalsServicePrincipalUpdateServicePrincipal(ctx, servicePrincipalId, microsoftGraphServicePrincipal)

Update entity in servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**microsoftGraphServicePrincipal** | [**MicrosoftGraphServicePrincipal**](MicrosoftGraphServicePrincipal.md)| New property values | 

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

