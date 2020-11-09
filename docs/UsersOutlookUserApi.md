# \UsersOutlookUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetOutlook**](UsersOutlookUserApi.md#UsersGetOutlook) | **Get** /users/{user-id}/outlook | Get outlook from users
[**UsersOutlookCreateMasterCategories**](UsersOutlookUserApi.md#UsersOutlookCreateMasterCategories) | **Post** /users/{user-id}/outlook/masterCategories | Create new navigation property to masterCategories for users
[**UsersOutlookGetMasterCategories**](UsersOutlookUserApi.md#UsersOutlookGetMasterCategories) | **Get** /users/{user-id}/outlook/masterCategories/{outlookCategory-id} | Get masterCategories from users
[**UsersOutlookListMasterCategories**](UsersOutlookUserApi.md#UsersOutlookListMasterCategories) | **Get** /users/{user-id}/outlook/masterCategories | Get masterCategories from users
[**UsersOutlookUpdateMasterCategories**](UsersOutlookUserApi.md#UsersOutlookUpdateMasterCategories) | **Patch** /users/{user-id}/outlook/masterCategories/{outlookCategory-id} | Update the navigation property masterCategories in users
[**UsersUpdateOutlook**](UsersOutlookUserApi.md#UsersUpdateOutlook) | **Patch** /users/{user-id}/outlook | Update the navigation property outlook in users



## UsersGetOutlook

> MicrosoftGraphOutlookUser UsersGetOutlook(ctx, userId, optional)

Get outlook from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetOutlookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetOutlookOpts struct


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


## UsersOutlookCreateMasterCategories

> MicrosoftGraphOutlookCategory UsersOutlookCreateMasterCategories(ctx, userId, microsoftGraphOutlookCategory)

Create new navigation property to masterCategories for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOutlookGetMasterCategories

> MicrosoftGraphOutlookCategory UsersOutlookGetMasterCategories(ctx, userId, outlookCategoryId, optional)

Get masterCategories from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**outlookCategoryId** | **string**| key: outlookCategory-id of outlookCategory | 
 **optional** | ***UsersOutlookGetMasterCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOutlookGetMasterCategoriesOpts struct


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


## UsersOutlookListMasterCategories

> CollectionOfOutlookCategory UsersOutlookListMasterCategories(ctx, userId, optional)

Get masterCategories from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersOutlookListMasterCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOutlookListMasterCategoriesOpts struct


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


## UsersOutlookUpdateMasterCategories

> UsersOutlookUpdateMasterCategories(ctx, userId, outlookCategoryId, microsoftGraphOutlookCategory)

Update the navigation property masterCategories in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersUpdateOutlook

> UsersUpdateOutlook(ctx, userId, microsoftGraphOutlookUser)

Update the navigation property outlook in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

