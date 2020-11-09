# \PoliciesPolicyRootApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesPolicyRootGetPolicyRoot**](PoliciesPolicyRootApi.md#PoliciesPolicyRootGetPolicyRoot) | **Get** /policies | Get policies
[**PoliciesPolicyRootUpdatePolicyRoot**](PoliciesPolicyRootApi.md#PoliciesPolicyRootUpdatePolicyRoot) | **Patch** /policies | Update policies



## PoliciesPolicyRootGetPolicyRoot

> MicrosoftGraphPolicyRoot PoliciesPolicyRootGetPolicyRoot(ctx, optional)

Get policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesPolicyRootGetPolicyRootOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyRootGetPolicyRootOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPolicyRoot**](microsoft.graph.policyRoot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyRootUpdatePolicyRoot

> PoliciesPolicyRootUpdatePolicyRoot(ctx, microsoftGraphPolicyRoot)

Update policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphPolicyRoot** | [**MicrosoftGraphPolicyRoot**](MicrosoftGraphPolicyRoot.md)| New property values | 

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

