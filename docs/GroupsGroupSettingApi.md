# \GroupsGroupSettingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateSettings**](GroupsGroupSettingApi.md#GroupsCreateSettings) | **Post** /groups/{group-id}/settings | Create new navigation property to settings for groups
[**GroupsGetSettings**](GroupsGroupSettingApi.md#GroupsGetSettings) | **Get** /groups/{group-id}/settings/{groupSetting-id} | Get settings from groups
[**GroupsListSettings**](GroupsGroupSettingApi.md#GroupsListSettings) | **Get** /groups/{group-id}/settings | Get settings from groups
[**GroupsUpdateSettings**](GroupsGroupSettingApi.md#GroupsUpdateSettings) | **Patch** /groups/{group-id}/settings/{groupSetting-id} | Update the navigation property settings in groups



## GroupsCreateSettings

> MicrosoftGraphGroupSetting GroupsCreateSettings(ctx, groupId, microsoftGraphGroupSetting)

Create new navigation property to settings for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphGroupSetting** | [**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md)| New navigation property | 

### Return type

[**MicrosoftGraphGroupSetting**](microsoft.graph.groupSetting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetSettings

> MicrosoftGraphGroupSetting GroupsGetSettings(ctx, groupId, groupSettingId, optional)

Get settings from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**groupSettingId** | **string**| key: groupSetting-id of groupSetting | 
 **optional** | ***GroupsGetSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphGroupSetting**](microsoft.graph.groupSetting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsListSettings

> CollectionOfGroupSetting GroupsListSettings(ctx, groupId, optional)

Get settings from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsListSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsListSettingsOpts struct


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

[**CollectionOfGroupSetting**](Collection_of_groupSetting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdateSettings

> GroupsUpdateSettings(ctx, groupId, groupSettingId, microsoftGraphGroupSetting)

Update the navigation property settings in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**groupSettingId** | **string**| key: groupSetting-id of groupSetting | 
**microsoftGraphGroupSetting** | [**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md)| New navigation property values | 

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

