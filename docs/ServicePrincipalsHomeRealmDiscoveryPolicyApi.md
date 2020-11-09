# \ServicePrincipalsHomeRealmDiscoveryPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsGetHomeRealmDiscoveryPolicies**](ServicePrincipalsHomeRealmDiscoveryPolicyApi.md#ServicePrincipalsGetHomeRealmDiscoveryPolicies) | **Get** /servicePrincipals/{servicePrincipal-id}/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id} | Get homeRealmDiscoveryPolicies from servicePrincipals
[**ServicePrincipalsListHomeRealmDiscoveryPolicies**](ServicePrincipalsHomeRealmDiscoveryPolicyApi.md#ServicePrincipalsListHomeRealmDiscoveryPolicies) | **Get** /servicePrincipals/{servicePrincipal-id}/homeRealmDiscoveryPolicies | Get homeRealmDiscoveryPolicies from servicePrincipals



## ServicePrincipalsGetHomeRealmDiscoveryPolicies

> MicrosoftGraphHomeRealmDiscoveryPolicy ServicePrincipalsGetHomeRealmDiscoveryPolicies(ctx, servicePrincipalId, homeRealmDiscoveryPolicyId, optional)

Get homeRealmDiscoveryPolicies from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**homeRealmDiscoveryPolicyId** | **string**| key: homeRealmDiscoveryPolicy-id of homeRealmDiscoveryPolicy | 
 **optional** | ***ServicePrincipalsGetHomeRealmDiscoveryPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetHomeRealmDiscoveryPoliciesOpts struct


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


## ServicePrincipalsListHomeRealmDiscoveryPolicies

> CollectionOfHomeRealmDiscoveryPolicy ServicePrincipalsListHomeRealmDiscoveryPolicies(ctx, servicePrincipalId, optional)

Get homeRealmDiscoveryPolicies from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListHomeRealmDiscoveryPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListHomeRealmDiscoveryPoliciesOpts struct


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

