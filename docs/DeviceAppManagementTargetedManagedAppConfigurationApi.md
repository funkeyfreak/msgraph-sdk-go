# \DeviceAppManagementTargetedManagedAppConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateTargetedManagedAppConfigurations**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementCreateTargetedManagedAppConfigurations) | **Post** /deviceAppManagement/targetedManagedAppConfigurations | Create new navigation property to targetedManagedAppConfigurations for deviceAppManagement
[**DeviceAppManagementGetTargetedManagedAppConfigurations**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementGetTargetedManagedAppConfigurations) | **Get** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id} | Get targetedManagedAppConfigurations from deviceAppManagement
[**DeviceAppManagementListTargetedManagedAppConfigurations**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementListTargetedManagedAppConfigurations) | **Get** /deviceAppManagement/targetedManagedAppConfigurations | Get targetedManagedAppConfigurations from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsCreateApps**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsCreateApps) | **Post** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/apps | Create new navigation property to apps for deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsCreateAssignments**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsCreateAssignments) | **Post** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/assignments | Create new navigation property to assignments for deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsGetApps**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsGetApps) | **Get** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/apps/{managedMobileApp-id} | Get apps from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsGetAssignments**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsGetAssignments) | **Get** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/assignments/{targetedManagedAppPolicyAssignment-id} | Get assignments from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummary**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummary) | **Get** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/deploymentSummary | Get deploymentSummary from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsListApps**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsListApps) | **Get** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/apps | Get apps from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsListAssignments**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsListAssignments) | **Get** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/assignments | Get assignments from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsUpdateApps**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsUpdateApps) | **Patch** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/apps/{managedMobileApp-id} | Update the navigation property apps in deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsUpdateAssignments**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsUpdateAssignments) | **Patch** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/assignments/{targetedManagedAppPolicyAssignment-id} | Update the navigation property assignments in deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsUpdateDeploymentSummary**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementTargetedManagedAppConfigurationsUpdateDeploymentSummary) | **Patch** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id}/deploymentSummary | Update the navigation property deploymentSummary in deviceAppManagement
[**DeviceAppManagementUpdateTargetedManagedAppConfigurations**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementUpdateTargetedManagedAppConfigurations) | **Patch** /deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration-id} | Update the navigation property targetedManagedAppConfigurations in deviceAppManagement



## DeviceAppManagementCreateTargetedManagedAppConfigurations

> MicrosoftGraphTargetedManagedAppConfiguration DeviceAppManagementCreateTargetedManagedAppConfigurations(ctx, microsoftGraphTargetedManagedAppConfiguration)

Create new navigation property to targetedManagedAppConfigurations for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTargetedManagedAppConfiguration** | [**MicrosoftGraphTargetedManagedAppConfiguration**](MicrosoftGraphTargetedManagedAppConfiguration.md)| New navigation property | 

### Return type

