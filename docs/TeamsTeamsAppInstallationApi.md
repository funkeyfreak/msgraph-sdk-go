# \TeamsTeamsAppInstallationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsCreateInstalledApps**](TeamsTeamsAppInstallationApi.md#TeamsCreateInstalledApps) | **Post** /teams/{team-id}/installedApps | Create new navigation property to installedApps for teams
[**TeamsGetInstalledApps**](TeamsTeamsAppInstallationApi.md#TeamsGetInstalledApps) | **Get** /teams/{team-id}/installedApps/{teamsAppInstallation-id} | Get installedApps from teams
[**TeamsInstalledAppsGetTeamsApp**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsGetTeamsApp) | **Get** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsApp | Get teamsApp from teams
[**TeamsInstalledAppsGetTeamsAppDefinition**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsGetTeamsAppDefinition) | **Get** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsAppDefinition | Get teamsAppDefinition from teams
[**TeamsListInstalledApps**](TeamsTeamsAppInstallationApi.md#TeamsListInstalledApps) | **Get** /teams/{team-id}/installedApps | Get installedApps from teams
[**TeamsUpdateInstalledApps**](TeamsTeamsAppInstallationApi.md#TeamsUpdateInstalledApps) | **Patch** /teams/{team-id}/installedApps/{teamsAppInstallation-id} | Update the navigation property installedApps in teams



## TeamsCreateInstalledApps

> MicrosoftGraphTeamsAppInstallation TeamsCreateInstalledApps(ctx, teamId, microsoftGraphTeamsAppInstallation)

Create new navigation property to installedApps for teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**microsoftGraphTeamsAppInstallation** | [**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md)| New navigation property | 

### Return type

[**MicrosoftGraphTeamsAppInstallation**](microsoft.graph.teamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetInstalledApps

> MicrosoftGraphTeamsAppInstallation TeamsGetInstalledApps(ctx, teamId, teamsAppInstallationId, optional)

Get installedApps from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**teamsAppInstallationId** | **string**| key: teamsAppInstallation-id of teamsAppInstallation | 
 **optional** | ***TeamsGetInstalledAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsGetInstalledAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTeamsAppInstallation**](microsoft.graph.teamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsInstalledAppsGetTeamsApp

> MicrosoftGraphTeamsApp TeamsInstalledAppsGetTeamsApp(ctx, teamId, teamsAppInstallationId, optional)

Get teamsApp from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**teamsAppInstallationId** | **string**| key: teamsAppInstallation-id of teamsAppInstallation | 
 **optional** | ***TeamsInstalledAppsGetTeamsAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsInstalledAppsGetTeamsAppOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTeamsApp**](microsoft.graph.teamsApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsInstalledAppsGetTeamsAppDefinition

> MicrosoftGraphTeamsAppDefinition TeamsInstalledAppsGetTeamsAppDefinition(ctx, teamId, teamsAppInstallationId, optional)

Get teamsAppDefinition from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**teamsAppInstallationId** | **string**| key: teamsAppInstallation-id of teamsAppInstallation | 
 **optional** | ***TeamsInstalledAppsGetTeamsAppDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsInstalledAppsGetTeamsAppDefinitionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTeamsAppDefinition**](microsoft.graph.teamsAppDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListInstalledApps

> CollectionOfTeamsAppInstallation TeamsListInstalledApps(ctx, teamId, optional)

Get installedApps from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
 **optional** | ***TeamsListInstalledAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsListInstalledAppsOpts struct


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

[**CollectionOfTeamsAppInstallation**](Collection_of_teamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateInstalledApps

> TeamsUpdateInstalledApps(ctx, teamId, teamsAppInstallationId, microsoftGraphTeamsAppInstallation)

Update the navigation property installedApps in teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**teamsAppInstallationId** | **string**| key: teamsAppInstallation-id of teamsAppInstallation | 
**microsoftGraphTeamsAppInstallation** | [**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md)| New navigation property values | 

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

