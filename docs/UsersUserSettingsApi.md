# \UsersUserSettingsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetSettings**](UsersUserSettingsApi.md#UsersGetSettings) | **Get** /users/{user-id}/settings | Get settings from users
[**UsersSettingsGetShiftPreferences**](UsersUserSettingsApi.md#UsersSettingsGetShiftPreferences) | **Get** /users/{user-id}/settings/shiftPreferences | Get shiftPreferences from users
[**UsersSettingsUpdateShiftPreferences**](UsersUserSettingsApi.md#UsersSettingsUpdateShiftPreferences) | **Patch** /users/{user-id}/settings/shiftPreferences | Update the navigation property shiftPreferences in users
[**UsersUpdateSettings**](UsersUserSettingsApi.md#UsersUpdateSettings) | **Patch** /users/{user-id}/settings | Update the navigation property settings in users



## UsersGetSettings

> MicrosoftGraphUserSettings UsersGetSettings(ctx, userId, optional)

Get settings from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUserSettings**](microsoft.graph.userSettings.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSettingsGetShiftPreferences

> MicrosoftGraphShiftPreferences UsersSettingsGetShiftPreferences(ctx, userId, optional)

Get shiftPreferences from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersSettingsGetShiftPreferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersSettingsGetShiftPreferencesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphShiftPreferences**](microsoft.graph.shiftPreferences.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSettingsUpdateShiftPreferences

> UsersSettingsUpdateShiftPreferences(ctx, userId, microsoftGraphShiftPreferences)

Update the navigation property shiftPreferences in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphShiftPreferences** | [**MicrosoftGraphShiftPreferences**](MicrosoftGraphShiftPreferences.md)| New navigation property values | 

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


## UsersUpdateSettings

> UsersUpdateSettings(ctx, userId, microsoftGraphUserSettings)

Update the navigation property settings in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphUserSettings** | [**MicrosoftGraphUserSettings**](MicrosoftGraphUserSettings.md)| New navigation property values | 

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

