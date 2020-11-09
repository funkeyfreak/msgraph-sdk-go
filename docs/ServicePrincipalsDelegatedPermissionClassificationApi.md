# \ServicePrincipalsDelegatedPermissionClassificationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsCreateDelegatedPermissionClassifications**](ServicePrincipalsDelegatedPermissionClassificationApi.md#ServicePrincipalsCreateDelegatedPermissionClassifications) | **Post** /servicePrincipals/{servicePrincipal-id}/delegatedPermissionClassifications | Create new navigation property to delegatedPermissionClassifications for servicePrincipals
[**ServicePrincipalsGetDelegatedPermissionClassifications**](ServicePrincipalsDelegatedPermissionClassificationApi.md#ServicePrincipalsGetDelegatedPermissionClassifications) | **Get** /servicePrincipals/{servicePrincipal-id}/delegatedPermissionClassifications/{delegatedPermissionClassification-id} | Get delegatedPermissionClassifications from servicePrincipals
[**ServicePrincipalsListDelegatedPermissionClassifications**](ServicePrincipalsDelegatedPermissionClassificationApi.md#ServicePrincipalsListDelegatedPermissionClassifications) | **Get** /servicePrincipals/{servicePrincipal-id}/delegatedPermissionClassifications | Get delegatedPermissionClassifications from servicePrincipals
[**ServicePrincipalsUpdateDelegatedPermissionClassifications**](ServicePrincipalsDelegatedPermissionClassificationApi.md#ServicePrincipalsUpdateDelegatedPermissionClassifications) | **Patch** /servicePrincipals/{servicePrincipal-id}/delegatedPermissionClassifications/{delegatedPermissionClassification-id} | Update the navigation property delegatedPermissionClassifications in servicePrincipals



## ServicePrincipalsCreateDelegatedPermissionClassifications

> MicrosoftGraphDelegatedPermissionClassification ServicePrincipalsCreateDelegatedPermissionClassifications(ctx, servicePrincipalId, microsoftGraphDelegatedPermissionClassification)

Create new navigation property to delegatedPermissionClassifications for servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**microsoftGraphDelegatedPermissionClassification** | [**MicrosoftGraphDelegatedPermissionClassification**](MicrosoftGraphDelegatedPermissionClassification.md)| New navigation property | 

### Return type

[**MicrosoftGraphDelegatedPermissionClassification**](microsoft.graph.delegatedPermissionClassification.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsGetDelegatedPermissionClassifications

> MicrosoftGraphDelegatedPermissionClassification ServicePrincipalsGetDelegatedPermissionClassifications(ctx, servicePrincipalId, delegatedPermissionClassificationId, optional)

Get delegatedPermissionClassifications from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**delegatedPermissionClassificationId** | **string**| key: delegatedPermissionClassification-id of delegatedPermissionClassification | 
 **optional** | ***ServicePrincipalsGetDelegatedPermissionClassificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetDelegatedPermissionClassificationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDelegatedPermissionClassification**](microsoft.graph.delegatedPermissionClassification.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListDelegatedPermissionClassifications

> CollectionOfDelegatedPermissionClassification ServicePrincipalsListDelegatedPermissionClassifications(ctx, servicePrincipalId, optional)

Get delegatedPermissionClassifications from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListDelegatedPermissionClassificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListDelegatedPermissionClassificationsOpts struct


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

[**CollectionOfDelegatedPermissionClassification**](Collection_of_delegatedPermissionClassification.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsUpdateDelegatedPermissionClassifications

> ServicePrincipalsUpdateDelegatedPermissionClassifications(ctx, servicePrincipalId, delegatedPermissionClassificationId, microsoftGraphDelegatedPermissionClassification)

Update the navigation property delegatedPermissionClassifications in servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**delegatedPermissionClassificationId** | **string**| key: delegatedPermissionClassification-id of delegatedPermissionClassification | 
**microsoftGraphDelegatedPermissionClassification** | [**MicrosoftGraphDelegatedPermissionClassification**](MicrosoftGraphDelegatedPermissionClassification.md)| New navigation property values | 

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

