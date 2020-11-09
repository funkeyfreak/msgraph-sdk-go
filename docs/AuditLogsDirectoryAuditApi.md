# \AuditLogsDirectoryAuditApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogsCreateDirectoryAudits**](AuditLogsDirectoryAuditApi.md#AuditLogsCreateDirectoryAudits) | **Post** /auditLogs/directoryAudits | Create new navigation property to directoryAudits for auditLogs
[**AuditLogsGetDirectoryAudits**](AuditLogsDirectoryAuditApi.md#AuditLogsGetDirectoryAudits) | **Get** /auditLogs/directoryAudits/{directoryAudit-id} | Get directoryAudits from auditLogs
[**AuditLogsListDirectoryAudits**](AuditLogsDirectoryAuditApi.md#AuditLogsListDirectoryAudits) | **Get** /auditLogs/directoryAudits | Get directoryAudits from auditLogs
[**AuditLogsUpdateDirectoryAudits**](AuditLogsDirectoryAuditApi.md#AuditLogsUpdateDirectoryAudits) | **Patch** /auditLogs/directoryAudits/{directoryAudit-id} | Update the navigation property directoryAudits in auditLogs



## AuditLogsCreateDirectoryAudits

> MicrosoftGraphDirectoryAudit AuditLogsCreateDirectoryAudits(ctx, microsoftGraphDirectoryAudit)

Create new navigation property to directoryAudits for auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDirectoryAudit** | [**MicrosoftGraphDirectoryAudit**](MicrosoftGraphDirectoryAudit.md)| New navigation property | 

### Return type

[**MicrosoftGraphDirectoryAudit**](microsoft.graph.directoryAudit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsGetDirectoryAudits

> MicrosoftGraphDirectoryAudit AuditLogsGetDirectoryAudits(ctx, directoryAuditId, optional)

Get directoryAudits from auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryAuditId** | **string**| key: directoryAudit-id of directoryAudit | 
 **optional** | ***AuditLogsGetDirectoryAuditsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditLogsGetDirectoryAuditsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryAudit**](microsoft.graph.directoryAudit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsListDirectoryAudits

> CollectionOfDirectoryAudit AuditLogsListDirectoryAudits(ctx, optional)

Get directoryAudits from auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditLogsListDirectoryAuditsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditLogsListDirectoryAuditsOpts struct


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

[**CollectionOfDirectoryAudit**](Collection_of_directoryAudit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsUpdateDirectoryAudits

> AuditLogsUpdateDirectoryAudits(ctx, directoryAuditId, microsoftGraphDirectoryAudit)

Update the navigation property directoryAudits in auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryAuditId** | **string**| key: directoryAudit-id of directoryAudit | 
**microsoftGraphDirectoryAudit** | [**MicrosoftGraphDirectoryAudit**](MicrosoftGraphDirectoryAudit.md)| New navigation property values | 

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

