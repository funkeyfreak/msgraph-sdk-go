# \GroupLifecyclePoliciesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupLifecyclePoliciesAddGroup**](GroupLifecyclePoliciesActionsApi.md#GroupLifecyclePoliciesAddGroup) | **Post** /groupLifecyclePolicies/{groupLifecyclePolicy-id}/microsoft.graph.addGroup | Invoke action addGroup
[**GroupLifecyclePoliciesRemoveGroup**](GroupLifecyclePoliciesActionsApi.md#GroupLifecyclePoliciesRemoveGroup) | **Post** /groupLifecyclePolicies/{groupLifecyclePolicy-id}/microsoft.graph.removeGroup | Invoke action removeGroup



## GroupLifecyclePoliciesAddGroup

> bool GroupLifecyclePoliciesAddGroup(ctx, groupLifecyclePolicyId, inlineObject52)

Invoke action addGroup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string**| key: groupLifecyclePolicy-id of groupLifecyclePolicy | 
**inlineObject52** | [**InlineObject52**](InlineObject52.md)|  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupLifecyclePoliciesRemoveGroup

> bool GroupLifecyclePoliciesRemoveGroup(ctx, groupLifecyclePolicyId, inlineObject53)

Invoke action removeGroup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string**| key: groupLifecyclePolicy-id of groupLifecyclePolicy | 
**inlineObject53** | [**InlineObject53**](InlineObject53.md)|  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

