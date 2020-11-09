# \ApplicationsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsAddKey**](ApplicationsActionsApi.md#ApplicationsAddKey) | **Post** /applications/{application-id}/microsoft.graph.addKey | Invoke action addKey
[**ApplicationsAddPassword**](ApplicationsActionsApi.md#ApplicationsAddPassword) | **Post** /applications/{application-id}/microsoft.graph.addPassword | Invoke action addPassword
[**ApplicationsRemoveKey**](ApplicationsActionsApi.md#ApplicationsRemoveKey) | **Post** /applications/{application-id}/microsoft.graph.removeKey | Invoke action removeKey
[**ApplicationsRemovePassword**](ApplicationsActionsApi.md#ApplicationsRemovePassword) | **Post** /applications/{application-id}/microsoft.graph.removePassword | Invoke action removePassword



## ApplicationsAddKey

> MicrosoftGraphKeyCredential ApplicationsAddKey(ctx, applicationId, inlineObject)

Invoke action addKey

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

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


## ApplicationsAddPassword

> MicrosoftGraphPasswordCredential ApplicationsAddPassword(ctx, applicationId, inlineObject1)

Invoke action addPassword

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | 

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


## ApplicationsRemoveKey

> ApplicationsRemoveKey(ctx, applicationId, inlineObject2)

Invoke action removeKey

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | 

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


## ApplicationsRemovePassword

> ApplicationsRemovePassword(ctx, applicationId, inlineObject3)

Invoke action removePassword

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**inlineObject3** | [**InlineObject3**](InlineObject3.md)|  | 

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

