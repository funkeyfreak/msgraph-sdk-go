# \GroupsSiteApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateSites**](GroupsSiteApi.md#GroupsCreateSites) | **Post** /groups/{group-id}/sites | Create new navigation property to sites for groups
[**GroupsGetSites**](GroupsSiteApi.md#GroupsGetSites) | **Get** /groups/{group-id}/sites/{site-id} | Get sites from groups
[**GroupsListSites**](GroupsSiteApi.md#GroupsListSites) | **Get** /groups/{group-id}/sites | Get sites from groups
[**GroupsUpdateSites**](GroupsSiteApi.md#GroupsUpdateSites) | **Patch** /groups/{group-id}/sites/{site-id} | Update the navigation property sites in groups



## GroupsCreateSites

> MicrosoftGraphSite GroupsCreateSites(ctx, groupId, microsoftGraphSite)

Create new navigation property to sites for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphSite** | [**MicrosoftGraphSite**](MicrosoftGraphSite.md)| New navigation property | 

### Return type

[**MicrosoftGraphSite**](microsoft.graph.site.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetSites

> MicrosoftGraphSite GroupsGetSites(ctx, groupId, siteId, optional)

Get sites from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**siteId** | **string**| key: site-id of site | 
 **optional** | ***GroupsGetSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetSitesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSite**](microsoft.graph.site.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsListSites

> CollectionOfSite GroupsListSites(ctx, groupId, optional)

Get sites from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsListSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsListSitesOpts struct


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

[**CollectionOfSite**](Collection_of_site.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdateSites

> GroupsUpdateSites(ctx, groupId, siteId, microsoftGraphSite)

Update the navigation property sites in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**siteId** | **string**| key: site-id of site | 
**microsoftGraphSite** | [**MicrosoftGraphSite**](MicrosoftGraphSite.md)| New navigation property values | 

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

