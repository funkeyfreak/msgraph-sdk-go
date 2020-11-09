# \PlannerPlannerTaskApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerCreateTasks**](PlannerPlannerTaskApi.md#PlannerCreateTasks) | **Post** /planner/tasks | Create new navigation property to tasks for planner
[**PlannerGetTasks**](PlannerPlannerTaskApi.md#PlannerGetTasks) | **Get** /planner/tasks/{plannerTask-id} | Get tasks from planner
[**PlannerListTasks**](PlannerPlannerTaskApi.md#PlannerListTasks) | **Get** /planner/tasks | Get tasks from planner
[**PlannerTasksGetAssignedToTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksGetAssignedToTaskBoardFormat) | **Get** /planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from planner
[**PlannerTasksGetBucketTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksGetBucketTaskBoardFormat) | **Get** /planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Get bucketTaskBoardFormat from planner
[**PlannerTasksGetDetails**](PlannerPlannerTaskApi.md#PlannerTasksGetDetails) | **Get** /planner/tasks/{plannerTask-id}/details | Get details from planner
[**PlannerTasksGetProgressTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksGetProgressTaskBoardFormat) | **Get** /planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Get progressTaskBoardFormat from planner
[**PlannerTasksUpdateAssignedToTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksUpdateAssignedToTaskBoardFormat) | **Patch** /planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in planner
[**PlannerTasksUpdateBucketTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksUpdateBucketTaskBoardFormat) | **Patch** /planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in planner
[**PlannerTasksUpdateDetails**](PlannerPlannerTaskApi.md#PlannerTasksUpdateDetails) | **Patch** /planner/tasks/{plannerTask-id}/details | Update the navigation property details in planner
[**PlannerTasksUpdateProgressTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksUpdateProgressTaskBoardFormat) | **Patch** /planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in planner
[**PlannerUpdateTasks**](PlannerPlannerTaskApi.md#PlannerUpdateTasks) | **Patch** /planner/tasks/{plannerTask-id} | Update the navigation property tasks in planner



## PlannerCreateTasks

> MicrosoftGraphPlannerTask PlannerCreateTasks(ctx, microsoftGraphPlannerTask)

Create new navigation property to tasks for planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphPlannerTask** | [**MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md)| New navigation property | 

### Return type

[**MicrosoftGraphPlannerTask**](microsoft.graph.plannerTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerGetTasks

> MicrosoftGraphPlannerTask PlannerGetTasks(ctx, plannerTaskId, optional)

Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerGetTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerGetTasksOpts struct


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


## PlannerListTasks

> CollectionOfPlannerTask PlannerListTasks(ctx, optional)

Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlannerListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerListTasksOpts struct


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


## PlannerTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat PlannerTasksGetAssignedToTaskBoardFormat(ctx, plannerTaskId, optional)

Get assignedToTaskBoardFormat from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerTasksGetAssignedToTaskBoardFormatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerTasksGetAssignedToTaskBoardFormatOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat**](microsoft.graph.plannerAssignedToTaskBoardTaskFormat.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat PlannerTasksGetBucketTaskBoardFormat(ctx, plannerTaskId, optional)

Get bucketTaskBoardFormat from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerTasksGetBucketTaskBoardFormatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerTasksGetBucketTaskBoardFormatOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerBucketTaskBoardTaskFormat**](microsoft.graph.plannerBucketTaskBoardTaskFormat.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksGetDetails

> MicrosoftGraphPlannerTaskDetails PlannerTasksGetDetails(ctx, plannerTaskId, optional)

Get details from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerTasksGetDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerTasksGetDetailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerTaskDetails**](microsoft.graph.plannerTaskDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat PlannerTasksGetProgressTaskBoardFormat(ctx, plannerTaskId, optional)

Get progressTaskBoardFormat from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerTasksGetProgressTaskBoardFormatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerTasksGetProgressTaskBoardFormatOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerProgressTaskBoardTaskFormat**](microsoft.graph.plannerProgressTaskBoardTaskFormat.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksUpdateAssignedToTaskBoardFormat

> PlannerTasksUpdateAssignedToTaskBoardFormat(ctx, plannerTaskId, microsoftGraphPlannerAssignedToTaskBoardTaskFormat)

Update the navigation property assignedToTaskBoardFormat in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerAssignedToTaskBoardTaskFormat** | [**MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat**](MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat.md)| New navigation property values | 

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


## PlannerTasksUpdateBucketTaskBoardFormat

> PlannerTasksUpdateBucketTaskBoardFormat(ctx, plannerTaskId, microsoftGraphPlannerBucketTaskBoardTaskFormat)

Update the navigation property bucketTaskBoardFormat in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerBucketTaskBoardTaskFormat** | [**MicrosoftGraphPlannerBucketTaskBoardTaskFormat**](MicrosoftGraphPlannerBucketTaskBoardTaskFormat.md)| New navigation property values | 

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


## PlannerTasksUpdateDetails

> PlannerTasksUpdateDetails(ctx, plannerTaskId, microsoftGraphPlannerTaskDetails)

Update the navigation property details in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerTaskDetails** | [**MicrosoftGraphPlannerTaskDetails**](MicrosoftGraphPlannerTaskDetails.md)| New navigation property values | 

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


## PlannerTasksUpdateProgressTaskBoardFormat

> PlannerTasksUpdateProgressTaskBoardFormat(ctx, plannerTaskId, microsoftGraphPlannerProgressTaskBoardTaskFormat)

Update the navigation property progressTaskBoardFormat in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerProgressTaskBoardTaskFormat** | [**MicrosoftGraphPlannerProgressTaskBoardTaskFormat**](MicrosoftGraphPlannerProgressTaskBoardTaskFormat.md)| New navigation property values | 

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


## PlannerUpdateTasks

> PlannerUpdateTasks(ctx, plannerTaskId, microsoftGraphPlannerTask)

Update the navigation property tasks in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerTask** | [**MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md)| New navigation property values | 

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

