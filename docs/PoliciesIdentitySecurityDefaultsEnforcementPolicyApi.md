# \PoliciesIdentitySecurityDefaultsEnforcementPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesGetIdentitySecurityDefaultsEnforcementPolicy**](PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.md#PoliciesGetIdentitySecurityDefaultsEnforcementPolicy) | **Get** /policies/identitySecurityDefaultsEnforcementPolicy | Get identitySecurityDefaultsEnforcementPolicy from policies
[**PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy**](PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.md#PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy) | **Patch** /policies/identitySecurityDefaultsEnforcementPolicy | Update the navigation property identitySecurityDefaultsEnforcementPolicy in policies



## PoliciesGetIdentitySecurityDefaultsEnforcementPolicy

> MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy PoliciesGetIdentitySecurityDefaultsEnforcementPolicy(ctx, optional)

Get identitySecurityDefaultsEnforcementPolicy from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesGetIdentitySecurityDefaultsEnforcementPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesGetIdentitySecurityDefaultsEnforcementPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy**](microsoft.graph.identitySecurityDefaultsEnforcementPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy

> PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy(ctx, microsoftGraphIdentitySecurityDefaultsEnforcementPolicy)

Update the navigation property identitySecurityDefaultsEnforcementPolicy in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphIdentitySecurityDefaultsEnforcementPolicy** | [**MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy**](MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy.md)| New navigation property values | 

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

