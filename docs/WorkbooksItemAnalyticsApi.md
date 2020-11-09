# \WorkbooksItemAnalyticsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksGetAnalytics**](WorkbooksItemAnalyticsApi.md#WorkbooksGetAnalytics) | **Get** /workbooks/{driveItem-id}/analytics | Get analytics from workbooks



## WorkbooksGetAnalytics

> MicrosoftGraphItemAnalytics WorkbooksGetAnalytics(ctx, driveItemId, optional)

Get analytics from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksGetAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksGetAnalyticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphItemAnalytics**](microsoft.graph.itemAnalytics.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

