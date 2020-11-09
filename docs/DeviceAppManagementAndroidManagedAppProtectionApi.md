# \DeviceAppManagementAndroidManagedAppProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementAndroidManagedAppProtectionsCreateApps**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsCreateApps) | **Post** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/apps | Create new navigation property to apps for deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsGetApps**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsGetApps) | **Get** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/apps/{managedMobileApp-id} | Get apps from deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary) | **Get** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/deploymentSummary | Get deploymentSummary from deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsListApps**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsListApps) | **Get** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/apps | Get apps from deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsUpdateApps**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsUpdateApps) | **Patch** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/apps/{managedMobileApp-id} | Update the navigation property apps in deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary) | **Patch** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/deploymentSummary | Update the navigation property deploymentSummary in deviceAppManagement
[**DeviceAppManagementCreateAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementCreateAndroidManagedAppProtections) | **Post** /deviceAppManagement/androidManagedAppProtections | Create new navigation property to androidManagedAppProtections for deviceAppManagement
[**DeviceAppManagementGetAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementGetAndroidManagedAppProtections) | **Get** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id} | Get androidManagedAppProtections from deviceAppManagement
[**DeviceAppManagementListAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementListAndroidManagedAppProtections) | **Get** /deviceAppManagement/androidManagedAppProtections | Get androidManagedAppProtections from deviceAppManagement
[**DeviceAppManagementUpdateAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementUpdateAndroidManagedAppProtections) | **Patch** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id} | Update the navigation property androidManagedAppProtections in deviceAppManagement



## DeviceAppManagementAndroidManagedAppProtectionsCreateApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementAndroidManagedAppProtectionsCreateApps(ctx, androidManagedAppProtectionId, microsoftGraphManagedMobileApp)

Create new navigation property to apps for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
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


## DeviceAppManagementAndroidManagedAppProtectionsGetApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementAndroidManagedAppProtectionsGetApps(ctx, androidManagedAppProtectionId, managedMobileAppId, optional)

Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
**managedMobileAppId** | **string**| key: managedMobileApp-id of managedMobileApp | 
 **optional** | ***DeviceAppManagementAndroidManagedAppProtectionsGetAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementAndroidManagedAppProtectionsGetAppsOpts struct


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


## DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary

> MicrosoftGraphManagedAppPolicyDeploymentSummary DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary(ctx, androidManagedAppProtectionId, optional)

Get deploymentSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
 **optional** | ***DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummaryOpts struct


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


## DeviceAppManagementAndroidManagedAppProtectionsListApps

> CollectionOfManagedMobileApp DeviceAppManagementAndroidManagedAppProtectionsListApps(ctx, androidManagedAppProtectionId, optional)

Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
 **optional** | ***DeviceAppManagementAndroidManagedAppProtectionsListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementAndroidManagedAppProtectionsListAppsOpts struct


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


## DeviceAppManagementAndroidManagedAppProtectionsUpdateApps

> DeviceAppManagementAndroidManagedAppProtectionsUpdateApps(ctx, androidManagedAppProtectionId, managedMobileAppId, microsoftGraphManagedMobileApp)

Update the navigation property apps in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
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


## DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary

> DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary(ctx, androidManagedAppProtectionId, microsoftGraphManagedAppPolicyDeploymentSummary)

Update the navigation property deploymentSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
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


## DeviceAppManagementCreateAndroidManagedAppProtections

> MicrosoftGraphAndroidManagedAppProtection DeviceAppManagementCreateAndroidManagedAppProtections(ctx, microsoftGraphAndroidManagedAppProtection)

Create new navigation property to androidManagedAppProtections for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphAndroidManagedAppProtection** | [**MicrosoftGraphAndroidManagedAppProtection**](MicrosoftGraphAndroidManagedAppProtection.md)| New navigation property | 

### Return type

[**MicrosoftGraphAndroidManagedAppProtection**](microsoft.graph.androidManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetAndroidManagedAppProtections

> MicrosoftGraphAndroidManagedAppProtection DeviceAppManagementGetAndroidManagedAppProtections(ctx, androidManagedAppProtectionId, optional)

Get androidManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
 **optional** | ***DeviceAppManagementGetAndroidManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetAndroidManagedAppProtectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAndroidManagedAppProtection**](microsoft.graph.androidManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListAndroidManagedAppProtections

> CollectionOfAndroidManagedAppProtection DeviceAppManagementListAndroidManagedAppProtections(ctx, optional)

Get androidManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListAndroidManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListAndroidManagedAppProtectionsOpts struct


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

[**CollectionOfAndroidManagedAppProtection**](Collection_of_androidManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateAndroidManagedAppProtections

> DeviceAppManagementUpdateAndroidManagedAppProtections(ctx, androidManagedAppProtectionId, microsoftGraphAndroidManagedAppProtection)

Update the navigation property androidManagedAppProtections in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
**microsoftGraphAndroidManagedAppProtection** | [**MicrosoftGraphAndroidManagedAppProtection**](MicrosoftGraphAndroidManagedAppProtection.md)| New navigation property values | 

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

