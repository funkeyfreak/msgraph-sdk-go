# \SharesListItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesGetListItem**](SharesListItemApi.md#SharesGetListItem) | **Get** /shares/{sharedDriveItem-id}/listItem | Get listItem from shares
[**SharesListItemCreateVersions**](SharesListItemApi.md#SharesListItemCreateVersions) | **Post** /shares/{sharedDriveItem-id}/listItem/versions | Create new navigation property to versions for shares
[**SharesListItemGetAnalytics**](SharesListItemApi.md#SharesListItemGetAnalytics) | **Get** /shares/{sharedDriveItem-id}/listItem/analytics | Get analytics from shares
[**SharesListItemGetDriveItem**](SharesListItemApi.md#SharesListItemGetDriveItem) | **Get** /shares/{sharedDriveItem-id}/listItem/driveItem | Get driveItem from shares
[**SharesListItemGetFields**](SharesListItemApi.md#SharesListItemGetFields) | **Get** /shares/{sharedDriveItem-id}/listItem/fields | Get fields from shares
[**SharesListItemGetVersions**](SharesListItemApi.md#SharesListItemGetVersions) | **Get** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id} | Get versions from shares
[**SharesListItemListVersions**](SharesListItemApi.md#SharesListItemListVersions) | **Get** /shares/{sharedDriveItem-id}/listItem/versions | Get versions from shares
[**SharesListItemUpdateDriveItem**](SharesListItemApi.md#SharesListItemUpdateDriveItem) | **Patch** /shares/{sharedDriveItem-id}/listItem/driveItem | Update the navigation property driveItem in shares
[**SharesListItemUpdateFields**](SharesListItemApi.md#SharesListItemUpdateFields) | **Patch** /shares/{sharedDriveItem-id}/listItem/fields | Update the navigation property fields in shares
[**SharesListItemUpdateVersions**](SharesListItemApi.md#SharesListItemUpdateVersions) | **Patch** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id} | Update the navigation property versions in shares
[**SharesListItemVersionsGetFields**](SharesListItemApi.md#SharesListItemVersionsGetFields) | **Get** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id}/fields | Get fields from shares
[**SharesListItemVersionsUpdateFields**](SharesListItemApi.md#SharesListItemVersionsUpdateFields) | **Patch** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id}/fields | Update the navigation property fields in shares
[**SharesUpdateListItem**](SharesListItemApi.md#SharesUpdateListItem) | **Patch** /shares/{sharedDriveItem-id}/listItem | Update the navigation property listItem in shares



## SharesGetListItem

> MicrosoftGraphListItem SharesGetListItem(ctx, sharedDriveItemId, optional)

Get listItem from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesGetListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesGetListItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphListItem**](microsoft.graph.listItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemCreateVersions

> MicrosoftGraphListItemVersion SharesListItemCreateVersions(ctx, sharedDriveItemId, microsoftGraphListItemVersion)

Create new navigation property to versions for shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)| New navigation property | 

### Return type

[**MicrosoftGraphListItemVersion**](microsoft.graph.listItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemGetAnalytics

> MicrosoftGraphItemAnalytics SharesListItemGetAnalytics(ctx, sharedDriveItemId, optional)

Get analytics from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesListItemGetAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemGetAnalyticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphItemAnalytics**](microsoft.graph.itemAnalytics.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemGetDriveItem

> MicrosoftGraphDriveItem SharesListItemGetDriveItem(ctx, sharedDriveItemId, optional)

Get driveItem from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesListItemGetDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemGetDriveItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemGetFields

> MicrosoftGraphFieldValueSet SharesListItemGetFields(ctx, sharedDriveItemId, optional)

Get fields from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesListItemGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemGetFieldsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphFieldValueSet**](microsoft.graph.fieldValueSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemGetVersions

> MicrosoftGraphListItemVersion SharesListItemGetVersions(ctx, sharedDriveItemId, listItemVersionId, optional)

Get versions from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***SharesListItemGetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemGetVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphListItemVersion**](microsoft.graph.listItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemListVersions

> CollectionOfListItemVersion SharesListItemListVersions(ctx, sharedDriveItemId, optional)

Get versions from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesListItemListVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemListVersionsOpts struct


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

[**CollectionOfListItemVersion**](Collection_of_listItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemUpdateDriveItem

> SharesListItemUpdateDriveItem(ctx, sharedDriveItemId, microsoftGraphDriveItem)

Update the navigation property driveItem in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property values | 

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


## SharesListItemUpdateFields

> SharesListItemUpdateFields(ctx, sharedDriveItemId, microsoftGraphFieldValueSet)

Update the navigation property fields in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**microsoftGraphFieldValueSet** | [**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md)| New navigation property values | 

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


## SharesListItemUpdateVersions

> SharesListItemUpdateVersions(ctx, sharedDriveItemId, listItemVersionId, microsoftGraphListItemVersion)

Update the navigation property versions in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
**microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)| New navigation property values | 

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


## SharesListItemVersionsGetFields

> MicrosoftGraphFieldValueSet SharesListItemVersionsGetFields(ctx, sharedDriveItemId, listItemVersionId, optional)

Get fields from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***SharesListItemVersionsGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemVersionsGetFieldsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphFieldValueSet**](microsoft.graph.fieldValueSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemVersionsUpdateFields

> SharesListItemVersionsUpdateFields(ctx, sharedDriveItemId, listItemVersionId, microsoftGraphFieldValueSet)

Update the navigation property fields in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
**microsoftGraphFieldValueSet** | [**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md)| New navigation property values | 

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


## SharesUpdateListItem

> SharesUpdateListItem(ctx, sharedDriveItemId, microsoftGraphListItem)

Update the navigation property listItem in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**microsoftGraphListItem** | [**MicrosoftGraphListItem**](MicrosoftGraphListItem.md)| New navigation property values | 

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

