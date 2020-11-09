# \DeviceAppManagementActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementManagedAppPoliciesTargetApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementManagedAppPoliciesTargetApps) | **Post** /deviceAppManagement/managedAppPolicies/{managedAppPolicy-id}/microsoft.graph.targetApps | Invoke action targetApps
[**DeviceAppManagementManagedAppRegistrationsAppliedPoliciesTargetApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementManagedAppRegistrationsAppliedPoliciesTargetApps) | **Post** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/appliedPolicies/{managedAppPolicy-id}/microsoft.graph.targetApps | Invoke action targetApps
[**DeviceAppManagementManagedAppRegistrationsIntendedPoliciesTargetApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementManagedAppRegistrationsIntendedPoliciesTargetApps) | **Post** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/intendedPolicies/{managedAppPolicy-id}/microsoft.graph.targetApps | Invoke action targetApps
[**DeviceAppManagementManagedEBooksAssign**](DeviceAppManagementActionsApi.md#DeviceAppManagementManagedEBooksAssign) | **Post** /deviceAppManagement/managedEBooks/{managedEBook-id}/microsoft.graph.assign | Invoke action assign
[**DeviceAppManagementMobileAppConfigurationsAssign**](DeviceAppManagementActionsApi.md#DeviceAppManagementMobileAppConfigurationsAssign) | **Post** /deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration-id}/microsoft.graph.assign | Invoke action assign
[**DeviceAppManagementMobileAppsAssign**](DeviceAppManagementActionsApi.md#DeviceAppManagementMobileAppsAssign) | **Post** /deviceAppManagement/mobileApps/{mobileApp-id}/microsoft.graph.assign | Invoke action assign
[**DeviceAppManagementSyncMicrosoftStoreForBusinessApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementSyncMicrosoftStoreForBusinessApps) | **Post** /deviceAppManagement/microsoft.graph.syncMicrosoftStoreForBusinessApps | Invoke action syncMicrosoftStoreForBusinessApps
[**DeviceAppManagementTargetedManagedAppConfigurationsAssign**](DeviceAppManagementActionsApi.md#DeviceAppManagementTargetedManagedAppConfigurationsAssign) | **Post** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/microsoft.graph.assign | Invoke action assign
[**DeviceAppManagementTargetedManagedAppConfigurationsTargetApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementTargetedManagedAppConfigurationsTargetApps) | **Post** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/microsoft.graph.targetApps | Invoke action targetApps
[**DeviceAppManagementVppTokensSyncLicenses**](DeviceAppManagementActionsApi.md#DeviceAppManagementVppTokensSyncLicenses) | **Post** /deviceAppManagement/vppTokens/{vppToken-id}/microsoft.graph.syncLicenses | Invoke action syncLicenses



## DeviceAppManagementManagedAppPoliciesTargetApps

> DeviceAppManagementManagedAppPoliciesTargetApps(ctx, managedAppPolicyId, inlineObject20)

Invoke action targetApps

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**inlineObject20** | [**InlineObject20**](InlineObject20.md)|  | 

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


## DeviceAppManagementManagedAppRegistrationsAppliedPoliciesTargetApps

> DeviceAppManagementManagedAppRegistrationsAppliedPoliciesTargetApps(ctx, managedAppRegistrationId, managedAppPolicyId, inlineObject21)

Invoke action targetApps

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**inlineObject21** | [**InlineObject21**](InlineObject21.md)|  | 

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


## DeviceAppManagementManagedAppRegistrationsIntendedPoliciesTargetApps

> DeviceAppManagementManagedAppRegistrationsIntendedPoliciesTargetApps(ctx, managedAppRegistrationId, managedAppPolicyId, inlineObject22)

Invoke action targetApps

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**inlineObject22** | [**InlineObject22**](InlineObject22.md)|  | 

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


## DeviceAppManagementManagedEBooksAssign

> DeviceAppManagementManagedEBooksAssign(ctx, managedEBookId, inlineObject23)

Invoke action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**inlineObject23** | [**InlineObject23**](InlineObject23.md)|  | 

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


## DeviceAppManagementMobileAppConfigurationsAssign

> DeviceAppManagementMobileAppConfigurationsAssign(ctx, managedDeviceMobileAppConfigurationId, inlineObject24)

Invoke action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**inlineObject24** | [**InlineObject24**](InlineObject24.md)|  | 

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


## DeviceAppManagementMobileAppsAssign

> DeviceAppManagementMobileAppsAssign(ctx, mobileAppId, inlineObject25)

Invoke action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**inlineObject25** | [**InlineObject25**](InlineObject25.md)|  | 

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


## DeviceAppManagementSyncMicrosoftStoreForBusinessApps

> DeviceAppManagementSyncMicrosoftStoreForBusinessApps(ctx, )

Invoke action syncMicrosoftStoreForBusinessApps

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsAssign

> DeviceAppManagementTargetedManagedAppConfigurationsAssign(ctx, targetedManagedAppConfigurationId, inlineObject26)

Invoke action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**inlineObject26** | [**InlineObject26**](InlineObject26.md)|  | 

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


## DeviceAppManagementTargetedManagedAppConfigurationsTargetApps

> DeviceAppManagementTargetedManagedAppConfigurationsTargetApps(ctx, targetedManagedAppConfigurationId, inlineObject27)

Invoke action targetApps

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**inlineObject27** | [**InlineObject27**](InlineObject27.md)|  | 

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


## DeviceAppManagementVppTokensSyncLicenses

> interface{} DeviceAppManagementVppTokensSyncLicenses(ctx, vppTokenId)

Invoke action syncLicenses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vppTokenId** | **string**| key: vppToken-id of vppToken | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

