# \UsersLicenseDetailsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateLicenseDetails**](UsersLicenseDetailsApi.md#UsersCreateLicenseDetails) | **Post** /users/{user-id}/licenseDetails | Create new navigation property to licenseDetails for users
[**UsersGetLicenseDetails**](UsersLicenseDetailsApi.md#UsersGetLicenseDetails) | **Get** /users/{user-id}/licenseDetails/{licenseDetails-id} | Get licenseDetails from users
[**UsersListLicenseDetails**](UsersLicenseDetailsApi.md#UsersListLicenseDetails) | **Get** /users/{user-id}/licenseDetails | Get licenseDetails from users
[**UsersUpdateLicenseDetails**](UsersLicenseDetailsApi.md#UsersUpdateLicenseDetails) | **Patch** /users/{user-id}/licenseDetails/{licenseDetails-id} | Update the navigation property licenseDetails in users



## UsersCreateLicenseDetails

> MicrosoftGraphLicenseDetails UsersCreateLicenseDetails(ctx, userId, microsoftGraphLicenseDetails)

Create new navigation property to licenseDetails for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphLicenseDetails** | [**MicrosoftGraphLicenseDetails**](MicrosoftGraphLicenseDetails.md)| New navigation property | 

### Return type

[**MicrosoftGraphLicenseDetails**](microsoft.graph.licenseDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetLicenseDetails

> MicrosoftGraphLicenseDetails UsersGetLicenseDetails(ctx, userId, licenseDetailsId, optional)

Get licenseDetails from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**licenseDetailsId** | **string**| key: licenseDetails-id of licenseDetails | 
 **optional** | ***UsersGetLicenseDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetLicenseDetailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphLicenseDetails**](microsoft.graph.licenseDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListLicenseDetails

> CollectionOfLicenseDetails UsersListLicenseDetails(ctx, userId, optional)

Get licenseDetails from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListLicenseDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListLicenseDetailsOpts struct


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

[**CollectionOfLicenseDetails**](Collection_of_licenseDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateLicenseDetails

> UsersUpdateLicenseDetails(ctx, userId, licenseDetailsId, microsoftGraphLicenseDetails)

Update the navigation property licenseDetails in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**licenseDetailsId** | **string**| key: licenseDetails-id of licenseDetails | 
**microsoftGraphLicenseDetails** | [**MicrosoftGraphLicenseDetails**](MicrosoftGraphLicenseDetails.md)| New navigation property values | 

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

