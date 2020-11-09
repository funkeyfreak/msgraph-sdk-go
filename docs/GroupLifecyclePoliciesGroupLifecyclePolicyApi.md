# \GroupLifecyclePoliciesGroupLifecyclePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy) | **Post** /groupLifecyclePolicies | Add new entity to groupLifecyclePolicies
[**GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy) | **Delete** /groupLifecyclePolicies/{groupLifecyclePolicy-id} | Delete entity from groupLifecyclePolicies
[**GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy) | **Get** /groupLifecyclePolicies/{groupLifecyclePolicy-id} | Get entity from groupLifecyclePolicies by key
[**GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy) | **Get** /groupLifecyclePolicies | Get entities from groupLifecyclePolicies
[**GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy) | **Patch** /groupLifecyclePolicies/{groupLifecyclePolicy-id} | Update entity in groupLifecyclePolicies



## GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy

> MicrosoftGraphGroupLifecyclePolicy GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy(ctx, microsoftGraphGroupLifecyclePolicy)

Add new entity to groupLifecyclePolicies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphGroupLifecyclePolicy** | [**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md)| New entity | 

### Return type

[**MicrosoftGraphGroupLifecyclePolicy**](microsoft.graph.groupLifecyclePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy

> GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy(ctx, groupLifecyclePolicyId, optional)

Delete entity from groupLifecyclePolicies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string**| key: groupLifecyclePolicy-id of groupLifecyclePolicy | 
 **optional** | ***GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicyOpts struct


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


## GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy

> MicrosoftGraphGroupLifecyclePolicy GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy(ctx, groupLifecyclePolicyId, optional)

Get entity from groupLifecyclePolicies by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string**| key: groupLifecyclePolicy-id of groupLifecyclePolicy | 
 **optional** | ***GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphGroupLifecyclePolicy**](microsoft.graph.groupLifecyclePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy

> CollectionOfGroupLifecyclePolicy GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy(ctx, optional)

Get entities from groupLifecyclePolicies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicyOpts struct


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

[**CollectionOfGroupLifecyclePolicy**](Collection_of_groupLifecyclePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy

> GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy(ctx, groupLifecyclePolicyId, microsoftGraphGroupLifecyclePolicy)

Update entity in groupLifecyclePolicies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string**| key: groupLifecyclePolicy-id of groupLifecyclePolicy | 
**microsoftGraphGroupLifecyclePolicy** | [**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md)| New property values | 

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

