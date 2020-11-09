# \PoliciesTokenIssuancePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateTokenIssuancePolicies**](PoliciesTokenIssuancePolicyApi.md#PoliciesCreateTokenIssuancePolicies) | **Post** /policies/tokenIssuancePolicies | Create new navigation property to tokenIssuancePolicies for policies
[**PoliciesGetTokenIssuancePolicies**](PoliciesTokenIssuancePolicyApi.md#PoliciesGetTokenIssuancePolicies) | **Get** /policies/tokenIssuancePolicies/{tokenIssuancePolicy-id} | Get tokenIssuancePolicies from policies
[**PoliciesListTokenIssuancePolicies**](PoliciesTokenIssuancePolicyApi.md#PoliciesListTokenIssuancePolicies) | **Get** /policies/tokenIssuancePolicies | Get tokenIssuancePolicies from policies
[**PoliciesUpdateTokenIssuancePolicies**](PoliciesTokenIssuancePolicyApi.md#PoliciesUpdateTokenIssuancePolicies) | **Patch** /policies/tokenIssuancePolicies/{tokenIssuancePolicy-id} | Update the navigation property tokenIssuancePolicies in policies



## PoliciesCreateTokenIssuancePolicies

> MicrosoftGraphTokenIssuancePolicy PoliciesCreateTokenIssuancePolicies(ctx, microsoftGraphTokenIssuancePolicy)

Create new navigation property to tokenIssuancePolicies for policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTokenIssuancePolicy** | [**MicrosoftGraphTokenIssuancePolicy**](MicrosoftGraphTokenIssuancePolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphTokenIssuancePolicy**](microsoft.graph.tokenIssuancePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetTokenIssuancePolicies

> MicrosoftGraphTokenIssuancePolicy PoliciesGetTokenIssuancePolicies(ctx, tokenIssuancePolicyId, optional)

Get tokenIssuancePolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenIssuancePolicyId** | **string**| key: tokenIssuancePolicy-id of tokenIssuancePolicy | 
 **optional** | ***PoliciesGetTokenIssuancePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesGetTokenIssuancePoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTokenIssuancePolicy**](microsoft.graph.tokenIssuancePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListTokenIssuancePolicies

> CollectionOfTokenIssuancePolicy PoliciesListTokenIssuancePolicies(ctx, optional)

Get tokenIssuancePolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesListTokenIssuancePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesListTokenIssuancePoliciesOpts struct


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

[**CollectionOfTokenIssuancePolicy**](Collection_of_tokenIssuancePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateTokenIssuancePolicies

> PoliciesUpdateTokenIssuancePolicies(ctx, tokenIssuancePolicyId, microsoftGraphTokenIssuancePolicy)

Update the navigation property tokenIssuancePolicies in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenIssuancePolicyId** | **string**| key: tokenIssuancePolicy-id of tokenIssuancePolicy | 
**microsoftGraphTokenIssuancePolicy** | [**MicrosoftGraphTokenIssuancePolicy**](MicrosoftGraphTokenIssuancePolicy.md)| New navigation property values | 

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

