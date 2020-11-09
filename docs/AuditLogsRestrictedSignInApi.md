# \AuditLogsRestrictedSignInApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogsCreateRestrictedSignIns**](AuditLogsRestrictedSignInApi.md#AuditLogsCreateRestrictedSignIns) | **Post** /auditLogs/restrictedSignIns | Create new navigation property to restrictedSignIns for auditLogs
[**AuditLogsGetRestrictedSignIns**](AuditLogsRestrictedSignInApi.md#AuditLogsGetRestrictedSignIns) | **Get** /auditLogs/restrictedSignIns/{restrictedSignIn-id} | Get restrictedSignIns from auditLogs
[**AuditLogsListRestrictedSignIns**](AuditLogsRestrictedSignInApi.md#AuditLogsListRestrictedSignIns) | **Get** /auditLogs/restrictedSignIns | Get restrictedSignIns from auditLogs
[**AuditLogsUpdateRestrictedSignIns**](AuditLogsRestrictedSignInApi.md#AuditLogsUpdateRestrictedSignIns) | **Patch** /auditLogs/restrictedSignIns/{restrictedSignIn-id} | Update the navigation property restrictedSignIns in auditLogs



## AuditLogsCreateRestrictedSignIns

> MicrosoftGraphRestrictedSignIn AuditLogsCreateRestrictedSignIns(ctx, microsoftGraphRestrictedSignIn)

Create new navigation property to restrictedSignIns for auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphRestrictedSignIn** | [**MicrosoftGraphRestrictedSignIn**](MicrosoftGraphRestrictedSignIn.md)| New navigation property | 

### Return type

[**MicrosoftGraphRestrictedSignIn**](microsoft.graph.restrictedSignIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsGetRestrictedSignIns

> MicrosoftGraphRestrictedSignIn AuditLogsGetRestrictedSignIns(ctx, restrictedSignInId, optional)

Get restrictedSignIns from auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restrictedSignInId** | **string**| key: restrictedSignIn-id of restrictedSignIn | 
 **optional** | ***AuditLogsGetRestrictedSignInsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditLogsGetRestrictedSignInsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphRestrictedSignIn**](microsoft.graph.restrictedSignIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsListRestrictedSignIns

> CollectionOfRestrictedSignIn AuditLogsListRestrictedSignIns(ctx, optional)

Get restrictedSignIns from auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditLogsListRestrictedSignInsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditLogsListRestrictedSignInsOpts struct


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

[**CollectionOfRestrictedSignIn**](Collection_of_restrictedSignIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsUpdateRestrictedSignIns

> AuditLogsUpdateRestrictedSignIns(ctx, restrictedSignInId, microsoftGraphRestrictedSignIn)

Update the navigation property restrictedSignIns in auditLogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restrictedSignInId** | **string**| key: restrictedSignIn-id of restrictedSignIn | 
**microsoftGraphRestrictedSignIn** | [**MicrosoftGraphRestrictedSignIn**](MicrosoftGraphRestrictedSignIn.md)| New navigation property values | 

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

