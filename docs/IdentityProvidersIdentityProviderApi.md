# \IdentityProvidersIdentityProviderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityProvidersIdentityProviderCreateIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderCreateIdentityProvider) | **Post** /identityProviders | Add new entity to identityProviders
[**IdentityProvidersIdentityProviderDeleteIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderDeleteIdentityProvider) | **Delete** /identityProviders/{identityProvider-id} | Delete entity from identityProviders
[**IdentityProvidersIdentityProviderGetIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderGetIdentityProvider) | **Get** /identityProviders/{identityProvider-id} | Get entity from identityProviders by key
[**IdentityProvidersIdentityProviderListIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderListIdentityProvider) | **Get** /identityProviders | Get entities from identityProviders
[**IdentityProvidersIdentityProviderUpdateIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderUpdateIdentityProvider) | **Patch** /identityProviders/{identityProvider-id} | Update entity in identityProviders



## IdentityProvidersIdentityProviderCreateIdentityProvider

> MicrosoftGraphIdentityProvider IdentityProvidersIdentityProviderCreateIdentityProvider(ctx, microsoftGraphIdentityProvider)

Add new entity to identityProviders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphIdentityProvider** | [**MicrosoftGraphIdentityProvider**](MicrosoftGraphIdentityProvider.md)| New entity | 

### Return type

[**MicrosoftGraphIdentityProvider**](microsoft.graph.identityProvider.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProvidersIdentityProviderDeleteIdentityProvider

> IdentityProvidersIdentityProviderDeleteIdentityProvider(ctx, identityProviderId, optional)

Delete entity from identityProviders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProviderId** | **string**| key: identityProvider-id of identityProvider | 
 **optional** | ***IdentityProvidersIdentityProviderDeleteIdentityProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IdentityProvidersIdentityProviderDeleteIdentityProviderOpts struct


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


## IdentityProvidersIdentityProviderGetIdentityProvider

> MicrosoftGraphIdentityProvider IdentityProvidersIdentityProviderGetIdentityProvider(ctx, identityProviderId, optional)

Get entity from identityProviders by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProviderId** | **string**| key: identityProvider-id of identityProvider | 
 **optional** | ***IdentityProvidersIdentityProviderGetIdentityProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IdentityProvidersIdentityProviderGetIdentityProviderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphIdentityProvider**](microsoft.graph.identityProvider.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProvidersIdentityProviderListIdentityProvider

> CollectionOfIdentityProvider IdentityProvidersIdentityProviderListIdentityProvider(ctx, optional)

Get entities from identityProviders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IdentityProvidersIdentityProviderListIdentityProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IdentityProvidersIdentityProviderListIdentityProviderOpts struct


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

[**CollectionOfIdentityProvider**](Collection_of_identityProvider.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProvidersIdentityProviderUpdateIdentityProvider

> IdentityProvidersIdentityProviderUpdateIdentityProvider(ctx, identityProviderId, microsoftGraphIdentityProvider)

Update entity in identityProviders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProviderId** | **string**| key: identityProvider-id of identityProvider | 
**microsoftGraphIdentityProvider** | [**MicrosoftGraphIdentityProvider**](MicrosoftGraphIdentityProvider.md)| New property values | 

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

