# \SitesColumnDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesCreateColumns**](SitesColumnDefinitionApi.md#SitesCreateColumns) | **Post** /sites/{site-id}/columns | Create new navigation property to columns for sites
[**SitesGetColumns**](SitesColumnDefinitionApi.md#SitesGetColumns) | **Get** /sites/{site-id}/columns/{columnDefinition-id} | Get columns from sites
[**SitesListColumns**](SitesColumnDefinitionApi.md#SitesListColumns) | **Get** /sites/{site-id}/columns | Get columns from sites
[**SitesUpdateColumns**](SitesColumnDefinitionApi.md#SitesUpdateColumns) | **Patch** /sites/{site-id}/columns/{columnDefinition-id} | Update the navigation property columns in sites



## SitesCreateColumns

> MicrosoftGraphColumnDefinition SitesCreateColumns(ctx, siteId, microsoftGraphColumnDefinition)

Create new navigation property to columns for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)| New navigation property | 

### Return type

[**MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetColumns

> MicrosoftGraphColumnDefinition SitesGetColumns(ctx, siteId, columnDefinitionId, optional)

Get columns from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**columnDefinitionId** | **string**| key: columnDefinition-id of columnDefinition | 
 **optional** | ***SitesGetColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesGetColumnsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListColumns

> CollectionOfColumnDefinition SitesListColumns(ctx, siteId, optional)

Get columns from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
 **optional** | ***SitesListColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListColumnsOpts struct


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

[**CollectionOfColumnDefinition**](Collection_of_columnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesUpdateColumns

> SitesUpdateColumns(ctx, siteId, columnDefinitionId, microsoftGraphColumnDefinition)

Update the navigation property columns in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**columnDefinitionId** | **string**| key: columnDefinition-id of columnDefinition | 
**microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)| New navigation property values | 

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

