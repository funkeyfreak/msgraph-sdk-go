# \MeOnlineMeetingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateOnlineMeetings**](MeOnlineMeetingApi.md#MeCreateOnlineMeetings) | **Post** /me/onlineMeetings | Create new navigation property to onlineMeetings for me
[**MeGetOnlineMeetings**](MeOnlineMeetingApi.md#MeGetOnlineMeetings) | **Get** /me/onlineMeetings/{onlineMeeting-id} | Get onlineMeetings from me
[**MeListOnlineMeetings**](MeOnlineMeetingApi.md#MeListOnlineMeetings) | **Get** /me/onlineMeetings | Get onlineMeetings from me
[**MeUpdateOnlineMeetings**](MeOnlineMeetingApi.md#MeUpdateOnlineMeetings) | **Patch** /me/onlineMeetings/{onlineMeeting-id} | Update the navigation property onlineMeetings in me



## MeCreateOnlineMeetings

> MicrosoftGraphOnlineMeeting MeCreateOnlineMeetings(ctx, microsoftGraphOnlineMeeting)

Create new navigation property to onlineMeetings for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeGetOnlineMeetings

> MicrosoftGraphOnlineMeeting MeGetOnlineMeetings(ctx, onlineMeetingId, optional)

Get onlineMeetings from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onlineMeetingId** | **string**| key: onlineMeeting-id of onlineMeeting | 
 **optional** | ***MeGetOnlineMeetingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetOnlineMeetingsOpts struct


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


## MeListOnlineMeetings

> CollectionOfOnlineMeeting MeListOnlineMeetings(ctx, optional)

Get onlineMeetings from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListOnlineMeetingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListOnlineMeetingsOpts struct


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


## MeUpdateOnlineMeetings

> MeUpdateOnlineMeetings(ctx, onlineMeetingId, microsoftGraphOnlineMeeting)

Update the navigation property onlineMeetings in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

