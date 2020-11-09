# \SubscriptionsSubscriptionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsSubscriptionCreateSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionCreateSubscription) | **Post** /subscriptions | Add new entity to subscriptions
[**SubscriptionsSubscriptionDeleteSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionDeleteSubscription) | **Delete** /subscriptions/{subscription-id} | Delete entity from subscriptions
[**SubscriptionsSubscriptionGetSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionGetSubscription) | **Get** /subscriptions/{subscription-id} | Get entity from subscriptions by key
[**SubscriptionsSubscriptionListSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionListSubscription) | **Get** /subscriptions | Get entities from subscriptions
[**SubscriptionsSubscriptionUpdateSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionUpdateSubscription) | **Patch** /subscriptions/{subscription-id} | Update entity in subscriptions



## SubscriptionsSubscriptionCreateSubscription

> MicrosoftGraphSubscription SubscriptionsSubscriptionCreateSubscription(ctx, microsoftGraphSubscription)

Add new entity to subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md)| New entity | 

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


## SubscriptionsSubscriptionDeleteSubscription

> SubscriptionsSubscriptionDeleteSubscription(ctx, subscriptionId, optional)

Delete entity from subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| key: subscription-id of subscription | 
 **optional** | ***SubscriptionsSubscriptionDeleteSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscriptionsSubscriptionDeleteSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| ETag | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionGetSubscription

> MicrosoftGraphSubscription SubscriptionsSubscriptionGetSubscription(ctx, subscriptionId, optional)

Get entity from subscriptions by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| key: subscription-id of subscription | 
 **optional** | ***SubscriptionsSubscriptionGetSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscriptionsSubscriptionGetSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## SubscriptionsSubscriptionListSubscription

> CollectionOfSubscription SubscriptionsSubscriptionListSubscription(ctx, optional)

Get entities from subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriptionsSubscriptionListSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscriptionsSubscriptionListSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Search items by search phrases | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## SubscriptionsSubscriptionUpdateSubscription

> SubscriptionsSubscriptionUpdateSubscription(ctx, subscriptionId, microsoftGraphSubscription)

Update entity in subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| key: subscription-id of subscription | 
**microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md)| New property values | 

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

