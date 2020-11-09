# \MeDeviceManagementTroubleshootingEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateDeviceManagementTroubleshootingEvents**](MeDeviceManagementTroubleshootingEventApi.md#MeCreateDeviceManagementTroubleshootingEvents) | **Post** /me/deviceManagementTroubleshootingEvents | Create new navigation property to deviceManagementTroubleshootingEvents for me
[**MeGetDeviceManagementTroubleshootingEvents**](MeDeviceManagementTroubleshootingEventApi.md#MeGetDeviceManagementTroubleshootingEvents) | **Get** /me/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Get deviceManagementTroubleshootingEvents from me
[**MeListDeviceManagementTroubleshootingEvents**](MeDeviceManagementTroubleshootingEventApi.md#MeListDeviceManagementTroubleshootingEvents) | **Get** /me/deviceManagementTroubleshootingEvents | Get deviceManagementTroubleshootingEvents from me
[**MeUpdateDeviceManagementTroubleshootingEvents**](MeDeviceManagementTroubleshootingEventApi.md#MeUpdateDeviceManagementTroubleshootingEvents) | **Patch** /me/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Update the navigation property deviceManagementTroubleshootingEvents in me



## MeCreateDeviceManagementTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent MeCreateDeviceManagementTroubleshootingEvents(ctx, microsoftGraphDeviceManagementTroubleshootingEvent)

Create new navigation property to deviceManagementTroubleshootingEvents for me

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


## MeGetDeviceManagementTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent MeGetDeviceManagementTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId, optional)

Get deviceManagementTroubleshootingEvents from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementTroubleshootingEventId** | **string**| key: deviceManagementTroubleshootingEvent-id of deviceManagementTroubleshootingEvent | 
 **optional** | ***MeGetDeviceManagementTroubleshootingEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetDeviceManagementTroubleshootingEventsOpts struct


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


## MeListDeviceManagementTroubleshootingEvents

> CollectionOfDeviceManagementTroubleshootingEvent MeListDeviceManagementTroubleshootingEvents(ctx, optional)

Get deviceManagementTroubleshootingEvents from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListDeviceManagementTroubleshootingEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListDeviceManagementTroubleshootingEventsOpts struct


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


## MeUpdateDeviceManagementTroubleshootingEvents

> MeUpdateDeviceManagementTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId, microsoftGraphDeviceManagementTroubleshootingEvent)

Update the navigation property deviceManagementTroubleshootingEvents in me

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

