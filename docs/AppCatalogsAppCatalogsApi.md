# \AppCatalogsAppCatalogsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCatalogsAppCatalogsGetAppCatalogs**](AppCatalogsAppCatalogsApi.md#AppCatalogsAppCatalogsGetAppCatalogs) | **Get** /appCatalogs | Get appCatalogs
[**AppCatalogsAppCatalogsUpdateAppCatalogs**](AppCatalogsAppCatalogsApi.md#AppCatalogsAppCatalogsUpdateAppCatalogs) | **Patch** /appCatalogs | Update appCatalogs



## AppCatalogsAppCatalogsGetAppCatalogs

> MicrosoftGraphAppCatalogs AppCatalogsAppCatalogsGetAppCatalogs(ctx, optional)

Get appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppCatalogsAppCatalogsGetAppCatalogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCatalogsAppCatalogsGetAppCatalogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAppCatalogs**](microsoft.graph.appCatalogs.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsAppCatalogsUpdateAppCatalogs

> AppCatalogsAppCatalogsUpdateAppCatalogs(ctx, microsoftGraphAppCatalogs)

Update appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphAppCatalogs** | [**MicrosoftGraphAppCatalogs**](MicrosoftGraphAppCatalogs.md)| New property values | 

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

