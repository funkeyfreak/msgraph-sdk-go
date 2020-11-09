# \ApplicationsExtensionPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsCreateExtensionProperties**](ApplicationsExtensionPropertyApi.md#ApplicationsCreateExtensionProperties) | **Post** /applications/{application-id}/extensionProperties | Create new navigation property to extensionProperties for applications
[**ApplicationsGetExtensionProperties**](ApplicationsExtensionPropertyApi.md#ApplicationsGetExtensionProperties) | **Get** /applications/{application-id}/extensionProperties/{extensionProperty-id} | Get extensionProperties from applications
[**ApplicationsListExtensionProperties**](ApplicationsExtensionPropertyApi.md#ApplicationsListExtensionProperties) | **Get** /applications/{application-id}/extensionProperties | Get extensionProperties from applications
[**ApplicationsUpdateExtensionProperties**](ApplicationsExtensionPropertyApi.md#ApplicationsUpdateExtensionProperties) | **Patch** /applications/{application-id}/extensionProperties/{extensionProperty-id} | Update the navigation property extensionProperties in applications



## ApplicationsCreateExtensionProperties

> MicrosoftGraphExtensionProperty ApplicationsCreateExtensionProperties(ctx, applicationId, microsoftGraphExtensionProperty)

Create new navigation property to extensionProperties for applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**microsoftGraphExtensionProperty** | [**MicrosoftGraphExtensionProperty**](MicrosoftGraphExtensionProperty.md)| New navigation property | 

### Return type

[**MicrosoftGraphExtensionProperty**](microsoft.graph.extensionProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsGetExtensionProperties

> MicrosoftGraphExtensionProperty ApplicationsGetExtensionProperties(ctx, applicationId, extensionPropertyId, optional)

Get extensionProperties from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**extensionPropertyId** | **string**| key: extensionProperty-id of extensionProperty | 
 **optional** | ***ApplicationsGetExtensionPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsGetExtensionPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphExtensionProperty**](microsoft.graph.extensionProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsListExtensionProperties

> CollectionOfExtensionProperty ApplicationsListExtensionProperties(ctx, applicationId, optional)

Get extensionProperties from applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
 **optional** | ***ApplicationsListExtensionPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsListExtensionPropertiesOpts struct


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

[**CollectionOfExtensionProperty**](Collection_of_extensionProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsUpdateExtensionProperties

> ApplicationsUpdateExtensionProperties(ctx, applicationId, extensionPropertyId, microsoftGraphExtensionProperty)

Update the navigation property extensionProperties in applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**| key: application-id of application | 
**extensionPropertyId** | **string**| key: extensionProperty-id of extensionProperty | 
**microsoftGraphExtensionProperty** | [**MicrosoftGraphExtensionProperty**](MicrosoftGraphExtensionProperty.md)| New navigation property values | 

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

