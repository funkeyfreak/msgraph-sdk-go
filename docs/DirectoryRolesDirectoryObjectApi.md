# \DirectoryRolesDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRolesGetMembers**](DirectoryRolesDirectoryObjectApi.md#DirectoryRolesGetMembers) | **Get** /directoryRoles/{directoryRole-id}/members/{directoryObject-id} | Get members from directoryRoles
[**DirectoryRolesListMembers**](DirectoryRolesDirectoryObjectApi.md#DirectoryRolesListMembers) | **Get** /directoryRoles/{directoryRole-id}/members | Get members from directoryRoles



## DirectoryRolesGetMembers

> MicrosoftGraphDirectoryObject DirectoryRolesGetMembers(ctx, directoryRoleId, directoryObjectId, optional)

Get members from directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string**| key: directoryRole-id of directoryRole | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***DirectoryRolesGetMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRolesGetMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesListMembers

> CollectionOfDirectoryObject DirectoryRolesListMembers(ctx, directoryRoleId, optional)

Get members from directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string**| key: directoryRole-id of directoryRole | 
 **optional** | ***DirectoryRolesListMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRolesListMembersOpts struct


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

[**CollectionOfDirectoryObject**](Collection_of_directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

