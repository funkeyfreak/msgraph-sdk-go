# \MeOutlookUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeGetOutlook**](MeOutlookUserApi.md#MeGetOutlook) | **Get** /me/outlook | Get outlook from me
[**MeOutlookCreateMasterCategories**](MeOutlookUserApi.md#MeOutlookCreateMasterCategories) | **Post** /me/outlook/masterCategories | Create new navigation property to masterCategories for me
[**MeOutlookGetMasterCategories**](MeOutlookUserApi.md#MeOutlookGetMasterCategories) | **Get** /me/outlook/masterCategories/{outlookCategory-id} | Get masterCategories from me
[**MeOutlookListMasterCategories**](MeOutlookUserApi.md#MeOutlookListMasterCategories) | **Get** /me/outlook/masterCategories | Get masterCategories from me
[**MeOutlookUpdateMasterCategories**](MeOutlookUserApi.md#MeOutlookUpdateMasterCategories) | **Patch** /me/outlook/masterCategories/{outlookCategory-id} | Update the navigation property masterCategories in me
[**MeUpdateOutlook**](MeOutlookUserApi.md#MeUpdateOutlook) | **Patch** /me/outlook | Update the navigation property outlook in me



## MeGetOutlook

> MicrosoftGraphOutlookUser MeGetOutlook(ctx, optional)

Get outlook from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeGetOutlookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetOutlookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphOutlookUser**](microsoft.graph.outlookUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookCreateMasterCategories

> MicrosoftGraphOutlookCategory MeOutlookCreateMasterCategories(ctx, microsoftGraphOutlookCategory)

Create new navigation property to masterCategories for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOutlookCategory** | [**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md)| New navigation property | 

### Return type

[**MicrosoftGraphOutlookCategory**](microsoft.graph.outlookCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookGetMasterCategories

> MicrosoftGraphOutlookCategory MeOutlookGetMasterCategories(ctx, outlookCategoryId, optional)

Get masterCategories from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlookCategoryId** | **string**| key: outlookCategory-id of outlookCategory | 
 **optional** | ***MeOutlookGetMasterCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOutlookGetMasterCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphOutlookCategory**](microsoft.graph.outlookCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookListMasterCategories

> CollectionOfOutlookCategory MeOutlookListMasterCategories(ctx, optional)

Get masterCategories from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeOutlookListMasterCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOutlookListMasterCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfOutlookCategory**](Collection_of_outlookCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookUpdateMasterCategories

> MeOutlookUpdateMasterCategories(ctx, outlookCategoryId, microsoftGraphOutlookCategory)

Update the navigation property masterCategories in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlookCategoryId** | **string**| key: outlookCategory-id of outlookCategory | 
**microsoftGraphOutlookCategory** | [**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md)| New navigation property values | 

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


## MeUpdateOutlook

> MeUpdateOutlook(ctx, microsoftGraphOutlookUser)

Update the navigation property outlook in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOutlookUser** | [**MicrosoftGraphOutlookUser**](MicrosoftGraphOutlookUser.md)| New navigation property values | 

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

