# \GroupsProfilePhotoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreatePhotos**](GroupsProfilePhotoApi.md#GroupsCreatePhotos) | **Post** /groups/{group-id}/photos | Create new navigation property to photos for groups
[**GroupsGetPhoto**](GroupsProfilePhotoApi.md#GroupsGetPhoto) | **Get** /groups/{group-id}/photo | Get photo from groups
[**GroupsGetPhotos**](GroupsProfilePhotoApi.md#GroupsGetPhotos) | **Get** /groups/{group-id}/photos/{profilePhoto-id} | Get photos from groups
[**GroupsListPhotos**](GroupsProfilePhotoApi.md#GroupsListPhotos) | **Get** /groups/{group-id}/photos | Get photos from groups
[**GroupsUpdatePhoto**](GroupsProfilePhotoApi.md#GroupsUpdatePhoto) | **Patch** /groups/{group-id}/photo | Update the navigation property photo in groups
[**GroupsUpdatePhotos**](GroupsProfilePhotoApi.md#GroupsUpdatePhotos) | **Patch** /groups/{group-id}/photos/{profilePhoto-id} | Update the navigation property photos in groups



## GroupsCreatePhotos

> MicrosoftGraphProfilePhoto GroupsCreatePhotos(ctx, groupId, microsoftGraphProfilePhoto)

Create new navigation property to photos for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsGetPhoto

> MicrosoftGraphProfilePhoto GroupsGetPhoto(ctx, groupId, optional)

Get photo from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsGetPhotoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetPhotoOpts struct


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


## GroupsGetPhotos

> MicrosoftGraphProfilePhoto GroupsGetPhotos(ctx, groupId, profilePhotoId, optional)

Get photos from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**profilePhotoId** | **string**| key: profilePhoto-id of profilePhoto | 
 **optional** | ***GroupsGetPhotosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetPhotosOpts struct


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


## GroupsListPhotos

> CollectionOfProfilePhoto GroupsListPhotos(ctx, groupId, optional)

Get photos from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsListPhotosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsListPhotosOpts struct


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


## GroupsUpdatePhoto

> GroupsUpdatePhoto(ctx, groupId, microsoftGraphProfilePhoto)

Update the navigation property photo in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsUpdatePhotos

> GroupsUpdatePhotos(ctx, groupId, profilePhotoId, microsoftGraphProfilePhoto)

Update the navigation property photos in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

