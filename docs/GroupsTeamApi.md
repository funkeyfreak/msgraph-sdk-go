# \GroupsTeamApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGetTeam**](GroupsTeamApi.md#GroupsGetTeam) | **Get** /groups/{group-id}/team | Get team from groups
[**GroupsUpdateTeam**](GroupsTeamApi.md#GroupsUpdateTeam) | **Patch** /groups/{group-id}/team | Update the navigation property team in groups



## GroupsGetTeam

> MicrosoftGraphTeam GroupsGetTeam(ctx, groupId, optional)

Get team from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsGetTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetTeamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTeam**](microsoft.graph.team.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdateTeam

> GroupsUpdateTeam(ctx, groupId, microsoftGraphTeam)

Update the navigation property team in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphTeam** | [**MicrosoftGraphTeam**](MicrosoftGraphTeam.md)| New navigation property values | 

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