[**MicrosoftGraphTargetedManagedAppConfiguration**](microsoft.graph.targetedManagedAppConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetTargetedManagedAppConfigurations

> MicrosoftGraphTargetedManagedAppConfiguration DeviceAppManagementGetTargetedManagedAppConfigurations(ctx, targetedManagedAppConfigurationId, optional)

Get targetedManagedAppConfigurations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
 **optional** | ***DeviceAppManagementGetTargetedManagedAppConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetTargetedManagedAppConfigurationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTargetedManagedAppConfiguration**](microsoft.graph.targetedManagedAppConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListTargetedManagedAppConfigurations

> CollectionOfTargetedManagedAppConfiguration DeviceAppManagementListTargetedManagedAppConfigurations(ctx, optional)

Get targetedManagedAppConfigurations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListTargetedManagedAppConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListTargetedManagedAppConfigurationsOpts struct


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

[**CollectionOfTargetedManagedAppConfiguration**](Collection_of_targetedManagedAppConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsCreateApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementTargetedManagedAppConfigurationsCreateApps(ctx, targetedManagedAppConfigurationId, microsoftGraphManagedMobileApp)

Create new navigation property to apps for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
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


## DeviceAppManagementTargetedManagedAppConfigurationsCreateAssignments

> MicrosoftGraphTargetedManagedAppPolicyAssignment DeviceAppManagementTargetedManagedAppConfigurationsCreateAssignments(ctx, targetedManagedAppConfigurationId, microsoftGraphTargetedManagedAppPolicyAssignment)

Create new navigation property to assignments for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**microsoftGraphTargetedManagedAppPolicyAssignment** | [**MicrosoftGraphTargetedManagedAppPolicyAssignment**](MicrosoftGraphTargetedManagedAppPolicyAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphTargetedManagedAppPolicyAssignment**](microsoft.graph.targetedManagedAppPolicyAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsGetApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementTargetedManagedAppConfigurationsGetApps(ctx, targetedManagedAppConfigurationId, managedMobileAppId, optional)

Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**managedMobileAppId** | **string**| key: managedMobileApp-id of managedMobileApp | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsGetAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsGetAppsOpts struct


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


## DeviceAppManagementTargetedManagedAppConfigurationsGetAssignments

> MicrosoftGraphTargetedManagedAppPolicyAssignment DeviceAppManagementTargetedManagedAppConfigurationsGetAssignments(ctx, targetedManagedAppConfigurationId, targetedManagedAppPolicyAssignmentId, optional)

Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**targetedManagedAppPolicyAssignmentId** | **string**| key: targetedManagedAppPolicyAssignment-id of targetedManagedAppPolicyAssignment | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTargetedManagedAppPolicyAssignment**](microsoft.graph.targetedManagedAppPolicyAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummary

> MicrosoftGraphManagedAppPolicyDeploymentSummary DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummary(ctx, targetedManagedAppConfigurationId, optional)

Get deploymentSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummaryOpts struct


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


## DeviceAppManagementTargetedManagedAppConfigurationsListApps

> CollectionOfManagedMobileApp DeviceAppManagementTargetedManagedAppConfigurationsListApps(ctx, targetedManagedAppConfigurationId, optional)

Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsListAppsOpts struct


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


## DeviceAppManagementTargetedManagedAppConfigurationsListAssignments

> CollectionOfTargetedManagedAppPolicyAssignment DeviceAppManagementTargetedManagedAppConfigurationsListAssignments(ctx, targetedManagedAppConfigurationId, optional)

Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsListAssignmentsOpts struct


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

[**CollectionOfTargetedManagedAppPolicyAssignment**](Collection_of_targetedManagedAppPolicyAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsUpdateApps

> DeviceAppManagementTargetedManagedAppConfigurationsUpdateApps(ctx, targetedManagedAppConfigurationId, managedMobileAppId, microsoftGraphManagedMobileApp)

Update the navigation property apps in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
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


## DeviceAppManagementTargetedManagedAppConfigurationsUpdateAssignments

> DeviceAppManagementTargetedManagedAppConfigurationsUpdateAssignments(ctx, targetedManagedAppConfigurationId, targetedManagedAppPolicyAssignmentId, microsoftGraphTargetedManagedAppPolicyAssignment)

Update the navigation property assignments in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**targetedManagedAppPolicyAssignmentId** | **string**| key: targetedManagedAppPolicyAssignment-id of targetedManagedAppPolicyAssignment | 
**microsoftGraphTargetedManagedAppPolicyAssignment** | [**MicrosoftGraphTargetedManagedAppPolicyAssignment**](MicrosoftGraphTargetedManagedAppPolicyAssignment.md)| New navigation property values | 

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


## DeviceAppManagementTargetedManagedAppConfigurationsUpdateDeploymentSummary

> DeviceAppManagementTargetedManagedAppConfigurationsUpdateDeploymentSummary(ctx, targetedManagedAppConfigurationId, microsoftGraphManagedAppPolicyDeploymentSummary)

Update the navigation property deploymentSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
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


## DeviceAppManagementUpdateTargetedManagedAppConfigurations

> DeviceAppManagementUpdateTargetedManagedAppConfigurations(ctx, targetedManagedAppConfigurationId, microsoftGraphTargetedManagedAppConfiguration)

Update the navigation property targetedManagedAppConfigurations in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**microsoftGraphTargetedManagedAppConfiguration** | [**MicrosoftGraphTargetedManagedAppConfiguration**](MicrosoftGraphTargetedManagedAppConfiguration.md)| New navigation property values | 

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

