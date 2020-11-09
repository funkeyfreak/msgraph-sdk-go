# \DeviceAppManagementIosManagedAppProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementCreateIosManagedAppProtections) | **Post** /deviceAppManagement/iosManagedAppProtections | Create new navigation property to iosManagedAppProtections for deviceAppManagement
[**DeviceAppManagementGetIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementGetIosManagedAppProtections) | **Get** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id} | Get iosManagedAppProtections from deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsCreateApps**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsCreateApps) | **Post** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/apps | Create new navigation property to apps for deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsGetApps**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsGetApps) | **Get** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/apps/{managedMobileApp-id} | Get apps from deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary) | **Get** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/deploymentSummary | Get deploymentSummary from deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsListApps**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsListApps) | **Get** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/apps | Get apps from deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsUpdateApps**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsUpdateApps) | **Patch** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/apps/{managedMobileApp-id} | Update the navigation property apps in deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary) | **Patch** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/deploymentSummary | Update the navigation property deploymentSummary in deviceAppManagement
[**DeviceAppManagementListIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementListIosManagedAppProtections) | **Get** /deviceAppManagement/iosManagedAppProtections | Get iosManagedAppProtections from deviceAppManagement
[**DeviceAppManagementUpdateIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementUpdateIosManagedAppProtections) | **Patch** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id} | Update the navigation property iosManagedAppProtections in deviceAppManagement



## DeviceAppManagementCreateIosManagedAppProtections

> MicrosoftGraphIosManagedAppProtection DeviceAppManagementCreateIosManagedAppProtections(ctx, microsoftGraphIosManagedAppProtection)

Create new navigation property to iosManagedAppProtections for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphIosManagedAppProtection** | [**MicrosoftGraphIosManagedAppProtection**](MicrosoftGraphIosManagedAppProtection.md)| New navigation property | 

### Return type

[**MicrosoftGraphIosManagedAppProtection**](microsoft.graph.iosManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetIosManagedAppProtections

> MicrosoftGraphIosManagedAppProtection DeviceAppManagementGetIosManagedAppProtections(ctx, iosManagedAppProtectionId, optional)

Get iosManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
 **optional** | ***DeviceAppManagementGetIosManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetIosManagedAppProtectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphIosManagedAppProtection**](microsoft.graph.iosManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementIosManagedAppProtectionsCreateApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementIosManagedAppProtectionsCreateApps(ctx, iosManagedAppProtectionId, microsoftGraphManagedMobileApp)

Create new navigation property to apps for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
**microsoftGraphManagedMobileApp** | [**MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementIosManagedAppProtectionsGetApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementIosManagedAppProtectionsGetApps(ctx, iosManagedAppProtectionId, managedMobileAppId, optional)

Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
**managedMobileAppId** | **string**| key: managedMobileApp-id of managedMobileApp | 
 **optional** | ***DeviceAppManagementIosManagedAppProtectionsGetAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementIosManagedAppProtectionsGetAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary

> MicrosoftGraphManagedAppPolicyDeploymentSummary DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary(ctx, iosManagedAppProtectionId, optional)

Get deploymentSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
 **optional** | ***DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppPolicyDeploymentSummary**](microsoft.graph.managedAppPolicyDeploymentSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementIosManagedAppProtectionsListApps

> CollectionOfManagedMobileApp DeviceAppManagementIosManagedAppProtectionsListApps(ctx, iosManagedAppProtectionId, optional)

Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
 **optional** | ***DeviceAppManagementIosManagedAppProtectionsListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementIosManagedAppProtectionsListAppsOpts struct


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

[**CollectionOfManagedMobileApp**](Collection_of_managedMobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementIosManagedAppProtectionsUpdateApps

> DeviceAppManagementIosManagedAppProtectionsUpdateApps(ctx, iosManagedAppProtectionId, managedMobileAppId, microsoftGraphManagedMobileApp)

Update the navigation property apps in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
**managedMobileAppId** | **string**| key: managedMobileApp-id of managedMobileApp | 
**microsoftGraphManagedMobileApp** | [**MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md)| New navigation property values | 

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


## DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary

> DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary(ctx, iosManagedAppProtectionId, microsoftGraphManagedAppPolicyDeploymentSummary)

Update the navigation property deploymentSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
**microsoftGraphManagedAppPolicyDeploymentSummary** | [**MicrosoftGraphManagedAppPolicyDeploymentSummary**](MicrosoftGraphManagedAppPolicyDeploymentSummary.md)| New navigation property values | 

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


## DeviceAppManagementListIosManagedAppProtections

> CollectionOfIosManagedAppProtection DeviceAppManagementListIosManagedAppProtections(ctx, optional)

Get iosManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListIosManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListIosManagedAppProtectionsOpts struct


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

[**CollectionOfIosManagedAppProtection**](Collection_of_iosManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateIosManagedAppProtections

> DeviceAppManagementUpdateIosManagedAppProtections(ctx, iosManagedAppProtectionId, microsoftGraphIosManagedAppProtection)

Update the navigation property iosManagedAppProtections in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
**microsoftGraphIosManagedAppProtection** | [**MicrosoftGraphIosManagedAppProtection**](MicrosoftGraphIosManagedAppProtection.md)| New navigation property values | 

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

