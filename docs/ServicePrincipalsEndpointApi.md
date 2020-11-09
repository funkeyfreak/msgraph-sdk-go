# \ServicePrincipalsEndpointApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsCreateEndpoints**](ServicePrincipalsEndpointApi.md#ServicePrincipalsCreateEndpoints) | **Post** /servicePrincipals/{servicePrincipal-id}/endpoints | Create new navigation property to endpoints for servicePrincipals
[**ServicePrincipalsGetEndpoints**](ServicePrincipalsEndpointApi.md#ServicePrincipalsGetEndpoints) | **Get** /servicePrincipals/{servicePrincipal-id}/endpoints/{endpoint-id} | Get endpoints from servicePrincipals
[**ServicePrincipalsListEndpoints**](ServicePrincipalsEndpointApi.md#ServicePrincipalsListEndpoints) | **Get** /servicePrincipals/{servicePrincipal-id}/endpoints | Get endpoints from servicePrincipals
[**ServicePrincipalsUpdateEndpoints**](ServicePrincipalsEndpointApi.md#ServicePrincipalsUpdateEndpoints) | **Patch** /servicePrincipals/{servicePrincipal-id}/endpoints/{endpoint-id} | Update the navigation property endpoints in servicePrincipals



## ServicePrincipalsCreateEndpoints

> MicrosoftGraphEndpoint ServicePrincipalsCreateEndpoints(ctx, servicePrincipalId, microsoftGraphEndpoint)

Create new navigation property to endpoints for servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**microsoftGraphEndpoint** | [**MicrosoftGraphEndpoint**](MicrosoftGraphEndpoint.md)| New navigation property | 

### Return type

[**MicrosoftGraphEndpoint**](microsoft.graph.endpoint.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsGetEndpoints

> MicrosoftGraphEndpoint ServicePrincipalsGetEndpoints(ctx, servicePrincipalId, endpointId, optional)

Get endpoints from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**endpointId** | **string**| key: endpoint-id of endpoint | 
 **optional** | ***ServicePrincipalsGetEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetEndpointsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEndpoint**](microsoft.graph.endpoint.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListEndpoints

> CollectionOfEndpoint ServicePrincipalsListEndpoints(ctx, servicePrincipalId, optional)

Get endpoints from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListEndpointsOpts struct


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

[**CollectionOfEndpoint**](Collection_of_endpoint.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsUpdateEndpoints

> ServicePrincipalsUpdateEndpoints(ctx, servicePrincipalId, endpointId, microsoftGraphEndpoint)

Update the navigation property endpoints in servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**endpointId** | **string**| key: endpoint-id of endpoint | 
**microsoftGraphEndpoint** | [**MicrosoftGraphEndpoint**](MicrosoftGraphEndpoint.md)| New navigation property values | 

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

