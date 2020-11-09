# \CommunicationsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCallsAnswer**](CommunicationsActionsApi.md#CommunicationsCallsAnswer) | **Post** /communications/calls/{call-id}/microsoft.graph.answer | Invoke action answer
[**CommunicationsCallsCancelMediaProcessing**](CommunicationsActionsApi.md#CommunicationsCallsCancelMediaProcessing) | **Post** /communications/calls/{call-id}/microsoft.graph.cancelMediaProcessing | Invoke action cancelMediaProcessing
[**CommunicationsCallsChangeScreenSharingRole**](CommunicationsActionsApi.md#CommunicationsCallsChangeScreenSharingRole) | **Post** /communications/calls/{call-id}/microsoft.graph.changeScreenSharingRole | Invoke action changeScreenSharingRole
[**CommunicationsCallsKeepAlive**](CommunicationsActionsApi.md#CommunicationsCallsKeepAlive) | **Post** /communications/calls/{call-id}/microsoft.graph.keepAlive | Invoke action keepAlive
[**CommunicationsCallsLogTeleconferenceDeviceQuality**](CommunicationsActionsApi.md#CommunicationsCallsLogTeleconferenceDeviceQuality) | **Post** /communications/calls/microsoft.graph.logTeleconferenceDeviceQuality | Invoke action logTeleconferenceDeviceQuality
[**CommunicationsCallsMute**](CommunicationsActionsApi.md#CommunicationsCallsMute) | **Post** /communications/calls/{call-id}/microsoft.graph.mute | Invoke action mute
[**CommunicationsCallsParticipantsInvite**](CommunicationsActionsApi.md#CommunicationsCallsParticipantsInvite) | **Post** /communications/calls/{call-id}/participants/microsoft.graph.invite | Invoke action invite
[**CommunicationsCallsParticipantsMute**](CommunicationsActionsApi.md#CommunicationsCallsParticipantsMute) | **Post** /communications/calls/{call-id}/participants/{participant-id}/microsoft.graph.mute | Invoke action mute
[**CommunicationsCallsPlayPrompt**](CommunicationsActionsApi.md#CommunicationsCallsPlayPrompt) | **Post** /communications/calls/{call-id}/microsoft.graph.playPrompt | Invoke action playPrompt
[**CommunicationsCallsRecordResponse**](CommunicationsActionsApi.md#CommunicationsCallsRecordResponse) | **Post** /communications/calls/{call-id}/microsoft.graph.recordResponse | Invoke action recordResponse
[**CommunicationsCallsRedirect**](CommunicationsActionsApi.md#CommunicationsCallsRedirect) | **Post** /communications/calls/{call-id}/microsoft.graph.redirect | Invoke action redirect
[**CommunicationsCallsReject**](CommunicationsActionsApi.md#CommunicationsCallsReject) | **Post** /communications/calls/{call-id}/microsoft.graph.reject | Invoke action reject
[**CommunicationsCallsSubscribeToTone**](CommunicationsActionsApi.md#CommunicationsCallsSubscribeToTone) | **Post** /communications/calls/{call-id}/microsoft.graph.subscribeToTone | Invoke action subscribeToTone
[**CommunicationsCallsTransfer**](CommunicationsActionsApi.md#CommunicationsCallsTransfer) | **Post** /communications/calls/{call-id}/microsoft.graph.transfer | Invoke action transfer
[**CommunicationsCallsUnmute**](CommunicationsActionsApi.md#CommunicationsCallsUnmute) | **Post** /communications/calls/{call-id}/microsoft.graph.unmute | Invoke action unmute
[**CommunicationsCallsUpdateRecordingStatus**](CommunicationsActionsApi.md#CommunicationsCallsUpdateRecordingStatus) | **Post** /communications/calls/{call-id}/microsoft.graph.updateRecordingStatus | Invoke action updateRecordingStatus
[**CommunicationsOnlineMeetingsCreateOrGet**](CommunicationsActionsApi.md#CommunicationsOnlineMeetingsCreateOrGet) | **Post** /communications/onlineMeetings/microsoft.graph.createOrGet | Invoke action createOrGet



## CommunicationsCallsAnswer

> CommunicationsCallsAnswer(ctx, callId, inlineObject4)

Invoke action answer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject4** | [**InlineObject4**](InlineObject4.md)|  | 

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


## CommunicationsCallsCancelMediaProcessing

> interface{} CommunicationsCallsCancelMediaProcessing(ctx, callId, inlineObject5)

Invoke action cancelMediaProcessing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject5** | [**InlineObject5**](InlineObject5.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsChangeScreenSharingRole

> CommunicationsCallsChangeScreenSharingRole(ctx, callId, inlineObject6)

Invoke action changeScreenSharingRole

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject6** | [**InlineObject6**](InlineObject6.md)|  | 

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


## CommunicationsCallsKeepAlive

> CommunicationsCallsKeepAlive(ctx, callId)

Invoke action keepAlive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsLogTeleconferenceDeviceQuality

> CommunicationsCallsLogTeleconferenceDeviceQuality(ctx, inlineObject18)

Invoke action logTeleconferenceDeviceQuality

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject18** | [**InlineObject18**](InlineObject18.md)|  | 

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


## CommunicationsCallsMute

> interface{} CommunicationsCallsMute(ctx, callId, inlineObject7)

Invoke action mute

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject7** | [**InlineObject7**](InlineObject7.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsParticipantsInvite

> interface{} CommunicationsCallsParticipantsInvite(ctx, callId, inlineObject17)

Invoke action invite

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject17** | [**InlineObject17**](InlineObject17.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsParticipantsMute

> interface{} CommunicationsCallsParticipantsMute(ctx, callId, participantId, inlineObject16)

Invoke action mute

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**participantId** | **string**| key: participant-id of participant | 
**inlineObject16** | [**InlineObject16**](InlineObject16.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsPlayPrompt

> interface{} CommunicationsCallsPlayPrompt(ctx, callId, inlineObject8)

Invoke action playPrompt

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject8** | [**InlineObject8**](InlineObject8.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsRecordResponse

> interface{} CommunicationsCallsRecordResponse(ctx, callId, inlineObject9)

Invoke action recordResponse

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject9** | [**InlineObject9**](InlineObject9.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsRedirect

> CommunicationsCallsRedirect(ctx, callId, inlineObject10)

Invoke action redirect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject10** | [**InlineObject10**](InlineObject10.md)|  | 

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


## CommunicationsCallsReject

> CommunicationsCallsReject(ctx, callId, inlineObject11)

Invoke action reject

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject11** | [**InlineObject11**](InlineObject11.md)|  | 

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


## CommunicationsCallsSubscribeToTone

> interface{} CommunicationsCallsSubscribeToTone(ctx, callId, inlineObject12)

Invoke action subscribeToTone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject12** | [**InlineObject12**](InlineObject12.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsTransfer

> CommunicationsCallsTransfer(ctx, callId, inlineObject13)

Invoke action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject13** | [**InlineObject13**](InlineObject13.md)|  | 

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


## CommunicationsCallsUnmute

> interface{} CommunicationsCallsUnmute(ctx, callId, inlineObject14)

Invoke action unmute

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject14** | [**InlineObject14**](InlineObject14.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsUpdateRecordingStatus

> interface{} CommunicationsCallsUpdateRecordingStatus(ctx, callId, inlineObject15)

Invoke action updateRecordingStatus

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string**| key: call-id of call | 
**inlineObject15** | [**InlineObject15**](InlineObject15.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsOnlineMeetingsCreateOrGet

> interface{} CommunicationsOnlineMeetingsCreateOrGet(ctx, inlineObject19)

Invoke action createOrGet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject19** | [**InlineObject19**](InlineObject19.md)|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

