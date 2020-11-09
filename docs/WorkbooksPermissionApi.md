# \WorkbooksPermissionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksCreatePermissions**](WorkbooksPermissionApi.md#WorkbooksCreatePermissions) | **Post** /workbooks/{driveItem-id}/permissions | Create new navigation property to permissions for workbooks
[**WorkbooksGetPermissions**](WorkbooksPermissionApi.md#WorkbooksGetPermissions) | **Get** /workbooks/{driveItem-id}/permissions/{permission-id} | Get permissions from workbooks
[**WorkbooksListPermissions**](WorkbooksPermissionApi.md#WorkbooksListPermissions) | **Get** /workbooks/{driveItem-id}/permissions | Get permissions from workbooks
[**WorkbooksUpdatePermissions**](WorkbooksPermissionApi.md#WorkbooksUpdatePermissions) | **Patch** /workbooks/{driveItem-id}/permissions/{permission-id} | Update the navigation property permissions in workbooks



## WorkbooksCreatePermissions

> MicrosoftGraphPermission WorkbooksCreatePermissions(ctx, driveItemId, microsoftGraphPermission)

Create new navigation property to permissions for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphPermission** | [**MicrosoftGraphPermission**](MicrosoftGraphPermission.md)| New navigation property | 

### Return type

[**MicrosoftGraphPermission**](microsoft.graph.permission.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksGetPermissions

> MicrosoftGraphPermission WorkbooksGetPermissions(ctx, driveItemId, permissionId, optional)

Get permissions from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**permissionId** | **string**| key: permission-id of permission | 
 **optional** | ***WorkbooksGetPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksGetPermissionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPermission**](microsoft.graph.permission.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListPermissions

> CollectionOfPermission WorkbooksListPermissions(ctx, driveItemId, optional)

Get permissions from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksListPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListPermissionsOpts struct


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

[**CollectionOfPermission**](Collection_of_permission.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksUpdatePermissions

> WorkbooksUpdatePermissions(ctx, driveItemId, permissionId, microsoftGraphPermission)

Update the navigation property permissions in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**permissionId** | **string**| key: permission-id of permission | 
**microsoftGraphPermission** | [**MicrosoftGraphPermission**](MicrosoftGraphPermission.md)| New navigation property values | 

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

