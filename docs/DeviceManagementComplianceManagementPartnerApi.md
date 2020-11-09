# \DeviceManagementComplianceManagementPartnerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateComplianceManagementPartners**](DeviceManagementComplianceManagementPartnerApi.md#DeviceManagementCreateComplianceManagementPartners) | **Post** /deviceManagement/complianceManagementPartners | Create new navigation property to complianceManagementPartners for deviceManagement
[**DeviceManagementGetComplianceManagementPartners**](DeviceManagementComplianceManagementPartnerApi.md#DeviceManagementGetComplianceManagementPartners) | **Get** /deviceManagement/complianceManagementPartners/{complianceManagementPartner-id} | Get complianceManagementPartners from deviceManagement
[**DeviceManagementListComplianceManagementPartners**](DeviceManagementComplianceManagementPartnerApi.md#DeviceManagementListComplianceManagementPartners) | **Get** /deviceManagement/complianceManagementPartners | Get complianceManagementPartners from deviceManagement
[**DeviceManagementUpdateComplianceManagementPartners**](DeviceManagementComplianceManagementPartnerApi.md#DeviceManagementUpdateComplianceManagementPartners) | **Patch** /deviceManagement/complianceManagementPartners/{complianceManagementPartner-id} | Update the navigation property complianceManagementPartners in deviceManagement



## DeviceManagementCreateComplianceManagementPartners

> MicrosoftGraphComplianceManagementPartner DeviceManagementCreateComplianceManagementPartners(ctx, microsoftGraphComplianceManagementPartner)

Create new navigation property to complianceManagementPartners for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphComplianceManagementPartner** | [**MicrosoftGraphComplianceManagementPartner**](MicrosoftGraphComplianceManagementPartner.md)| New navigation property | 

### Return type

[**MicrosoftGraphComplianceManagementPartner**](microsoft.graph.complianceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetComplianceManagementPartners

> MicrosoftGraphComplianceManagementPartner DeviceManagementGetComplianceManagementPartners(ctx, complianceManagementPartnerId, optional)

Get complianceManagementPartners from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**complianceManagementPartnerId** | **string**| key: complianceManagementPartner-id of complianceManagementPartner | 
 **optional** | ***DeviceManagementGetComplianceManagementPartnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetComplianceManagementPartnersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphComplianceManagementPartner**](microsoft.graph.complianceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListComplianceManagementPartners

> CollectionOfComplianceManagementPartner DeviceManagementListComplianceManagementPartners(ctx, optional)

Get complianceManagementPartners from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListComplianceManagementPartnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListComplianceManagementPartnersOpts struct


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

[**CollectionOfComplianceManagementPartner**](Collection_of_complianceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateComplianceManagementPartners

> DeviceManagementUpdateComplianceManagementPartners(ctx, complianceManagementPartnerId, microsoftGraphComplianceManagementPartner)

Update the navigation property complianceManagementPartners in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**complianceManagementPartnerId** | **string**| key: complianceManagementPartner-id of complianceManagementPartner | 
**microsoftGraphComplianceManagementPartner** | [**MicrosoftGraphComplianceManagementPartner**](MicrosoftGraphComplianceManagementPartner.md)| New navigation property values | 

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

