# \DrivesDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesCreateFollowing**](DrivesDriveItemApi.md#DrivesCreateFollowing) | **Post** /drives/{drive-id}/following | Create new navigation property to following for drives
[**DrivesCreateItems**](DrivesDriveItemApi.md#DrivesCreateItems) | **Post** /drives/{drive-id}/items | Create new navigation property to items for drives
[**DrivesCreateSpecial**](DrivesDriveItemApi.md#DrivesCreateSpecial) | **Post** /drives/{drive-id}/special | Create new navigation property to special for drives
[**DrivesGetFollowing**](DrivesDriveItemApi.md#DrivesGetFollowing) | **Get** /drives/{drive-id}/following/{driveItem-id} | Get following from drives
[**DrivesGetItems**](DrivesDriveItemApi.md#DrivesGetItems) | **Get** /drives/{drive-id}/items/{driveItem-id} | Get items from drives
[**DrivesGetRoot**](DrivesDriveItemApi.md#DrivesGetRoot) | **Get** /drives/{drive-id}/root | Get root from drives
[**DrivesGetSpecial**](DrivesDriveItemApi.md#DrivesGetSpecial) | **Get** /drives/{drive-id}/special/{driveItem-id} | Get special from drives
[**DrivesListFollowing**](DrivesDriveItemApi.md#DrivesListFollowing) | **Get** /drives/{drive-id}/following | Get following from drives
[**DrivesListItems**](DrivesDriveItemApi.md#DrivesListItems) | **Get** /drives/{drive-id}/items | Get items from drives
[**DrivesListSpecial**](DrivesDriveItemApi.md#DrivesListSpecial) | **Get** /drives/{drive-id}/special | Get special from drives
[**DrivesUpdateFollowing**](DrivesDriveItemApi.md#DrivesUpdateFollowing) | **Patch** /drives/{drive-id}/following/{driveItem-id} | Update the navigation property following in drives
[**DrivesUpdateItems**](DrivesDriveItemApi.md#DrivesUpdateItems) | **Patch** /drives/{drive-id}/items/{driveItem-id} | Update the navigation property items in drives
[**DrivesUpdateRoot**](DrivesDriveItemApi.md#DrivesUpdateRoot) | **Patch** /drives/{drive-id}/root | Update the navigation property root in drives
[**DrivesUpdateSpecial**](DrivesDriveItemApi.md#DrivesUpdateSpecial) | **Patch** /drives/{drive-id}/special/{driveItem-id} | Update the navigation property special in drives



## DrivesCreateFollowing

> MicrosoftGraphDriveItem DrivesCreateFollowing(ctx, driveId, microsoftGraphDriveItem)

Create new navigation property to following for drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesCreateItems

> MicrosoftGraphDriveItem DrivesCreateItems(ctx, driveId, microsoftGraphDriveItem)

Create new navigation property to items for drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesCreateSpecial

> MicrosoftGraphDriveItem DrivesCreateSpecial(ctx, driveId, microsoftGraphDriveItem)

Create new navigation property to special for drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesGetFollowing

> MicrosoftGraphDriveItem DrivesGetFollowing(ctx, driveId, driveItemId, optional)

Get following from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***DrivesGetFollowingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesGetFollowingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesGetItems

> MicrosoftGraphDriveItem DrivesGetItems(ctx, driveId, driveItemId, optional)

Get items from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***DrivesGetItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesGetItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesGetRoot

> MicrosoftGraphDriveItem DrivesGetRoot(ctx, driveId, optional)

Get root from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
 **optional** | ***DrivesGetRootOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesGetRootOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesGetSpecial

> MicrosoftGraphDriveItem DrivesGetSpecial(ctx, driveId, driveItemId, optional)

Get special from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***DrivesGetSpecialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesGetSpecialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesListFollowing

> CollectionOfDriveItem DrivesListFollowing(ctx, driveId, optional)

Get following from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
 **optional** | ***DrivesListFollowingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListFollowingOpts struct


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

[**CollectionOfDriveItem**](Collection_of_driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesListItems

> CollectionOfDriveItem DrivesListItems(ctx, driveId, optional)

Get items from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
 **optional** | ***DrivesListItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListItemsOpts struct


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

[**CollectionOfDriveItem**](Collection_of_driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesListSpecial

> CollectionOfDriveItem DrivesListSpecial(ctx, driveId, optional)

Get special from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
 **optional** | ***DrivesListSpecialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListSpecialOpts struct


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

[**CollectionOfDriveItem**](Collection_of_driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesUpdateFollowing

> DrivesUpdateFollowing(ctx, driveId, driveItemId, microsoftGraphDriveItem)

Update the navigation property following in drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property values | 

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


## DrivesUpdateItems

> DrivesUpdateItems(ctx, driveId, driveItemId, microsoftGraphDriveItem)

Update the navigation property items in drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property values | 

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


## DrivesUpdateRoot

> DrivesUpdateRoot(ctx, driveId, microsoftGraphDriveItem)

Update the navigation property root in drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property values | 

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


## DrivesUpdateSpecial

> DrivesUpdateSpecial(ctx, driveId, driveItemId, microsoftGraphDriveItem)

Update the navigation property special in drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property values | 

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

