# \MeOfficeGraphInsightsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeGetInsights**](MeOfficeGraphInsightsApi.md#MeGetInsights) | **Get** /me/insights | Get insights from me
[**MeInsightsCreateShared**](MeOfficeGraphInsightsApi.md#MeInsightsCreateShared) | **Post** /me/insights/shared | Create new navigation property to shared for me
[**MeInsightsCreateTrending**](MeOfficeGraphInsightsApi.md#MeInsightsCreateTrending) | **Post** /me/insights/trending | Create new navigation property to trending for me
[**MeInsightsCreateUsed**](MeOfficeGraphInsightsApi.md#MeInsightsCreateUsed) | **Post** /me/insights/used | Create new navigation property to used for me
[**MeInsightsGetShared**](MeOfficeGraphInsightsApi.md#MeInsightsGetShared) | **Get** /me/insights/shared/{sharedInsight-id} | Get shared from me
[**MeInsightsGetTrending**](MeOfficeGraphInsightsApi.md#MeInsightsGetTrending) | **Get** /me/insights/trending/{trending-id} | Get trending from me
[**MeInsightsGetUsed**](MeOfficeGraphInsightsApi.md#MeInsightsGetUsed) | **Get** /me/insights/used/{usedInsight-id} | Get used from me
[**MeInsightsListShared**](MeOfficeGraphInsightsApi.md#MeInsightsListShared) | **Get** /me/insights/shared | Get shared from me
[**MeInsightsListTrending**](MeOfficeGraphInsightsApi.md#MeInsightsListTrending) | **Get** /me/insights/trending | Get trending from me
[**MeInsightsListUsed**](MeOfficeGraphInsightsApi.md#MeInsightsListUsed) | **Get** /me/insights/used | Get used from me
[**MeInsightsSharedGetLastSharedMethod**](MeOfficeGraphInsightsApi.md#MeInsightsSharedGetLastSharedMethod) | **Get** /me/insights/shared/{sharedInsight-id}/lastSharedMethod | Get lastSharedMethod from me
[**MeInsightsSharedGetResource**](MeOfficeGraphInsightsApi.md#MeInsightsSharedGetResource) | **Get** /me/insights/shared/{sharedInsight-id}/resource | Get resource from me
[**MeInsightsTrendingGetResource**](MeOfficeGraphInsightsApi.md#MeInsightsTrendingGetResource) | **Get** /me/insights/trending/{trending-id}/resource | Get resource from me
[**MeInsightsUpdateShared**](MeOfficeGraphInsightsApi.md#MeInsightsUpdateShared) | **Patch** /me/insights/shared/{sharedInsight-id} | Update the navigation property shared in me
[**MeInsightsUpdateTrending**](MeOfficeGraphInsightsApi.md#MeInsightsUpdateTrending) | **Patch** /me/insights/trending/{trending-id} | Update the navigation property trending in me
[**MeInsightsUpdateUsed**](MeOfficeGraphInsightsApi.md#MeInsightsUpdateUsed) | **Patch** /me/insights/used/{usedInsight-id} | Update the navigation property used in me
[**MeInsightsUsedGetResource**](MeOfficeGraphInsightsApi.md#MeInsightsUsedGetResource) | **Get** /me/insights/used/{usedInsight-id}/resource | Get resource from me
[**MeUpdateInsights**](MeOfficeGraphInsightsApi.md#MeUpdateInsights) | **Patch** /me/insights | Update the navigation property insights in me



## MeGetInsights

> MicrosoftGraphOfficeGraphInsights MeGetInsights(ctx, optional)

Get insights from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeGetInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetInsightsOpts struct


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


## MeInsightsCreateShared

> MicrosoftGraphSharedInsight MeInsightsCreateShared(ctx, microsoftGraphSharedInsight)

Create new navigation property to shared for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeInsightsCreateTrending

> MicrosoftGraphTrending MeInsightsCreateTrending(ctx, microsoftGraphTrending)

Create new navigation property to trending for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeInsightsCreateUsed

> MicrosoftGraphUsedInsight MeInsightsCreateUsed(ctx, microsoftGraphUsedInsight)

Create new navigation property to used for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeInsightsGetShared

> MicrosoftGraphSharedInsight MeInsightsGetShared(ctx, sharedInsightId, optional)

Get shared from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***MeInsightsGetSharedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsGetSharedOpts struct


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


## MeInsightsGetTrending

> MicrosoftGraphTrending MeInsightsGetTrending(ctx, trendingId, optional)

Get trending from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string**| key: trending-id of trending | 
 **optional** | ***MeInsightsGetTrendingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsGetTrendingOpts struct


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


## MeInsightsGetUsed

> MicrosoftGraphUsedInsight MeInsightsGetUsed(ctx, usedInsightId, optional)

Get used from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
 **optional** | ***MeInsightsGetUsedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsGetUsedOpts struct


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


## MeInsightsListShared

> CollectionOfSharedInsight MeInsightsListShared(ctx, optional)

Get shared from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeInsightsListSharedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsListSharedOpts struct


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


## MeInsightsListTrending

> CollectionOfTrending MeInsightsListTrending(ctx, optional)

Get trending from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeInsightsListTrendingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsListTrendingOpts struct


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


## MeInsightsListUsed

> CollectionOfUsedInsight MeInsightsListUsed(ctx, optional)

Get used from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeInsightsListUsedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsListUsedOpts struct


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


## MeInsightsSharedGetLastSharedMethod

> MicrosoftGraphEntity MeInsightsSharedGetLastSharedMethod(ctx, sharedInsightId, optional)

Get lastSharedMethod from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***MeInsightsSharedGetLastSharedMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsSharedGetLastSharedMethodOpts struct


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


## MeInsightsSharedGetResource

> MicrosoftGraphEntity MeInsightsSharedGetResource(ctx, sharedInsightId, optional)

Get resource from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***MeInsightsSharedGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsSharedGetResourceOpts struct


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


## MeInsightsTrendingGetResource

> MicrosoftGraphEntity MeInsightsTrendingGetResource(ctx, trendingId, optional)

Get resource from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string**| key: trending-id of trending | 
 **optional** | ***MeInsightsTrendingGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsTrendingGetResourceOpts struct


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


## MeInsightsUpdateShared

> MeInsightsUpdateShared(ctx, sharedInsightId, microsoftGraphSharedInsight)

Update the navigation property shared in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeInsightsUpdateTrending

> MeInsightsUpdateTrending(ctx, trendingId, microsoftGraphTrending)

Update the navigation property trending in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeInsightsUpdateUsed

> MeInsightsUpdateUsed(ctx, usedInsightId, microsoftGraphUsedInsight)

Update the navigation property used in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeInsightsUsedGetResource

> MicrosoftGraphEntity MeInsightsUsedGetResource(ctx, usedInsightId, optional)

Get resource from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
 **optional** | ***MeInsightsUsedGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsUsedGetResourceOpts struct


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


## MeUpdateInsights

> MeUpdateInsights(ctx, microsoftGraphOfficeGraphInsights)

Update the navigation property insights in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

