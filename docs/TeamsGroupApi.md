# \TeamsGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsGetGroup**](TeamsGroupApi.md#TeamsGetGroup) | **Get** /teams/{team-id}/group | Get group from teams



## TeamsGetGroup

> MicrosoftGraphGroup TeamsGetGroup(ctx, teamId, optional)

Get group from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
 **optional** | ***TeamsGetGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsGetGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphGroup**](microsoft.graph.group.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

