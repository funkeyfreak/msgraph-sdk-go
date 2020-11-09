# \SecuritySecureScoreApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityCreateSecureScores**](SecuritySecureScoreApi.md#SecurityCreateSecureScores) | **Post** /Security/secureScores | Create new navigation property to secureScores for Security
[**SecurityGetSecureScores**](SecuritySecureScoreApi.md#SecurityGetSecureScores) | **Get** /Security/secureScores/{secureScore-id} | Get secureScores from Security
[**SecurityListSecureScores**](SecuritySecureScoreApi.md#SecurityListSecureScores) | **Get** /Security/secureScores | Get secureScores from Security
[**SecurityUpdateSecureScores**](SecuritySecureScoreApi.md#SecurityUpdateSecureScores) | **Patch** /Security/secureScores/{secureScore-id} | Update the navigation property secureScores in Security



## SecurityCreateSecureScores

> MicrosoftGraphSecureScore SecurityCreateSecureScores(ctx, microsoftGraphSecureScore)

Create new navigation property to secureScores for Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSecureScore** | [**MicrosoftGraphSecureScore**](MicrosoftGraphSecureScore.md)| New navigation property | 

### Return type

[**MicrosoftGraphSecureScore**](microsoft.graph.secureScore.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityGetSecureScores

> MicrosoftGraphSecureScore SecurityGetSecureScores(ctx, secureScoreId, optional)

Get secureScores from Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreId** | **string**| key: secureScore-id of secureScore | 
 **optional** | ***SecurityGetSecureScoresOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityGetSecureScoresOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSecureScore**](microsoft.graph.secureScore.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityListSecureScores

> CollectionOfSecureScore SecurityListSecureScores(ctx, optional)

Get secureScores from Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityListSecureScoresOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityListSecureScoresOpts struct


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

[**CollectionOfSecureScore**](Collection_of_secureScore.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityUpdateSecureScores

> SecurityUpdateSecureScores(ctx, secureScoreId, microsoftGraphSecureScore)

Update the navigation property secureScores in Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreId** | **string**| key: secureScore-id of secureScore | 
**microsoftGraphSecureScore** | [**MicrosoftGraphSecureScore**](MicrosoftGraphSecureScore.md)| New navigation property values | 

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

