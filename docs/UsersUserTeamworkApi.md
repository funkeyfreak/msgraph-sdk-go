# \UsersUserTeamworkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetTeamwork**](UsersUserTeamworkApi.md#UsersGetTeamwork) | **Get** /users/{user-id}/teamwork | Get teamwork from users
[**UsersTeamworkCreateInstalledApps**](UsersUserTeamworkApi.md#UsersTeamworkCreateInstalledApps) | **Post** /users/{user-id}/teamwork/installedApps | Create new navigation property to installedApps for users
[**UsersTeamworkGetInstalledApps**](UsersUserTeamworkApi.md#UsersTeamworkGetInstalledApps) | **Get** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Get installedApps from users
[**UsersTeamworkInstalledAppsGetChat**](UsersUserTeamworkApi.md#UsersTeamworkInstalledAppsGetChat) | **Get** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat | Get chat from users
[**UsersTeamworkListInstalledApps**](UsersUserTeamworkApi.md#UsersTeamworkListInstalledApps) | **Get** /users/{user-id}/teamwork/installedApps | Get installedApps from users
[**UsersTeamworkUpdateInstalledApps**](UsersUserTeamworkApi.md#UsersTeamworkUpdateInstalledApps) | **Patch** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Update the navigation property installedApps in users
[**UsersUpdateTeamwork**](UsersUserTeamworkApi.md#UsersUpdateTeamwork) | **Patch** /users/{user-id}/teamwork | Update the navigation property teamwork in users



## UsersGetTeamwork

> MicrosoftGraphUserTeamwork UsersGetTeamwork(ctx, userId, optional)

Get teamwork from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetTeamworkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetTeamworkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUserTeamwork**](microsoft.graph.userTeamwork.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTeamworkCreateInstalledApps

> MicrosoftGraphUserScopeTeamsAppInstallation UsersTeamworkCreateInstalledApps(ctx, userId, microsoftGraphUserScopeTeamsAppInstallation)

Create new navigation property to installedApps for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphUserScopeTeamsAppInstallation** | [**MicrosoftGraphUserScopeTeamsAppInstallation**](MicrosoftGraphUserScopeTeamsAppInstallation.md)| New navigation property | 

### Return type

[**MicrosoftGraphUserScopeTeamsAppInstallation**](microsoft.graph.userScopeTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTeamworkGetInstalledApps

> MicrosoftGraphUserScopeTeamsAppInstallation UsersTeamworkGetInstalledApps(ctx, userId, userScopeTeamsAppInstallationId, optional)

Get installedApps from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userScopeTeamsAppInstallationId** | **string**| key: userScopeTeamsAppInstallation-id of userScopeTeamsAppInstallation | 
 **optional** | ***UsersTeamworkGetInstalledAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTeamworkGetInstalledAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUserScopeTeamsAppInstallation**](microsoft.graph.userScopeTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTeamworkInstalledAppsGetChat

> MicrosoftGraphChat UsersTeamworkInstalledAppsGetChat(ctx, userId, userScopeTeamsAppInstallationId, optional)

Get chat from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userScopeTeamsAppInstallationId** | **string**| key: userScopeTeamsAppInstallation-id of userScopeTeamsAppInstallation | 
 **optional** | ***UsersTeamworkInstalledAppsGetChatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTeamworkInstalledAppsGetChatOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphChat**](microsoft.graph.chat.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTeamworkListInstalledApps

> CollectionOfUserScopeTeamsAppInstallation UsersTeamworkListInstalledApps(ctx, userId, optional)

Get installedApps from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersTeamworkListInstalledAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTeamworkListInstalledAppsOpts struct


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

[**CollectionOfUserScopeTeamsAppInstallation**](Collection_of_userScopeTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTeamworkUpdateInstalledApps

> UsersTeamworkUpdateInstalledApps(ctx, userId, userScopeTeamsAppInstallationId, microsoftGraphUserScopeTeamsAppInstallation)

Update the navigation property installedApps in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userScopeTeamsAppInstallationId** | **string**| key: userScopeTeamsAppInstallation-id of userScopeTeamsAppInstallation | 
**microsoftGraphUserScopeTeamsAppInstallation** | [**MicrosoftGraphUserScopeTeamsAppInstallation**](MicrosoftGraphUserScopeTeamsAppInstallation.md)| New navigation property values | 

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


## UsersUpdateTeamwork

> UsersUpdateTeamwork(ctx, userId, microsoftGraphUserTeamwork)

Update the navigation property teamwork in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphUserTeamwork** | [**MicrosoftGraphUserTeamwork**](MicrosoftGraphUserTeamwork.md)| New navigation property values | 

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

