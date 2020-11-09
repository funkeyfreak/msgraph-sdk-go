# \UsersInferenceClassificationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetInferenceClassification**](UsersInferenceClassificationApi.md#UsersGetInferenceClassification) | **Get** /users/{user-id}/inferenceClassification | Get inferenceClassification from users
[**UsersInferenceClassificationCreateOverrides**](UsersInferenceClassificationApi.md#UsersInferenceClassificationCreateOverrides) | **Post** /users/{user-id}/inferenceClassification/overrides | Create new navigation property to overrides for users
[**UsersInferenceClassificationGetOverrides**](UsersInferenceClassificationApi.md#UsersInferenceClassificationGetOverrides) | **Get** /users/{user-id}/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Get overrides from users
[**UsersInferenceClassificationListOverrides**](UsersInferenceClassificationApi.md#UsersInferenceClassificationListOverrides) | **Get** /users/{user-id}/inferenceClassification/overrides | Get overrides from users
[**UsersInferenceClassificationUpdateOverrides**](UsersInferenceClassificationApi.md#UsersInferenceClassificationUpdateOverrides) | **Patch** /users/{user-id}/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Update the navigation property overrides in users
[**UsersUpdateInferenceClassification**](UsersInferenceClassificationApi.md#UsersUpdateInferenceClassification) | **Patch** /users/{user-id}/inferenceClassification | Update the navigation property inferenceClassification in users



## UsersGetInferenceClassification

> MicrosoftGraphInferenceClassification UsersGetInferenceClassification(ctx, userId, optional)

Get inferenceClassification from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetInferenceClassificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetInferenceClassificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphInferenceClassification**](microsoft.graph.inferenceClassification.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInferenceClassificationCreateOverrides

> MicrosoftGraphInferenceClassificationOverride UsersInferenceClassificationCreateOverrides(ctx, userId, microsoftGraphInferenceClassificationOverride)

Create new navigation property to overrides for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphInferenceClassificationOverride** | [**MicrosoftGraphInferenceClassificationOverride**](MicrosoftGraphInferenceClassificationOverride.md)| New navigation property | 

### Return type

[**MicrosoftGraphInferenceClassificationOverride**](microsoft.graph.inferenceClassificationOverride.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInferenceClassificationGetOverrides

> MicrosoftGraphInferenceClassificationOverride UsersInferenceClassificationGetOverrides(ctx, userId, inferenceClassificationOverrideId, optional)

Get overrides from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**inferenceClassificationOverrideId** | **string**| key: inferenceClassificationOverride-id of inferenceClassificationOverride | 
 **optional** | ***UsersInferenceClassificationGetOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInferenceClassificationGetOverridesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphInferenceClassificationOverride**](microsoft.graph.inferenceClassificationOverride.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInferenceClassificationListOverrides

> CollectionOfInferenceClassificationOverride UsersInferenceClassificationListOverrides(ctx, userId, optional)

Get overrides from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersInferenceClassificationListOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInferenceClassificationListOverridesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfInferenceClassificationOverride**](Collection_of_inferenceClassificationOverride.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersInferenceClassificationUpdateOverrides

> UsersInferenceClassificationUpdateOverrides(ctx, userId, inferenceClassificationOverrideId, microsoftGraphInferenceClassificationOverride)

Update the navigation property overrides in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**inferenceClassificationOverrideId** | **string**| key: inferenceClassificationOverride-id of inferenceClassificationOverride | 
**microsoftGraphInferenceClassificationOverride** | [**MicrosoftGraphInferenceClassificationOverride**](MicrosoftGraphInferenceClassificationOverride.md)| New navigation property values | 

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


## UsersUpdateInferenceClassification

> UsersUpdateInferenceClassification(ctx, userId, microsoftGraphInferenceClassification)

Update the navigation property inferenceClassification in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphInferenceClassification** | [**MicrosoftGraphInferenceClassification**](MicrosoftGraphInferenceClassification.md)| New navigation property values | 

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

