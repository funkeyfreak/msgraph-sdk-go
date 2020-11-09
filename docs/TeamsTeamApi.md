# \TeamsTeamApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsTeamCreateTeam**](TeamsTeamApi.md#TeamsTeamCreateTeam) | **Post** /teams | Add new entity to teams
[**TeamsTeamDeleteTeam**](TeamsTeamApi.md#TeamsTeamDeleteTeam) | **Delete** /teams/{team-id} | Delete entity from teams
[**TeamsTeamGetTeam**](TeamsTeamApi.md#TeamsTeamGetTeam) | **Get** /teams/{team-id} | Get entity from teams by key
[**TeamsTeamListTeam**](TeamsTeamApi.md#TeamsTeamListTeam) | **Get** /teams | Get entities from teams
[**TeamsTeamUpdateTeam**](TeamsTeamApi.md#TeamsTeamUpdateTeam) | **Patch** /teams/{team-id} | Update entity in teams



## TeamsTeamCreateTeam

> MicrosoftGraphTeam TeamsTeamCreateTeam(ctx, microsoftGraphTeam)

Add new entity to teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTeam** | [**MicrosoftGraphTeam**](MicrosoftGraphTeam.md)| New entity | 

### Return type

[**MicrosoftGraphTeam**](microsoft.graph.team.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamDeleteTeam

> TeamsTeamDeleteTeam(ctx, teamId, optional)

Delete entity from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
 **optional** | ***TeamsTeamDeleteTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsTeamDeleteTeamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| ETag | 

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


## TeamsTeamGetTeam

> MicrosoftGraphTeam TeamsTeamGetTeam(ctx, teamId, optional)

Get entity from teams by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
 **optional** | ***TeamsTeamGetTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsTeamGetTeamOpts struct


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


## TeamsTeamListTeam

> CollectionOfTeam TeamsTeamListTeam(ctx, optional)

Get entities from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsTeamListTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsTeamListTeamOpts struct


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

[**CollectionOfTeam**](Collection_of_team.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamUpdateTeam

> TeamsTeamUpdateTeam(ctx, teamId, microsoftGraphTeam)

Update entity in teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**microsoftGraphTeam** | [**MicrosoftGraphTeam**](MicrosoftGraphTeam.md)| New property values | 

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

