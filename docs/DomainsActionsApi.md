# \DomainsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsForceDelete**](DomainsActionsApi.md#DomainsForceDelete) | **Post** /domains/{domain-id}/microsoft.graph.forceDelete | Invoke action forceDelete
[**DomainsVerify**](DomainsActionsApi.md#DomainsVerify) | **Post** /domains/{domain-id}/microsoft.graph.verify | Invoke action verify



## DomainsForceDelete

> DomainsForceDelete(ctx, domainId, inlineObject51)

Invoke action forceDelete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**inlineObject51** | [**InlineObject51**](InlineObject51.md)|  | 

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


## DomainsVerify

> interface{} DomainsVerify(ctx, domainId)

Invoke action verify

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

