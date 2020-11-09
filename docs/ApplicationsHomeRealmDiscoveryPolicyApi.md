# \ApplicationsHomeRealmDiscoveryPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsGetHomeRealmDiscoveryPolicies**](ApplicationsHomeRealmDiscoveryPolicyApi.md#ApplicationsGetHomeRealmDiscoveryPolicies) | **Get** /applications/{application-id}/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id} | Get homeRealmDiscoveryPolicies from applications
[**ApplicationsListHomeRealmDiscoveryPolicies**](ApplicationsHomeRealmDiscoveryPolicyApi.md#ApplicationsListHomeRealmDiscoveryPolicies) | **Get** /applications/{application-id}/homeRealmDiscoveryPolicies | Get homeRealmDiscoveryPolicies from applications



## ApplicationsGetHomeRealmDiscoveryPolicies

> MicrosoftGraphHomeRealmDiscoveryPolicy ApplicationsGetHomeRealmDiscoveryPolicies(ctx, applicationId, homeRealmDiscoveryPolicyId, optional)

Get homeRealmDiscoveryPolicies from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**homeRealmDiscoveryPolicyId** | **string**| key: homeRealmDiscoveryPolicy-id of homeRealmDiscoveryPolicy | 
 **optional** | ***ApplicationsGetHomeRealmDiscoveryPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsGetHomeRealmDiscoveryPoliciesOpts struct


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


## ApplicationsListHomeRealmDiscoveryPolicies

> CollectionOfHomeRealmDiscoveryPolicy ApplicationsListHomeRealmDiscoveryPolicies(ctx, applicationId, optional)

Get homeRealmDiscoveryPolicies from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
 **optional** | ***ApplicationsListHomeRealmDiscoveryPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsListHomeRealmDiscoveryPoliciesOpts struct


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

