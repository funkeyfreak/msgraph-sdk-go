# \MeProfilePhotoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreatePhotos**](MeProfilePhotoApi.md#MeCreatePhotos) | **Post** /me/photos | Create new navigation property to photos for me
[**MeGetPhoto**](MeProfilePhotoApi.md#MeGetPhoto) | **Get** /me/photo | Get photo from me
[**MeGetPhotos**](MeProfilePhotoApi.md#MeGetPhotos) | **Get** /me/photos/{profilePhoto-id} | Get photos from me
[**MeListPhotos**](MeProfilePhotoApi.md#MeListPhotos) | **Get** /me/photos | Get photos from me
[**MeUpdatePhoto**](MeProfilePhotoApi.md#MeUpdatePhoto) | **Patch** /me/photo | Update the navigation property photo in me
[**MeUpdatePhotos**](MeProfilePhotoApi.md#MeUpdatePhotos) | **Patch** /me/photos/{profilePhoto-id} | Update the navigation property photos in me



## MeCreatePhotos

> MicrosoftGraphProfilePhoto MeCreatePhotos(ctx, microsoftGraphProfilePhoto)

Create new navigation property to photos for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)| New navigation property | 

### Return type

[**MicrosoftGraphProfilePhoto**](microsoft.graph.profilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeGetPhoto

> MicrosoftGraphProfilePhoto MeGetPhoto(ctx, optional)

Get photo from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeGetPhotoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetPhotoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphProfilePhoto**](microsoft.graph.profilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeGetPhotos

> MicrosoftGraphProfilePhoto MeGetPhotos(ctx, profilePhotoId, optional)

Get photos from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profilePhotoId** | **string**| key: profilePhoto-id of profilePhoto | 
 **optional** | ***MeGetPhotosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetPhotosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphProfilePhoto**](microsoft.graph.profilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListPhotos

> CollectionOfProfilePhoto MeListPhotos(ctx, optional)

Get photos from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListPhotosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListPhotosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfProfilePhoto**](Collection_of_profilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdatePhoto

> MeUpdatePhoto(ctx, microsoftGraphProfilePhoto)

Update the navigation property photo in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)| New navigation property values | 

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


## MeUpdatePhotos

> MeUpdatePhotos(ctx, profilePhotoId, microsoftGraphProfilePhoto)

Update the navigation property photos in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profilePhotoId** | **string**| key: profilePhoto-id of profilePhoto | 
**microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)| New navigation property values | 

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

