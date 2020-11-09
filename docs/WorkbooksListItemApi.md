# \WorkbooksListItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksGetListItem**](WorkbooksListItemApi.md#WorkbooksGetListItem) | **Get** /workbooks/{driveItem-id}/listItem | Get listItem from workbooks
[**WorkbooksListItemCreateVersions**](WorkbooksListItemApi.md#WorkbooksListItemCreateVersions) | **Post** /workbooks/{driveItem-id}/listItem/versions | Create new navigation property to versions for workbooks
[**WorkbooksListItemGetAnalytics**](WorkbooksListItemApi.md#WorkbooksListItemGetAnalytics) | **Get** /workbooks/{driveItem-id}/listItem/analytics | Get analytics from workbooks
[**WorkbooksListItemGetDriveItem**](WorkbooksListItemApi.md#WorkbooksListItemGetDriveItem) | **Get** /workbooks/{driveItem-id}/listItem/driveItem | Get driveItem from workbooks
[**WorkbooksListItemGetFields**](WorkbooksListItemApi.md#WorkbooksListItemGetFields) | **Get** /workbooks/{driveItem-id}/listItem/fields | Get fields from workbooks
[**WorkbooksListItemGetVersions**](WorkbooksListItemApi.md#WorkbooksListItemGetVersions) | **Get** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id} | Get versions from workbooks
[**WorkbooksListItemListVersions**](WorkbooksListItemApi.md#WorkbooksListItemListVersions) | **Get** /workbooks/{driveItem-id}/listItem/versions | Get versions from workbooks
[**WorkbooksListItemUpdateDriveItem**](WorkbooksListItemApi.md#WorkbooksListItemUpdateDriveItem) | **Patch** /workbooks/{driveItem-id}/listItem/driveItem | Update the navigation property driveItem in workbooks
[**WorkbooksListItemUpdateFields**](WorkbooksListItemApi.md#WorkbooksListItemUpdateFields) | **Patch** /workbooks/{driveItem-id}/listItem/fields | Update the navigation property fields in workbooks
[**WorkbooksListItemUpdateVersions**](WorkbooksListItemApi.md#WorkbooksListItemUpdateVersions) | **Patch** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id} | Update the navigation property versions in workbooks
[**WorkbooksListItemVersionsGetFields**](WorkbooksListItemApi.md#WorkbooksListItemVersionsGetFields) | **Get** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id}/fields | Get fields from workbooks
[**WorkbooksListItemVersionsUpdateFields**](WorkbooksListItemApi.md#WorkbooksListItemVersionsUpdateFields) | **Patch** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id}/fields | Update the navigation property fields in workbooks
[**WorkbooksUpdateListItem**](WorkbooksListItemApi.md#WorkbooksUpdateListItem) | **Patch** /workbooks/{driveItem-id}/listItem | Update the navigation property listItem in workbooks



## WorkbooksGetListItem

> MicrosoftGraphListItem WorkbooksGetListItem(ctx, driveItemId, optional)

Get listItem from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksGetListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksGetListItemOpts struct


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


## WorkbooksListItemCreateVersions

> MicrosoftGraphListItemVersion WorkbooksListItemCreateVersions(ctx, driveItemId, microsoftGraphListItemVersion)

Create new navigation property to versions for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksListItemGetAnalytics

> MicrosoftGraphItemAnalytics WorkbooksListItemGetAnalytics(ctx, driveItemId, optional)

Get analytics from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksListItemGetAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListItemGetAnalyticsOpts struct


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


## WorkbooksListItemGetDriveItem

> MicrosoftGraphDriveItem WorkbooksListItemGetDriveItem(ctx, driveItemId, optional)

Get driveItem from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksListItemGetDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListItemGetDriveItemOpts struct


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


## WorkbooksListItemGetFields

> MicrosoftGraphFieldValueSet WorkbooksListItemGetFields(ctx, driveItemId, optional)

Get fields from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksListItemGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListItemGetFieldsOpts struct


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


## WorkbooksListItemGetVersions

> MicrosoftGraphListItemVersion WorkbooksListItemGetVersions(ctx, driveItemId, listItemVersionId, optional)

Get versions from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***WorkbooksListItemGetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListItemGetVersionsOpts struct


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


## WorkbooksListItemListVersions

> CollectionOfListItemVersion WorkbooksListItemListVersions(ctx, driveItemId, optional)

Get versions from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksListItemListVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListItemListVersionsOpts struct


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


## WorkbooksListItemUpdateDriveItem

> WorkbooksListItemUpdateDriveItem(ctx, driveItemId, microsoftGraphDriveItem)

Update the navigation property driveItem in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksListItemUpdateFields

> WorkbooksListItemUpdateFields(ctx, driveItemId, microsoftGraphFieldValueSet)

Update the navigation property fields in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksListItemUpdateVersions

> WorkbooksListItemUpdateVersions(ctx, driveItemId, listItemVersionId, microsoftGraphListItemVersion)

Update the navigation property versions in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksListItemVersionsGetFields

> MicrosoftGraphFieldValueSet WorkbooksListItemVersionsGetFields(ctx, driveItemId, listItemVersionId, optional)

Get fields from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***WorkbooksListItemVersionsGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListItemVersionsGetFieldsOpts struct


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


## WorkbooksListItemVersionsUpdateFields

> WorkbooksListItemVersionsUpdateFields(ctx, driveItemId, listItemVersionId, microsoftGraphFieldValueSet)

Update the navigation property fields in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksUpdateListItem

> WorkbooksUpdateListItem(ctx, driveItemId, microsoftGraphListItem)

Update the navigation property listItem in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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

