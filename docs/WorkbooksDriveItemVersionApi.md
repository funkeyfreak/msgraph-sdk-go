# \WorkbooksDriveItemVersionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksCreateVersions**](WorkbooksDriveItemVersionApi.md#WorkbooksCreateVersions) | **Post** /workbooks/{driveItem-id}/versions | Create new navigation property to versions for workbooks
[**WorkbooksGetVersions**](WorkbooksDriveItemVersionApi.md#WorkbooksGetVersions) | **Get** /workbooks/{driveItem-id}/versions/{driveItemVersion-id} | Get versions from workbooks
[**WorkbooksListVersions**](WorkbooksDriveItemVersionApi.md#WorkbooksListVersions) | **Get** /workbooks/{driveItem-id}/versions | Get versions from workbooks
[**WorkbooksUpdateVersions**](WorkbooksDriveItemVersionApi.md#WorkbooksUpdateVersions) | **Patch** /workbooks/{driveItem-id}/versions/{driveItemVersion-id} | Update the navigation property versions in workbooks



## WorkbooksCreateVersions

> MicrosoftGraphDriveItemVersion WorkbooksCreateVersions(ctx, driveItemId, microsoftGraphDriveItemVersion)

Create new navigation property to versions for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphDriveItemVersion** | [**MicrosoftGraphDriveItemVersion**](MicrosoftGraphDriveItemVersion.md)| New navigation property | 

### Return type

[**MicrosoftGraphDriveItemVersion**](microsoft.graph.driveItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksGetVersions

> MicrosoftGraphDriveItemVersion WorkbooksGetVersions(ctx, driveItemId, driveItemVersionId, optional)

Get versions from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**driveItemVersionId** | **string**| key: driveItemVersion-id of driveItemVersion | 
 **optional** | ***WorkbooksGetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksGetVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDriveItemVersion**](microsoft.graph.driveItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListVersions

> CollectionOfDriveItemVersion WorkbooksListVersions(ctx, driveItemId, optional)

Get versions from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksListVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListVersionsOpts struct


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

[**CollectionOfDriveItemVersion**](Collection_of_driveItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksUpdateVersions

> WorkbooksUpdateVersions(ctx, driveItemId, driveItemVersionId, microsoftGraphDriveItemVersion)

Update the navigation property versions in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**driveItemVersionId** | **string**| key: driveItemVersion-id of driveItemVersion | 
**microsoftGraphDriveItemVersion** | [**MicrosoftGraphDriveItemVersion**](MicrosoftGraphDriveItemVersion.md)| New navigation property values | 

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

