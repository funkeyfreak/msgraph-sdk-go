# \PoliciesConditionalAccessPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateConditionalAccessPolicies**](PoliciesConditionalAccessPolicyApi.md#PoliciesCreateConditionalAccessPolicies) | **Post** /policies/conditionalAccessPolicies | Create new navigation property to conditionalAccessPolicies for policies
[**PoliciesGetConditionalAccessPolicies**](PoliciesConditionalAccessPolicyApi.md#PoliciesGetConditionalAccessPolicies) | **Get** /policies/conditionalAccessPolicies/{conditionalAccessPolicy-id} | Get conditionalAccessPolicies from policies
[**PoliciesListConditionalAccessPolicies**](PoliciesConditionalAccessPolicyApi.md#PoliciesListConditionalAccessPolicies) | **Get** /policies/conditionalAccessPolicies | Get conditionalAccessPolicies from policies
[**PoliciesUpdateConditionalAccessPolicies**](PoliciesConditionalAccessPolicyApi.md#PoliciesUpdateConditionalAccessPolicies) | **Patch** /policies/conditionalAccessPolicies/{conditionalAccessPolicy-id} | Update the navigation property conditionalAccessPolicies in policies



## PoliciesCreateConditionalAccessPolicies

> MicrosoftGraphConditionalAccessPolicy PoliciesCreateConditionalAccessPolicies(ctx, microsoftGraphConditionalAccessPolicy)

Create new navigation property to conditionalAccessPolicies for policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphConditionalAccessPolicy** | [**MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphConditionalAccessPolicy**](microsoft.graph.conditionalAccessPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetConditionalAccessPolicies

> MicrosoftGraphConditionalAccessPolicy PoliciesGetConditionalAccessPolicies(ctx, conditionalAccessPolicyId, optional)

Get conditionalAccessPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionalAccessPolicyId** | **string**| key: conditionalAccessPolicy-id of conditionalAccessPolicy | 
 **optional** | ***PoliciesGetConditionalAccessPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesGetConditionalAccessPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphConditionalAccessPolicy**](microsoft.graph.conditionalAccessPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListConditionalAccessPolicies

> CollectionOfConditionalAccessPolicy PoliciesListConditionalAccessPolicies(ctx, optional)

Get conditionalAccessPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesListConditionalAccessPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesListConditionalAccessPoliciesOpts struct


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

[**CollectionOfConditionalAccessPolicy**](Collection_of_conditionalAccessPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateConditionalAccessPolicies

> PoliciesUpdateConditionalAccessPolicies(ctx, conditionalAccessPolicyId, microsoftGraphConditionalAccessPolicy)

Update the navigation property conditionalAccessPolicies in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionalAccessPolicyId** | **string**| key: conditionalAccessPolicy-id of conditionalAccessPolicy | 
**microsoftGraphConditionalAccessPolicy** | [**MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md)| New navigation property values | 

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

