# \ApplicationsTokenLifetimePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsGetTokenLifetimePolicies**](ApplicationsTokenLifetimePolicyApi.md#ApplicationsGetTokenLifetimePolicies) | **Get** /applications/{application-id}/tokenLifetimePolicies/{tokenLifetimePolicy-id} | Get tokenLifetimePolicies from applications
[**ApplicationsListTokenLifetimePolicies**](ApplicationsTokenLifetimePolicyApi.md#ApplicationsListTokenLifetimePolicies) | **Get** /applications/{application-id}/tokenLifetimePolicies | Get tokenLifetimePolicies from applications



## ApplicationsGetTokenLifetimePolicies

> MicrosoftGraphTokenLifetimePolicy ApplicationsGetTokenLifetimePolicies(ctx, applicationId, tokenLifetimePolicyId, optional)

Get tokenLifetimePolicies from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**tokenLifetimePolicyId** | **string**| key: tokenLifetimePolicy-id of tokenLifetimePolicy | 
 **optional** | ***ApplicationsGetTokenLifetimePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsGetTokenLifetimePoliciesOpts struct


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


## ApplicationsListTokenLifetimePolicies

> CollectionOfTokenLifetimePolicy ApplicationsListTokenLifetimePolicies(ctx, applicationId, optional)

Get tokenLifetimePolicies from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
 **optional** | ***ApplicationsListTokenLifetimePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsListTokenLifetimePoliciesOpts struct


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

