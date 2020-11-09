# \DrivesDriveApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesDriveCreateDrive**](DrivesDriveApi.md#DrivesDriveCreateDrive) | **Post** /drives | Add new entity to drives
[**DrivesDriveDeleteDrive**](DrivesDriveApi.md#DrivesDriveDeleteDrive) | **Delete** /drives/{drive-id} | Delete entity from drives
[**DrivesDriveGetDrive**](DrivesDriveApi.md#DrivesDriveGetDrive) | **Get** /drives/{drive-id} | Get entity from drives by key
[**DrivesDriveListDrive**](DrivesDriveApi.md#DrivesDriveListDrive) | **Get** /drives | Get entities from drives
[**DrivesDriveUpdateDrive**](DrivesDriveApi.md#DrivesDriveUpdateDrive) | **Patch** /drives/{drive-id} | Update entity in drives



## DrivesDriveCreateDrive

> MicrosoftGraphDrive DrivesDriveCreateDrive(ctx, microsoftGraphDrive)

Add new entity to drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)| New entity | 

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


## DrivesDriveDeleteDrive

> DrivesDriveDeleteDrive(ctx, driveId, optional)

Delete entity from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
 **optional** | ***DrivesDriveDeleteDriveOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesDriveDeleteDriveOpts struct


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


## DrivesDriveGetDrive

> MicrosoftGraphDrive DrivesDriveGetDrive(ctx, driveId, optional)

Get entity from drives by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
 **optional** | ***DrivesDriveGetDriveOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesDriveGetDriveOpts struct


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


## DrivesDriveListDrive

> CollectionOfDrive DrivesDriveListDrive(ctx, optional)

Get entities from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DrivesDriveListDriveOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesDriveListDriveOpts struct


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


## DrivesDriveUpdateDrive

> DrivesDriveUpdateDrive(ctx, driveId, microsoftGraphDrive)

Update entity in drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)| New property values | 

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

