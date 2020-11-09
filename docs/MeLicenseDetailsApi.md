# \MeLicenseDetailsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateLicenseDetails**](MeLicenseDetailsApi.md#MeCreateLicenseDetails) | **Post** /me/licenseDetails | Create new navigation property to licenseDetails for me
[**MeGetLicenseDetails**](MeLicenseDetailsApi.md#MeGetLicenseDetails) | **Get** /me/licenseDetails/{licenseDetails-id} | Get licenseDetails from me
[**MeListLicenseDetails**](MeLicenseDetailsApi.md#MeListLicenseDetails) | **Get** /me/licenseDetails | Get licenseDetails from me
[**MeUpdateLicenseDetails**](MeLicenseDetailsApi.md#MeUpdateLicenseDetails) | **Patch** /me/licenseDetails/{licenseDetails-id} | Update the navigation property licenseDetails in me



## MeCreateLicenseDetails

> MicrosoftGraphLicenseDetails MeCreateLicenseDetails(ctx, microsoftGraphLicenseDetails)

Create new navigation property to licenseDetails for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeGetLicenseDetails

> MicrosoftGraphLicenseDetails MeGetLicenseDetails(ctx, licenseDetailsId, optional)

Get licenseDetails from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseDetailsId** | **string**| key: licenseDetails-id of licenseDetails | 
 **optional** | ***MeGetLicenseDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetLicenseDetailsOpts struct


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


## MeListLicenseDetails

> CollectionOfLicenseDetails MeListLicenseDetails(ctx, optional)

Get licenseDetails from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListLicenseDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListLicenseDetailsOpts struct


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


## MeUpdateLicenseDetails

> MeUpdateLicenseDetails(ctx, licenseDetailsId, microsoftGraphLicenseDetails)

Update the navigation property licenseDetails in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

