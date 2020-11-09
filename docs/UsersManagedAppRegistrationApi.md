# \UsersManagedAppRegistrationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetManagedAppRegistrations**](UsersManagedAppRegistrationApi.md#UsersGetManagedAppRegistrations) | **Get** /users/{user-id}/managedAppRegistrations/{managedAppRegistration-id} | Get managedAppRegistrations from users
[**UsersListManagedAppRegistrations**](UsersManagedAppRegistrationApi.md#UsersListManagedAppRegistrations) | **Get** /users/{user-id}/managedAppRegistrations | Get managedAppRegistrations from users



## UsersGetManagedAppRegistrations

> MicrosoftGraphManagedAppRegistration UsersGetManagedAppRegistrations(ctx, userId, managedAppRegistrationId, optional)

Get managedAppRegistrations from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
 **optional** | ***UsersGetManagedAppRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetManagedAppRegistrationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppRegistration**](microsoft.graph.managedAppRegistration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListManagedAppRegistrations

> CollectionOfManagedAppRegistration UsersListManagedAppRegistrations(ctx, userId, optional)

Get managedAppRegistrations from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListManagedAppRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListManagedAppRegistrationsOpts struct


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

[**CollectionOfManagedAppRegistration**](Collection_of_managedAppRegistration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

