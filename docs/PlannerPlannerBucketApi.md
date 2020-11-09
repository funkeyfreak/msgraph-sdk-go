# \PlannerPlannerBucketApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerBucketsGetTasks**](PlannerPlannerBucketApi.md#PlannerBucketsGetTasks) | **Get** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Get tasks from planner
[**PlannerBucketsListTasks**](PlannerPlannerBucketApi.md#PlannerBucketsListTasks) | **Get** /planner/buckets/{plannerBucket-id}/tasks | Get tasks from planner
[**PlannerCreateBuckets**](PlannerPlannerBucketApi.md#PlannerCreateBuckets) | **Post** /planner/buckets | Create new navigation property to buckets for planner
[**PlannerGetBuckets**](PlannerPlannerBucketApi.md#PlannerGetBuckets) | **Get** /planner/buckets/{plannerBucket-id} | Get buckets from planner
[**PlannerListBuckets**](PlannerPlannerBucketApi.md#PlannerListBuckets) | **Get** /planner/buckets | Get buckets from planner
[**PlannerUpdateBuckets**](PlannerPlannerBucketApi.md#PlannerUpdateBuckets) | **Patch** /planner/buckets/{plannerBucket-id} | Update the navigation property buckets in planner



## PlannerBucketsGetTasks

> MicrosoftGraphPlannerTask PlannerBucketsGetTasks(ctx, plannerBucketId, plannerTaskId, optional)

Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string**| key: plannerBucket-id of plannerBucket | 
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerBucketsGetTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerBucketsGetTasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerTask**](microsoft.graph.plannerTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerBucketsListTasks

> CollectionOfPlannerTask PlannerBucketsListTasks(ctx, plannerBucketId, optional)

Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string**| key: plannerBucket-id of plannerBucket | 
 **optional** | ***PlannerBucketsListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerBucketsListTasksOpts struct


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

[**CollectionOfPlannerTask**](Collection_of_plannerTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerCreateBuckets

> MicrosoftGraphPlannerBucket PlannerCreateBuckets(ctx, microsoftGraphPlannerBucket)

Create new navigation property to buckets for planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphPlannerBucket** | [**MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md)| New navigation property | 

### Return type

[**MicrosoftGraphPlannerBucket**](microsoft.graph.plannerBucket.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerGetBuckets

> MicrosoftGraphPlannerBucket PlannerGetBuckets(ctx, plannerBucketId, optional)

Get buckets from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string**| key: plannerBucket-id of plannerBucket | 
 **optional** | ***PlannerGetBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerGetBucketsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerBucket**](microsoft.graph.plannerBucket.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerListBuckets

> CollectionOfPlannerBucket PlannerListBuckets(ctx, optional)

Get buckets from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlannerListBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerListBucketsOpts struct


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

[**CollectionOfPlannerBucket**](Collection_of_plannerBucket.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerUpdateBuckets

> PlannerUpdateBuckets(ctx, plannerBucketId, microsoftGraphPlannerBucket)

Update the navigation property buckets in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string**| key: plannerBucket-id of plannerBucket | 
**microsoftGraphPlannerBucket** | [**MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md)| New navigation property values | 

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

