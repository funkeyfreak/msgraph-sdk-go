# \GroupsPlannerGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGetPlanner**](GroupsPlannerGroupApi.md#GroupsGetPlanner) | **Get** /groups/{group-id}/planner | Get planner from groups
[**GroupsPlannerGetPlans**](GroupsPlannerGroupApi.md#GroupsPlannerGetPlans) | **Get** /groups/{group-id}/planner/plans/{plannerPlan-id} | Get plans from groups
[**GroupsPlannerListPlans**](GroupsPlannerGroupApi.md#GroupsPlannerListPlans) | **Get** /groups/{group-id}/planner/plans | Get plans from groups
[**GroupsUpdatePlanner**](GroupsPlannerGroupApi.md#GroupsUpdatePlanner) | **Patch** /groups/{group-id}/planner | Update the navigation property planner in groups



## GroupsGetPlanner

> MicrosoftGraphPlannerGroup GroupsGetPlanner(ctx, groupId, optional)

Get planner from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsGetPlannerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetPlannerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerGroup**](microsoft.graph.plannerGroup.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsPlannerGetPlans

> MicrosoftGraphPlannerPlan GroupsPlannerGetPlans(ctx, groupId, plannerPlanId, optional)

Get plans from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
 **optional** | ***GroupsPlannerGetPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsPlannerGetPlansOpts struct


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


## GroupsPlannerListPlans

> CollectionOfPlannerPlan GroupsPlannerListPlans(ctx, groupId, optional)

Get plans from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsPlannerListPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsPlannerListPlansOpts struct


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


## GroupsUpdatePlanner

> GroupsUpdatePlanner(ctx, groupId, microsoftGraphPlannerGroup)

Update the navigation property planner in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphPlannerGroup** | [**MicrosoftGraphPlannerGroup**](MicrosoftGraphPlannerGroup.md)| New navigation property values | 

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

