# \ApplicationsApplicationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsApplicationCreateApplication**](ApplicationsApplicationApi.md#ApplicationsApplicationCreateApplication) | **Post** /applications | Add new entity to applications
[**ApplicationsApplicationDeleteApplication**](ApplicationsApplicationApi.md#ApplicationsApplicationDeleteApplication) | **Delete** /applications/{application-id} | Delete entity from applications
[**ApplicationsApplicationGetApplication**](ApplicationsApplicationApi.md#ApplicationsApplicationGetApplication) | **Get** /applications/{application-id} | Get entity from applications by key
[**ApplicationsApplicationListApplication**](ApplicationsApplicationApi.md#ApplicationsApplicationListApplication) | **Get** /applications | Get entities from applications
[**ApplicationsApplicationUpdateApplication**](ApplicationsApplicationApi.md#ApplicationsApplicationUpdateApplication) | **Patch** /applications/{application-id} | Update entity in applications



## ApplicationsApplicationCreateApplication

> MicrosoftGraphApplication ApplicationsApplicationCreateApplication(ctx, microsoftGraphApplication)

Add new entity to applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphApplication** | [**MicrosoftGraphApplication**](MicrosoftGraphApplication.md)| New entity | 

### Return type

[**MicrosoftGraphApplication**](microsoft.graph.application.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsApplicationDeleteApplication

> ApplicationsApplicationDeleteApplication(ctx, applicationId, optional)

Delete entity from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
 **optional** | ***ApplicationsApplicationDeleteApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsApplicationDeleteApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| ETag | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsApplicationGetApplication

> MicrosoftGraphApplication ApplicationsApplicationGetApplication(ctx, applicationId, optional)

Get entity from applications by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
 **optional** | ***ApplicationsApplicationGetApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsApplicationGetApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphApplication**](microsoft.graph.application.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsApplicationListApplication

> CollectionOfApplication ApplicationsApplicationListApplication(ctx, optional)

Get entities from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsApplicationListApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsApplicationListApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfApplication**](Collection_of_application.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsApplicationUpdateApplication

> ApplicationsApplicationUpdateApplication(ctx, applicationId, microsoftGraphApplication)

Update entity in applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**microsoftGraphApplication** | [**MicrosoftGraphApplication**](MicrosoftGraphApplication.md)| New property values | 

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

