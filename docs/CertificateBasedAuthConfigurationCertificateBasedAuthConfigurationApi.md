# \CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration) | **Post** /certificateBasedAuthConfiguration | Add new entity to certificateBasedAuthConfiguration
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration) | **Delete** /certificateBasedAuthConfiguration/{certificateBasedAuthConfiguration-id} | Delete entity from certificateBasedAuthConfiguration
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration) | **Get** /certificateBasedAuthConfiguration/{certificateBasedAuthConfiguration-id} | Get entity from certificateBasedAuthConfiguration by key
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration) | **Get** /certificateBasedAuthConfiguration | Get entities from certificateBasedAuthConfiguration
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration) | **Patch** /certificateBasedAuthConfiguration/{certificateBasedAuthConfiguration-id} | Update entity in certificateBasedAuthConfiguration



## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration

> MicrosoftGraphCertificateBasedAuthConfiguration CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration(ctx, microsoftGraphCertificateBasedAuthConfiguration)

Add new entity to certificateBasedAuthConfiguration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphCertificateBasedAuthConfiguration** | [**MicrosoftGraphCertificateBasedAuthConfiguration**](MicrosoftGraphCertificateBasedAuthConfiguration.md)| New entity | 

### Return type

[**MicrosoftGraphCertificateBasedAuthConfiguration**](microsoft.graph.certificateBasedAuthConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration

> CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration(ctx, certificateBasedAuthConfigurationId, optional)

Delete entity from certificateBasedAuthConfiguration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateBasedAuthConfigurationId** | **string**| key: certificateBasedAuthConfiguration-id of certificateBasedAuthConfiguration | 
 **optional** | ***CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| ETag | 

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


## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration

> MicrosoftGraphCertificateBasedAuthConfiguration CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration(ctx, certificateBasedAuthConfigurationId, optional)

Get entity from certificateBasedAuthConfiguration by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateBasedAuthConfigurationId** | **string**| key: certificateBasedAuthConfiguration-id of certificateBasedAuthConfiguration | 
 **optional** | ***CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphCertificateBasedAuthConfiguration**](microsoft.graph.certificateBasedAuthConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration

> CollectionOfCertificateBasedAuthConfiguration CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration(ctx, optional)

Get entities from certificateBasedAuthConfiguration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfigurationOpts struct


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

[**CollectionOfCertificateBasedAuthConfiguration**](Collection_of_certificateBasedAuthConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration

> CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration(ctx, certificateBasedAuthConfigurationId, microsoftGraphCertificateBasedAuthConfiguration)

Update entity in certificateBasedAuthConfiguration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateBasedAuthConfigurationId** | **string**| key: certificateBasedAuthConfiguration-id of certificateBasedAuthConfiguration | 
**microsoftGraphCertificateBasedAuthConfiguration** | [**MicrosoftGraphCertificateBasedAuthConfiguration**](MicrosoftGraphCertificateBasedAuthConfiguration.md)| New property values | 

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

