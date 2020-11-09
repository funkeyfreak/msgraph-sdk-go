# \MeManagedDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateManagedDevices**](MeManagedDeviceApi.md#MeCreateManagedDevices) | **Post** /me/managedDevices | Create new navigation property to managedDevices for me
[**MeGetManagedDevices**](MeManagedDeviceApi.md#MeGetManagedDevices) | **Get** /me/managedDevices/{managedDevice-id} | Get managedDevices from me
[**MeListManagedDevices**](MeManagedDeviceApi.md#MeListManagedDevices) | **Get** /me/managedDevices | Get managedDevices from me
[**MeManagedDevicesCreateDeviceCompliancePolicyStates**](MeManagedDeviceApi.md#MeManagedDevicesCreateDeviceCompliancePolicyStates) | **Post** /me/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Create new navigation property to deviceCompliancePolicyStates for me
[**MeManagedDevicesCreateDeviceConfigurationStates**](MeManagedDeviceApi.md#MeManagedDevicesCreateDeviceConfigurationStates) | **Post** /me/managedDevices/{managedDevice-id}/deviceConfigurationStates | Create new navigation property to deviceConfigurationStates for me
[**MeManagedDevicesGetDeviceCategory**](MeManagedDeviceApi.md#MeManagedDevicesGetDeviceCategory) | **Get** /me/managedDevices/{managedDevice-id}/deviceCategory | Get deviceCategory from me
[**MeManagedDevicesGetDeviceCompliancePolicyStates**](MeManagedDeviceApi.md#MeManagedDevicesGetDeviceCompliancePolicyStates) | **Get** /me/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Get deviceCompliancePolicyStates from me
[**MeManagedDevicesGetDeviceConfigurationStates**](MeManagedDeviceApi.md#MeManagedDevicesGetDeviceConfigurationStates) | **Get** /me/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Get deviceConfigurationStates from me
[**MeManagedDevicesListDeviceCompliancePolicyStates**](MeManagedDeviceApi.md#MeManagedDevicesListDeviceCompliancePolicyStates) | **Get** /me/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Get deviceCompliancePolicyStates from me
[**MeManagedDevicesListDeviceConfigurationStates**](MeManagedDeviceApi.md#MeManagedDevicesListDeviceConfigurationStates) | **Get** /me/managedDevices/{managedDevice-id}/deviceConfigurationStates | Get deviceConfigurationStates from me
[**MeManagedDevicesUpdateDeviceCategory**](MeManagedDeviceApi.md#MeManagedDevicesUpdateDeviceCategory) | **Patch** /me/managedDevices/{managedDevice-id}/deviceCategory | Update the navigation property deviceCategory in me
[**MeManagedDevicesUpdateDeviceCompliancePolicyStates**](MeManagedDeviceApi.md#MeManagedDevicesUpdateDeviceCompliancePolicyStates) | **Patch** /me/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Update the navigation property deviceCompliancePolicyStates in me
[**MeManagedDevicesUpdateDeviceConfigurationStates**](MeManagedDeviceApi.md#MeManagedDevicesUpdateDeviceConfigurationStates) | **Patch** /me/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Update the navigation property deviceConfigurationStates in me
[**MeUpdateManagedDevices**](MeManagedDeviceApi.md#MeUpdateManagedDevices) | **Patch** /me/managedDevices/{managedDevice-id} | Update the navigation property managedDevices in me



## MeCreateManagedDevices

> MicrosoftGraphManagedDevice MeCreateManagedDevices(ctx, microsoftGraphManagedDevice)

Create new navigation property to managedDevices for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeGetManagedDevices

> MicrosoftGraphManagedDevice MeGetManagedDevices(ctx, managedDeviceId, optional)

Get managedDevices from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***MeGetManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetManagedDevicesOpts struct


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


## MeListManagedDevices

> CollectionOfManagedDevice MeListManagedDevices(ctx, optional)

Get managedDevices from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListManagedDevicesOpts struct


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


## MeManagedDevicesCreateDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState MeManagedDevicesCreateDeviceCompliancePolicyStates(ctx, managedDeviceId, microsoftGraphDeviceCompliancePolicyState)

Create new navigation property to deviceCompliancePolicyStates for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeManagedDevicesCreateDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState MeManagedDevicesCreateDeviceConfigurationStates(ctx, managedDeviceId, microsoftGraphDeviceConfigurationState)

Create new navigation property to deviceConfigurationStates for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeManagedDevicesGetDeviceCategory

> MicrosoftGraphDeviceCategory MeManagedDevicesGetDeviceCategory(ctx, managedDeviceId, optional)

Get deviceCategory from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***MeManagedDevicesGetDeviceCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeManagedDevicesGetDeviceCategoryOpts struct


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


## MeManagedDevicesGetDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState MeManagedDevicesGetDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId, optional)

Get deviceCompliancePolicyStates from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceCompliancePolicyStateId** | **string**| key: deviceCompliancePolicyState-id of deviceCompliancePolicyState | 
 **optional** | ***MeManagedDevicesGetDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeManagedDevicesGetDeviceCompliancePolicyStatesOpts struct


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


## MeManagedDevicesGetDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState MeManagedDevicesGetDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId, optional)

Get deviceConfigurationStates from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceConfigurationStateId** | **string**| key: deviceConfigurationState-id of deviceConfigurationState | 
 **optional** | ***MeManagedDevicesGetDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeManagedDevicesGetDeviceConfigurationStatesOpts struct


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


## MeManagedDevicesListDeviceCompliancePolicyStates

> CollectionOfDeviceCompliancePolicyState MeManagedDevicesListDeviceCompliancePolicyStates(ctx, managedDeviceId, optional)

Get deviceCompliancePolicyStates from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***MeManagedDevicesListDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeManagedDevicesListDeviceCompliancePolicyStatesOpts struct


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


## MeManagedDevicesListDeviceConfigurationStates

> CollectionOfDeviceConfigurationState MeManagedDevicesListDeviceConfigurationStates(ctx, managedDeviceId, optional)

Get deviceConfigurationStates from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***MeManagedDevicesListDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeManagedDevicesListDeviceConfigurationStatesOpts struct


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


## MeManagedDevicesUpdateDeviceCategory

> MeManagedDevicesUpdateDeviceCategory(ctx, managedDeviceId, microsoftGraphDeviceCategory)

Update the navigation property deviceCategory in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeManagedDevicesUpdateDeviceCompliancePolicyStates

> MeManagedDevicesUpdateDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId, microsoftGraphDeviceCompliancePolicyState)

Update the navigation property deviceCompliancePolicyStates in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeManagedDevicesUpdateDeviceConfigurationStates

> MeManagedDevicesUpdateDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId, microsoftGraphDeviceConfigurationState)

Update the navigation property deviceConfigurationStates in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeUpdateManagedDevices

> MeUpdateManagedDevices(ctx, managedDeviceId, microsoftGraphManagedDevice)

Update the navigation property managedDevices in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

