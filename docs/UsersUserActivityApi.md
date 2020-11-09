# \UsersUserActivityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersActivitiesCreateHistoryItems**](UsersUserActivityApi.md#UsersActivitiesCreateHistoryItems) | **Post** /users/{user-id}/activities/{userActivity-id}/historyItems | Create new navigation property to historyItems for users
[**UsersActivitiesGetHistoryItems**](UsersUserActivityApi.md#UsersActivitiesGetHistoryItems) | **Get** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id} | Get historyItems from users
[**UsersActivitiesHistoryItemsGetActivity**](UsersUserActivityApi.md#UsersActivitiesHistoryItemsGetActivity) | **Get** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id}/activity | Get activity from users
[**UsersActivitiesListHistoryItems**](UsersUserActivityApi.md#UsersActivitiesListHistoryItems) | **Get** /users/{user-id}/activities/{userActivity-id}/historyItems | Get historyItems from users
[**UsersActivitiesUpdateHistoryItems**](UsersUserActivityApi.md#UsersActivitiesUpdateHistoryItems) | **Patch** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id} | Update the navigation property historyItems in users
[**UsersCreateActivities**](UsersUserActivityApi.md#UsersCreateActivities) | **Post** /users/{user-id}/activities | Create new navigation property to activities for users
[**UsersGetActivities**](UsersUserActivityApi.md#UsersGetActivities) | **Get** /users/{user-id}/activities/{userActivity-id} | Get activities from users
[**UsersListActivities**](UsersUserActivityApi.md#UsersListActivities) | **Get** /users/{user-id}/activities | Get activities from users
[**UsersUpdateActivities**](UsersUserActivityApi.md#UsersUpdateActivities) | **Patch** /users/{user-id}/activities/{userActivity-id} | Update the navigation property activities in users



## UsersActivitiesCreateHistoryItems

> MicrosoftGraphActivityHistoryItem UsersActivitiesCreateHistoryItems(ctx, userId, userActivityId, microsoftGraphActivityHistoryItem)

Create new navigation property to historyItems for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**microsoftGraphActivityHistoryItem** | [**MicrosoftGraphActivityHistoryItem**](MicrosoftGraphActivityHistoryItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphActivityHistoryItem**](microsoft.graph.activityHistoryItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersActivitiesGetHistoryItems

> MicrosoftGraphActivityHistoryItem UsersActivitiesGetHistoryItems(ctx, userId, userActivityId, activityHistoryItemId, optional)

Get historyItems from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**activityHistoryItemId** | **string**| key: activityHistoryItem-id of activityHistoryItem | 
 **optional** | ***UsersActivitiesGetHistoryItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersActivitiesGetHistoryItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphActivityHistoryItem**](microsoft.graph.activityHistoryItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersActivitiesHistoryItemsGetActivity

> MicrosoftGraphUserActivity UsersActivitiesHistoryItemsGetActivity(ctx, userId, userActivityId, activityHistoryItemId, optional)

Get activity from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**activityHistoryItemId** | **string**| key: activityHistoryItem-id of activityHistoryItem | 
 **optional** | ***UsersActivitiesHistoryItemsGetActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersActivitiesHistoryItemsGetActivityOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUserActivity**](microsoft.graph.userActivity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersActivitiesListHistoryItems

> CollectionOfActivityHistoryItem UsersActivitiesListHistoryItems(ctx, userId, userActivityId, optional)

Get historyItems from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
 **optional** | ***UsersActivitiesListHistoryItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersActivitiesListHistoryItemsOpts struct


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

[**CollectionOfActivityHistoryItem**](Collection_of_activityHistoryItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersActivitiesUpdateHistoryItems

> UsersActivitiesUpdateHistoryItems(ctx, userId, userActivityId, activityHistoryItemId, microsoftGraphActivityHistoryItem)

Update the navigation property historyItems in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**activityHistoryItemId** | **string**| key: activityHistoryItem-id of activityHistoryItem | 
**microsoftGraphActivityHistoryItem** | [**MicrosoftGraphActivityHistoryItem**](MicrosoftGraphActivityHistoryItem.md)| New navigation property values | 

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


## UsersCreateActivities

> MicrosoftGraphUserActivity UsersCreateActivities(ctx, userId, microsoftGraphUserActivity)

Create new navigation property to activities for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphUserActivity** | [**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md)| New navigation property | 

### Return type

[**MicrosoftGraphUserActivity**](microsoft.graph.userActivity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetActivities

> MicrosoftGraphUserActivity UsersGetActivities(ctx, userId, userActivityId, optional)

Get activities from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
 **optional** | ***UsersGetActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetActivitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUserActivity**](microsoft.graph.userActivity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListActivities

> CollectionOfUserActivity UsersListActivities(ctx, userId, optional)

Get activities from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListActivitiesOpts struct


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

[**CollectionOfUserActivity**](Collection_of_userActivity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateActivities

> UsersUpdateActivities(ctx, userId, userActivityId, microsoftGraphUserActivity)

Update the navigation property activities in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**microsoftGraphUserActivity** | [**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md)| New navigation property values | 

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

