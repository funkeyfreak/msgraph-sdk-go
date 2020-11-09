# \ServicePrincipalsOAuth2PermissionGrantApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsGetOauth2PermissionGrants**](ServicePrincipalsOAuth2PermissionGrantApi.md#ServicePrincipalsGetOauth2PermissionGrants) | **Get** /servicePrincipals/{servicePrincipal-id}/oauth2PermissionGrants/{oAuth2PermissionGrant-id} | Get oauth2PermissionGrants from servicePrincipals
[**ServicePrincipalsListOauth2PermissionGrants**](ServicePrincipalsOAuth2PermissionGrantApi.md#ServicePrincipalsListOauth2PermissionGrants) | **Get** /servicePrincipals/{servicePrincipal-id}/oauth2PermissionGrants | Get oauth2PermissionGrants from servicePrincipals



## ServicePrincipalsGetOauth2PermissionGrants

> MicrosoftGraphOAuth2PermissionGrant ServicePrincipalsGetOauth2PermissionGrants(ctx, servicePrincipalId, oAuth2PermissionGrantId, optional)

Get oauth2PermissionGrants from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**oAuth2PermissionGrantId** | **string**| key: oAuth2PermissionGrant-id of oAuth2PermissionGrant | 
 **optional** | ***ServicePrincipalsGetOauth2PermissionGrantsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetOauth2PermissionGrantsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOAuth2PermissionGrant**](microsoft.graph.oAuth2PermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListOauth2PermissionGrants

> CollectionOfOAuth2PermissionGrant ServicePrincipalsListOauth2PermissionGrants(ctx, servicePrincipalId, optional)

Get oauth2PermissionGrants from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListOauth2PermissionGrantsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListOauth2PermissionGrantsOpts struct


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

[**CollectionOfOAuth2PermissionGrant**](Collection_of_oAuth2PermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

