# \GroupSettingTemplatesGroupSettingTemplateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate) | **Post** /groupSettingTemplates | Add new entity to groupSettingTemplates
[**GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate) | **Delete** /groupSettingTemplates/{groupSettingTemplate-id} | Delete entity from groupSettingTemplates
[**GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate) | **Get** /groupSettingTemplates/{groupSettingTemplate-id} | Get entity from groupSettingTemplates by key
[**GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate) | **Get** /groupSettingTemplates | Get entities from groupSettingTemplates
[**GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate) | **Patch** /groupSettingTemplates/{groupSettingTemplate-id} | Update entity in groupSettingTemplates



## GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate

> MicrosoftGraphGroupSettingTemplate GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate(ctx, microsoftGraphGroupSettingTemplate)

Add new entity to groupSettingTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphGroupSettingTemplate** | [**MicrosoftGraphGroupSettingTemplate**](MicrosoftGraphGroupSettingTemplate.md)| New entity | 

### Return type

[**MicrosoftGraphGroupSettingTemplate**](microsoft.graph.groupSettingTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate

> GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate(ctx, groupSettingTemplateId, optional)

Delete entity from groupSettingTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string**| key: groupSettingTemplate-id of groupSettingTemplate | 
 **optional** | ***GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplateOpts struct


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


## GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate

> MicrosoftGraphGroupSettingTemplate GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate(ctx, groupSettingTemplateId, optional)

Get entity from groupSettingTemplates by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string**| key: groupSettingTemplate-id of groupSettingTemplate | 
 **optional** | ***GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphGroupSettingTemplate**](microsoft.graph.groupSettingTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate

> CollectionOfGroupSettingTemplate GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate(ctx, optional)

Get entities from groupSettingTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfGroupSettingTemplate**](Collection_of_groupSettingTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate

> GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate(ctx, groupSettingTemplateId, microsoftGraphGroupSettingTemplate)

Update entity in groupSettingTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string**| key: groupSettingTemplate-id of groupSettingTemplate | 
**microsoftGraphGroupSettingTemplate** | [**MicrosoftGraphGroupSettingTemplate**](MicrosoftGraphGroupSettingTemplate.md)| New property values | 

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

