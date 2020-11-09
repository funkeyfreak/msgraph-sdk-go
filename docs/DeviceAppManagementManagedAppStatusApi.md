# \DeviceAppManagementManagedAppStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateManagedAppStatuses**](DeviceAppManagementManagedAppStatusApi.md#DeviceAppManagementCreateManagedAppStatuses) | **Post** /deviceAppManagement/managedAppStatuses | Create new navigation property to managedAppStatuses for deviceAppManagement
[**DeviceAppManagementGetManagedAppStatuses**](DeviceAppManagementManagedAppStatusApi.md#DeviceAppManagementGetManagedAppStatuses) | **Get** /deviceAppManagement/managedAppStatuses/{managedAppStatus-id} | Get managedAppStatuses from deviceAppManagement
[**DeviceAppManagementListManagedAppStatuses**](DeviceAppManagementManagedAppStatusApi.md#DeviceAppManagementListManagedAppStatuses) | **Get** /deviceAppManagement/managedAppStatuses | Get managedAppStatuses from deviceAppManagement
[**DeviceAppManagementUpdateManagedAppStatuses**](DeviceAppManagementManagedAppStatusApi.md#DeviceAppManagementUpdateManagedAppStatuses) | **Patch** /deviceAppManagement/managedAppStatuses/{managedAppStatus-id} | Update the navigation property managedAppStatuses in deviceAppManagement



## DeviceAppManagementCreateManagedAppStatuses

> MicrosoftGraphManagedAppStatus DeviceAppManagementCreateManagedAppStatuses(ctx, microsoftGraphManagedAppStatus)

Create new navigation property to managedAppStatuses for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphManagedAppStatus** | [**MicrosoftGraphManagedAppStatus**](MicrosoftGraphManagedAppStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppStatus**](microsoft.graph.managedAppStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetManagedAppStatuses

> MicrosoftGraphManagedAppStatus DeviceAppManagementGetManagedAppStatuses(ctx, managedAppStatusId, optional)

Get managedAppStatuses from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppStatusId** | **string**| key: managedAppStatus-id of managedAppStatus | 
 **optional** | ***DeviceAppManagementGetManagedAppStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetManagedAppStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppStatus**](microsoft.graph.managedAppStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListManagedAppStatuses

> CollectionOfManagedAppStatus DeviceAppManagementListManagedAppStatuses(ctx, optional)

Get managedAppStatuses from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListManagedAppStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListManagedAppStatusesOpts struct


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

[**CollectionOfManagedAppStatus**](Collection_of_managedAppStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateManagedAppStatuses

> DeviceAppManagementUpdateManagedAppStatuses(ctx, managedAppStatusId, microsoftGraphManagedAppStatus)

Update the navigation property managedAppStatuses in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppStatusId** | **string**| key: managedAppStatus-id of managedAppStatus | 
**microsoftGraphManagedAppStatus** | [**MicrosoftGraphManagedAppStatus**](MicrosoftGraphManagedAppStatus.md)| New navigation property values | 

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

