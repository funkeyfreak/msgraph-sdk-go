# \UsersOfficeGraphInsightsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetInsights**](UsersOfficeGraphInsightsApi.md#UsersGetInsights) | **Get** /users/{user-id}/insights | Get insights from users
[**UsersInsightsCreateShared**](UsersOfficeGraphInsightsApi.md#UsersInsightsCreateShared) | **Post** /users/{user-id}/insights/shared | Create new navigation property to shared for users
[**UsersInsightsCreateTrending**](UsersOfficeGraphInsightsApi.md#UsersInsightsCreateTrending) | **Post** /users/{user-id}/insights/trending | Create new navigation property to trending for users
[**UsersInsightsCreateUsed**](UsersOfficeGraphInsightsApi.md#UsersInsightsCreateUsed) | **Post** /users/{user-id}/insights/used | Create new navigation property to used for users
[**UsersInsightsGetShared**](UsersOfficeGraphInsightsApi.md#UsersInsightsGetShared) | **Get** /users/{user-id}/insights/shared/{sharedInsight-id} | Get shared from users
[**UsersInsightsGetTrending**](UsersOfficeGraphInsightsApi.md#UsersInsightsGetTrending) | **Get** /users/{user-id}/insights/trending/{trending-id} | Get trending from users
[**UsersInsightsGetUsed**](UsersOfficeGraphInsightsApi.md#UsersInsightsGetUsed) | **Get** /users/{user-id}/insights/used/{usedInsight-id} | Get used from users
[**UsersInsightsListShared**](UsersOfficeGraphInsightsApi.md#UsersInsightsListShared) | **Get** /users/{user-id}/insights/shared | Get shared from users
[**UsersInsightsListTrending**](UsersOfficeGraphInsightsApi.md#UsersInsightsListTrending) | **Get** /users/{user-id}/insights/trending | Get trending from users
[**UsersInsightsListUsed**](UsersOfficeGraphInsightsApi.md#UsersInsightsListUsed) | **Get** /users/{user-id}/insights/used | Get used from users
[**UsersInsightsSharedGetLastSharedMethod**](UsersOfficeGraphInsightsApi.md#UsersInsightsSharedGetLastSharedMethod) | **Get** /users/{user-id}/insights/shared/{sharedInsight-id}/lastSharedMethod | Get lastSharedMethod from users
[**UsersInsightsSharedGetResource**](UsersOfficeGraphInsightsApi.md#UsersInsightsSharedGetResource) | **Get** /users/{user-id}/insights/shared/{sharedInsight-id}/resource | Get resource from users
[**UsersInsightsTrendingGetResource**](UsersOfficeGraphInsightsApi.md#UsersInsightsTrendingGetResource) | **Get** /users/{user-id}/insights/trending/{trending-id}/resource | Get resource from users
[**UsersInsightsUpdateShared**](UsersOfficeGraphInsightsApi.md#UsersInsightsUpdateShared) | **Patch** /users/{user-id}/insights/shared/{sharedInsight-id} | Update the navigation property shared in users
[**UsersInsightsUpdateTrending**](UsersOfficeGraphInsightsApi.md#UsersInsightsUpdateTrending) | **Patch** /users/{user-id}/insights/trending/{trending-id} | Update the navigation property trending in users
[**UsersInsightsUpdateUsed**](UsersOfficeGraphInsightsApi.md#UsersInsightsUpdateUsed) | **Patch** /users/{user-id}/insights/used/{usedInsight-id} | Update the navigation property used in users
[**UsersInsightsUsedGetResource**](UsersOfficeGraphInsightsApi.md#UsersInsightsUsedGetResource) | **Get** /users/{user-id}/insights/used/{usedInsight-id}/resource | Get resource from users
[**UsersUpdateInsights**](UsersOfficeGraphInsightsApi.md#UsersUpdateInsights) | **Patch** /users/{user-id}/insights | Update the navigation property insights in users



## UsersGetInsights

> MicrosoftGraphOfficeGraphInsights UsersGetInsights(ctx, userId, optional)

Get insights from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetInsightsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOfficeGraphInsights**](microsoft.graph.officeGraphInsights.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsCreateShared

> MicrosoftGraphSharedInsight UsersInsightsCreateShared(ctx, userId, microsoftGraphSharedInsight)

Create new navigation property to shared for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphSharedInsight** | [**MicrosoftGraphSharedInsight**](MicrosoftGraphSharedInsight.md)| New navigation property | 

### Return type

[**MicrosoftGraphSharedInsight**](microsoft.graph.sharedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsCreateTrending

> MicrosoftGraphTrending UsersInsightsCreateTrending(ctx, userId, microsoftGraphTrending)

Create new navigation property to trending for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphTrending** | [**MicrosoftGraphTrending**](MicrosoftGraphTrending.md)| New navigation property | 

### Return type

[**MicrosoftGraphTrending**](microsoft.graph.trending.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsCreateUsed

> MicrosoftGraphUsedInsight UsersInsightsCreateUsed(ctx, userId, microsoftGraphUsedInsight)

Create new navigation property to used for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphUsedInsight** | [**MicrosoftGraphUsedInsight**](MicrosoftGraphUsedInsight.md)| New navigation property | 

### Return type

[**MicrosoftGraphUsedInsight**](microsoft.graph.usedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsGetShared

> MicrosoftGraphSharedInsight UsersInsightsGetShared(ctx, userId, sharedInsightId, optional)

Get shared from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***UsersInsightsGetSharedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsGetSharedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSharedInsight**](microsoft.graph.sharedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsGetTrending

> MicrosoftGraphTrending UsersInsightsGetTrending(ctx, userId, trendingId, optional)

Get trending from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**trendingId** | **string**| key: trending-id of trending | 
 **optional** | ***UsersInsightsGetTrendingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsGetTrendingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTrending**](microsoft.graph.trending.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsGetUsed

> MicrosoftGraphUsedInsight UsersInsightsGetUsed(ctx, userId, usedInsightId, optional)

Get used from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
 **optional** | ***UsersInsightsGetUsedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsGetUsedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUsedInsight**](microsoft.graph.usedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsListShared

> CollectionOfSharedInsight UsersInsightsListShared(ctx, userId, optional)

Get shared from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersInsightsListSharedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsListSharedOpts struct


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

[**CollectionOfSharedInsight**](Collection_of_sharedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsListTrending

> CollectionOfTrending UsersInsightsListTrending(ctx, userId, optional)

Get trending from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersInsightsListTrendingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsListTrendingOpts struct


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

[**CollectionOfTrending**](Collection_of_trending.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsListUsed

> CollectionOfUsedInsight UsersInsightsListUsed(ctx, userId, optional)

Get used from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersInsightsListUsedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsListUsedOpts struct


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

[**CollectionOfUsedInsight**](Collection_of_usedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsSharedGetLastSharedMethod

> MicrosoftGraphEntity UsersInsightsSharedGetLastSharedMethod(ctx, userId, sharedInsightId, optional)

Get lastSharedMethod from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***UsersInsightsSharedGetLastSharedMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsSharedGetLastSharedMethodOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEntity**](microsoft.graph.entity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsSharedGetResource

> MicrosoftGraphEntity UsersInsightsSharedGetResource(ctx, userId, sharedInsightId, optional)

Get resource from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***UsersInsightsSharedGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsSharedGetResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEntity**](microsoft.graph.entity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsTrendingGetResource

> MicrosoftGraphEntity UsersInsightsTrendingGetResource(ctx, userId, trendingId, optional)

Get resource from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**trendingId** | **string**| key: trending-id of trending | 
 **optional** | ***UsersInsightsTrendingGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsTrendingGetResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEntity**](microsoft.graph.entity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInsightsUpdateShared

> UsersInsightsUpdateShared(ctx, userId, sharedInsightId, microsoftGraphSharedInsight)

Update the navigation property shared in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
**microsoftGraphSharedInsight** | [**MicrosoftGraphSharedInsight**](MicrosoftGraphSharedInsight.md)| New navigation property values | 

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


## UsersInsightsUpdateTrending

> UsersInsightsUpdateTrending(ctx, userId, trendingId, microsoftGraphTrending)

Update the navigation property trending in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**trendingId** | **string**| key: trending-id of trending | 
**microsoftGraphTrending** | [**MicrosoftGraphTrending**](MicrosoftGraphTrending.md)| New navigation property values | 

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


## UsersInsightsUpdateUsed

> UsersInsightsUpdateUsed(ctx, userId, usedInsightId, microsoftGraphUsedInsight)

Update the navigation property used in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
**microsoftGraphUsedInsight** | [**MicrosoftGraphUsedInsight**](MicrosoftGraphUsedInsight.md)| New navigation property values | 

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


## UsersInsightsUsedGetResource

> MicrosoftGraphEntity UsersInsightsUsedGetResource(ctx, userId, usedInsightId, optional)

Get resource from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
 **optional** | ***UsersInsightsUsedGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsUsedGetResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEntity**](microsoft.graph.entity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateInsights

> UsersUpdateInsights(ctx, userId, microsoftGraphOfficeGraphInsights)

Update the navigation property insights in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphOfficeGraphInsights** | [**MicrosoftGraphOfficeGraphInsights**](MicrosoftGraphOfficeGraphInsights.md)| New navigation property values | 

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

