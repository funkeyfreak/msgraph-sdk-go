# \DirectoryRolesScopedRoleMembershipApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRolesCreateScopedMembers**](DirectoryRolesScopedRoleMembershipApi.md#DirectoryRolesCreateScopedMembers) | **Post** /directoryRoles/{directoryRole-id}/scopedMembers | Create new navigation property to scopedMembers for directoryRoles
[**DirectoryRolesGetScopedMembers**](DirectoryRolesScopedRoleMembershipApi.md#DirectoryRolesGetScopedMembers) | **Get** /directoryRoles/{directoryRole-id}/scopedMembers/{scopedRoleMembership-id} | Get scopedMembers from directoryRoles
[**DirectoryRolesListScopedMembers**](DirectoryRolesScopedRoleMembershipApi.md#DirectoryRolesListScopedMembers) | **Get** /directoryRoles/{directoryRole-id}/scopedMembers | Get scopedMembers from directoryRoles
[**DirectoryRolesUpdateScopedMembers**](DirectoryRolesScopedRoleMembershipApi.md#DirectoryRolesUpdateScopedMembers) | **Patch** /directoryRoles/{directoryRole-id}/scopedMembers/{scopedRoleMembership-id} | Update the navigation property scopedMembers in directoryRoles



## DirectoryRolesCreateScopedMembers

> MicrosoftGraphScopedRoleMembership DirectoryRolesCreateScopedMembers(ctx, directoryRoleId, microsoftGraphScopedRoleMembership)

Create new navigation property to scopedMembers for directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string**| key: directoryRole-id of directoryRole | 
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New navigation property | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](microsoft.graph.scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesGetScopedMembers

> MicrosoftGraphScopedRoleMembership DirectoryRolesGetScopedMembers(ctx, directoryRoleId, scopedRoleMembershipId, optional)

Get scopedMembers from directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string**| key: directoryRole-id of directoryRole | 
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
 **optional** | ***DirectoryRolesGetScopedMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRolesGetScopedMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](microsoft.graph.scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesListScopedMembers

> CollectionOfScopedRoleMembership DirectoryRolesListScopedMembers(ctx, directoryRoleId, optional)

Get scopedMembers from directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string**| key: directoryRole-id of directoryRole | 
 **optional** | ***DirectoryRolesListScopedMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryRolesListScopedMembersOpts struct


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

[**CollectionOfScopedRoleMembership**](Collection_of_scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesUpdateScopedMembers

> DirectoryRolesUpdateScopedMembers(ctx, directoryRoleId, scopedRoleMembershipId, microsoftGraphScopedRoleMembership)

Update the navigation property scopedMembers in directoryRoles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string**| key: directoryRole-id of directoryRole | 
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New navigation property values | 

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

