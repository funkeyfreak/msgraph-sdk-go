# \UsersManagedDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateManagedDevices**](UsersManagedDeviceApi.md#UsersCreateManagedDevices) | **Post** /users/{user-id}/managedDevices | Create new navigation property to managedDevices for users
[**UsersGetManagedDevices**](UsersManagedDeviceApi.md#UsersGetManagedDevices) | **Get** /users/{user-id}/managedDevices/{managedDevice-id} | Get managedDevices from users
[**UsersListManagedDevices**](UsersManagedDeviceApi.md#UsersListManagedDevices) | **Get** /users/{user-id}/managedDevices | Get managedDevices from users
[**UsersManagedDevicesCreateDeviceCompliancePolicyStates**](UsersManagedDeviceApi.md#UsersManagedDevicesCreateDeviceCompliancePolicyStates) | **Post** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Create new navigation property to deviceCompliancePolicyStates for users
[**UsersManagedDevicesCreateDeviceConfigurationStates**](UsersManagedDeviceApi.md#UsersManagedDevicesCreateDeviceConfigurationStates) | **Post** /users/{user-id}/managedDevices/{managedDevice-id}/deviceConfigurationStates | Create new navigation property to deviceConfigurationStates for users
[**UsersManagedDevicesGetDeviceCategory**](UsersManagedDeviceApi.md#UsersManagedDevicesGetDeviceCategory) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCategory | Get deviceCategory from users
[**UsersManagedDevicesGetDeviceCompliancePolicyStates**](UsersManagedDeviceApi.md#UsersManagedDevicesGetDeviceCompliancePolicyStates) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Get deviceCompliancePolicyStates from users
[**UsersManagedDevicesGetDeviceConfigurationStates**](UsersManagedDeviceApi.md#UsersManagedDevicesGetDeviceConfigurationStates) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Get deviceConfigurationStates from users
[**UsersManagedDevicesListDeviceCompliancePolicyStates**](UsersManagedDeviceApi.md#UsersManagedDevicesListDeviceCompliancePolicyStates) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Get deviceCompliancePolicyStates from users
[**UsersManagedDevicesListDeviceConfigurationStates**](UsersManagedDeviceApi.md#UsersManagedDevicesListDeviceConfigurationStates) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceConfigurationStates | Get deviceConfigurationStates from users
[**UsersManagedDevicesUpdateDeviceCategory**](UsersManagedDeviceApi.md#UsersManagedDevicesUpdateDeviceCategory) | **Patch** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCategory | Update the navigation property deviceCategory in users
[**UsersManagedDevicesUpdateDeviceCompliancePolicyStates**](UsersManagedDeviceApi.md#UsersManagedDevicesUpdateDeviceCompliancePolicyStates) | **Patch** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Update the navigation property deviceCompliancePolicyStates in users
[**UsersManagedDevicesUpdateDeviceConfigurationStates**](UsersManagedDeviceApi.md#UsersManagedDevicesUpdateDeviceConfigurationStates) | **Patch** /users/{user-id}/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Update the navigation property deviceConfigurationStates in users
[**UsersUpdateManagedDevices**](UsersManagedDeviceApi.md#UsersUpdateManagedDevices) | **Patch** /users/{user-id}/managedDevices/{managedDevice-id} | Update the navigation property managedDevices in users



## UsersCreateManagedDevices

> MicrosoftGraphManagedDevice UsersCreateManagedDevices(ctx, userId, microsoftGraphManagedDevice)

Create new navigation property to managedDevices for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphManagedDevice** | [**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetManagedDevices

> MicrosoftGraphManagedDevice UsersGetManagedDevices(ctx, userId, managedDeviceId, optional)

Get managedDevices from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***UsersGetManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetManagedDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListManagedDevices

> CollectionOfManagedDevice UsersListManagedDevices(ctx, userId, optional)

Get managedDevices from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListManagedDevicesOpts struct


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

[**CollectionOfManagedDevice**](Collection_of_managedDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesCreateDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState UsersManagedDevicesCreateDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, microsoftGraphDeviceCompliancePolicyState)

Create new navigation property to deviceCompliancePolicyStates for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**microsoftGraphDeviceCompliancePolicyState** | [**MicrosoftGraphDeviceCompliancePolicyState**](MicrosoftGraphDeviceCompliancePolicyState.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyState**](microsoft.graph.deviceCompliancePolicyState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesCreateDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState UsersManagedDevicesCreateDeviceConfigurationStates(ctx, userId, managedDeviceId, microsoftGraphDeviceConfigurationState)

Create new navigation property to deviceConfigurationStates for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**microsoftGraphDeviceConfigurationState** | [**MicrosoftGraphDeviceConfigurationState**](MicrosoftGraphDeviceConfigurationState.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceConfigurationState**](microsoft.graph.deviceConfigurationState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesGetDeviceCategory

> MicrosoftGraphDeviceCategory UsersManagedDevicesGetDeviceCategory(ctx, userId, managedDeviceId, optional)

Get deviceCategory from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***UsersManagedDevicesGetDeviceCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesGetDeviceCategoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCategory**](microsoft.graph.deviceCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesGetDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState UsersManagedDevicesGetDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, deviceCompliancePolicyStateId, optional)

Get deviceCompliancePolicyStates from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceCompliancePolicyStateId** | **string**| key: deviceCompliancePolicyState-id of deviceCompliancePolicyState | 
 **optional** | ***UsersManagedDevicesGetDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesGetDeviceCompliancePolicyStatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyState**](microsoft.graph.deviceCompliancePolicyState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesGetDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState UsersManagedDevicesGetDeviceConfigurationStates(ctx, userId, managedDeviceId, deviceConfigurationStateId, optional)

Get deviceConfigurationStates from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceConfigurationStateId** | **string**| key: deviceConfigurationState-id of deviceConfigurationState | 
 **optional** | ***UsersManagedDevicesGetDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesGetDeviceConfigurationStatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationState**](microsoft.graph.deviceConfigurationState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesListDeviceCompliancePolicyStates

> CollectionOfDeviceCompliancePolicyState UsersManagedDevicesListDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, optional)

Get deviceCompliancePolicyStates from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***UsersManagedDevicesListDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesListDeviceCompliancePolicyStatesOpts struct


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

[**CollectionOfDeviceCompliancePolicyState**](Collection_of_deviceCompliancePolicyState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesListDeviceConfigurationStates

> CollectionOfDeviceConfigurationState UsersManagedDevicesListDeviceConfigurationStates(ctx, userId, managedDeviceId, optional)

Get deviceConfigurationStates from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***UsersManagedDevicesListDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesListDeviceConfigurationStatesOpts struct


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

[**CollectionOfDeviceConfigurationState**](Collection_of_deviceConfigurationState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesUpdateDeviceCategory

> UsersManagedDevicesUpdateDeviceCategory(ctx, userId, managedDeviceId, microsoftGraphDeviceCategory)

Update the navigation property deviceCategory in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**microsoftGraphDeviceCategory** | [**MicrosoftGraphDeviceCategory**](MicrosoftGraphDeviceCategory.md)| New navigation property values | 

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


## UsersManagedDevicesUpdateDeviceCompliancePolicyStates

> UsersManagedDevicesUpdateDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, deviceCompliancePolicyStateId, microsoftGraphDeviceCompliancePolicyState)

Update the navigation property deviceCompliancePolicyStates in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceCompliancePolicyStateId** | **string**| key: deviceCompliancePolicyState-id of deviceCompliancePolicyState | 
**microsoftGraphDeviceCompliancePolicyState** | [**MicrosoftGraphDeviceCompliancePolicyState**](MicrosoftGraphDeviceCompliancePolicyState.md)| New navigation property values | 

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


## UsersManagedDevicesUpdateDeviceConfigurationStates

> UsersManagedDevicesUpdateDeviceConfigurationStates(ctx, userId, managedDeviceId, deviceConfigurationStateId, microsoftGraphDeviceConfigurationState)

Update the navigation property deviceConfigurationStates in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceConfigurationStateId** | **string**| key: deviceConfigurationState-id of deviceConfigurationState | 
**microsoftGraphDeviceConfigurationState** | [**MicrosoftGraphDeviceConfigurationState**](MicrosoftGraphDeviceConfigurationState.md)| New navigation property values | 

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


## UsersUpdateManagedDevices

> UsersUpdateManagedDevices(ctx, userId, managedDeviceId, microsoftGraphManagedDevice)

Update the navigation property managedDevices in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**microsoftGraphManagedDevice** | [**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md)| New navigation property values | 

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

