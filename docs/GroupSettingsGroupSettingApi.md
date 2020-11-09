# \GroupSettingsGroupSettingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupSettingsGroupSettingCreateGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingCreateGroupSetting) | **Post** /groupSettings | Add new entity to groupSettings
[**GroupSettingsGroupSettingDeleteGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingDeleteGroupSetting) | **Delete** /groupSettings/{groupSetting-id} | Delete entity from groupSettings
[**GroupSettingsGroupSettingGetGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingGetGroupSetting) | **Get** /groupSettings/{groupSetting-id} | Get entity from groupSettings by key
[**GroupSettingsGroupSettingListGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingListGroupSetting) | **Get** /groupSettings | Get entities from groupSettings
[**GroupSettingsGroupSettingUpdateGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingUpdateGroupSetting) | **Patch** /groupSettings/{groupSetting-id} | Update entity in groupSettings



## GroupSettingsGroupSettingCreateGroupSetting

> MicrosoftGraphGroupSetting GroupSettingsGroupSettingCreateGroupSetting(ctx, microsoftGraphGroupSetting)

Add new entity to groupSettings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphGroupSetting** | [**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md)| New entity | 

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


## GroupSettingsGroupSettingDeleteGroupSetting

> GroupSettingsGroupSettingDeleteGroupSetting(ctx, groupSettingId, optional)

Delete entity from groupSettings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingId** | **string**| key: groupSetting-id of groupSetting | 
 **optional** | ***GroupSettingsGroupSettingDeleteGroupSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupSettingsGroupSettingDeleteGroupSettingOpts struct


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


## GroupSettingsGroupSettingGetGroupSetting

> MicrosoftGraphGroupSetting GroupSettingsGroupSettingGetGroupSetting(ctx, groupSettingId, optional)

Get entity from groupSettings by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingId** | **string**| key: groupSetting-id of groupSetting | 
 **optional** | ***GroupSettingsGroupSettingGetGroupSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupSettingsGroupSettingGetGroupSettingOpts struct


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


## GroupSettingsGroupSettingListGroupSetting

> CollectionOfGroupSetting GroupSettingsGroupSettingListGroupSetting(ctx, optional)

Get entities from groupSettings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupSettingsGroupSettingListGroupSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupSettingsGroupSettingListGroupSettingOpts struct


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


## GroupSettingsGroupSettingUpdateGroupSetting

> GroupSettingsGroupSettingUpdateGroupSetting(ctx, groupSettingId, microsoftGraphGroupSetting)

Update entity in groupSettings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingId** | **string**| key: groupSetting-id of groupSetting | 
**microsoftGraphGroupSetting** | [**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md)| New property values | 

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

