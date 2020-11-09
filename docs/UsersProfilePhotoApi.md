# \UsersProfilePhotoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreatePhotos**](UsersProfilePhotoApi.md#UsersCreatePhotos) | **Post** /users/{user-id}/photos | Create new navigation property to photos for users
[**UsersGetPhoto**](UsersProfilePhotoApi.md#UsersGetPhoto) | **Get** /users/{user-id}/photo | Get photo from users
[**UsersGetPhotos**](UsersProfilePhotoApi.md#UsersGetPhotos) | **Get** /users/{user-id}/photos/{profilePhoto-id} | Get photos from users
[**UsersListPhotos**](UsersProfilePhotoApi.md#UsersListPhotos) | **Get** /users/{user-id}/photos | Get photos from users
[**UsersUpdatePhoto**](UsersProfilePhotoApi.md#UsersUpdatePhoto) | **Patch** /users/{user-id}/photo | Update the navigation property photo in users
[**UsersUpdatePhotos**](UsersProfilePhotoApi.md#UsersUpdatePhotos) | **Patch** /users/{user-id}/photos/{profilePhoto-id} | Update the navigation property photos in users



## UsersCreatePhotos

> MicrosoftGraphProfilePhoto UsersCreatePhotos(ctx, userId, microsoftGraphProfilePhoto)

Create new navigation property to photos for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersGetPhoto

> MicrosoftGraphProfilePhoto UsersGetPhoto(ctx, userId, optional)

Get photo from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetPhotoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetPhotoOpts struct


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


## UsersGetPhotos

> MicrosoftGraphProfilePhoto UsersGetPhotos(ctx, userId, profilePhotoId, optional)

Get photos from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**profilePhotoId** | **string**| key: profilePhoto-id of profilePhoto | 
 **optional** | ***UsersGetPhotosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetPhotosOpts struct


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


## UsersListPhotos

> CollectionOfProfilePhoto UsersListPhotos(ctx, userId, optional)

Get photos from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListPhotosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListPhotosOpts struct


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


## UsersUpdatePhoto

> UsersUpdatePhoto(ctx, userId, microsoftGraphProfilePhoto)

Update the navigation property photo in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersUpdatePhotos

> UsersUpdatePhotos(ctx, userId, profilePhotoId, microsoftGraphProfilePhoto)

Update the navigation property photos in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

