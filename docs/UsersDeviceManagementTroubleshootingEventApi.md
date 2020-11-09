# \UsersDeviceManagementTroubleshootingEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateDeviceManagementTroubleshootingEvents**](UsersDeviceManagementTroubleshootingEventApi.md#UsersCreateDeviceManagementTroubleshootingEvents) | **Post** /users/{user-id}/deviceManagementTroubleshootingEvents | Create new navigation property to deviceManagementTroubleshootingEvents for users
[**UsersGetDeviceManagementTroubleshootingEvents**](UsersDeviceManagementTroubleshootingEventApi.md#UsersGetDeviceManagementTroubleshootingEvents) | **Get** /users/{user-id}/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Get deviceManagementTroubleshootingEvents from users
[**UsersListDeviceManagementTroubleshootingEvents**](UsersDeviceManagementTroubleshootingEventApi.md#UsersListDeviceManagementTroubleshootingEvents) | **Get** /users/{user-id}/deviceManagementTroubleshootingEvents | Get deviceManagementTroubleshootingEvents from users
[**UsersUpdateDeviceManagementTroubleshootingEvents**](UsersDeviceManagementTroubleshootingEventApi.md#UsersUpdateDeviceManagementTroubleshootingEvents) | **Patch** /users/{user-id}/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Update the navigation property deviceManagementTroubleshootingEvents in users



## UsersCreateDeviceManagementTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent UsersCreateDeviceManagementTroubleshootingEvents(ctx, userId, microsoftGraphDeviceManagementTroubleshootingEvent)

Create new navigation property to deviceManagementTroubleshootingEvents for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersGetDeviceManagementTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent UsersGetDeviceManagementTroubleshootingEvents(ctx, userId, deviceManagementTroubleshootingEventId, optional)

Get deviceManagementTroubleshootingEvents from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**deviceManagementTroubleshootingEventId** | **string**| key: deviceManagementTroubleshootingEvent-id of deviceManagementTroubleshootingEvent | 
 **optional** | ***UsersGetDeviceManagementTroubleshootingEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetDeviceManagementTroubleshootingEventsOpts struct


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


## UsersListDeviceManagementTroubleshootingEvents

> CollectionOfDeviceManagementTroubleshootingEvent UsersListDeviceManagementTroubleshootingEvents(ctx, userId, optional)

Get deviceManagementTroubleshootingEvents from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListDeviceManagementTroubleshootingEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListDeviceManagementTroubleshootingEventsOpts struct


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


## UsersUpdateDeviceManagementTroubleshootingEvents

> UsersUpdateDeviceManagementTroubleshootingEvents(ctx, userId, deviceManagementTroubleshootingEventId, microsoftGraphDeviceManagementTroubleshootingEvent)

Update the navigation property deviceManagementTroubleshootingEvents in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

