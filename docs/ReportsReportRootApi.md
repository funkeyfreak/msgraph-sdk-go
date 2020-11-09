# \ReportsReportRootApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportsReportRootGetReportRoot**](ReportsReportRootApi.md#ReportsReportRootGetReportRoot) | **Get** /reports | Get reports
[**ReportsReportRootUpdateReportRoot**](ReportsReportRootApi.md#ReportsReportRootUpdateReportRoot) | **Patch** /reports | Update reports



## ReportsReportRootGetReportRoot

> MicrosoftGraphReportRoot ReportsReportRootGetReportRoot(ctx, optional)

Get reports

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsReportRootGetReportRootOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReportsReportRootGetReportRootOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphReportRoot**](microsoft.graph.reportRoot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsReportRootUpdateReportRoot

> ReportsReportRootUpdateReportRoot(ctx, microsoftGraphReportRoot)

Update reports

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphReportRoot** | [**MicrosoftGraphReportRoot**](MicrosoftGraphReportRoot.md)| New property values | 

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

