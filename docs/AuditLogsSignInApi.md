# \AuditLogsSignInApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogsCreateSignIns**](AuditLogsSignInApi.md#AuditLogsCreateSignIns) | **Post** /auditLogs/signIns | Create new navigation property to signIns for auditLogs
[**AuditLogsGetSignIns**](AuditLogsSignInApi.md#AuditLogsGetSignIns) | **Get** /auditLogs/signIns/{signIn-id} | Get signIns from auditLogs
[**AuditLogsListSignIns**](AuditLogsSignInApi.md#AuditLogsListSignIns) | **Get** /auditLogs/signIns | Get signIns from auditLogs
[**AuditLogsUpdateSignIns**](AuditLogsSignInApi.md#AuditLogsUpdateSignIns) | **Patch** /auditLogs/signIns/{signIn-id} | Update the navigation property signIns in auditLogs



## AuditLogsCreateSignIns

> MicrosoftGraphSignIn AuditLogsCreateSignIns(ctx, microsoftGraphSignIn)

Create new navigation property to signIns for auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSignIn** | [**MicrosoftGraphSignIn**](MicrosoftGraphSignIn.md)| New navigation property | 

### Return type

[**MicrosoftGraphSignIn**](microsoft.graph.signIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsGetSignIns

> MicrosoftGraphSignIn AuditLogsGetSignIns(ctx, signInId, optional)

Get signIns from auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signInId** | **string**| key: signIn-id of signIn | 
 **optional** | ***AuditLogsGetSignInsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditLogsGetSignInsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSignIn**](microsoft.graph.signIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsListSignIns

> CollectionOfSignIn AuditLogsListSignIns(ctx, optional)

Get signIns from auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditLogsListSignInsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditLogsListSignInsOpts struct


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

[**CollectionOfSignIn**](Collection_of_signIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsUpdateSignIns

> AuditLogsUpdateSignIns(ctx, signInId, microsoftGraphSignIn)

Update the navigation property signIns in auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signInId** | **string**| key: signIn-id of signIn | 
**microsoftGraphSignIn** | [**MicrosoftGraphSignIn**](MicrosoftGraphSignIn.md)| New navigation property values | 

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

