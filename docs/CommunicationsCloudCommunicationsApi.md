# \CommunicationsCloudCommunicationsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCloudCommunicationsGetCloudCommunications**](CommunicationsCloudCommunicationsApi.md#CommunicationsCloudCommunicationsGetCloudCommunications) | **Get** /communications | Get communications
[**CommunicationsCloudCommunicationsUpdateCloudCommunications**](CommunicationsCloudCommunicationsApi.md#CommunicationsCloudCommunicationsUpdateCloudCommunications) | **Patch** /communications | Update communications



## CommunicationsCloudCommunicationsGetCloudCommunications

> MicrosoftGraphCloudCommunications CommunicationsCloudCommunicationsGetCloudCommunications(ctx, optional)

Get communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommunicationsCloudCommunicationsGetCloudCommunicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsCloudCommunicationsGetCloudCommunicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphCloudCommunications**](microsoft.graph.cloudCommunications.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCloudCommunicationsUpdateCloudCommunications

> CommunicationsCloudCommunicationsUpdateCloudCommunications(ctx, microsoftGraphCloudCommunications)

Update communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphCloudCommunications** | [**MicrosoftGraphCloudCommunications**](MicrosoftGraphCloudCommunications.md)| New property values | 

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

