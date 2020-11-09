# \DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary**](DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.md#DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary) | **Get** /deviceManagement/deviceCompliancePolicyDeviceStateSummary | Get deviceCompliancePolicyDeviceStateSummary from deviceManagement
[**DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary**](DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.md#DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary) | **Patch** /deviceManagement/deviceCompliancePolicyDeviceStateSummary | Update the navigation property deviceCompliancePolicyDeviceStateSummary in deviceManagement



## DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary

> MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary(ctx, optional)

Get deviceCompliancePolicyDeviceStateSummary from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary**](microsoft.graph.deviceCompliancePolicyDeviceStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary

> DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary(ctx, microsoftGraphDeviceCompliancePolicyDeviceStateSummary)

Update the navigation property deviceCompliancePolicyDeviceStateSummary in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceCompliancePolicyDeviceStateSummary** | [**MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary**](MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary.md)| New navigation property values | 

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

