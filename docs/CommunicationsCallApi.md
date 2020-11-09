# \CommunicationsCallApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCallsCreateOperations**](CommunicationsCallApi.md#CommunicationsCallsCreateOperations) | **Post** /communications/calls/{call-id}/operations | Create new navigation property to operations for communications
[**CommunicationsCallsCreateParticipants**](CommunicationsCallApi.md#CommunicationsCallsCreateParticipants) | **Post** /communications/calls/{call-id}/participants | Create new navigation property to participants for communications
[**CommunicationsCallsGetOperations**](CommunicationsCallApi.md#CommunicationsCallsGetOperations) | **Get** /communications/calls/{call-id}/operations/{commsOperation-id} | Get operations from communications
[**CommunicationsCallsGetParticipants**](CommunicationsCallApi.md#CommunicationsCallsGetParticipants) | **Get** /communications/calls/{call-id}/participants/{participant-id} | Get participants from communications
[**CommunicationsCallsListOperations**](CommunicationsCallApi.md#CommunicationsCallsListOperations) | **Get** /communications/calls/{call-id}/operations | Get operations from communications
[**CommunicationsCallsListParticipants**](CommunicationsCallApi.md#CommunicationsCallsListParticipants) | **Get** /communications/calls/{call-id}/participants | Get participants from communications
[**CommunicationsCallsUpdateOperations**](CommunicationsCallApi.md#CommunicationsCallsUpdateOperations) | **Patch** /communications/calls/{call-id}/operations/{commsOperation-id} | Update the navigation property operations in communications
[**CommunicationsCallsUpdateParticipants**](CommunicationsCallApi.md#CommunicationsCallsUpdateParticipants) | **Patch** /communications/calls/{call-id}/participants/{participant-id} | Update the navigation property participants in communications
[**CommunicationsCreateCalls**](CommunicationsCallApi.md#CommunicationsCreateCalls) | **Post** /communications/calls | Create new navigation property to calls for communications
[**CommunicationsGetCalls**](CommunicationsCallApi.md#CommunicationsGetCalls) | **Get** /communications/calls/{call-id} | Get calls from communications
[**CommunicationsListCalls**](CommunicationsCallApi.md#CommunicationsListCalls) | **Get** /communications/calls | Get calls from communications
[**CommunicationsUpdateCalls**](CommunicationsCallApi.md#CommunicationsUpdateCalls) | **Patch** /communications/calls/{call-id} | Update the navigation property calls in communications



## CommunicationsCallsCreateOperations

> MicrosoftGraphCommsOperation CommunicationsCallsCreateOperations(ctx, callId, microsoftGraphCommsOperation)

Create new navigation property to operations for communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**microsoftGraphCommsOperation** | [**MicrosoftGraphCommsOperation**](MicrosoftGraphCommsOperation.md)| New navigation property | 

### Return type

[**MicrosoftGraphCommsOperation**](microsoft.graph.commsOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCreateParticipants

> MicrosoftGraphParticipant CommunicationsCallsCreateParticipants(ctx, callId, microsoftGraphParticipant)

Create new navigation property to participants for communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**microsoftGraphParticipant** | [**MicrosoftGraphParticipant**](MicrosoftGraphParticipant.md)| New navigation property | 

### Return type

[**MicrosoftGraphParticipant**](microsoft.graph.participant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsGetOperations

> MicrosoftGraphCommsOperation CommunicationsCallsGetOperations(ctx, callId, commsOperationId, optional)

Get operations from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**commsOperationId** | **string**| key: commsOperation-id of commsOperation | 
 **optional** | ***CommunicationsCallsGetOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsCallsGetOperationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphCommsOperation**](microsoft.graph.commsOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsGetParticipants

> MicrosoftGraphParticipant CommunicationsCallsGetParticipants(ctx, callId, participantId, optional)

Get participants from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**participantId** | **string**| key: participant-id of participant | 
 **optional** | ***CommunicationsCallsGetParticipantsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsCallsGetParticipantsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphParticipant**](microsoft.graph.participant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsListOperations

> CollectionOfCommsOperation CommunicationsCallsListOperations(ctx, callId, optional)

Get operations from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
 **optional** | ***CommunicationsCallsListOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsCallsListOperationsOpts struct


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

[**CollectionOfCommsOperation**](Collection_of_commsOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsListParticipants

> CollectionOfParticipant CommunicationsCallsListParticipants(ctx, callId, optional)

Get participants from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
 **optional** | ***CommunicationsCallsListParticipantsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsCallsListParticipantsOpts struct


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

[**CollectionOfParticipant**](Collection_of_participant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsUpdateOperations

> CommunicationsCallsUpdateOperations(ctx, callId, commsOperationId, microsoftGraphCommsOperation)

Update the navigation property operations in communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**commsOperationId** | **string**| key: commsOperation-id of commsOperation | 
**microsoftGraphCommsOperation** | [**MicrosoftGraphCommsOperation**](MicrosoftGraphCommsOperation.md)| New navigation property values | 

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


## CommunicationsCallsUpdateParticipants

> CommunicationsCallsUpdateParticipants(ctx, callId, participantId, microsoftGraphParticipant)

Update the navigation property participants in communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**participantId** | **string**| key: participant-id of participant | 
**microsoftGraphParticipant** | [**MicrosoftGraphParticipant**](MicrosoftGraphParticipant.md)| New navigation property values | 

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


## CommunicationsCreateCalls

> MicrosoftGraphCall CommunicationsCreateCalls(ctx, microsoftGraphCall)

Create new navigation property to calls for communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphCall** | [**MicrosoftGraphCall**](MicrosoftGraphCall.md)| New navigation property | 

### Return type

[**MicrosoftGraphCall**](microsoft.graph.call.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsGetCalls

> MicrosoftGraphCall CommunicationsGetCalls(ctx, callId, optional)

Get calls from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
 **optional** | ***CommunicationsGetCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsGetCallsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphCall**](microsoft.graph.call.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsListCalls

> CollectionOfCall CommunicationsListCalls(ctx, optional)

Get calls from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommunicationsListCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsListCallsOpts struct


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

[**CollectionOfCall**](Collection_of_call.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsUpdateCalls

> CommunicationsUpdateCalls(ctx, callId, microsoftGraphCall)

Update the navigation property calls in communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**microsoftGraphCall** | [**MicrosoftGraphCall**](MicrosoftGraphCall.md)| New navigation property values | 

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

