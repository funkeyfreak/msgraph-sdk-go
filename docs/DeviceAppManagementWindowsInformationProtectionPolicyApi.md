# \DeviceAppManagementWindowsInformationProtectionPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateWindowsInformationProtectionPolicies**](DeviceAppManagementWindowsInformationProtectionPolicyApi.md#DeviceAppManagementCreateWindowsInformationProtectionPolicies) | **Post** /deviceAppManagement/windowsInformationProtectionPolicies | Create new navigation property to windowsInformationProtectionPolicies for deviceAppManagement
[**DeviceAppManagementGetWindowsInformationProtectionPolicies**](DeviceAppManagementWindowsInformationProtectionPolicyApi.md#DeviceAppManagementGetWindowsInformationProtectionPolicies) | **Get** /deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy-id} | Get windowsInformationProtectionPolicies from deviceAppManagement
[**DeviceAppManagementListWindowsInformationProtectionPolicies**](DeviceAppManagementWindowsInformationProtectionPolicyApi.md#DeviceAppManagementListWindowsInformationProtectionPolicies) | **Get** /deviceAppManagement/windowsInformationProtectionPolicies | Get windowsInformationProtectionPolicies from deviceAppManagement
[**DeviceAppManagementUpdateWindowsInformationProtectionPolicies**](DeviceAppManagementWindowsInformationProtectionPolicyApi.md#DeviceAppManagementUpdateWindowsInformationProtectionPolicies) | **Patch** /deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy-id} | Update the navigation property windowsInformationProtectionPolicies in deviceAppManagement



## DeviceAppManagementCreateWindowsInformationProtectionPolicies

> MicrosoftGraphWindowsInformationProtectionPolicy DeviceAppManagementCreateWindowsInformationProtectionPolicies(ctx, microsoftGraphWindowsInformationProtectionPolicy)

Create new navigation property to windowsInformationProtectionPolicies for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphWindowsInformationProtectionPolicy** | [**MicrosoftGraphWindowsInformationProtectionPolicy**](MicrosoftGraphWindowsInformationProtectionPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionPolicy**](microsoft.graph.windowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetWindowsInformationProtectionPolicies

> MicrosoftGraphWindowsInformationProtectionPolicy DeviceAppManagementGetWindowsInformationProtectionPolicies(ctx, windowsInformationProtectionPolicyId, optional)

Get windowsInformationProtectionPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionPolicyId** | **string**| key: windowsInformationProtectionPolicy-id of windowsInformationProtectionPolicy | 
 **optional** | ***DeviceAppManagementGetWindowsInformationProtectionPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetWindowsInformationProtectionPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionPolicy**](microsoft.graph.windowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListWindowsInformationProtectionPolicies

> CollectionOfWindowsInformationProtectionPolicy DeviceAppManagementListWindowsInformationProtectionPolicies(ctx, optional)

Get windowsInformationProtectionPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListWindowsInformationProtectionPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListWindowsInformationProtectionPoliciesOpts struct


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

[**CollectionOfWindowsInformationProtectionPolicy**](Collection_of_windowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateWindowsInformationProtectionPolicies

> DeviceAppManagementUpdateWindowsInformationProtectionPolicies(ctx, windowsInformationProtectionPolicyId, microsoftGraphWindowsInformationProtectionPolicy)

Update the navigation property windowsInformationProtectionPolicies in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionPolicyId** | **string**| key: windowsInformationProtectionPolicy-id of windowsInformationProtectionPolicy | 
**microsoftGraphWindowsInformationProtectionPolicy** | [**MicrosoftGraphWindowsInformationProtectionPolicy**](MicrosoftGraphWindowsInformationProtectionPolicy.md)| New navigation property values | 

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

