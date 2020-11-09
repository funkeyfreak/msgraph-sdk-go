# \DeviceAppManagementMobileAppCategoryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateMobileAppCategories**](DeviceAppManagementMobileAppCategoryApi.md#DeviceAppManagementCreateMobileAppCategories) | **Post** /deviceAppManagement/mobileAppCategories | Create new navigation property to mobileAppCategories for deviceAppManagement
[**DeviceAppManagementGetMobileAppCategories**](DeviceAppManagementMobileAppCategoryApi.md#DeviceAppManagementGetMobileAppCategories) | **Get** /deviceAppManagement/mobileAppCategories/{mobileAppCategory-id} | Get mobileAppCategories from deviceAppManagement
[**DeviceAppManagementListMobileAppCategories**](DeviceAppManagementMobileAppCategoryApi.md#DeviceAppManagementListMobileAppCategories) | **Get** /deviceAppManagement/mobileAppCategories | Get mobileAppCategories from deviceAppManagement
[**DeviceAppManagementUpdateMobileAppCategories**](DeviceAppManagementMobileAppCategoryApi.md#DeviceAppManagementUpdateMobileAppCategories) | **Patch** /deviceAppManagement/mobileAppCategories/{mobileAppCategory-id} | Update the navigation property mobileAppCategories in deviceAppManagement



## DeviceAppManagementCreateMobileAppCategories

> MicrosoftGraphMobileAppCategory DeviceAppManagementCreateMobileAppCategories(ctx, microsoftGraphMobileAppCategory)

Create new navigation property to mobileAppCategories for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphMobileAppCategory** | [**MicrosoftGraphMobileAppCategory**](MicrosoftGraphMobileAppCategory.md)| New navigation property | 

### Return type

[**MicrosoftGraphMobileAppCategory**](microsoft.graph.mobileAppCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetMobileAppCategories

> MicrosoftGraphMobileAppCategory DeviceAppManagementGetMobileAppCategories(ctx, mobileAppCategoryId, optional)

Get mobileAppCategories from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppCategoryId** | **string**| key: mobileAppCategory-id of mobileAppCategory | 
 **optional** | ***DeviceAppManagementGetMobileAppCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetMobileAppCategoriesOpts struct


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


## DeviceAppManagementListMobileAppCategories

> CollectionOfMobileAppCategory DeviceAppManagementListMobileAppCategories(ctx, optional)

Get mobileAppCategories from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListMobileAppCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListMobileAppCategoriesOpts struct


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


## DeviceAppManagementUpdateMobileAppCategories

> DeviceAppManagementUpdateMobileAppCategories(ctx, mobileAppCategoryId, microsoftGraphMobileAppCategory)

Update the navigation property mobileAppCategories in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppCategoryId** | **string**| key: mobileAppCategory-id of mobileAppCategory | 
**microsoftGraphMobileAppCategory** | [**MicrosoftGraphMobileAppCategory**](MicrosoftGraphMobileAppCategory.md)| New navigation property values | 

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

