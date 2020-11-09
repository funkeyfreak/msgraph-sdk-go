# \PlannerPlannerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerPlannerGetPlanner**](PlannerPlannerApi.md#PlannerPlannerGetPlanner) | **Get** /planner | Get planner
[**PlannerPlannerUpdatePlanner**](PlannerPlannerApi.md#PlannerPlannerUpdatePlanner) | **Patch** /planner | Update planner



## PlannerPlannerGetPlanner

> MicrosoftGraphPlanner PlannerPlannerGetPlanner(ctx, optional)

Get planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlannerPlannerGetPlannerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerPlannerGetPlannerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlanner**](microsoft.graph.planner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerPlannerUpdatePlanner

> PlannerPlannerUpdatePlanner(ctx, microsoftGraphPlanner)

Update planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphPlanner** | [**MicrosoftGraphPlanner**](MicrosoftGraphPlanner.md)| New property values | 

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

