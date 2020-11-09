# \DeviceManagementDeviceManagementTroubleshootingEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateTroubleshootingEvents**](DeviceManagementDeviceManagementTroubleshootingEventApi.md#DeviceManagementCreateTroubleshootingEvents) | **Post** /deviceManagement/troubleshootingEvents | Create new navigation property to troubleshootingEvents for deviceManagement
[**DeviceManagementGetTroubleshootingEvents**](DeviceManagementDeviceManagementTroubleshootingEventApi.md#DeviceManagementGetTroubleshootingEvents) | **Get** /deviceManagement/troubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Get troubleshootingEvents from deviceManagement
[**DeviceManagementListTroubleshootingEvents**](DeviceManagementDeviceManagementTroubleshootingEventApi.md#DeviceManagementListTroubleshootingEvents) | **Get** /deviceManagement/troubleshootingEvents | Get troubleshootingEvents from deviceManagement
[**DeviceManagementUpdateTroubleshootingEvents**](DeviceManagementDeviceManagementTroubleshootingEventApi.md#DeviceManagementUpdateTroubleshootingEvents) | **Patch** /deviceManagement/troubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Update the navigation property troubleshootingEvents in deviceManagement



## DeviceManagementCreateTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent DeviceManagementCreateTroubleshootingEvents(ctx, microsoftGraphDeviceManagementTroubleshootingEvent)

Create new navigation property to troubleshootingEvents for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceManagementTroubleshootingEvent** | [**MicrosoftGraphDeviceManagementTroubleshootingEvent**](MicrosoftGraphDeviceManagementTroubleshootingEvent.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceManagementTroubleshootingEvent**](microsoft.graph.deviceManagementTroubleshootingEvent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent DeviceManagementGetTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId, optional)

Get troubleshootingEvents from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementTroubleshootingEventId** | **string**| key: deviceManagementTroubleshootingEvent-id of deviceManagementTroubleshootingEvent | 
 **optional** | ***DeviceManagementGetTroubleshootingEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetTroubleshootingEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagementTroubleshootingEvent**](microsoft.graph.deviceManagementTroubleshootingEvent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListTroubleshootingEvents

> CollectionOfDeviceManagementTroubleshootingEvent DeviceManagementListTroubleshootingEvents(ctx, optional)

Get troubleshootingEvents from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListTroubleshootingEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListTroubleshootingEventsOpts struct


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

[**CollectionOfDeviceManagementTroubleshootingEvent**](Collection_of_deviceManagementTroubleshootingEvent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateTroubleshootingEvents

> DeviceManagementUpdateTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId, microsoftGraphDeviceManagementTroubleshootingEvent)

Update the navigation property troubleshootingEvents in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementTroubleshootingEventId** | **string**| key: deviceManagementTroubleshootingEvent-id of deviceManagementTroubleshootingEvent | 
**microsoftGraphDeviceManagementTroubleshootingEvent** | [**MicrosoftGraphDeviceManagementTroubleshootingEvent**](MicrosoftGraphDeviceManagementTroubleshootingEvent.md)| New navigation property values | 

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

