# \DirectoryRoleTemplatesDirectoryRoleTemplateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate) | **Post** /directoryRoleTemplates | Add new entity to directoryRoleTemplates
[**DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate) | **Delete** /directoryRoleTemplates/{directoryRoleTemplate-id} | Delete entity from directoryRoleTemplates
[**DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate) | **Get** /directoryRoleTemplates/{directoryRoleTemplate-id} | Get entity from directoryRoleTemplates by key
[**DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate) | **Get** /directoryRoleTemplates | Get entities from directoryRoleTemplates
[**DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate) | **Patch** /directoryRoleTemplates/{directoryRoleTemplate-id} | Update entity in directoryRoleTemplates



## DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate

> MicrosoftGraphDirectoryRoleTemplate DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate(ctx, microsoftGraphDirectoryRoleTemplate)

Add new entity to directoryRoleTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDirectoryRoleTemplate** | [**MicrosoftGraphDirectoryRoleTemplate**](MicrosoftGraphDirectoryRoleTemplate.md)| New entity | 

### Return type

[**MicrosoftGraphDirectoryRoleTemplate**](microsoft.graph.directoryRoleTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate

> DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate(ctx, directoryRoleTemplateId, optional)

Delete entity from directoryRoleTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string**| key: directoryRoleTemplate-id of directoryRoleTemplate | 
 **optional** | ***DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplateOpts struct


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


## DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate

> MicrosoftGraphDirectoryRoleTemplate DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate(ctx, directoryRoleTemplateId, optional)

Get entity from directoryRoleTemplates by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string**| key: directoryRoleTemplate-id of directoryRoleTemplate | 
 **optional** | ***DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphDirectoryRoleTemplate**](microsoft.graph.directoryRoleTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate

> CollectionOfDirectoryRoleTemplate DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate(ctx, optional)

Get entities from directoryRoleTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfDirectoryRoleTemplate**](Collection_of_directoryRoleTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate

> DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate(ctx, directoryRoleTemplateId, microsoftGraphDirectoryRoleTemplate)

Update entity in directoryRoleTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string**| key: directoryRoleTemplate-id of directoryRoleTemplate | 
**microsoftGraphDirectoryRoleTemplate** | [**MicrosoftGraphDirectoryRoleTemplate**](MicrosoftGraphDirectoryRoleTemplate.md)| New property values | 

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

