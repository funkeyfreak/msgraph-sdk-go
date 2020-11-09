# \ServicePrincipalsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsAddKey**](ServicePrincipalsActionsApi.md#ServicePrincipalsAddKey) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.addKey | Invoke action addKey
[**ServicePrincipalsAddPassword**](ServicePrincipalsActionsApi.md#ServicePrincipalsAddPassword) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.addPassword | Invoke action addPassword
[**ServicePrincipalsRemoveKey**](ServicePrincipalsActionsApi.md#ServicePrincipalsRemoveKey) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.removeKey | Invoke action removeKey
[**ServicePrincipalsRemovePassword**](ServicePrincipalsActionsApi.md#ServicePrincipalsRemovePassword) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.removePassword | Invoke action removePassword



## ServicePrincipalsAddKey

> MicrosoftGraphKeyCredential ServicePrincipalsAddKey(ctx, servicePrincipalId, inlineObject520)

Invoke action addKey

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**inlineObject520** | [**InlineObject520**](InlineObject520.md)|  | 

### Return type

[**MicrosoftGraphKeyCredential**](microsoft.graph.keyCredential.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsAddPassword

> MicrosoftGraphPasswordCredential ServicePrincipalsAddPassword(ctx, servicePrincipalId, inlineObject521)

Invoke action addPassword

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**inlineObject521** | [**InlineObject521**](InlineObject521.md)|  | 

### Return type

[**MicrosoftGraphPasswordCredential**](microsoft.graph.passwordCredential.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsRemoveKey

> ServicePrincipalsRemoveKey(ctx, servicePrincipalId, inlineObject522)

Invoke action removeKey

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**inlineObject522** | [**InlineObject522**](InlineObject522.md)|  | 

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


## ServicePrincipalsRemovePassword

> ServicePrincipalsRemovePassword(ctx, servicePrincipalId, inlineObject523)

Invoke action removePassword

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**inlineObject523** | [**InlineObject523**](InlineObject523.md)|  | 

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

