# \UsersScopedRoleMembershipApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateScopedRoleMemberOf**](UsersScopedRoleMembershipApi.md#UsersCreateScopedRoleMemberOf) | **Post** /users/{user-id}/scopedRoleMemberOf | Create new navigation property to scopedRoleMemberOf for users
[**UsersGetScopedRoleMemberOf**](UsersScopedRoleMembershipApi.md#UsersGetScopedRoleMemberOf) | **Get** /users/{user-id}/scopedRoleMemberOf/{scopedRoleMembership-id} | Get scopedRoleMemberOf from users
[**UsersListScopedRoleMemberOf**](UsersScopedRoleMembershipApi.md#UsersListScopedRoleMemberOf) | **Get** /users/{user-id}/scopedRoleMemberOf | Get scopedRoleMemberOf from users
[**UsersUpdateScopedRoleMemberOf**](UsersScopedRoleMembershipApi.md#UsersUpdateScopedRoleMemberOf) | **Patch** /users/{user-id}/scopedRoleMemberOf/{scopedRoleMembership-id} | Update the navigation property scopedRoleMemberOf in users



## UsersCreateScopedRoleMemberOf

> MicrosoftGraphScopedRoleMembership UsersCreateScopedRoleMemberOf(ctx, userId, microsoftGraphScopedRoleMembership)

Create new navigation property to scopedRoleMemberOf for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New navigation property | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](microsoft.graph.scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetScopedRoleMemberOf

> MicrosoftGraphScopedRoleMembership UsersGetScopedRoleMemberOf(ctx, userId, scopedRoleMembershipId, optional)

Get scopedRoleMemberOf from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
 **optional** | ***UsersGetScopedRoleMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetScopedRoleMemberOfOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](microsoft.graph.scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListScopedRoleMemberOf

> CollectionOfScopedRoleMembership UsersListScopedRoleMemberOf(ctx, userId, optional)

Get scopedRoleMemberOf from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListScopedRoleMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListScopedRoleMemberOfOpts struct


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

[**CollectionOfScopedRoleMembership**](Collection_of_scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateScopedRoleMemberOf

> UsersUpdateScopedRoleMemberOf(ctx, userId, scopedRoleMembershipId, microsoftGraphScopedRoleMembership)

Update the navigation property scopedRoleMemberOf in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New navigation property values | 

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

