# \DeviceManagementSoftwareUpdateStatusSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementGetSoftwareUpdateStatusSummary**](DeviceManagementSoftwareUpdateStatusSummaryApi.md#DeviceManagementGetSoftwareUpdateStatusSummary) | **Get** /deviceManagement/softwareUpdateStatusSummary | Get softwareUpdateStatusSummary from deviceManagement



## DeviceManagementGetSoftwareUpdateStatusSummary

> MicrosoftGraphSoftwareUpdateStatusSummary DeviceManagementGetSoftwareUpdateStatusSummary(ctx, optional)

Get softwareUpdateStatusSummary from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementGetSoftwareUpdateStatusSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetSoftwareUpdateStatusSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSoftwareUpdateStatusSummary**](microsoft.graph.softwareUpdateStatusSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

