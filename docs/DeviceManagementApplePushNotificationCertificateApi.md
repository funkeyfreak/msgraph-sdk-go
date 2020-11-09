# \DeviceManagementApplePushNotificationCertificateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementGetApplePushNotificationCertificate**](DeviceManagementApplePushNotificationCertificateApi.md#DeviceManagementGetApplePushNotificationCertificate) | **Get** /deviceManagement/applePushNotificationCertificate | Get applePushNotificationCertificate from deviceManagement
[**DeviceManagementUpdateApplePushNotificationCertificate**](DeviceManagementApplePushNotificationCertificateApi.md#DeviceManagementUpdateApplePushNotificationCertificate) | **Patch** /deviceManagement/applePushNotificationCertificate | Update the navigation property applePushNotificationCertificate in deviceManagement



## DeviceManagementGetApplePushNotificationCertificate

> MicrosoftGraphApplePushNotificationCertificate DeviceManagementGetApplePushNotificationCertificate(ctx, optional)

Get applePushNotificationCertificate from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementGetApplePushNotificationCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetApplePushNotificationCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphApplePushNotificationCertificate**](microsoft.graph.applePushNotificationCertificate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateApplePushNotificationCertificate

> DeviceManagementUpdateApplePushNotificationCertificate(ctx, microsoftGraphApplePushNotificationCertificate)

Update the navigation property applePushNotificationCertificate in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphApplePushNotificationCertificate** | [**MicrosoftGraphApplePushNotificationCertificate**](MicrosoftGraphApplePushNotificationCertificate.md)| New navigation property values | 

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

