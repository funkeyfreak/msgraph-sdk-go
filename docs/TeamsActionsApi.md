# \TeamsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsArchive**](TeamsActionsApi.md#TeamsArchive) | **Post** /teams/{team-id}/microsoft.graph.archive | Invoke action archive
[**TeamsClone**](TeamsActionsApi.md#TeamsClone) | **Post** /teams/{team-id}/microsoft.graph.clone | Invoke action clone
[**TeamsInstalledAppsUpgrade**](TeamsActionsApi.md#TeamsInstalledAppsUpgrade) | **Post** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/microsoft.graph.upgrade | Invoke action upgrade
[**TeamsScheduleShare**](TeamsActionsApi.md#TeamsScheduleShare) | **Post** /teams/{team-id}/schedule/microsoft.graph.share | Invoke action share
[**TeamsUnarchive**](TeamsActionsApi.md#TeamsUnarchive) | **Post** /teams/{team-id}/microsoft.graph.unarchive | Invoke action unarchive



## TeamsArchive

> TeamsArchive(ctx, teamId, inlineObject625)

Invoke action archive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**inlineObject625** | [**InlineObject625**](InlineObject625.md)|  | 

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


## TeamsClone

> TeamsClone(ctx, teamId, inlineObject626)

Invoke action clone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**inlineObject626** | [**InlineObject626**](InlineObject626.md)|  | 

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


## TeamsInstalledAppsUpgrade

> TeamsInstalledAppsUpgrade(ctx, teamId, teamsAppInstallationId)

Invoke action upgrade

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**teamsAppInstallationId** | **string**| key: teamsAppInstallation-id of teamsAppInstallation | 

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


## TeamsScheduleShare

> TeamsScheduleShare(ctx, teamId, inlineObject627)

Invoke action share

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**inlineObject627** | [**InlineObject627**](InlineObject627.md)|  | 

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


## TeamsUnarchive

> TeamsUnarchive(ctx, teamId)

Invoke action unarchive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 

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

