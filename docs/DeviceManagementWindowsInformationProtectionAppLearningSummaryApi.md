# \DeviceManagementWindowsInformationProtectionAppLearningSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries**](DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.md#DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries) | **Post** /deviceManagement/windowsInformationProtectionAppLearningSummaries | Create new navigation property to windowsInformationProtectionAppLearningSummaries for deviceManagement
[**DeviceManagementGetWindowsInformationProtectionAppLearningSummaries**](DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.md#DeviceManagementGetWindowsInformationProtectionAppLearningSummaries) | **Get** /deviceManagement/windowsInformationProtectionAppLearningSummaries/{windowsInformationProtectionAppLearningSummary-id} | Get windowsInformationProtectionAppLearningSummaries from deviceManagement
[**DeviceManagementListWindowsInformationProtectionAppLearningSummaries**](DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.md#DeviceManagementListWindowsInformationProtectionAppLearningSummaries) | **Get** /deviceManagement/windowsInformationProtectionAppLearningSummaries | Get windowsInformationProtectionAppLearningSummaries from deviceManagement
[**DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries**](DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.md#DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries) | **Patch** /deviceManagement/windowsInformationProtectionAppLearningSummaries/{windowsInformationProtectionAppLearningSummary-id} | Update the navigation property windowsInformationProtectionAppLearningSummaries in deviceManagement



## DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries

> MicrosoftGraphWindowsInformationProtectionAppLearningSummary DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries(ctx, microsoftGraphWindowsInformationProtectionAppLearningSummary)

Create new navigation property to windowsInformationProtectionAppLearningSummaries for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphWindowsInformationProtectionAppLearningSummary** | [**MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](MicrosoftGraphWindowsInformationProtectionAppLearningSummary.md)| New navigation property | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](microsoft.graph.windowsInformationProtectionAppLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetWindowsInformationProtectionAppLearningSummaries

> MicrosoftGraphWindowsInformationProtectionAppLearningSummary DeviceManagementGetWindowsInformationProtectionAppLearningSummaries(ctx, windowsInformationProtectionAppLearningSummaryId, optional)

Get windowsInformationProtectionAppLearningSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionAppLearningSummaryId** | **string**| key: windowsInformationProtectionAppLearningSummary-id of windowsInformationProtectionAppLearningSummary | 
 **optional** | ***DeviceManagementGetWindowsInformationProtectionAppLearningSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetWindowsInformationProtectionAppLearningSummariesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](microsoft.graph.windowsInformationProtectionAppLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListWindowsInformationProtectionAppLearningSummaries

> CollectionOfWindowsInformationProtectionAppLearningSummary DeviceManagementListWindowsInformationProtectionAppLearningSummaries(ctx, optional)

Get windowsInformationProtectionAppLearningSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListWindowsInformationProtectionAppLearningSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListWindowsInformationProtectionAppLearningSummariesOpts struct


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

[**CollectionOfWindowsInformationProtectionAppLearningSummary**](Collection_of_windowsInformationProtectionAppLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries

> DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries(ctx, windowsInformationProtectionAppLearningSummaryId, microsoftGraphWindowsInformationProtectionAppLearningSummary)

Update the navigation property windowsInformationProtectionAppLearningSummaries in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionAppLearningSummaryId** | **string**| key: windowsInformationProtectionAppLearningSummary-id of windowsInformationProtectionAppLearningSummary | 
**microsoftGraphWindowsInformationProtectionAppLearningSummary** | [**MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](MicrosoftGraphWindowsInformationProtectionAppLearningSummary.md)| New navigation property values | 

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

