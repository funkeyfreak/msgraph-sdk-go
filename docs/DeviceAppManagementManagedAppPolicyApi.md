# \DeviceAppManagementManagedAppPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateManagedAppPolicies**](DeviceAppManagementManagedAppPolicyApi.md#DeviceAppManagementCreateManagedAppPolicies) | **Post** /deviceAppManagement/managedAppPolicies | Create new navigation property to managedAppPolicies for deviceAppManagement
[**DeviceAppManagementGetManagedAppPolicies**](DeviceAppManagementManagedAppPolicyApi.md#DeviceAppManagementGetManagedAppPolicies) | **Get** /deviceAppManagement/managedAppPolicies/{managedAppPolicy-id} | Get managedAppPolicies from deviceAppManagement
[**DeviceAppManagementListManagedAppPolicies**](DeviceAppManagementManagedAppPolicyApi.md#DeviceAppManagementListManagedAppPolicies) | **Get** /deviceAppManagement/managedAppPolicies | Get managedAppPolicies from deviceAppManagement
[**DeviceAppManagementUpdateManagedAppPolicies**](DeviceAppManagementManagedAppPolicyApi.md#DeviceAppManagementUpdateManagedAppPolicies) | **Patch** /deviceAppManagement/managedAppPolicies/{managedAppPolicy-id} | Update the navigation property managedAppPolicies in deviceAppManagement



## DeviceAppManagementCreateManagedAppPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementCreateManagedAppPolicies(ctx, microsoftGraphManagedAppPolicy)

Create new navigation property to managedAppPolicies for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetManagedAppPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementGetManagedAppPolicies(ctx, managedAppPolicyId, optional)

Get managedAppPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
 **optional** | ***DeviceAppManagementGetManagedAppPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetManagedAppPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListManagedAppPolicies

> CollectionOfManagedAppPolicy DeviceAppManagementListManagedAppPolicies(ctx, optional)

Get managedAppPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListManagedAppPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListManagedAppPoliciesOpts struct


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

[**CollectionOfManagedAppPolicy**](Collection_of_managedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateManagedAppPolicies

> DeviceAppManagementUpdateManagedAppPolicies(ctx, managedAppPolicyId, microsoftGraphManagedAppPolicy)

Update the navigation property managedAppPolicies in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property values | 

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

