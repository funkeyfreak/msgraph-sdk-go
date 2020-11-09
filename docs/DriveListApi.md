# \DriveListApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveGetList**](DriveListApi.md#DriveGetList) | **Get** /drive/list | Get list from drive
[**DriveListContentTypesCreateColumnLinks**](DriveListApi.md#DriveListContentTypesCreateColumnLinks) | **Post** /drive/list/contentTypes/{contentType-id}/columnLinks | Create new navigation property to columnLinks for drive
[**DriveListContentTypesGetColumnLinks**](DriveListApi.md#DriveListContentTypesGetColumnLinks) | **Get** /drive/list/contentTypes/{contentType-id}/columnLinks/{columnLink-id} | Get columnLinks from drive
[**DriveListContentTypesListColumnLinks**](DriveListApi.md#DriveListContentTypesListColumnLinks) | **Get** /drive/list/contentTypes/{contentType-id}/columnLinks | Get columnLinks from drive
[**DriveListContentTypesUpdateColumnLinks**](DriveListApi.md#DriveListContentTypesUpdateColumnLinks) | **Patch** /drive/list/contentTypes/{contentType-id}/columnLinks/{columnLink-id} | Update the navigation property columnLinks in drive
[**DriveListCreateColumns**](DriveListApi.md#DriveListCreateColumns) | **Post** /drive/list/columns | Create new navigation property to columns for drive
[**DriveListCreateContentTypes**](DriveListApi.md#DriveListCreateContentTypes) | **Post** /drive/list/contentTypes | Create new navigation property to contentTypes for drive
[**DriveListCreateItems**](DriveListApi.md#DriveListCreateItems) | **Post** /drive/list/items | Create new navigation property to items for drive
[**DriveListCreateSubscriptions**](DriveListApi.md#DriveListCreateSubscriptions) | **Post** /drive/list/subscriptions | Create new navigation property to subscriptions for drive
[**DriveListGetColumns**](DriveListApi.md#DriveListGetColumns) | **Get** /drive/list/columns/{columnDefinition-id} | Get columns from drive
[**DriveListGetContentTypes**](DriveListApi.md#DriveListGetContentTypes) | **Get** /drive/list/contentTypes/{contentType-id} | Get contentTypes from drive
[**DriveListGetDrive**](DriveListApi.md#DriveListGetDrive) | **Get** /drive/list/drive | Get drive from drive
[**DriveListGetItems**](DriveListApi.md#DriveListGetItems) | **Get** /drive/list/items/{listItem-id} | Get items from drive
[**DriveListGetSubscriptions**](DriveListApi.md#DriveListGetSubscriptions) | **Get** /drive/list/subscriptions/{subscription-id} | Get subscriptions from drive
[**DriveListItemsCreateVersions**](DriveListApi.md#DriveListItemsCreateVersions) | **Post** /drive/list/items/{listItem-id}/versions | Create new navigation property to versions for drive
[**DriveListItemsGetAnalytics**](DriveListApi.md#DriveListItemsGetAnalytics) | **Get** /drive/list/items/{listItem-id}/analytics | Get analytics from drive
[**DriveListItemsGetDriveItem**](DriveListApi.md#DriveListItemsGetDriveItem) | **Get** /drive/list/items/{listItem-id}/driveItem | Get driveItem from drive
[**DriveListItemsGetFields**](DriveListApi.md#DriveListItemsGetFields) | **Get** /drive/list/items/{listItem-id}/fields | Get fields from drive
[**DriveListItemsGetVersions**](DriveListApi.md#DriveListItemsGetVersions) | **Get** /drive/list/items/{listItem-id}/versions/{listItemVersion-id} | Get versions from drive
[**DriveListItemsListVersions**](DriveListApi.md#DriveListItemsListVersions) | **Get** /drive/list/items/{listItem-id}/versions | Get versions from drive
[**DriveListItemsUpdateDriveItem**](DriveListApi.md#DriveListItemsUpdateDriveItem) | **Patch** /drive/list/items/{listItem-id}/driveItem | Update the navigation property driveItem in drive
[**DriveListItemsUpdateFields**](DriveListApi.md#DriveListItemsUpdateFields) | **Patch** /drive/list/items/{listItem-id}/fields | Update the navigation property fields in drive
[**DriveListItemsUpdateVersions**](DriveListApi.md#DriveListItemsUpdateVersions) | **Patch** /drive/list/items/{listItem-id}/versions/{listItemVersion-id} | Update the navigation property versions in drive
[**DriveListItemsVersionsGetFields**](DriveListApi.md#DriveListItemsVersionsGetFields) | **Get** /drive/list/items/{listItem-id}/versions/{listItemVersion-id}/fields | Get fields from drive
[**DriveListItemsVersionsUpdateFields**](DriveListApi.md#DriveListItemsVersionsUpdateFields) | **Patch** /drive/list/items/{listItem-id}/versions/{listItemVersion-id}/fields | Update the navigation property fields in drive
[**DriveListListColumns**](DriveListApi.md#DriveListListColumns) | **Get** /drive/list/columns | Get columns from drive
[**DriveListListContentTypes**](DriveListApi.md#DriveListListContentTypes) | **Get** /drive/list/contentTypes | Get contentTypes from drive
[**DriveListListItems**](DriveListApi.md#DriveListListItems) | **Get** /drive/list/items | Get items from drive
[**DriveListListSubscriptions**](DriveListApi.md#DriveListListSubscriptions) | **Get** /drive/list/subscriptions | Get subscriptions from drive
[**DriveListUpdateColumns**](DriveListApi.md#DriveListUpdateColumns) | **Patch** /drive/list/columns/{columnDefinition-id} | Update the navigation property columns in drive
[**DriveListUpdateContentTypes**](DriveListApi.md#DriveListUpdateContentTypes) | **Patch** /drive/list/contentTypes/{contentType-id} | Update the navigation property contentTypes in drive
[**DriveListUpdateDrive**](DriveListApi.md#DriveListUpdateDrive) | **Patch** /drive/list/drive | Update the navigation property drive in drive
[**DriveListUpdateItems**](DriveListApi.md#DriveListUpdateItems) | **Patch** /drive/list/items/{listItem-id} | Update the navigation property items in drive
[**DriveListUpdateSubscriptions**](DriveListApi.md#DriveListUpdateSubscriptions) | **Patch** /drive/list/subscriptions/{subscription-id} | Update the navigation property subscriptions in drive
[**DriveUpdateList**](DriveListApi.md#DriveUpdateList) | **Patch** /drive/list | Update the navigation property list in drive



## DriveGetList

> MicrosoftGraphList DriveGetList(ctx, optional)

Get list from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriveGetListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveGetListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphList**](microsoft.graph.list.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListContentTypesCreateColumnLinks

> MicrosoftGraphColumnLink DriveListContentTypesCreateColumnLinks(ctx, contentTypeId, microsoftGraphColumnLink)

Create new navigation property to columnLinks for drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
**microsoftGraphColumnLink** | [**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md)| New navigation property | 

### Return type

[**MicrosoftGraphColumnLink**](microsoft.graph.columnLink.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListContentTypesGetColumnLinks

> MicrosoftGraphColumnLink DriveListContentTypesGetColumnLinks(ctx, contentTypeId, columnLinkId, optional)

Get columnLinks from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
**columnLinkId** | **string**| key: columnLink-id of columnLink | 
 **optional** | ***DriveListContentTypesGetColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListContentTypesGetColumnLinksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphColumnLink**](microsoft.graph.columnLink.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListContentTypesListColumnLinks

> CollectionOfColumnLink DriveListContentTypesListColumnLinks(ctx, contentTypeId, optional)

Get columnLinks from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***DriveListContentTypesListColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListContentTypesListColumnLinksOpts struct


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

[**CollectionOfColumnLink**](Collection_of_columnLink.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListContentTypesUpdateColumnLinks

> DriveListContentTypesUpdateColumnLinks(ctx, contentTypeId, columnLinkId, microsoftGraphColumnLink)

Update the navigation property columnLinks in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
**columnLinkId** | **string**| key: columnLink-id of columnLink | 
**microsoftGraphColumnLink** | [**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md)| New navigation property values | 

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


## DriveListCreateColumns

> MicrosoftGraphColumnDefinition DriveListCreateColumns(ctx, microsoftGraphColumnDefinition)

Create new navigation property to columns for drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)| New navigation property | 

### Return type

[**MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListCreateContentTypes

> MicrosoftGraphContentType DriveListCreateContentTypes(ctx, microsoftGraphContentType)

Create new navigation property to contentTypes for drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphContentType** | [**MicrosoftGraphContentType**](MicrosoftGraphContentType.md)| New navigation property | 

### Return type

[**MicrosoftGraphContentType**](microsoft.graph.contentType.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListCreateItems

> MicrosoftGraphListItem DriveListCreateItems(ctx, microsoftGraphListItem)

Create new navigation property to items for drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphListItem** | [**MicrosoftGraphListItem**](MicrosoftGraphListItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphListItem**](microsoft.graph.listItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListCreateSubscriptions

> MicrosoftGraphSubscription DriveListCreateSubscriptions(ctx, microsoftGraphSubscription)

Create new navigation property to subscriptions for drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md)| New navigation property | 

### Return type

[**MicrosoftGraphSubscription**](microsoft.graph.subscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListGetColumns

> MicrosoftGraphColumnDefinition DriveListGetColumns(ctx, columnDefinitionId, optional)

Get columns from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnDefinitionId** | **string**| key: columnDefinition-id of columnDefinition | 
 **optional** | ***DriveListGetColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListGetColumnsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListGetContentTypes

> MicrosoftGraphContentType DriveListGetContentTypes(ctx, contentTypeId, optional)

Get contentTypes from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***DriveListGetContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListGetContentTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphContentType**](microsoft.graph.contentType.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListGetDrive

> MicrosoftGraphDrive DriveListGetDrive(ctx, optional)

Get drive from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriveListGetDriveOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListGetDriveOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDrive**](microsoft.graph.drive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListGetItems

> MicrosoftGraphListItem DriveListGetItems(ctx, listItemId, optional)

Get items from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***DriveListGetItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListGetItemsOpts struct


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


## DriveListGetSubscriptions

> MicrosoftGraphSubscription DriveListGetSubscriptions(ctx, subscriptionId, optional)

Get subscriptions from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| key: subscription-id of subscription | 
 **optional** | ***DriveListGetSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListGetSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSubscription**](microsoft.graph.subscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListItemsCreateVersions

> MicrosoftGraphListItemVersion DriveListItemsCreateVersions(ctx, listItemId, microsoftGraphListItemVersion)

Create new navigation property to versions for drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
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


## DriveListItemsGetAnalytics

> MicrosoftGraphItemAnalytics DriveListItemsGetAnalytics(ctx, listItemId, optional)

Get analytics from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***DriveListItemsGetAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListItemsGetAnalyticsOpts struct


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


## DriveListItemsGetDriveItem

> MicrosoftGraphDriveItem DriveListItemsGetDriveItem(ctx, listItemId, optional)

Get driveItem from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***DriveListItemsGetDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListItemsGetDriveItemOpts struct


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


## DriveListItemsGetFields

> MicrosoftGraphFieldValueSet DriveListItemsGetFields(ctx, listItemId, optional)

Get fields from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***DriveListItemsGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListItemsGetFieldsOpts struct


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


## DriveListItemsGetVersions

> MicrosoftGraphListItemVersion DriveListItemsGetVersions(ctx, listItemId, listItemVersionId, optional)

Get versions from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***DriveListItemsGetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListItemsGetVersionsOpts struct


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


## DriveListItemsListVersions

> CollectionOfListItemVersion DriveListItemsListVersions(ctx, listItemId, optional)

Get versions from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***DriveListItemsListVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListItemsListVersionsOpts struct


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


## DriveListItemsUpdateDriveItem

> DriveListItemsUpdateDriveItem(ctx, listItemId, microsoftGraphDriveItem)

Update the navigation property driveItem in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
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


## DriveListItemsUpdateFields

> DriveListItemsUpdateFields(ctx, listItemId, microsoftGraphFieldValueSet)

Update the navigation property fields in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
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


## DriveListItemsUpdateVersions

> DriveListItemsUpdateVersions(ctx, listItemId, listItemVersionId, microsoftGraphListItemVersion)

Update the navigation property versions in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
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


## DriveListItemsVersionsGetFields

> MicrosoftGraphFieldValueSet DriveListItemsVersionsGetFields(ctx, listItemId, listItemVersionId, optional)

Get fields from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***DriveListItemsVersionsGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListItemsVersionsGetFieldsOpts struct


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


## DriveListItemsVersionsUpdateFields

> DriveListItemsVersionsUpdateFields(ctx, listItemId, listItemVersionId, microsoftGraphFieldValueSet)

Update the navigation property fields in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
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


## DriveListListColumns

> CollectionOfColumnDefinition DriveListListColumns(ctx, optional)

Get columns from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriveListListColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListListColumnsOpts struct


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

[**CollectionOfColumnDefinition**](Collection_of_columnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListListContentTypes

> CollectionOfContentType DriveListListContentTypes(ctx, optional)

Get contentTypes from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriveListListContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListListContentTypesOpts struct


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

[**CollectionOfContentType**](Collection_of_contentType.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListListItems

> CollectionOfListItem DriveListListItems(ctx, optional)

Get items from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriveListListItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListListItemsOpts struct


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

[**CollectionOfListItem**](Collection_of_listItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListListSubscriptions

> CollectionOfSubscription DriveListListSubscriptions(ctx, optional)

Get subscriptions from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriveListListSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListListSubscriptionsOpts struct


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

[**CollectionOfSubscription**](Collection_of_subscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListUpdateColumns

> DriveListUpdateColumns(ctx, columnDefinitionId, microsoftGraphColumnDefinition)

Update the navigation property columns in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnDefinitionId** | **string**| key: columnDefinition-id of columnDefinition | 
**microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)| New navigation property values | 

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


## DriveListUpdateContentTypes

> DriveListUpdateContentTypes(ctx, contentTypeId, microsoftGraphContentType)

Update the navigation property contentTypes in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
**microsoftGraphContentType** | [**MicrosoftGraphContentType**](MicrosoftGraphContentType.md)| New navigation property values | 

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


## DriveListUpdateDrive

> DriveListUpdateDrive(ctx, microsoftGraphDrive)

Update the navigation property drive in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)| New navigation property values | 

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


## DriveListUpdateItems

> DriveListUpdateItems(ctx, listItemId, microsoftGraphListItem)

Update the navigation property items in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
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


## DriveListUpdateSubscriptions

> DriveListUpdateSubscriptions(ctx, subscriptionId, microsoftGraphSubscription)

Update the navigation property subscriptions in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| key: subscription-id of subscription | 
**microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md)| New navigation property values | 

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


## DriveUpdateList

> DriveUpdateList(ctx, microsoftGraphList)

Update the navigation property list in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphList** | [**MicrosoftGraphList**](MicrosoftGraphList.md)| New navigation property values | 

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

