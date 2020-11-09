# \SecuritySecurityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecuritySecurityGetSecurity**](SecuritySecurityApi.md#SecuritySecurityGetSecurity) | **Get** /Security | Get Security
[**SecuritySecurityUpdateSecurity**](SecuritySecurityApi.md#SecuritySecurityUpdateSecurity) | **Patch** /Security | Update Security



## SecuritySecurityGetSecurity

> MicrosoftGraphSecurity SecuritySecurityGetSecurity(ctx, optional)

Get Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecuritySecurityGetSecurityOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecuritySecurityGetSecurityOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSecurity**](microsoft.graph.security.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySecurityUpdateSecurity

> SecuritySecurityUpdateSecurity(ctx, microsoftGraphSecurity)

Update Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSecurity** | [**MicrosoftGraphSecurity**](MicrosoftGraphSecurity.md)| New property values | 

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

