# \PlannerPlannerPlanApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerCreatePlans**](PlannerPlannerPlanApi.md#PlannerCreatePlans) | **Post** /planner/plans | Create new navigation property to plans for planner
[**PlannerGetPlans**](PlannerPlannerPlanApi.md#PlannerGetPlans) | **Get** /planner/plans/{plannerPlan-id} | Get plans from planner
[**PlannerListPlans**](PlannerPlannerPlanApi.md#PlannerListPlans) | **Get** /planner/plans | Get plans from planner
[**PlannerPlansGetBuckets**](PlannerPlannerPlanApi.md#PlannerPlansGetBuckets) | **Get** /planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id} | Get buckets from planner
[**PlannerPlansGetDetails**](PlannerPlannerPlanApi.md#PlannerPlansGetDetails) | **Get** /planner/plans/{plannerPlan-id}/details | Get details from planner
[**PlannerPlansGetTasks**](PlannerPlannerPlanApi.md#PlannerPlansGetTasks) | **Get** /planner/plans/{plannerPlan-id}/tasks/{plannerTask-id} | Get tasks from planner
[**PlannerPlansListBuckets**](PlannerPlannerPlanApi.md#PlannerPlansListBuckets) | **Get** /planner/plans/{plannerPlan-id}/buckets | Get buckets from planner
[**PlannerPlansListTasks**](PlannerPlannerPlanApi.md#PlannerPlansListTasks) | **Get** /planner/plans/{plannerPlan-id}/tasks | Get tasks from planner
[**PlannerPlansUpdateDetails**](PlannerPlannerPlanApi.md#PlannerPlansUpdateDetails) | **Patch** /planner/plans/{plannerPlan-id}/details | Update the navigation property details in planner
[**PlannerUpdatePlans**](PlannerPlannerPlanApi.md#PlannerUpdatePlans) | **Patch** /planner/plans/{plannerPlan-id} | Update the navigation property plans in planner



## PlannerCreatePlans

> MicrosoftGraphPlannerPlan PlannerCreatePlans(ctx, microsoftGraphPlannerPlan)

Create new navigation property to plans for planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphPlannerPlan** | [**MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md)| New navigation property | 

### Return type

[**MicrosoftGraphPlannerPlan**](microsoft.graph.plannerPlan.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerGetPlans

> MicrosoftGraphPlannerPlan PlannerGetPlans(ctx, plannerPlanId, optional)

Get plans from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
 **optional** | ***PlannerGetPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerGetPlansOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerPlan**](microsoft.graph.plannerPlan.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerListPlans

> CollectionOfPlannerPlan PlannerListPlans(ctx, optional)

Get plans from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlannerListPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerListPlansOpts struct


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

[**CollectionOfPlannerPlan**](Collection_of_plannerPlan.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerPlansGetBuckets

> MicrosoftGraphPlannerBucket PlannerPlansGetBuckets(ctx, plannerPlanId, plannerBucketId, optional)

Get buckets from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
**plannerBucketId** | **string**| key: plannerBucket-id of plannerBucket | 
 **optional** | ***PlannerPlansGetBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerPlansGetBucketsOpts struct


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


## PlannerPlansGetDetails

> MicrosoftGraphPlannerPlanDetails PlannerPlansGetDetails(ctx, plannerPlanId, optional)

Get details from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
 **optional** | ***PlannerPlansGetDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerPlansGetDetailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerPlanDetails**](microsoft.graph.plannerPlanDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerPlansGetTasks

> MicrosoftGraphPlannerTask PlannerPlansGetTasks(ctx, plannerPlanId, plannerTaskId, optional)

Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerPlansGetTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerPlansGetTasksOpts struct


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


## PlannerPlansListBuckets

> CollectionOfPlannerBucket PlannerPlansListBuckets(ctx, plannerPlanId, optional)

Get buckets from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
 **optional** | ***PlannerPlansListBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerPlansListBucketsOpts struct


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


## PlannerPlansListTasks

> CollectionOfPlannerTask PlannerPlansListTasks(ctx, plannerPlanId, optional)

Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
 **optional** | ***PlannerPlansListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerPlansListTasksOpts struct


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


## PlannerPlansUpdateDetails

> PlannerPlansUpdateDetails(ctx, plannerPlanId, microsoftGraphPlannerPlanDetails)

Update the navigation property details in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
**microsoftGraphPlannerPlanDetails** | [**MicrosoftGraphPlannerPlanDetails**](MicrosoftGraphPlannerPlanDetails.md)| New navigation property values | 

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


## PlannerUpdatePlans

> PlannerUpdatePlans(ctx, plannerPlanId, microsoftGraphPlannerPlan)

Update the navigation property plans in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
**microsoftGraphPlannerPlan** | [**MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md)| New navigation property values | 

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

