# \ServicePrincipalsTokenLifetimePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsGetTokenLifetimePolicies**](ServicePrincipalsTokenLifetimePolicyApi.md#ServicePrincipalsGetTokenLifetimePolicies) | **Get** /servicePrincipals/{servicePrincipal-id}/tokenLifetimePolicies/{tokenLifetimePolicy-id} | Get tokenLifetimePolicies from servicePrincipals
[**ServicePrincipalsListTokenLifetimePolicies**](ServicePrincipalsTokenLifetimePolicyApi.md#ServicePrincipalsListTokenLifetimePolicies) | **Get** /servicePrincipals/{servicePrincipal-id}/tokenLifetimePolicies | Get tokenLifetimePolicies from servicePrincipals



## ServicePrincipalsGetTokenLifetimePolicies

> MicrosoftGraphTokenLifetimePolicy ServicePrincipalsGetTokenLifetimePolicies(ctx, servicePrincipalId, tokenLifetimePolicyId, optional)

Get tokenLifetimePolicies from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**tokenLifetimePolicyId** | **string**| key: tokenLifetimePolicy-id of tokenLifetimePolicy | 
 **optional** | ***ServicePrincipalsGetTokenLifetimePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetTokenLifetimePoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTokenLifetimePolicy**](microsoft.graph.tokenLifetimePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListTokenLifetimePolicies

> CollectionOfTokenLifetimePolicy ServicePrincipalsListTokenLifetimePolicies(ctx, servicePrincipalId, optional)

Get tokenLifetimePolicies from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListTokenLifetimePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListTokenLifetimePoliciesOpts struct


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

[**CollectionOfTokenLifetimePolicy**](Collection_of_tokenLifetimePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

