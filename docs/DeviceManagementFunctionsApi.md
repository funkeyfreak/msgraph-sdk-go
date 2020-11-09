# \DeviceManagementFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest**](DeviceManagementFunctionsApi.md#DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest) | **Get** /deviceManagement/applePushNotificationCertificate/microsoft.graph.downloadApplePushNotificationCertificateSigningRequest() | Invoke function downloadApplePushNotificationCertificateSigningRequest
[**DeviceManagementGetEffectivePermissions**](DeviceManagementFunctionsApi.md#DeviceManagementGetEffectivePermissions) | **Get** /deviceManagement/microsoft.graph.getEffectivePermissions(scope&#x3D;{scope}) | Invoke function getEffectivePermissions
[**DeviceManagementVerifyWindowsEnrollmentAutoDiscovery**](DeviceManagementFunctionsApi.md#DeviceManagementVerifyWindowsEnrollmentAutoDiscovery) | **Get** /deviceManagement/microsoft.graph.verifyWindowsEnrollmentAutoDiscovery(domainName&#x3D;{domainName}) | Invoke function verifyWindowsEnrollmentAutoDiscovery



## DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest

> string DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest(ctx, )

Invoke function downloadApplePushNotificationCertificateSigningRequest

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetEffectivePermissions

> []interface{} DeviceManagementGetEffectivePermissions(ctx, scope)

Invoke function getEffectivePermissions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string**|  | 

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementVerifyWindowsEnrollmentAutoDiscovery

> bool DeviceManagementVerifyWindowsEnrollmentAutoDiscovery(ctx, domainName)

Invoke function verifyWindowsEnrollmentAutoDiscovery

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**|  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

