# \UsersTeamApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateJoinedTeams**](UsersTeamApi.md#UsersCreateJoinedTeams) | **Post** /users/{user-id}/joinedTeams | Create new navigation property to joinedTeams for users
[**UsersGetJoinedTeams**](UsersTeamApi.md#UsersGetJoinedTeams) | **Get** /users/{user-id}/joinedTeams/{team-id} | Get joinedTeams from users
[**UsersListJoinedTeams**](UsersTeamApi.md#UsersListJoinedTeams) | **Get** /users/{user-id}/joinedTeams | Get joinedTeams from users
[**UsersUpdateJoinedTeams**](UsersTeamApi.md#UsersUpdateJoinedTeams) | **Patch** /users/{user-id}/joinedTeams/{team-id} | Update the navigation property joinedTeams in users



## UsersCreateJoinedTeams

> MicrosoftGraphTeam UsersCreateJoinedTeams(ctx, userId, microsoftGraphTeam)

Create new navigation property to joinedTeams for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphTeam** | [**MicrosoftGraphTeam**](MicrosoftGraphTeam.md)| New navigation property | 

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


## UsersGetJoinedTeams

> MicrosoftGraphTeam UsersGetJoinedTeams(ctx, userId, teamId, optional)

Get joinedTeams from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**teamId** | **string**| key: team-id of team | 
 **optional** | ***UsersGetJoinedTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetJoinedTeamsOpts struct


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


## UsersListJoinedTeams

> CollectionOfTeam UsersListJoinedTeams(ctx, userId, optional)

Get joinedTeams from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListJoinedTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListJoinedTeamsOpts struct


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


## UsersUpdateJoinedTeams

> UsersUpdateJoinedTeams(ctx, userId, teamId, microsoftGraphTeam)

Update the navigation property joinedTeams in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**teamId** | **string**| key: team-id of team | 
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

