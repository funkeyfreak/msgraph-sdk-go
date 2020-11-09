# \DeviceManagementManagedDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementCreateManagedDevices) | **Post** /deviceManagement/managedDevices | Create new navigation property to managedDevices for deviceManagement
[**DeviceManagementGetManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementGetManagedDevices) | **Get** /deviceManagement/managedDevices/{managedDevice-id} | Get managedDevices from deviceManagement
[**DeviceManagementListManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementListManagedDevices) | **Get** /deviceManagement/managedDevices | Get managedDevices from deviceManagement
[**DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates) | **Post** /deviceManagement/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Create new navigation property to deviceCompliancePolicyStates for deviceManagement
[**DeviceManagementManagedDevicesCreateDeviceConfigurationStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesCreateDeviceConfigurationStates) | **Post** /deviceManagement/managedDevices/{managedDevice-id}/deviceConfigurationStates | Create new navigation property to deviceConfigurationStates for deviceManagement
[**DeviceManagementManagedDevicesGetDeviceCategory**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesGetDeviceCategory) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceCategory | Get deviceCategory from deviceManagement
[**DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Get deviceCompliancePolicyStates from deviceManagement
[**DeviceManagementManagedDevicesGetDeviceConfigurationStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesGetDeviceConfigurationStates) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Get deviceConfigurationStates from deviceManagement
[**DeviceManagementManagedDevicesListDeviceCompliancePolicyStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesListDeviceCompliancePolicyStates) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Get deviceCompliancePolicyStates from deviceManagement
[**DeviceManagementManagedDevicesListDeviceConfigurationStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesListDeviceConfigurationStates) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceConfigurationStates | Get deviceConfigurationStates from deviceManagement
[**DeviceManagementManagedDevicesUpdateDeviceCategory**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesUpdateDeviceCategory) | **Patch** /deviceManagement/managedDevices/{managedDevice-id}/deviceCategory | Update the navigation property deviceCategory in deviceManagement
[**DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates) | **Patch** /deviceManagement/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Update the navigation property deviceCompliancePolicyStates in deviceManagement
[**DeviceManagementManagedDevicesUpdateDeviceConfigurationStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesUpdateDeviceConfigurationStates) | **Patch** /deviceManagement/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Update the navigation property deviceConfigurationStates in deviceManagement
[**DeviceManagementUpdateManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementUpdateManagedDevices) | **Patch** /deviceManagement/managedDevices/{managedDevice-id} | Update the navigation property managedDevices in deviceManagement



## DeviceManagementCreateManagedDevices

> MicrosoftGraphManagedDevice DeviceManagementCreateManagedDevices(ctx, microsoftGraphManagedDevice)

Create new navigation property to managedDevices for deviceManagement

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


## DeviceManagementGetManagedDevices

> MicrosoftGraphManagedDevice DeviceManagementGetManagedDevices(ctx, managedDeviceId, optional)

Get managedDevices from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***DeviceManagementGetManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetManagedDevicesOpts struct


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


## DeviceManagementListManagedDevices

> CollectionOfManagedDevice DeviceManagementListManagedDevices(ctx, optional)

Get managedDevices from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListManagedDevicesOpts struct


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


## DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates(ctx, managedDeviceId, microsoftGraphDeviceCompliancePolicyState)

Create new navigation property to deviceCompliancePolicyStates for deviceManagement

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


## DeviceManagementManagedDevicesCreateDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState DeviceManagementManagedDevicesCreateDeviceConfigurationStates(ctx, managedDeviceId, microsoftGraphDeviceConfigurationState)

Create new navigation property to deviceConfigurationStates for deviceManagement

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


## DeviceManagementManagedDevicesGetDeviceCategory

> MicrosoftGraphDeviceCategory DeviceManagementManagedDevicesGetDeviceCategory(ctx, managedDeviceId, optional)

Get deviceCategory from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***DeviceManagementManagedDevicesGetDeviceCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementManagedDevicesGetDeviceCategoryOpts struct


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


## DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId, optional)

Get deviceCompliancePolicyStates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceCompliancePolicyStateId** | **string**| key: deviceCompliancePolicyState-id of deviceCompliancePolicyState | 
 **optional** | ***DeviceManagementManagedDevicesGetDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementManagedDevicesGetDeviceCompliancePolicyStatesOpts struct


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


## DeviceManagementManagedDevicesGetDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState DeviceManagementManagedDevicesGetDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId, optional)

Get deviceConfigurationStates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceConfigurationStateId** | **string**| key: deviceConfigurationState-id of deviceConfigurationState | 
 **optional** | ***DeviceManagementManagedDevicesGetDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementManagedDevicesGetDeviceConfigurationStatesOpts struct


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


## DeviceManagementManagedDevicesListDeviceCompliancePolicyStates

> CollectionOfDeviceCompliancePolicyState DeviceManagementManagedDevicesListDeviceCompliancePolicyStates(ctx, managedDeviceId, optional)

Get deviceCompliancePolicyStates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***DeviceManagementManagedDevicesListDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementManagedDevicesListDeviceCompliancePolicyStatesOpts struct


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


## DeviceManagementManagedDevicesListDeviceConfigurationStates

> CollectionOfDeviceConfigurationState DeviceManagementManagedDevicesListDeviceConfigurationStates(ctx, managedDeviceId, optional)

Get deviceConfigurationStates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***DeviceManagementManagedDevicesListDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementManagedDevicesListDeviceConfigurationStatesOpts struct


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


## DeviceManagementManagedDevicesUpdateDeviceCategory

> DeviceManagementManagedDevicesUpdateDeviceCategory(ctx, managedDeviceId, microsoftGraphDeviceCategory)

Update the navigation property deviceCategory in deviceManagement

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


## DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates

> DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId, microsoftGraphDeviceCompliancePolicyState)

Update the navigation property deviceCompliancePolicyStates in deviceManagement

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


## DeviceManagementManagedDevicesUpdateDeviceConfigurationStates

> DeviceManagementManagedDevicesUpdateDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId, microsoftGraphDeviceConfigurationState)

Update the navigation property deviceConfigurationStates in deviceManagement

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


## DeviceManagementUpdateManagedDevices

> DeviceManagementUpdateManagedDevices(ctx, managedDeviceId, microsoftGraphManagedDevice)

Update the navigation property managedDevices in deviceManagement

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

