# \GroupsGroupLifecyclePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateGroupLifecyclePolicies**](GroupsGroupLifecyclePolicyApi.md#GroupsCreateGroupLifecyclePolicies) | **Post** /groups/{group-id}/groupLifecyclePolicies | Create new navigation property to groupLifecyclePolicies for groups
[**GroupsGetGroupLifecyclePolicies**](GroupsGroupLifecyclePolicyApi.md#GroupsGetGroupLifecyclePolicies) | **Get** /groups/{group-id}/groupLifecyclePolicies/{groupLifecyclePolicy-id} | Get groupLifecyclePolicies from groups
[**GroupsListGroupLifecyclePolicies**](GroupsGroupLifecyclePolicyApi.md#GroupsListGroupLifecyclePolicies) | **Get** /groups/{group-id}/groupLifecyclePolicies | Get groupLifecyclePolicies from groups
[**GroupsUpdateGroupLifecyclePolicies**](GroupsGroupLifecyclePolicyApi.md#GroupsUpdateGroupLifecyclePolicies) | **Patch** /groups/{group-id}/groupLifecyclePolicies/{groupLifecyclePolicy-id} | Update the navigation property groupLifecyclePolicies in groups



## GroupsCreateGroupLifecyclePolicies

> MicrosoftGraphGroupLifecyclePolicy GroupsCreateGroupLifecyclePolicies(ctx, groupId, microsoftGraphGroupLifecyclePolicy)

Create new navigation property to groupLifecyclePolicies for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphGroupLifecyclePolicy** | [**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md)| New navigation property | 

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


## GroupsGetGroupLifecyclePolicies

> MicrosoftGraphGroupLifecyclePolicy GroupsGetGroupLifecyclePolicies(ctx, groupId, groupLifecyclePolicyId, optional)

Get groupLifecyclePolicies from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**groupLifecyclePolicyId** | **string**| key: groupLifecyclePolicy-id of groupLifecyclePolicy | 
 **optional** | ***GroupsGetGroupLifecyclePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetGroupLifecyclePoliciesOpts struct


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


## GroupsListGroupLifecyclePolicies

> CollectionOfGroupLifecyclePolicy GroupsListGroupLifecyclePolicies(ctx, groupId, optional)

Get groupLifecyclePolicies from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsListGroupLifecyclePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsListGroupLifecyclePoliciesOpts struct


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


## GroupsUpdateGroupLifecyclePolicies

> GroupsUpdateGroupLifecyclePolicies(ctx, groupId, groupLifecyclePolicyId, microsoftGraphGroupLifecyclePolicy)

Update the navigation property groupLifecyclePolicies in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**groupLifecyclePolicyId** | **string**| key: groupLifecyclePolicy-id of groupLifecyclePolicy | 
**microsoftGraphGroupLifecyclePolicy** | [**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md)| New navigation property values | 

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

