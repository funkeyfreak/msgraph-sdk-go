# \UsersExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateExtensions**](UsersExtensionApi.md#UsersCreateExtensions) | **Post** /users/{user-id}/extensions | Create new navigation property to extensions for users
[**UsersGetExtensions**](UsersExtensionApi.md#UsersGetExtensions) | **Get** /users/{user-id}/extensions/{extension-id} | Get extensions from users
[**UsersListExtensions**](UsersExtensionApi.md#UsersListExtensions) | **Get** /users/{user-id}/extensions | Get extensions from users
[**UsersUpdateExtensions**](UsersExtensionApi.md#UsersUpdateExtensions) | **Patch** /users/{user-id}/extensions/{extension-id} | Update the navigation property extensions in users



## UsersCreateExtensions

> MicrosoftGraphExtension UsersCreateExtensions(ctx, userId, microsoftGraphExtension)

Create new navigation property to extensions for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)| New navigation property | 

### Return type

[**MicrosoftGraphExtension**](microsoft.graph.extension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetExtensions

> MicrosoftGraphExtension UsersGetExtensions(ctx, userId, extensionId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***UsersGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetExtensionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphExtension**](microsoft.graph.extension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListExtensions

> CollectionOfExtension UsersListExtensions(ctx, userId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListExtensionsOpts struct


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

[**CollectionOfExtension**](Collection_of_extension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateExtensions

> UsersUpdateExtensions(ctx, userId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**extensionId** | **string**| key: extension-id of extension | 
**microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)| New navigation property values | 

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

