# \DeviceAppManagementMdmWindowsInformationProtectionPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies**](DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.md#DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies) | **Post** /deviceAppManagement/mdmWindowsInformationProtectionPolicies | Create new navigation property to mdmWindowsInformationProtectionPolicies for deviceAppManagement
[**DeviceAppManagementGetMdmWindowsInformationProtectionPolicies**](DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.md#DeviceAppManagementGetMdmWindowsInformationProtectionPolicies) | **Get** /deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy-id} | Get mdmWindowsInformationProtectionPolicies from deviceAppManagement
[**DeviceAppManagementListMdmWindowsInformationProtectionPolicies**](DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.md#DeviceAppManagementListMdmWindowsInformationProtectionPolicies) | **Get** /deviceAppManagement/mdmWindowsInformationProtectionPolicies | Get mdmWindowsInformationProtectionPolicies from deviceAppManagement
[**DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies**](DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.md#DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies) | **Patch** /deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy-id} | Update the navigation property mdmWindowsInformationProtectionPolicies in deviceAppManagement



## DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies

> MicrosoftGraphMdmWindowsInformationProtectionPolicy DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies(ctx, microsoftGraphMdmWindowsInformationProtectionPolicy)

Create new navigation property to mdmWindowsInformationProtectionPolicies for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphMdmWindowsInformationProtectionPolicy** | [**MicrosoftGraphMdmWindowsInformationProtectionPolicy**](MicrosoftGraphMdmWindowsInformationProtectionPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphMdmWindowsInformationProtectionPolicy**](microsoft.graph.mdmWindowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetMdmWindowsInformationProtectionPolicies

> MicrosoftGraphMdmWindowsInformationProtectionPolicy DeviceAppManagementGetMdmWindowsInformationProtectionPolicies(ctx, mdmWindowsInformationProtectionPolicyId, optional)

Get mdmWindowsInformationProtectionPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mdmWindowsInformationProtectionPolicyId** | **string**| key: mdmWindowsInformationProtectionPolicy-id of mdmWindowsInformationProtectionPolicy | 
 **optional** | ***DeviceAppManagementGetMdmWindowsInformationProtectionPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetMdmWindowsInformationProtectionPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMdmWindowsInformationProtectionPolicy**](microsoft.graph.mdmWindowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListMdmWindowsInformationProtectionPolicies

> CollectionOfMdmWindowsInformationProtectionPolicy DeviceAppManagementListMdmWindowsInformationProtectionPolicies(ctx, optional)

Get mdmWindowsInformationProtectionPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListMdmWindowsInformationProtectionPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListMdmWindowsInformationProtectionPoliciesOpts struct


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

[**CollectionOfMdmWindowsInformationProtectionPolicy**](Collection_of_mdmWindowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies

> DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies(ctx, mdmWindowsInformationProtectionPolicyId, microsoftGraphMdmWindowsInformationProtectionPolicy)

Update the navigation property mdmWindowsInformationProtectionPolicies in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mdmWindowsInformationProtectionPolicyId** | **string**| key: mdmWindowsInformationProtectionPolicy-id of mdmWindowsInformationProtectionPolicy | 
**microsoftGraphMdmWindowsInformationProtectionPolicy** | [**MicrosoftGraphMdmWindowsInformationProtectionPolicy**](MicrosoftGraphMdmWindowsInformationProtectionPolicy.md)| New navigation property values | 

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

