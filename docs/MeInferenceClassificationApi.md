# \MeInferenceClassificationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeGetInferenceClassification**](MeInferenceClassificationApi.md#MeGetInferenceClassification) | **Get** /me/inferenceClassification | Get inferenceClassification from me
[**MeInferenceClassificationCreateOverrides**](MeInferenceClassificationApi.md#MeInferenceClassificationCreateOverrides) | **Post** /me/inferenceClassification/overrides | Create new navigation property to overrides for me
[**MeInferenceClassificationGetOverrides**](MeInferenceClassificationApi.md#MeInferenceClassificationGetOverrides) | **Get** /me/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Get overrides from me
[**MeInferenceClassificationListOverrides**](MeInferenceClassificationApi.md#MeInferenceClassificationListOverrides) | **Get** /me/inferenceClassification/overrides | Get overrides from me
[**MeInferenceClassificationUpdateOverrides**](MeInferenceClassificationApi.md#MeInferenceClassificationUpdateOverrides) | **Patch** /me/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Update the navigation property overrides in me
[**MeUpdateInferenceClassification**](MeInferenceClassificationApi.md#MeUpdateInferenceClassification) | **Patch** /me/inferenceClassification | Update the navigation property inferenceClassification in me



## MeGetInferenceClassification

> MicrosoftGraphInferenceClassification MeGetInferenceClassification(ctx, optional)

Get inferenceClassification from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeGetInferenceClassificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetInferenceClassificationOpts struct


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


## MeInferenceClassificationCreateOverrides

> MicrosoftGraphInferenceClassificationOverride MeInferenceClassificationCreateOverrides(ctx, microsoftGraphInferenceClassificationOverride)

Create new navigation property to overrides for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeInferenceClassificationGetOverrides

> MicrosoftGraphInferenceClassificationOverride MeInferenceClassificationGetOverrides(ctx, inferenceClassificationOverrideId, optional)

Get overrides from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceClassificationOverrideId** | **string**| key: inferenceClassificationOverride-id of inferenceClassificationOverride | 
 **optional** | ***MeInferenceClassificationGetOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInferenceClassificationGetOverridesOpts struct


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


## MeInferenceClassificationListOverrides

> CollectionOfInferenceClassificationOverride MeInferenceClassificationListOverrides(ctx, optional)

Get overrides from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeInferenceClassificationListOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInferenceClassificationListOverridesOpts struct


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


## MeInferenceClassificationUpdateOverrides

> MeInferenceClassificationUpdateOverrides(ctx, inferenceClassificationOverrideId, microsoftGraphInferenceClassificationOverride)

Update the navigation property overrides in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeUpdateInferenceClassification

> MeUpdateInferenceClassification(ctx, microsoftGraphInferenceClassification)

Update the navigation property inferenceClassification in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

