# \DeviceAppManagementMobileAppApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementCreateMobileApps) | **Post** /deviceAppManagement/mobileApps | Create new navigation property to mobileApps for deviceAppManagement
[**DeviceAppManagementGetMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementGetMobileApps) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id} | Get mobileApps from deviceAppManagement
[**DeviceAppManagementListMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementListMobileApps) | **Get** /deviceAppManagement/mobileApps | Get mobileApps from deviceAppManagement
[**DeviceAppManagementMobileAppsCreateAssignments**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsCreateAssignments) | **Post** /deviceAppManagement/mobileApps/{mobileApp-id}/assignments | Create new navigation property to assignments for deviceAppManagement
[**DeviceAppManagementMobileAppsGetAssignments**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsGetAssignments) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id}/assignments/{mobileAppAssignment-id} | Get assignments from deviceAppManagement
[**DeviceAppManagementMobileAppsGetCategories**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsGetCategories) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id}/categories/{mobileAppCategory-id} | Get categories from deviceAppManagement
[**DeviceAppManagementMobileAppsListAssignments**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsListAssignments) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id}/assignments | Get assignments from deviceAppManagement
[**DeviceAppManagementMobileAppsListCategories**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsListCategories) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id}/categories | Get categories from deviceAppManagement
[**DeviceAppManagementMobileAppsUpdateAssignments**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsUpdateAssignments) | **Patch** /deviceAppManagement/mobileApps/{mobileApp-id}/assignments/{mobileAppAssignment-id} | Update the navigation property assignments in deviceAppManagement
[**DeviceAppManagementUpdateMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementUpdateMobileApps) | **Patch** /deviceAppManagement/mobileApps/{mobileApp-id} | Update the navigation property mobileApps in deviceAppManagement



## DeviceAppManagementCreateMobileApps

> MicrosoftGraphMobileApp DeviceAppManagementCreateMobileApps(ctx, microsoftGraphMobileApp)

Create new navigation property to mobileApps for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphMobileApp** | [**MicrosoftGraphMobileApp**](MicrosoftGraphMobileApp.md)| New navigation property | 

### Return type

[**MicrosoftGraphMobileApp**](microsoft.graph.mobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetMobileApps

> MicrosoftGraphMobileApp DeviceAppManagementGetMobileApps(ctx, mobileAppId, optional)

Get mobileApps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
 **optional** | ***DeviceAppManagementGetMobileAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetMobileAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMobileApp**](microsoft.graph.mobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListMobileApps

> CollectionOfMobileApp DeviceAppManagementListMobileApps(ctx, optional)

Get mobileApps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListMobileAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListMobileAppsOpts struct


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

[**CollectionOfMobileApp**](Collection_of_mobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsCreateAssignments

> MicrosoftGraphMobileAppAssignment DeviceAppManagementMobileAppsCreateAssignments(ctx, mobileAppId, microsoftGraphMobileAppAssignment)

Create new navigation property to assignments for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**microsoftGraphMobileAppAssignment** | [**MicrosoftGraphMobileAppAssignment**](MicrosoftGraphMobileAppAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphMobileAppAssignment**](microsoft.graph.mobileAppAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsGetAssignments

> MicrosoftGraphMobileAppAssignment DeviceAppManagementMobileAppsGetAssignments(ctx, mobileAppId, mobileAppAssignmentId, optional)

Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**mobileAppAssignmentId** | **string**| key: mobileAppAssignment-id of mobileAppAssignment | 
 **optional** | ***DeviceAppManagementMobileAppsGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppsGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMobileAppAssignment**](microsoft.graph.mobileAppAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsGetCategories

> MicrosoftGraphMobileAppCategory DeviceAppManagementMobileAppsGetCategories(ctx, mobileAppId, mobileAppCategoryId, optional)

Get categories from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**mobileAppCategoryId** | **string**| key: mobileAppCategory-id of mobileAppCategory | 
 **optional** | ***DeviceAppManagementMobileAppsGetCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppsGetCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMobileAppCategory**](microsoft.graph.mobileAppCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsListAssignments

> CollectionOfMobileAppAssignment DeviceAppManagementMobileAppsListAssignments(ctx, mobileAppId, optional)

Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
 **optional** | ***DeviceAppManagementMobileAppsListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppsListAssignmentsOpts struct


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

[**CollectionOfMobileAppAssignment**](Collection_of_mobileAppAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsListCategories

> CollectionOfMobileAppCategory DeviceAppManagementMobileAppsListCategories(ctx, mobileAppId, optional)

Get categories from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
 **optional** | ***DeviceAppManagementMobileAppsListCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppsListCategoriesOpts struct


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

[**CollectionOfMobileAppCategory**](Collection_of_mobileAppCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsUpdateAssignments

> DeviceAppManagementMobileAppsUpdateAssignments(ctx, mobileAppId, mobileAppAssignmentId, microsoftGraphMobileAppAssignment)

Update the navigation property assignments in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**mobileAppAssignmentId** | **string**| key: mobileAppAssignment-id of mobileAppAssignment | 
**microsoftGraphMobileAppAssignment** | [**MicrosoftGraphMobileAppAssignment**](MicrosoftGraphMobileAppAssignment.md)| New navigation property values | 

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


## DeviceAppManagementUpdateMobileApps

> DeviceAppManagementUpdateMobileApps(ctx, mobileAppId, microsoftGraphMobileApp)

Update the navigation property mobileApps in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**microsoftGraphMobileApp** | [**MicrosoftGraphMobileApp**](MicrosoftGraphMobileApp.md)| New navigation property values | 

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

