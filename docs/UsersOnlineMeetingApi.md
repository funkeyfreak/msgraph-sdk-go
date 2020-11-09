# \UsersOnlineMeetingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateOnlineMeetings**](UsersOnlineMeetingApi.md#UsersCreateOnlineMeetings) | **Post** /users/{user-id}/onlineMeetings | Create new navigation property to onlineMeetings for users
[**UsersGetOnlineMeetings**](UsersOnlineMeetingApi.md#UsersGetOnlineMeetings) | **Get** /users/{user-id}/onlineMeetings/{onlineMeeting-id} | Get onlineMeetings from users
[**UsersListOnlineMeetings**](UsersOnlineMeetingApi.md#UsersListOnlineMeetings) | **Get** /users/{user-id}/onlineMeetings | Get onlineMeetings from users
[**UsersUpdateOnlineMeetings**](UsersOnlineMeetingApi.md#UsersUpdateOnlineMeetings) | **Patch** /users/{user-id}/onlineMeetings/{onlineMeeting-id} | Update the navigation property onlineMeetings in users



## UsersCreateOnlineMeetings

> MicrosoftGraphOnlineMeeting UsersCreateOnlineMeetings(ctx, userId, microsoftGraphOnlineMeeting)

Create new navigation property to onlineMeetings for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphOnlineMeeting** | [**MicrosoftGraphOnlineMeeting**](MicrosoftGraphOnlineMeeting.md)| New navigation property | 

### Return type

[**MicrosoftGraphOnlineMeeting**](microsoft.graph.onlineMeeting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetOnlineMeetings

> MicrosoftGraphOnlineMeeting UsersGetOnlineMeetings(ctx, userId, onlineMeetingId, optional)

Get onlineMeetings from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**onlineMeetingId** | **string**| key: onlineMeeting-id of onlineMeeting | 
 **optional** | ***UsersGetOnlineMeetingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetOnlineMeetingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnlineMeeting**](microsoft.graph.onlineMeeting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListOnlineMeetings

> CollectionOfOnlineMeeting UsersListOnlineMeetings(ctx, userId, optional)

Get onlineMeetings from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListOnlineMeetingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListOnlineMeetingsOpts struct


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

[**CollectionOfOnlineMeeting**](Collection_of_onlineMeeting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateOnlineMeetings

> UsersUpdateOnlineMeetings(ctx, userId, onlineMeetingId, microsoftGraphOnlineMeeting)

Update the navigation property onlineMeetings in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**onlineMeetingId** | **string**| key: onlineMeeting-id of onlineMeeting | 
**microsoftGraphOnlineMeeting** | [**MicrosoftGraphOnlineMeeting**](MicrosoftGraphOnlineMeeting.md)| New navigation property values | 

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

