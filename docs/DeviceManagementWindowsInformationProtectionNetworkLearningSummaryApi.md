# \DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries**](DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.md#DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries) | **Post** /deviceManagement/windowsInformationProtectionNetworkLearningSummaries | Create new navigation property to windowsInformationProtectionNetworkLearningSummaries for deviceManagement
[**DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries**](DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.md#DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries) | **Get** /deviceManagement/windowsInformationProtectionNetworkLearningSummaries/{windowsInformationProtectionNetworkLearningSummary-id} | Get windowsInformationProtectionNetworkLearningSummaries from deviceManagement
[**DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries**](DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.md#DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries) | **Get** /deviceManagement/windowsInformationProtectionNetworkLearningSummaries | Get windowsInformationProtectionNetworkLearningSummaries from deviceManagement
[**DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries**](DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.md#DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries) | **Patch** /deviceManagement/windowsInformationProtectionNetworkLearningSummaries/{windowsInformationProtectionNetworkLearningSummary-id} | Update the navigation property windowsInformationProtectionNetworkLearningSummaries in deviceManagement



## DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries

> MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries(ctx, microsoftGraphWindowsInformationProtectionNetworkLearningSummary)

Create new navigation property to windowsInformationProtectionNetworkLearningSummaries for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphWindowsInformationProtectionNetworkLearningSummary** | [**MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary.md)| New navigation property | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](microsoft.graph.windowsInformationProtectionNetworkLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries

> MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries(ctx, windowsInformationProtectionNetworkLearningSummaryId, optional)

Get windowsInformationProtectionNetworkLearningSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionNetworkLearningSummaryId** | **string**| key: windowsInformationProtectionNetworkLearningSummary-id of windowsInformationProtectionNetworkLearningSummary | 
 **optional** | ***DeviceManagementGetWindowsInformationProtectionNetworkLearningSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetWindowsInformationProtectionNetworkLearningSummariesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](microsoft.graph.windowsInformationProtectionNetworkLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries

> CollectionOfWindowsInformationProtectionNetworkLearningSummary DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries(ctx, optional)

Get windowsInformationProtectionNetworkLearningSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListWindowsInformationProtectionNetworkLearningSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListWindowsInformationProtectionNetworkLearningSummariesOpts struct


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

[**CollectionOfWindowsInformationProtectionNetworkLearningSummary**](Collection_of_windowsInformationProtectionNetworkLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries

> DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries(ctx, windowsInformationProtectionNetworkLearningSummaryId, microsoftGraphWindowsInformationProtectionNetworkLearningSummary)

Update the navigation property windowsInformationProtectionNetworkLearningSummaries in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionNetworkLearningSummaryId** | **string**| key: windowsInformationProtectionNetworkLearningSummary-id of windowsInformationProtectionNetworkLearningSummary | 
**microsoftGraphWindowsInformationProtectionNetworkLearningSummary** | [**MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary.md)| New navigation property values | 

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

