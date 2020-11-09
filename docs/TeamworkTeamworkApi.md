# \TeamworkTeamworkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamworkTeamworkGetTeamwork**](TeamworkTeamworkApi.md#TeamworkTeamworkGetTeamwork) | **Get** /teamwork | Get teamwork
[**TeamworkTeamworkUpdateTeamwork**](TeamworkTeamworkApi.md#TeamworkTeamworkUpdateTeamwork) | **Patch** /teamwork | Update teamwork



## TeamworkTeamworkGetTeamwork

> MicrosoftGraphTeamwork TeamworkTeamworkGetTeamwork(ctx, optional)

Get teamwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamworkTeamworkGetTeamworkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamworkTeamworkGetTeamworkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTeamwork**](microsoft.graph.teamwork.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamworkTeamworkUpdateTeamwork

> TeamworkTeamworkUpdateTeamwork(ctx, microsoftGraphTeamwork)

Update teamwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTeamwork** | [**MicrosoftGraphTeamwork**](MicrosoftGraphTeamwork.md)| New property values | 

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

