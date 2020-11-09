# \PoliciesTokenLifetimePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateTokenLifetimePolicies**](PoliciesTokenLifetimePolicyApi.md#PoliciesCreateTokenLifetimePolicies) | **Post** /policies/tokenLifetimePolicies | Create new navigation property to tokenLifetimePolicies for policies
[**PoliciesGetTokenLifetimePolicies**](PoliciesTokenLifetimePolicyApi.md#PoliciesGetTokenLifetimePolicies) | **Get** /policies/tokenLifetimePolicies/{tokenLifetimePolicy-id} | Get tokenLifetimePolicies from policies
[**PoliciesListTokenLifetimePolicies**](PoliciesTokenLifetimePolicyApi.md#PoliciesListTokenLifetimePolicies) | **Get** /policies/tokenLifetimePolicies | Get tokenLifetimePolicies from policies
[**PoliciesUpdateTokenLifetimePolicies**](PoliciesTokenLifetimePolicyApi.md#PoliciesUpdateTokenLifetimePolicies) | **Patch** /policies/tokenLifetimePolicies/{tokenLifetimePolicy-id} | Update the navigation property tokenLifetimePolicies in policies



## PoliciesCreateTokenLifetimePolicies

> MicrosoftGraphTokenLifetimePolicy PoliciesCreateTokenLifetimePolicies(ctx, microsoftGraphTokenLifetimePolicy)

Create new navigation property to tokenLifetimePolicies for policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTokenLifetimePolicy** | [**MicrosoftGraphTokenLifetimePolicy**](MicrosoftGraphTokenLifetimePolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphTokenLifetimePolicy**](microsoft.graph.tokenLifetimePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetTokenLifetimePolicies

> MicrosoftGraphTokenLifetimePolicy PoliciesGetTokenLifetimePolicies(ctx, tokenLifetimePolicyId, optional)

Get tokenLifetimePolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenLifetimePolicyId** | **string**| key: tokenLifetimePolicy-id of tokenLifetimePolicy | 
 **optional** | ***PoliciesGetTokenLifetimePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesGetTokenLifetimePoliciesOpts struct


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


## PoliciesListTokenLifetimePolicies

> CollectionOfTokenLifetimePolicy PoliciesListTokenLifetimePolicies(ctx, optional)

Get tokenLifetimePolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesListTokenLifetimePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesListTokenLifetimePoliciesOpts struct


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


## PoliciesUpdateTokenLifetimePolicies

> PoliciesUpdateTokenLifetimePolicies(ctx, tokenLifetimePolicyId, microsoftGraphTokenLifetimePolicy)

Update the navigation property tokenLifetimePolicies in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenLifetimePolicyId** | **string**| key: tokenLifetimePolicy-id of tokenLifetimePolicy | 
**microsoftGraphTokenLifetimePolicy** | [**MicrosoftGraphTokenLifetimePolicy**](MicrosoftGraphTokenLifetimePolicy.md)| New navigation property values | 

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

