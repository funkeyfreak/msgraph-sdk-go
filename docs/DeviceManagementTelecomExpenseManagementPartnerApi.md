# \DeviceManagementTelecomExpenseManagementPartnerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateTelecomExpenseManagementPartners**](DeviceManagementTelecomExpenseManagementPartnerApi.md#DeviceManagementCreateTelecomExpenseManagementPartners) | **Post** /deviceManagement/telecomExpenseManagementPartners | Create new navigation property to telecomExpenseManagementPartners for deviceManagement
[**DeviceManagementGetTelecomExpenseManagementPartners**](DeviceManagementTelecomExpenseManagementPartnerApi.md#DeviceManagementGetTelecomExpenseManagementPartners) | **Get** /deviceManagement/telecomExpenseManagementPartners/{telecomExpenseManagementPartner-id} | Get telecomExpenseManagementPartners from deviceManagement
[**DeviceManagementListTelecomExpenseManagementPartners**](DeviceManagementTelecomExpenseManagementPartnerApi.md#DeviceManagementListTelecomExpenseManagementPartners) | **Get** /deviceManagement/telecomExpenseManagementPartners | Get telecomExpenseManagementPartners from deviceManagement
[**DeviceManagementUpdateTelecomExpenseManagementPartners**](DeviceManagementTelecomExpenseManagementPartnerApi.md#DeviceManagementUpdateTelecomExpenseManagementPartners) | **Patch** /deviceManagement/telecomExpenseManagementPartners/{telecomExpenseManagementPartner-id} | Update the navigation property telecomExpenseManagementPartners in deviceManagement



## DeviceManagementCreateTelecomExpenseManagementPartners

> MicrosoftGraphTelecomExpenseManagementPartner DeviceManagementCreateTelecomExpenseManagementPartners(ctx, microsoftGraphTelecomExpenseManagementPartner)

Create new navigation property to telecomExpenseManagementPartners for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTelecomExpenseManagementPartner** | [**MicrosoftGraphTelecomExpenseManagementPartner**](MicrosoftGraphTelecomExpenseManagementPartner.md)| New navigation property | 

### Return type

[**MicrosoftGraphTelecomExpenseManagementPartner**](microsoft.graph.telecomExpenseManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetTelecomExpenseManagementPartners

> MicrosoftGraphTelecomExpenseManagementPartner DeviceManagementGetTelecomExpenseManagementPartners(ctx, telecomExpenseManagementPartnerId, optional)

Get telecomExpenseManagementPartners from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**telecomExpenseManagementPartnerId** | **string**| key: telecomExpenseManagementPartner-id of telecomExpenseManagementPartner | 
 **optional** | ***DeviceManagementGetTelecomExpenseManagementPartnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetTelecomExpenseManagementPartnersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTelecomExpenseManagementPartner**](microsoft.graph.telecomExpenseManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListTelecomExpenseManagementPartners

> CollectionOfTelecomExpenseManagementPartner DeviceManagementListTelecomExpenseManagementPartners(ctx, optional)

Get telecomExpenseManagementPartners from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListTelecomExpenseManagementPartnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListTelecomExpenseManagementPartnersOpts struct


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

[**CollectionOfTelecomExpenseManagementPartner**](Collection_of_telecomExpenseManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateTelecomExpenseManagementPartners

> DeviceManagementUpdateTelecomExpenseManagementPartners(ctx, telecomExpenseManagementPartnerId, microsoftGraphTelecomExpenseManagementPartner)

Update the navigation property telecomExpenseManagementPartners in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**telecomExpenseManagementPartnerId** | **string**| key: telecomExpenseManagementPartner-id of telecomExpenseManagementPartner | 
**microsoftGraphTelecomExpenseManagementPartner** | [**MicrosoftGraphTelecomExpenseManagementPartner**](MicrosoftGraphTelecomExpenseManagementPartner.md)| New navigation property values | 

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

