# \MeUserTeamworkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeGetTeamwork**](MeUserTeamworkApi.md#MeGetTeamwork) | **Get** /me/teamwork | Get teamwork from me
[**MeTeamworkCreateInstalledApps**](MeUserTeamworkApi.md#MeTeamworkCreateInstalledApps) | **Post** /me/teamwork/installedApps | Create new navigation property to installedApps for me
[**MeTeamworkGetInstalledApps**](MeUserTeamworkApi.md#MeTeamworkGetInstalledApps) | **Get** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Get installedApps from me
[**MeTeamworkInstalledAppsGetChat**](MeUserTeamworkApi.md#MeTeamworkInstalledAppsGetChat) | **Get** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat | Get chat from me
[**MeTeamworkListInstalledApps**](MeUserTeamworkApi.md#MeTeamworkListInstalledApps) | **Get** /me/teamwork/installedApps | Get installedApps from me
[**MeTeamworkUpdateInstalledApps**](MeUserTeamworkApi.md#MeTeamworkUpdateInstalledApps) | **Patch** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Update the navigation property installedApps in me
[**MeUpdateTeamwork**](MeUserTeamworkApi.md#MeUpdateTeamwork) | **Patch** /me/teamwork | Update the navigation property teamwork in me



## MeGetTeamwork

> MicrosoftGraphUserTeamwork MeGetTeamwork(ctx, optional)

Get teamwork from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeGetTeamworkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetTeamworkOpts struct


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


## MeTeamworkCreateInstalledApps

> MicrosoftGraphUserScopeTeamsAppInstallation MeTeamworkCreateInstalledApps(ctx, microsoftGraphUserScopeTeamsAppInstallation)

Create new navigation property to installedApps for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTeamworkGetInstalledApps

> MicrosoftGraphUserScopeTeamsAppInstallation MeTeamworkGetInstalledApps(ctx, userScopeTeamsAppInstallationId, optional)

Get installedApps from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userScopeTeamsAppInstallationId** | **string**| key: userScopeTeamsAppInstallation-id of userScopeTeamsAppInstallation | 
 **optional** | ***MeTeamworkGetInstalledAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTeamworkGetInstalledAppsOpts struct


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


## MeTeamworkInstalledAppsGetChat

> MicrosoftGraphChat MeTeamworkInstalledAppsGetChat(ctx, userScopeTeamsAppInstallationId, optional)

Get chat from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userScopeTeamsAppInstallationId** | **string**| key: userScopeTeamsAppInstallation-id of userScopeTeamsAppInstallation | 
 **optional** | ***MeTeamworkInstalledAppsGetChatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTeamworkInstalledAppsGetChatOpts struct


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


## MeTeamworkListInstalledApps

> CollectionOfUserScopeTeamsAppInstallation MeTeamworkListInstalledApps(ctx, optional)

Get installedApps from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeTeamworkListInstalledAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTeamworkListInstalledAppsOpts struct


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


## MeTeamworkUpdateInstalledApps

> MeTeamworkUpdateInstalledApps(ctx, userScopeTeamsAppInstallationId, microsoftGraphUserScopeTeamsAppInstallation)

Update the navigation property installedApps in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeUpdateTeamwork

> MeUpdateTeamwork(ctx, microsoftGraphUserTeamwork)

Update the navigation property teamwork in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

