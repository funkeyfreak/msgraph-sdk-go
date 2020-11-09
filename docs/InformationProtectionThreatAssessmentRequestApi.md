# \InformationProtectionThreatAssessmentRequestApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InformationProtectionCreateThreatAssessmentRequests**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionCreateThreatAssessmentRequests) | **Post** /informationProtection/threatAssessmentRequests | Create new navigation property to threatAssessmentRequests for informationProtection
[**InformationProtectionGetThreatAssessmentRequests**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionGetThreatAssessmentRequests) | **Get** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id} | Get threatAssessmentRequests from informationProtection
[**InformationProtectionListThreatAssessmentRequests**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionListThreatAssessmentRequests) | **Get** /informationProtection/threatAssessmentRequests | Get threatAssessmentRequests from informationProtection
[**InformationProtectionThreatAssessmentRequestsCreateResults**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionThreatAssessmentRequestsCreateResults) | **Post** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id}/results | Create new navigation property to results for informationProtection
[**InformationProtectionThreatAssessmentRequestsGetResults**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionThreatAssessmentRequestsGetResults) | **Get** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id}/results/{threatAssessmentResult-id} | Get results from informationProtection
[**InformationProtectionThreatAssessmentRequestsListResults**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionThreatAssessmentRequestsListResults) | **Get** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id}/results | Get results from informationProtection
[**InformationProtectionThreatAssessmentRequestsUpdateResults**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionThreatAssessmentRequestsUpdateResults) | **Patch** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id}/results/{threatAssessmentResult-id} | Update the navigation property results in informationProtection
[**InformationProtectionUpdateThreatAssessmentRequests**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionUpdateThreatAssessmentRequests) | **Patch** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id} | Update the navigation property threatAssessmentRequests in informationProtection



## InformationProtectionCreateThreatAssessmentRequests

> MicrosoftGraphThreatAssessmentRequest InformationProtectionCreateThreatAssessmentRequests(ctx, microsoftGraphThreatAssessmentRequest)

Create new navigation property to threatAssessmentRequests for informationProtection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphThreatAssessmentRequest** | [**MicrosoftGraphThreatAssessmentRequest**](MicrosoftGraphThreatAssessmentRequest.md)| New navigation property | 

### Return type

[**MicrosoftGraphThreatAssessmentRequest**](microsoft.graph.threatAssessmentRequest.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionGetThreatAssessmentRequests

> MicrosoftGraphThreatAssessmentRequest InformationProtectionGetThreatAssessmentRequests(ctx, threatAssessmentRequestId, optional)

Get threatAssessmentRequests from informationProtection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string**| key: threatAssessmentRequest-id of threatAssessmentRequest | 
 **optional** | ***InformationProtectionGetThreatAssessmentRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InformationProtectionGetThreatAssessmentRequestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphThreatAssessmentRequest**](microsoft.graph.threatAssessmentRequest.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionListThreatAssessmentRequests

> CollectionOfThreatAssessmentRequest InformationProtectionListThreatAssessmentRequests(ctx, optional)

Get threatAssessmentRequests from informationProtection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InformationProtectionListThreatAssessmentRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InformationProtectionListThreatAssessmentRequestsOpts struct


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

[**CollectionOfThreatAssessmentRequest**](Collection_of_threatAssessmentRequest.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionThreatAssessmentRequestsCreateResults

> MicrosoftGraphThreatAssessmentResult InformationProtectionThreatAssessmentRequestsCreateResults(ctx, threatAssessmentRequestId, microsoftGraphThreatAssessmentResult)

Create new navigation property to results for informationProtection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string**| key: threatAssessmentRequest-id of threatAssessmentRequest | 
**microsoftGraphThreatAssessmentResult** | [**MicrosoftGraphThreatAssessmentResult**](MicrosoftGraphThreatAssessmentResult.md)| New navigation property | 

### Return type

[**MicrosoftGraphThreatAssessmentResult**](microsoft.graph.threatAssessmentResult.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionThreatAssessmentRequestsGetResults

> MicrosoftGraphThreatAssessmentResult InformationProtectionThreatAssessmentRequestsGetResults(ctx, threatAssessmentRequestId, threatAssessmentResultId, optional)

Get results from informationProtection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string**| key: threatAssessmentRequest-id of threatAssessmentRequest | 
**threatAssessmentResultId** | **string**| key: threatAssessmentResult-id of threatAssessmentResult | 
 **optional** | ***InformationProtectionThreatAssessmentRequestsGetResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InformationProtectionThreatAssessmentRequestsGetResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphThreatAssessmentResult**](microsoft.graph.threatAssessmentResult.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionThreatAssessmentRequestsListResults

> CollectionOfThreatAssessmentResult InformationProtectionThreatAssessmentRequestsListResults(ctx, threatAssessmentRequestId, optional)

Get results from informationProtection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string**| key: threatAssessmentRequest-id of threatAssessmentRequest | 
 **optional** | ***InformationProtectionThreatAssessmentRequestsListResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InformationProtectionThreatAssessmentRequestsListResultsOpts struct


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

[**CollectionOfThreatAssessmentResult**](Collection_of_threatAssessmentResult.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionThreatAssessmentRequestsUpdateResults

> InformationProtectionThreatAssessmentRequestsUpdateResults(ctx, threatAssessmentRequestId, threatAssessmentResultId, microsoftGraphThreatAssessmentResult)

Update the navigation property results in informationProtection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string**| key: threatAssessmentRequest-id of threatAssessmentRequest | 
**threatAssessmentResultId** | **string**| key: threatAssessmentResult-id of threatAssessmentResult | 
**microsoftGraphThreatAssessmentResult** | [**MicrosoftGraphThreatAssessmentResult**](MicrosoftGraphThreatAssessmentResult.md)| New navigation property values | 

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


## InformationProtectionUpdateThreatAssessmentRequests

> InformationProtectionUpdateThreatAssessmentRequests(ctx, threatAssessmentRequestId, microsoftGraphThreatAssessmentRequest)

Update the navigation property threatAssessmentRequests in informationProtection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string**| key: threatAssessmentRequest-id of threatAssessmentRequest | 
**microsoftGraphThreatAssessmentRequest** | [**MicrosoftGraphThreatAssessmentRequest**](MicrosoftGraphThreatAssessmentRequest.md)| New navigation property values | 

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

