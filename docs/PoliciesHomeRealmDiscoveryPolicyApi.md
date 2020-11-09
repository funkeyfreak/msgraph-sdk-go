# \PoliciesHomeRealmDiscoveryPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateHomeRealmDiscoveryPolicies**](PoliciesHomeRealmDiscoveryPolicyApi.md#PoliciesCreateHomeRealmDiscoveryPolicies) | **Post** /policies/homeRealmDiscoveryPolicies | Create new navigation property to homeRealmDiscoveryPolicies for policies
[**PoliciesGetHomeRealmDiscoveryPolicies**](PoliciesHomeRealmDiscoveryPolicyApi.md#PoliciesGetHomeRealmDiscoveryPolicies) | **Get** /policies/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id} | Get homeRealmDiscoveryPolicies from policies
[**PoliciesListHomeRealmDiscoveryPolicies**](PoliciesHomeRealmDiscoveryPolicyApi.md#PoliciesListHomeRealmDiscoveryPolicies) | **Get** /policies/homeRealmDiscoveryPolicies | Get homeRealmDiscoveryPolicies from policies
[**PoliciesUpdateHomeRealmDiscoveryPolicies**](PoliciesHomeRealmDiscoveryPolicyApi.md#PoliciesUpdateHomeRealmDiscoveryPolicies) | **Patch** /policies/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id} | Update the navigation property homeRealmDiscoveryPolicies in policies



## PoliciesCreateHomeRealmDiscoveryPolicies

> MicrosoftGraphHomeRealmDiscoveryPolicy PoliciesCreateHomeRealmDiscoveryPolicies(ctx, microsoftGraphHomeRealmDiscoveryPolicy)

Create new navigation property to homeRealmDiscoveryPolicies for policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphHomeRealmDiscoveryPolicy** | [**MicrosoftGraphHomeRealmDiscoveryPolicy**](MicrosoftGraphHomeRealmDiscoveryPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphHomeRealmDiscoveryPolicy**](microsoft.graph.homeRealmDiscoveryPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetHomeRealmDiscoveryPolicies

> MicrosoftGraphHomeRealmDiscoveryPolicy PoliciesGetHomeRealmDiscoveryPolicies(ctx, homeRealmDiscoveryPolicyId, optional)

Get homeRealmDiscoveryPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homeRealmDiscoveryPolicyId** | **string**| key: homeRealmDiscoveryPolicy-id of homeRealmDiscoveryPolicy | 
 **optional** | ***PoliciesGetHomeRealmDiscoveryPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesGetHomeRealmDiscoveryPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphHomeRealmDiscoveryPolicy**](microsoft.graph.homeRealmDiscoveryPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListHomeRealmDiscoveryPolicies

> CollectionOfHomeRealmDiscoveryPolicy PoliciesListHomeRealmDiscoveryPolicies(ctx, optional)

Get homeRealmDiscoveryPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesListHomeRealmDiscoveryPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesListHomeRealmDiscoveryPoliciesOpts struct


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

[**CollectionOfHomeRealmDiscoveryPolicy**](Collection_of_homeRealmDiscoveryPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateHomeRealmDiscoveryPolicies

> PoliciesUpdateHomeRealmDiscoveryPolicies(ctx, homeRealmDiscoveryPolicyId, microsoftGraphHomeRealmDiscoveryPolicy)

Update the navigation property homeRealmDiscoveryPolicies in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homeRealmDiscoveryPolicyId** | **string**| key: homeRealmDiscoveryPolicy-id of homeRealmDiscoveryPolicy | 
**microsoftGraphHomeRealmDiscoveryPolicy** | [**MicrosoftGraphHomeRealmDiscoveryPolicy**](MicrosoftGraphHomeRealmDiscoveryPolicy.md)| New navigation property values | 

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

