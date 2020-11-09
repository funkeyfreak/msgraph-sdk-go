# \EducationEducationRootApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationEducationRootGetEducationRoot**](EducationEducationRootApi.md#EducationEducationRootGetEducationRoot) | **Get** /education | Get education
[**EducationEducationRootUpdateEducationRoot**](EducationEducationRootApi.md#EducationEducationRootUpdateEducationRoot) | **Patch** /education | Update education



## EducationEducationRootGetEducationRoot

> MicrosoftGraphEducationRoot EducationEducationRootGetEducationRoot(ctx, optional)

Get education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EducationEducationRootGetEducationRootOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationEducationRootGetEducationRootOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEducationRoot**](microsoft.graph.educationRoot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationEducationRootUpdateEducationRoot

> EducationEducationRootUpdateEducationRoot(ctx, microsoftGraphEducationRoot)

Update education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphEducationRoot** | [**MicrosoftGraphEducationRoot**](MicrosoftGraphEducationRoot.md)| New property values | 

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

