# \TeamsTemplatesTeamsTemplateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsTemplatesTeamsTemplateCreateTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateCreateTeamsTemplate) | **Post** /teamsTemplates | Add new entity to teamsTemplates
[**TeamsTemplatesTeamsTemplateDeleteTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateDeleteTeamsTemplate) | **Delete** /teamsTemplates/{teamsTemplate-id} | Delete entity from teamsTemplates
[**TeamsTemplatesTeamsTemplateGetTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateGetTeamsTemplate) | **Get** /teamsTemplates/{teamsTemplate-id} | Get entity from teamsTemplates by key
[**TeamsTemplatesTeamsTemplateListTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateListTeamsTemplate) | **Get** /teamsTemplates | Get entities from teamsTemplates
[**TeamsTemplatesTeamsTemplateUpdateTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateUpdateTeamsTemplate) | **Patch** /teamsTemplates/{teamsTemplate-id} | Update entity in teamsTemplates



## TeamsTemplatesTeamsTemplateCreateTeamsTemplate

> MicrosoftGraphTeamsTemplate TeamsTemplatesTeamsTemplateCreateTeamsTemplate(ctx, microsoftGraphTeamsTemplate)

Add new entity to teamsTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTeamsTemplate** | [**MicrosoftGraphTeamsTemplate**](MicrosoftGraphTeamsTemplate.md)| New entity | 

### Return type

[**MicrosoftGraphTeamsTemplate**](microsoft.graph.teamsTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTemplatesTeamsTemplateDeleteTeamsTemplate

> TeamsTemplatesTeamsTemplateDeleteTeamsTemplate(ctx, teamsTemplateId, optional)

Delete entity from teamsTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsTemplateId** | **string**| key: teamsTemplate-id of teamsTemplate | 
 **optional** | ***TeamsTemplatesTeamsTemplateDeleteTeamsTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsTemplatesTeamsTemplateDeleteTeamsTemplateOpts struct


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


## TeamsTemplatesTeamsTemplateGetTeamsTemplate

> MicrosoftGraphTeamsTemplate TeamsTemplatesTeamsTemplateGetTeamsTemplate(ctx, teamsTemplateId, optional)

Get entity from teamsTemplates by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsTemplateId** | **string**| key: teamsTemplate-id of teamsTemplate | 
 **optional** | ***TeamsTemplatesTeamsTemplateGetTeamsTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsTemplatesTeamsTemplateGetTeamsTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTeamsTemplate**](microsoft.graph.teamsTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTemplatesTeamsTemplateListTeamsTemplate

> CollectionOfTeamsTemplate TeamsTemplatesTeamsTemplateListTeamsTemplate(ctx, optional)

Get entities from teamsTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsTemplatesTeamsTemplateListTeamsTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsTemplatesTeamsTemplateListTeamsTemplateOpts struct


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

[**CollectionOfTeamsTemplate**](Collection_of_teamsTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTemplatesTeamsTemplateUpdateTeamsTemplate

> TeamsTemplatesTeamsTemplateUpdateTeamsTemplate(ctx, teamsTemplateId, microsoftGraphTeamsTemplate)

Update entity in teamsTemplates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsTemplateId** | **string**| key: teamsTemplate-id of teamsTemplate | 
**microsoftGraphTeamsTemplate** | [**MicrosoftGraphTeamsTemplate**](MicrosoftGraphTeamsTemplate.md)| New property values | 

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

