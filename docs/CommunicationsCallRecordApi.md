# \CommunicationsCallRecordApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCallRecordsCreateSessions**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsCreateSessions) | **Post** /communications/callRecords/{callRecord-id}/sessions | Create new navigation property to sessions for communications
[**CommunicationsCallRecordsGetSessions**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsGetSessions) | **Get** /communications/callRecords/{callRecord-id}/sessions/{session-id} | Get sessions from communications
[**CommunicationsCallRecordsListSessions**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsListSessions) | **Get** /communications/callRecords/{callRecord-id}/sessions | Get sessions from communications
[**CommunicationsCallRecordsSessionsCreateSegments**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsSessionsCreateSegments) | **Post** /communications/callRecords/{callRecord-id}/sessions/{session-id}/segments | Create new navigation property to segments for communications
[**CommunicationsCallRecordsSessionsGetSegments**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsSessionsGetSegments) | **Get** /communications/callRecords/{callRecord-id}/sessions/{session-id}/segments/{segment-id} | Get segments from communications
[**CommunicationsCallRecordsSessionsListSegments**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsSessionsListSegments) | **Get** /communications/callRecords/{callRecord-id}/sessions/{session-id}/segments | Get segments from communications
[**CommunicationsCallRecordsSessionsUpdateSegments**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsSessionsUpdateSegments) | **Patch** /communications/callRecords/{callRecord-id}/sessions/{session-id}/segments/{segment-id} | Update the navigation property segments in communications
[**CommunicationsCallRecordsUpdateSessions**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsUpdateSessions) | **Patch** /communications/callRecords/{callRecord-id}/sessions/{session-id} | Update the navigation property sessions in communications
[**CommunicationsCreateCallRecords**](CommunicationsCallRecordApi.md#CommunicationsCreateCallRecords) | **Post** /communications/callRecords | Create new navigation property to callRecords for communications
[**CommunicationsGetCallRecords**](CommunicationsCallRecordApi.md#CommunicationsGetCallRecords) | **Get** /communications/callRecords/{callRecord-id} | Get callRecords from communications
[**CommunicationsListCallRecords**](CommunicationsCallRecordApi.md#CommunicationsListCallRecords) | **Get** /communications/callRecords | Get callRecords from communications
[**CommunicationsUpdateCallRecords**](CommunicationsCallRecordApi.md#CommunicationsUpdateCallRecords) | **Patch** /communications/callRecords/{callRecord-id} | Update the navigation property callRecords in communications



## CommunicationsCallRecordsCreateSessions

> MicrosoftGraphCallRecordsSession CommunicationsCallRecordsCreateSessions(ctx, callRecordId, microsoftGraphCallRecordsSession)

Create new navigation property to sessions for communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
**microsoftGraphCallRecordsSession** | [**MicrosoftGraphCallRecordsSession**](MicrosoftGraphCallRecordsSession.md)| New navigation property | 

### Return type

[**MicrosoftGraphCallRecordsSession**](microsoft.graph.callRecords.session.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsGetSessions

> MicrosoftGraphCallRecordsSession CommunicationsCallRecordsGetSessions(ctx, callRecordId, sessionId, optional)

Get sessions from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
**sessionId** | **string**| key: session-id of session | 
 **optional** | ***CommunicationsCallRecordsGetSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsCallRecordsGetSessionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphCallRecordsSession**](microsoft.graph.callRecords.session.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsListSessions

> CollectionOfSession CommunicationsCallRecordsListSessions(ctx, callRecordId, optional)

Get sessions from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
 **optional** | ***CommunicationsCallRecordsListSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsCallRecordsListSessionsOpts struct


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

[**CollectionOfSession**](Collection_of_session.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsSessionsCreateSegments

> MicrosoftGraphCallRecordsSegment CommunicationsCallRecordsSessionsCreateSegments(ctx, callRecordId, sessionId, microsoftGraphCallRecordsSegment)

Create new navigation property to segments for communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
**sessionId** | **string**| key: session-id of session | 
**microsoftGraphCallRecordsSegment** | [**MicrosoftGraphCallRecordsSegment**](MicrosoftGraphCallRecordsSegment.md)| New navigation property | 

### Return type

[**MicrosoftGraphCallRecordsSegment**](microsoft.graph.callRecords.segment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsSessionsGetSegments

> MicrosoftGraphCallRecordsSegment CommunicationsCallRecordsSessionsGetSegments(ctx, callRecordId, sessionId, segmentId, optional)

Get segments from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
**sessionId** | **string**| key: session-id of session | 
**segmentId** | **string**| key: segment-id of segment | 
 **optional** | ***CommunicationsCallRecordsSessionsGetSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsCallRecordsSessionsGetSegmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphCallRecordsSegment**](microsoft.graph.callRecords.segment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsSessionsListSegments

> CollectionOfSegment CommunicationsCallRecordsSessionsListSegments(ctx, callRecordId, sessionId, optional)

Get segments from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
**sessionId** | **string**| key: session-id of session | 
 **optional** | ***CommunicationsCallRecordsSessionsListSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsCallRecordsSessionsListSegmentsOpts struct


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

[**CollectionOfSegment**](Collection_of_segment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsSessionsUpdateSegments

> CommunicationsCallRecordsSessionsUpdateSegments(ctx, callRecordId, sessionId, segmentId, microsoftGraphCallRecordsSegment)

Update the navigation property segments in communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
**sessionId** | **string**| key: session-id of session | 
**segmentId** | **string**| key: segment-id of segment | 
**microsoftGraphCallRecordsSegment** | [**MicrosoftGraphCallRecordsSegment**](MicrosoftGraphCallRecordsSegment.md)| New navigation property values | 

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


## CommunicationsCallRecordsUpdateSessions

> CommunicationsCallRecordsUpdateSessions(ctx, callRecordId, sessionId, microsoftGraphCallRecordsSession)

Update the navigation property sessions in communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
**sessionId** | **string**| key: session-id of session | 
**microsoftGraphCallRecordsSession** | [**MicrosoftGraphCallRecordsSession**](MicrosoftGraphCallRecordsSession.md)| New navigation property values | 

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


## CommunicationsCreateCallRecords

> MicrosoftGraphCallRecordsCallRecord CommunicationsCreateCallRecords(ctx, microsoftGraphCallRecordsCallRecord)

Create new navigation property to callRecords for communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphCallRecordsCallRecord** | [**MicrosoftGraphCallRecordsCallRecord**](MicrosoftGraphCallRecordsCallRecord.md)| New navigation property | 

### Return type

[**MicrosoftGraphCallRecordsCallRecord**](microsoft.graph.callRecords.callRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsGetCallRecords

> MicrosoftGraphCallRecordsCallRecord CommunicationsGetCallRecords(ctx, callRecordId, optional)

Get callRecords from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
 **optional** | ***CommunicationsGetCallRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsGetCallRecordsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphCallRecordsCallRecord**](microsoft.graph.callRecords.callRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsListCallRecords

> CollectionOfCallRecord CommunicationsListCallRecords(ctx, optional)

Get callRecords from communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommunicationsListCallRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommunicationsListCallRecordsOpts struct


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

[**CollectionOfCallRecord**](Collection_of_callRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsUpdateCallRecords

> CommunicationsUpdateCallRecords(ctx, callRecordId, microsoftGraphCallRecordsCallRecord)

Update the navigation property callRecords in communications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string**| key: callRecord-id of callRecord | 
**microsoftGraphCallRecordsCallRecord** | [**MicrosoftGraphCallRecordsCallRecord**](MicrosoftGraphCallRecordsCallRecord.md)| New navigation property values | 

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

