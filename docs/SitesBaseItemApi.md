# \SitesBaseItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesCreateItems**](SitesBaseItemApi.md#SitesCreateItems) | **Post** /sites/{site-id}/items | Create new navigation property to items for sites
[**SitesGetItems**](SitesBaseItemApi.md#SitesGetItems) | **Get** /sites/{site-id}/items/{baseItem-id} | Get items from sites
[**SitesListItems**](SitesBaseItemApi.md#SitesListItems) | **Get** /sites/{site-id}/items | Get items from sites
[**SitesUpdateItems**](SitesBaseItemApi.md#SitesUpdateItems) | **Patch** /sites/{site-id}/items/{baseItem-id} | Update the navigation property items in sites



## SitesCreateItems

> MicrosoftGraphBaseItem SitesCreateItems(ctx, siteId, microsoftGraphBaseItem)

Create new navigation property to items for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**microsoftGraphBaseItem** | [**MicrosoftGraphBaseItem**](MicrosoftGraphBaseItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphBaseItem**](microsoft.graph.baseItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetItems

> MicrosoftGraphBaseItem SitesGetItems(ctx, siteId, baseItemId, optional)

Get items from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**baseItemId** | **string**| key: baseItem-id of baseItem | 
 **optional** | ***SitesGetItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesGetItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphBaseItem**](microsoft.graph.baseItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListItems

> CollectionOfBaseItem SitesListItems(ctx, siteId, optional)

Get items from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
 **optional** | ***SitesListItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListItemsOpts struct


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

[**CollectionOfBaseItem**](Collection_of_baseItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesUpdateItems

> SitesUpdateItems(ctx, siteId, baseItemId, microsoftGraphBaseItem)

Update the navigation property items in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**baseItemId** | **string**| key: baseItem-id of baseItem | 
**microsoftGraphBaseItem** | [**MicrosoftGraphBaseItem**](MicrosoftGraphBaseItem.md)| New navigation property values | 

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

