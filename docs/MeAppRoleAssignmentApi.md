# \MeAppRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateAppRoleAssignments**](MeAppRoleAssignmentApi.md#MeCreateAppRoleAssignments) | **Post** /me/appRoleAssignments | Create new navigation property to appRoleAssignments for me
[**MeGetAppRoleAssignments**](MeAppRoleAssignmentApi.md#MeGetAppRoleAssignments) | **Get** /me/appRoleAssignments/{appRoleAssignment-id} | Get appRoleAssignments from me
[**MeListAppRoleAssignments**](MeAppRoleAssignmentApi.md#MeListAppRoleAssignments) | **Get** /me/appRoleAssignments | Get appRoleAssignments from me
[**MeUpdateAppRoleAssignments**](MeAppRoleAssignmentApi.md#MeUpdateAppRoleAssignments) | **Patch** /me/appRoleAssignments/{appRoleAssignment-id} | Update the navigation property appRoleAssignments in me



## MeCreateAppRoleAssignments

> MicrosoftGraphAppRoleAssignment MeCreateAppRoleAssignments(ctx, microsoftGraphAppRoleAssignment)

Create new navigation property to appRoleAssignments for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeGetAppRoleAssignments

> MicrosoftGraphAppRoleAssignment MeGetAppRoleAssignments(ctx, appRoleAssignmentId, optional)

Get appRoleAssignments from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appRoleAssignmentId** | **string**| key: appRoleAssignment-id of appRoleAssignment | 
 **optional** | ***MeGetAppRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetAppRoleAssignmentsOpts struct


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


## MeListAppRoleAssignments

> CollectionOfAppRoleAssignment MeListAppRoleAssignments(ctx, optional)

Get appRoleAssignments from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListAppRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListAppRoleAssignmentsOpts struct


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


## MeUpdateAppRoleAssignments

> MeUpdateAppRoleAssignments(ctx, appRoleAssignmentId, microsoftGraphAppRoleAssignment)

Update the navigation property appRoleAssignments in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

