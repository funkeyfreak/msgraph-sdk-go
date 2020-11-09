# \GroupsAppRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateAppRoleAssignments**](GroupsAppRoleAssignmentApi.md#GroupsCreateAppRoleAssignments) | **Post** /groups/{group-id}/appRoleAssignments | Create new navigation property to appRoleAssignments for groups
[**GroupsGetAppRoleAssignments**](GroupsAppRoleAssignmentApi.md#GroupsGetAppRoleAssignments) | **Get** /groups/{group-id}/appRoleAssignments/{appRoleAssignment-id} | Get appRoleAssignments from groups
[**GroupsListAppRoleAssignments**](GroupsAppRoleAssignmentApi.md#GroupsListAppRoleAssignments) | **Get** /groups/{group-id}/appRoleAssignments | Get appRoleAssignments from groups
[**GroupsUpdateAppRoleAssignments**](GroupsAppRoleAssignmentApi.md#GroupsUpdateAppRoleAssignments) | **Patch** /groups/{group-id}/appRoleAssignments/{appRoleAssignment-id} | Update the navigation property appRoleAssignments in groups



## GroupsCreateAppRoleAssignments

> MicrosoftGraphAppRoleAssignment GroupsCreateAppRoleAssignments(ctx, groupId, microsoftGraphAppRoleAssignment)

Create new navigation property to appRoleAssignments for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](microsoft.graph.appRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetAppRoleAssignments

> MicrosoftGraphAppRoleAssignment GroupsGetAppRoleAssignments(ctx, groupId, appRoleAssignmentId, optional)

Get appRoleAssignments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**appRoleAssignmentId** | **string**| key: appRoleAssignment-id of appRoleAssignment | 
 **optional** | ***GroupsGetAppRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetAppRoleAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](microsoft.graph.appRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsListAppRoleAssignments

> CollectionOfAppRoleAssignment GroupsListAppRoleAssignments(ctx, groupId, optional)

Get appRoleAssignments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsListAppRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsListAppRoleAssignmentsOpts struct


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

[**CollectionOfAppRoleAssignment**](Collection_of_appRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdateAppRoleAssignments

> GroupsUpdateAppRoleAssignments(ctx, groupId, appRoleAssignmentId, microsoftGraphAppRoleAssignment)

Update the navigation property appRoleAssignments in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**appRoleAssignmentId** | **string**| key: appRoleAssignment-id of appRoleAssignment | 
**microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)| New navigation property values | 

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

