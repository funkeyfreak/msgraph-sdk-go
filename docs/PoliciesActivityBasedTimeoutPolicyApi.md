# \PoliciesActivityBasedTimeoutPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateActivityBasedTimeoutPolicies**](PoliciesActivityBasedTimeoutPolicyApi.md#PoliciesCreateActivityBasedTimeoutPolicies) | **Post** /policies/activityBasedTimeoutPolicies | Create new navigation property to activityBasedTimeoutPolicies for policies
[**PoliciesGetActivityBasedTimeoutPolicies**](PoliciesActivityBasedTimeoutPolicyApi.md#PoliciesGetActivityBasedTimeoutPolicies) | **Get** /policies/activityBasedTimeoutPolicies/{activityBasedTimeoutPolicy-id} | Get activityBasedTimeoutPolicies from policies
[**PoliciesListActivityBasedTimeoutPolicies**](PoliciesActivityBasedTimeoutPolicyApi.md#PoliciesListActivityBasedTimeoutPolicies) | **Get** /policies/activityBasedTimeoutPolicies | Get activityBasedTimeoutPolicies from policies
[**PoliciesUpdateActivityBasedTimeoutPolicies**](PoliciesActivityBasedTimeoutPolicyApi.md#PoliciesUpdateActivityBasedTimeoutPolicies) | **Patch** /policies/activityBasedTimeoutPolicies/{activityBasedTimeoutPolicy-id} | Update the navigation property activityBasedTimeoutPolicies in policies



## PoliciesCreateActivityBasedTimeoutPolicies

> MicrosoftGraphActivityBasedTimeoutPolicy PoliciesCreateActivityBasedTimeoutPolicies(ctx, microsoftGraphActivityBasedTimeoutPolicy)

Create new navigation property to activityBasedTimeoutPolicies for policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphActivityBasedTimeoutPolicy** | [**MicrosoftGraphActivityBasedTimeoutPolicy**](MicrosoftGraphActivityBasedTimeoutPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphActivityBasedTimeoutPolicy**](microsoft.graph.activityBasedTimeoutPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetActivityBasedTimeoutPolicies

> MicrosoftGraphActivityBasedTimeoutPolicy PoliciesGetActivityBasedTimeoutPolicies(ctx, activityBasedTimeoutPolicyId, optional)

Get activityBasedTimeoutPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityBasedTimeoutPolicyId** | **string**| key: activityBasedTimeoutPolicy-id of activityBasedTimeoutPolicy | 
 **optional** | ***PoliciesGetActivityBasedTimeoutPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesGetActivityBasedTimeoutPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphActivityBasedTimeoutPolicy**](microsoft.graph.activityBasedTimeoutPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListActivityBasedTimeoutPolicies

> CollectionOfActivityBasedTimeoutPolicy PoliciesListActivityBasedTimeoutPolicies(ctx, optional)

Get activityBasedTimeoutPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesListActivityBasedTimeoutPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesListActivityBasedTimeoutPoliciesOpts struct


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

[**CollectionOfActivityBasedTimeoutPolicy**](Collection_of_activityBasedTimeoutPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateActivityBasedTimeoutPolicies

> PoliciesUpdateActivityBasedTimeoutPolicies(ctx, activityBasedTimeoutPolicyId, microsoftGraphActivityBasedTimeoutPolicy)

Update the navigation property activityBasedTimeoutPolicies in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityBasedTimeoutPolicyId** | **string**| key: activityBasedTimeoutPolicy-id of activityBasedTimeoutPolicy | 
**microsoftGraphActivityBasedTimeoutPolicy** | [**MicrosoftGraphActivityBasedTimeoutPolicy**](MicrosoftGraphActivityBasedTimeoutPolicy.md)| New navigation property values | 

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

