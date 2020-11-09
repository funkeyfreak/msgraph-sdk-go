# \CommunicationsOnlineMeetingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCreateOnlineMeetings**](CommunicationsOnlineMeetingApi.md#CommunicationsCreateOnlineMeetings) | **Post** /communications/onlineMeetings | Create new navigation property to onlineMeetings for communications
[**CommunicationsGetOnlineMeetings**](CommunicationsOnlineMeetingApi.md#CommunicationsGetOnlineMeetings) | **Get** /communications/onlineMeetings/{onlineMeeting-id} | Get onlineMeetings from communications
[**CommunicationsListOnlineMeetings**](CommunicationsOnlineMeetingApi.md#CommunicationsListOnlineMeetings) | **Get** /communications/onlineMeetings | Get onlineMeetings from communications
[**CommunicationsUpdateOnlineMeetings**](CommunicationsOnlineMeetingApi.md#CommunicationsUpdateOnlineMeetings) | **Patch** /communications/onlineMeetings/{onlineMeeting-id} | Update the navigation property onlineMeetings in communications



## CommunicationsCreateOnlineMeetings

> MicrosoftGraphOnlineMeeting CommunicationsCreateOnlineMeetings(ctx, microsoftGraphOnlineMeeting)

Create new navigation property to onlineMeetings for communications

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


## CommunicationsGetOnlineMeetings

> MicrosoftGraphOnlineMeeting CommunicationsGetOnlineMeetings(ctx, onlineMeetingId, optional)

Get onlineMeetings from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onlineMeetingId** | **string**| key: onlineMeeting-id of onlineMeeting | 
 **optional** | ***CommunicationsGetOnlineMeetingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsGetOnlineMeetingsOpts struct


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


## CommunicationsListOnlineMeetings

> CollectionOfOnlineMeeting CommunicationsListOnlineMeetings(ctx, optional)

Get onlineMeetings from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommunicationsListOnlineMeetingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsListOnlineMeetingsOpts struct


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


## CommunicationsUpdateOnlineMeetings

> CommunicationsUpdateOnlineMeetings(ctx, onlineMeetingId, microsoftGraphOnlineMeeting)

Update the navigation property onlineMeetings in communications

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

