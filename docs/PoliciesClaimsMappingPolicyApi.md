# \PoliciesClaimsMappingPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateClaimsMappingPolicies**](PoliciesClaimsMappingPolicyApi.md#PoliciesCreateClaimsMappingPolicies) | **Post** /policies/claimsMappingPolicies | Create new navigation property to claimsMappingPolicies for policies
[**PoliciesGetClaimsMappingPolicies**](PoliciesClaimsMappingPolicyApi.md#PoliciesGetClaimsMappingPolicies) | **Get** /policies/claimsMappingPolicies/{claimsMappingPolicy-id} | Get claimsMappingPolicies from policies
[**PoliciesListClaimsMappingPolicies**](PoliciesClaimsMappingPolicyApi.md#PoliciesListClaimsMappingPolicies) | **Get** /policies/claimsMappingPolicies | Get claimsMappingPolicies from policies
[**PoliciesUpdateClaimsMappingPolicies**](PoliciesClaimsMappingPolicyApi.md#PoliciesUpdateClaimsMappingPolicies) | **Patch** /policies/claimsMappingPolicies/{claimsMappingPolicy-id} | Update the navigation property claimsMappingPolicies in policies



## PoliciesCreateClaimsMappingPolicies

> MicrosoftGraphClaimsMappingPolicy PoliciesCreateClaimsMappingPolicies(ctx, microsoftGraphClaimsMappingPolicy)

Create new navigation property to claimsMappingPolicies for policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphClaimsMappingPolicy** | [**MicrosoftGraphClaimsMappingPolicy**](MicrosoftGraphClaimsMappingPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphClaimsMappingPolicy**](microsoft.graph.claimsMappingPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetClaimsMappingPolicies

> MicrosoftGraphClaimsMappingPolicy PoliciesGetClaimsMappingPolicies(ctx, claimsMappingPolicyId, optional)

Get claimsMappingPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**claimsMappingPolicyId** | **string**| key: claimsMappingPolicy-id of claimsMappingPolicy | 
 **optional** | ***PoliciesGetClaimsMappingPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesGetClaimsMappingPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphClaimsMappingPolicy**](microsoft.graph.claimsMappingPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListClaimsMappingPolicies

> CollectionOfClaimsMappingPolicy PoliciesListClaimsMappingPolicies(ctx, optional)

Get claimsMappingPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesListClaimsMappingPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesListClaimsMappingPoliciesOpts struct


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

[**CollectionOfClaimsMappingPolicy**](Collection_of_claimsMappingPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateClaimsMappingPolicies

> PoliciesUpdateClaimsMappingPolicies(ctx, claimsMappingPolicyId, microsoftGraphClaimsMappingPolicy)

Update the navigation property claimsMappingPolicies in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**claimsMappingPolicyId** | **string**| key: claimsMappingPolicy-id of claimsMappingPolicy | 
**microsoftGraphClaimsMappingPolicy** | [**MicrosoftGraphClaimsMappingPolicy**](MicrosoftGraphClaimsMappingPolicy.md)| New navigation property values | 

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

