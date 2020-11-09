# \AuditLogsAuditLogRootApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogsAuditLogRootGetAuditLogRoot**](AuditLogsAuditLogRootApi.md#AuditLogsAuditLogRootGetAuditLogRoot) | **Get** /auditLogs | Get auditLogs
[**AuditLogsAuditLogRootUpdateAuditLogRoot**](AuditLogsAuditLogRootApi.md#AuditLogsAuditLogRootUpdateAuditLogRoot) | **Patch** /auditLogs | Update auditLogs



## AuditLogsAuditLogRootGetAuditLogRoot

> MicrosoftGraphAuditLogRoot AuditLogsAuditLogRootGetAuditLogRoot(ctx, optional)

Get auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditLogsAuditLogRootGetAuditLogRootOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditLogsAuditLogRootGetAuditLogRootOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAuditLogRoot**](microsoft.graph.auditLogRoot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsAuditLogRootUpdateAuditLogRoot

> AuditLogsAuditLogRootUpdateAuditLogRoot(ctx, microsoftGraphAuditLogRoot)

Update auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphAuditLogRoot** | [**MicrosoftGraphAuditLogRoot**](MicrosoftGraphAuditLogRoot.md)| New property values | 

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

