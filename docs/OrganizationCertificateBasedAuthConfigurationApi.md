# \OrganizationCertificateBasedAuthConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationGetCertificateBasedAuthConfiguration**](OrganizationCertificateBasedAuthConfigurationApi.md#OrganizationGetCertificateBasedAuthConfiguration) | **Get** /organization/{organization-id}/certificateBasedAuthConfiguration/{certificateBasedAuthConfiguration-id} | Get certificateBasedAuthConfiguration from organization
[**OrganizationListCertificateBasedAuthConfiguration**](OrganizationCertificateBasedAuthConfigurationApi.md#OrganizationListCertificateBasedAuthConfiguration) | **Get** /organization/{organization-id}/certificateBasedAuthConfiguration | Get certificateBasedAuthConfiguration from organization



## OrganizationGetCertificateBasedAuthConfiguration

> MicrosoftGraphCertificateBasedAuthConfiguration OrganizationGetCertificateBasedAuthConfiguration(ctx, organizationId, certificateBasedAuthConfigurationId, optional)

Get certificateBasedAuthConfiguration from organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| key: organization-id of organization | 
**certificateBasedAuthConfigurationId** | **string**| key: certificateBasedAuthConfiguration-id of certificateBasedAuthConfiguration | 
 **optional** | ***OrganizationGetCertificateBasedAuthConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationGetCertificateBasedAuthConfigurationOpts struct


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


## OrganizationListCertificateBasedAuthConfiguration

> CollectionOfCertificateBasedAuthConfiguration OrganizationListCertificateBasedAuthConfiguration(ctx, organizationId, optional)

Get certificateBasedAuthConfiguration from organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| key: organization-id of organization | 
 **optional** | ***OrganizationListCertificateBasedAuthConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationListCertificateBasedAuthConfigurationOpts struct


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

