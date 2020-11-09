# \MeUserSettingsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeGetSettings**](MeUserSettingsApi.md#MeGetSettings) | **Get** /me/settings | Get settings from me
[**MeSettingsGetShiftPreferences**](MeUserSettingsApi.md#MeSettingsGetShiftPreferences) | **Get** /me/settings/shiftPreferences | Get shiftPreferences from me
[**MeSettingsUpdateShiftPreferences**](MeUserSettingsApi.md#MeSettingsUpdateShiftPreferences) | **Patch** /me/settings/shiftPreferences | Update the navigation property shiftPreferences in me
[**MeUpdateSettings**](MeUserSettingsApi.md#MeUpdateSettings) | **Patch** /me/settings | Update the navigation property settings in me



## MeGetSettings

> MicrosoftGraphUserSettings MeGetSettings(ctx, optional)

Get settings from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeGetSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetSettingsOpts struct


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


## MeSettingsGetShiftPreferences

> MicrosoftGraphShiftPreferences MeSettingsGetShiftPreferences(ctx, optional)

Get shiftPreferences from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeSettingsGetShiftPreferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeSettingsGetShiftPreferencesOpts struct


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


## MeSettingsUpdateShiftPreferences

> MeSettingsUpdateShiftPreferences(ctx, microsoftGraphShiftPreferences)

Update the navigation property shiftPreferences in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeUpdateSettings

> MeUpdateSettings(ctx, microsoftGraphUserSettings)

Update the navigation property settings in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

