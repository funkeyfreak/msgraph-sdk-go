# \SitesDriveApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesCreateDrives**](SitesDriveApi.md#SitesCreateDrives) | **Post** /sites/{site-id}/drives | Create new navigation property to drives for sites
[**SitesGetDrive**](SitesDriveApi.md#SitesGetDrive) | **Get** /sites/{site-id}/drive | Get drive from sites
[**SitesGetDrives**](SitesDriveApi.md#SitesGetDrives) | **Get** /sites/{site-id}/drives/{drive-id} | Get drives from sites
[**SitesListDrives**](SitesDriveApi.md#SitesListDrives) | **Get** /sites/{site-id}/drives | Get drives from sites
[**SitesUpdateDrive**](SitesDriveApi.md#SitesUpdateDrive) | **Patch** /sites/{site-id}/drive | Update the navigation property drive in sites
[**SitesUpdateDrives**](SitesDriveApi.md#SitesUpdateDrives) | **Patch** /sites/{site-id}/drives/{drive-id} | Update the navigation property drives in sites



## SitesCreateDrives

> MicrosoftGraphDrive SitesCreateDrives(ctx, siteId, microsoftGraphDrive)

Create new navigation property to drives for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)| New navigation property | 

### Return type

[**MicrosoftGraphDrive**](microsoft.graph.drive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetDrive

> MicrosoftGraphDrive SitesGetDrive(ctx, siteId, optional)

Get drive from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
 **optional** | ***SitesGetDriveOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesGetDriveOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDrive**](microsoft.graph.drive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetDrives

> MicrosoftGraphDrive SitesGetDrives(ctx, siteId, driveId, optional)

Get drives from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**driveId** | **string**| key: drive-id of drive | 
 **optional** | ***SitesGetDrivesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesGetDrivesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDrive**](microsoft.graph.drive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListDrives

> CollectionOfDrive SitesListDrives(ctx, siteId, optional)

Get drives from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
 **optional** | ***SitesListDrivesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListDrivesOpts struct


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

[**CollectionOfDrive**](Collection_of_drive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesUpdateDrive

> SitesUpdateDrive(ctx, siteId, microsoftGraphDrive)

Update the navigation property drive in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)| New navigation property values | 

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


## SitesUpdateDrives

> SitesUpdateDrives(ctx, siteId, driveId, microsoftGraphDrive)

Update the navigation property drives in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**driveId** | **string**| key: drive-id of drive | 
**microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)| New navigation property values | 

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

