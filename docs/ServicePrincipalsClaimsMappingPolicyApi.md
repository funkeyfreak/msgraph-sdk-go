# \ServicePrincipalsClaimsMappingPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsGetClaimsMappingPolicies**](ServicePrincipalsClaimsMappingPolicyApi.md#ServicePrincipalsGetClaimsMappingPolicies) | **Get** /servicePrincipals/{servicePrincipal-id}/claimsMappingPolicies/{claimsMappingPolicy-id} | Get claimsMappingPolicies from servicePrincipals
[**ServicePrincipalsListClaimsMappingPolicies**](ServicePrincipalsClaimsMappingPolicyApi.md#ServicePrincipalsListClaimsMappingPolicies) | **Get** /servicePrincipals/{servicePrincipal-id}/claimsMappingPolicies | Get claimsMappingPolicies from servicePrincipals



## ServicePrincipalsGetClaimsMappingPolicies

> MicrosoftGraphClaimsMappingPolicy ServicePrincipalsGetClaimsMappingPolicies(ctx, servicePrincipalId, claimsMappingPolicyId, optional)

Get claimsMappingPolicies from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**claimsMappingPolicyId** | **string**| key: claimsMappingPolicy-id of claimsMappingPolicy | 
 **optional** | ***ServicePrincipalsGetClaimsMappingPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetClaimsMappingPoliciesOpts struct


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


## ServicePrincipalsListClaimsMappingPolicies

> CollectionOfClaimsMappingPolicy ServicePrincipalsListClaimsMappingPolicies(ctx, servicePrincipalId, optional)

Get claimsMappingPolicies from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListClaimsMappingPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListClaimsMappingPoliciesOpts struct


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

