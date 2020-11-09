# \DirectoryRolesDirectoryRoleApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRolesDirectoryRoleCreateDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleCreateDirectoryRole) | **Post** /directoryRoles | Add new entity to directoryRoles
[**DirectoryRolesDirectoryRoleDeleteDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleDeleteDirectoryRole) | **Delete** /directoryRoles/{directoryRole-id} | Delete entity from directoryRoles
[**DirectoryRolesDirectoryRoleGetDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleGetDirectoryRole) | **Get** /directoryRoles/{directoryRole-id} | Get entity from directoryRoles by key
[**DirectoryRolesDirectoryRoleListDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleListDirectoryRole) | **Get** /directoryRoles | Get entities from directoryRoles
[**DirectoryRolesDirectoryRoleUpdateDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleUpdateDirectoryRole) | **Patch** /directoryRoles/{directoryRole-id} | Update entity in directoryRoles



## DirectoryRolesDirectoryRoleCreateDirectoryRole

> MicrosoftGraphDirectoryRole DirectoryRolesDirectoryRoleCreateDirectoryRole(ctx, microsoftGraphDirectoryRole)

Add new entity to directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDirectoryRole** | [**MicrosoftGraphDirectoryRole**](MicrosoftGraphDirectoryRole.md)| New entity | 

### Return type

[**MicrosoftGraphDirectoryRole**](microsoft.graph.directoryRole.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesDirectoryRoleDeleteDirectoryRole

> DirectoryRolesDirectoryRoleDeleteDirectoryRole(ctx, directoryRoleId, optional)

Delete entity from directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string**| key: directoryRole-id of directoryRole | 
 **optional** | ***DirectoryRolesDirectoryRoleDeleteDirectoryRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRolesDirectoryRoleDeleteDirectoryRoleOpts struct


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


## DirectoryRolesDirectoryRoleGetDirectoryRole

> MicrosoftGraphDirectoryRole DirectoryRolesDirectoryRoleGetDirectoryRole(ctx, directoryRoleId, optional)

Get entity from directoryRoles by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string**| key: directoryRole-id of directoryRole | 
 **optional** | ***DirectoryRolesDirectoryRoleGetDirectoryRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRolesDirectoryRoleGetDirectoryRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphDirectoryRole**](microsoft.graph.directoryRole.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesDirectoryRoleListDirectoryRole

> CollectionOfDirectoryRole DirectoryRolesDirectoryRoleListDirectoryRole(ctx, optional)

Get entities from directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryRolesDirectoryRoleListDirectoryRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRolesDirectoryRoleListDirectoryRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfDirectoryRole**](Collection_of_directoryRole.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesDirectoryRoleUpdateDirectoryRole

> DirectoryRolesDirectoryRoleUpdateDirectoryRole(ctx, directoryRoleId, microsoftGraphDirectoryRole)

Update entity in directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string**| key: directoryRole-id of directoryRole | 
**microsoftGraphDirectoryRole** | [**MicrosoftGraphDirectoryRole**](MicrosoftGraphDirectoryRole.md)| New property values | 

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

