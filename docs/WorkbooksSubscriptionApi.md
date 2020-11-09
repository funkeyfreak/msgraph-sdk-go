# \WorkbooksSubscriptionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksCreateSubscriptions**](WorkbooksSubscriptionApi.md#WorkbooksCreateSubscriptions) | **Post** /workbooks/{driveItem-id}/subscriptions | Create new navigation property to subscriptions for workbooks
[**WorkbooksGetSubscriptions**](WorkbooksSubscriptionApi.md#WorkbooksGetSubscriptions) | **Get** /workbooks/{driveItem-id}/subscriptions/{subscription-id} | Get subscriptions from workbooks
[**WorkbooksListSubscriptions**](WorkbooksSubscriptionApi.md#WorkbooksListSubscriptions) | **Get** /workbooks/{driveItem-id}/subscriptions | Get subscriptions from workbooks
[**WorkbooksUpdateSubscriptions**](WorkbooksSubscriptionApi.md#WorkbooksUpdateSubscriptions) | **Patch** /workbooks/{driveItem-id}/subscriptions/{subscription-id} | Update the navigation property subscriptions in workbooks



## WorkbooksCreateSubscriptions

> MicrosoftGraphSubscription WorkbooksCreateSubscriptions(ctx, driveItemId, microsoftGraphSubscription)

Create new navigation property to subscriptions for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksGetSubscriptions

> MicrosoftGraphSubscription WorkbooksGetSubscriptions(ctx, driveItemId, subscriptionId, optional)

Get subscriptions from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**subscriptionId** | **string**| key: subscription-id of subscription | 
 **optional** | ***WorkbooksGetSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksGetSubscriptionsOpts struct


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


## WorkbooksListSubscriptions

> CollectionOfSubscription WorkbooksListSubscriptions(ctx, driveItemId, optional)

Get subscriptions from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksListSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListSubscriptionsOpts struct


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


## WorkbooksUpdateSubscriptions

> WorkbooksUpdateSubscriptions(ctx, driveItemId, subscriptionId, microsoftGraphSubscription)

Update the navigation property subscriptions in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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

