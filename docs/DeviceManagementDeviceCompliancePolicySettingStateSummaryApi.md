# \DeviceManagementDeviceCompliancePolicySettingStateSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries) | **Post** /deviceManagement/deviceCompliancePolicySettingStateSummaries | Create new navigation property to deviceCompliancePolicySettingStateSummaries for deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates) | **Post** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id}/deviceComplianceSettingStates | Create new navigation property to deviceComplianceSettingStates for deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id}/deviceComplianceSettingStates/{deviceComplianceSettingState-id} | Get deviceComplianceSettingStates from deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id}/deviceComplianceSettingStates | Get deviceComplianceSettingStates from deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates) | **Patch** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id}/deviceComplianceSettingStates/{deviceComplianceSettingState-id} | Update the navigation property deviceComplianceSettingStates in deviceManagement
[**DeviceManagementGetDeviceCompliancePolicySettingStateSummaries**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementGetDeviceCompliancePolicySettingStateSummaries) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id} | Get deviceCompliancePolicySettingStateSummaries from deviceManagement
[**DeviceManagementListDeviceCompliancePolicySettingStateSummaries**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementListDeviceCompliancePolicySettingStateSummaries) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries | Get deviceCompliancePolicySettingStateSummaries from deviceManagement
[**DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries) | **Patch** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id} | Update the navigation property deviceCompliancePolicySettingStateSummaries in deviceManagement



## DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries

> MicrosoftGraphDeviceCompliancePolicySettingStateSummary DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries(ctx, microsoftGraphDeviceCompliancePolicySettingStateSummary)

Create new navigation property to deviceCompliancePolicySettingStateSummaries for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceCompliancePolicySettingStateSummary** | [**MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](MicrosoftGraphDeviceCompliancePolicySettingStateSummary.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](microsoft.graph.deviceCompliancePolicySettingStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates

> MicrosoftGraphDeviceComplianceSettingState DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, microsoftGraphDeviceComplianceSettingState)

Create new navigation property to deviceComplianceSettingStates for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
**microsoftGraphDeviceComplianceSettingState** | [**MicrosoftGraphDeviceComplianceSettingState**](MicrosoftGraphDeviceComplianceSettingState.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceSettingState**](microsoft.graph.deviceComplianceSettingState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates

> MicrosoftGraphDeviceComplianceSettingState DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId, optional)

Get deviceComplianceSettingStates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
**deviceComplianceSettingStateId** | **string**| key: deviceComplianceSettingState-id of deviceComplianceSettingState | 
 **optional** | ***DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceSettingState**](microsoft.graph.deviceComplianceSettingState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates

> CollectionOfDeviceComplianceSettingState DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, optional)

Get deviceComplianceSettingStates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
 **optional** | ***DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStatesOpts struct


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

[**CollectionOfDeviceComplianceSettingState**](Collection_of_deviceComplianceSettingState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates

> DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId, microsoftGraphDeviceComplianceSettingState)

Update the navigation property deviceComplianceSettingStates in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
**deviceComplianceSettingStateId** | **string**| key: deviceComplianceSettingState-id of deviceComplianceSettingState | 
**microsoftGraphDeviceComplianceSettingState** | [**MicrosoftGraphDeviceComplianceSettingState**](MicrosoftGraphDeviceComplianceSettingState.md)| New navigation property values | 

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


## DeviceManagementGetDeviceCompliancePolicySettingStateSummaries

> MicrosoftGraphDeviceCompliancePolicySettingStateSummary DeviceManagementGetDeviceCompliancePolicySettingStateSummaries(ctx, deviceCompliancePolicySettingStateSummaryId, optional)

Get deviceCompliancePolicySettingStateSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
 **optional** | ***DeviceManagementGetDeviceCompliancePolicySettingStateSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDeviceCompliancePolicySettingStateSummariesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](microsoft.graph.deviceCompliancePolicySettingStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceCompliancePolicySettingStateSummaries

> CollectionOfDeviceCompliancePolicySettingStateSummary DeviceManagementListDeviceCompliancePolicySettingStateSummaries(ctx, optional)

Get deviceCompliancePolicySettingStateSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListDeviceCompliancePolicySettingStateSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListDeviceCompliancePolicySettingStateSummariesOpts struct


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

[**CollectionOfDeviceCompliancePolicySettingStateSummary**](Collection_of_deviceCompliancePolicySettingStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries

> DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries(ctx, deviceCompliancePolicySettingStateSummaryId, microsoftGraphDeviceCompliancePolicySettingStateSummary)

Update the navigation property deviceCompliancePolicySettingStateSummaries in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
**microsoftGraphDeviceCompliancePolicySettingStateSummary** | [**MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](MicrosoftGraphDeviceCompliancePolicySettingStateSummary.md)| New navigation property values | 

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

