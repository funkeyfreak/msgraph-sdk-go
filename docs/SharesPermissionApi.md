# \SharesPermissionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesGetPermission**](SharesPermissionApi.md#SharesGetPermission) | **Get** /shares/{sharedDriveItem-id}/permission | Get permission from shares
[**SharesUpdatePermission**](SharesPermissionApi.md#SharesUpdatePermission) | **Patch** /shares/{sharedDriveItem-id}/permission | Update the navigation property permission in shares



## SharesGetPermission

> MicrosoftGraphPermission SharesGetPermission(ctx, sharedDriveItemId, optional)

Get permission from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesGetPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesGetPermissionOpts struct


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


## SharesUpdatePermission

> SharesUpdatePermission(ctx, sharedDriveItemId, microsoftGraphPermission)

Update the navigation property permission in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
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

