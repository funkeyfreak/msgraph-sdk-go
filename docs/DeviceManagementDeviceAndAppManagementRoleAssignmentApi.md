# \DeviceManagementDeviceAndAppManagementRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateRoleAssignments**](DeviceManagementDeviceAndAppManagementRoleAssignmentApi.md#DeviceManagementCreateRoleAssignments) | **Post** /deviceManagement/roleAssignments | Create new navigation property to roleAssignments for deviceManagement
[**DeviceManagementGetRoleAssignments**](DeviceManagementDeviceAndAppManagementRoleAssignmentApi.md#DeviceManagementGetRoleAssignments) | **Get** /deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment-id} | Get roleAssignments from deviceManagement
[**DeviceManagementListRoleAssignments**](DeviceManagementDeviceAndAppManagementRoleAssignmentApi.md#DeviceManagementListRoleAssignments) | **Get** /deviceManagement/roleAssignments | Get roleAssignments from deviceManagement
[**DeviceManagementUpdateRoleAssignments**](DeviceManagementDeviceAndAppManagementRoleAssignmentApi.md#DeviceManagementUpdateRoleAssignments) | **Patch** /deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment-id} | Update the navigation property roleAssignments in deviceManagement



## DeviceManagementCreateRoleAssignments

> MicrosoftGraphDeviceAndAppManagementRoleAssignment DeviceManagementCreateRoleAssignments(ctx, microsoftGraphDeviceAndAppManagementRoleAssignment)

Create new navigation property to roleAssignments for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceAndAppManagementRoleAssignment** | [**MicrosoftGraphDeviceAndAppManagementRoleAssignment**](MicrosoftGraphDeviceAndAppManagementRoleAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceAndAppManagementRoleAssignment**](microsoft.graph.deviceAndAppManagementRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetRoleAssignments

> MicrosoftGraphDeviceAndAppManagementRoleAssignment DeviceManagementGetRoleAssignments(ctx, deviceAndAppManagementRoleAssignmentId, optional)

Get roleAssignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAndAppManagementRoleAssignmentId** | **string**| key: deviceAndAppManagementRoleAssignment-id of deviceAndAppManagementRoleAssignment | 
 **optional** | ***DeviceManagementGetRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetRoleAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceAndAppManagementRoleAssignment**](microsoft.graph.deviceAndAppManagementRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListRoleAssignments

> CollectionOfDeviceAndAppManagementRoleAssignment DeviceManagementListRoleAssignments(ctx, optional)

Get roleAssignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListRoleAssignmentsOpts struct


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

[**CollectionOfDeviceAndAppManagementRoleAssignment**](Collection_of_deviceAndAppManagementRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateRoleAssignments

> DeviceManagementUpdateRoleAssignments(ctx, deviceAndAppManagementRoleAssignmentId, microsoftGraphDeviceAndAppManagementRoleAssignment)

Update the navigation property roleAssignments in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAndAppManagementRoleAssignmentId** | **string**| key: deviceAndAppManagementRoleAssignment-id of deviceAndAppManagementRoleAssignment | 
**microsoftGraphDeviceAndAppManagementRoleAssignment** | [**MicrosoftGraphDeviceAndAppManagementRoleAssignment**](MicrosoftGraphDeviceAndAppManagementRoleAssignment.md)| New navigation property values | 

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

