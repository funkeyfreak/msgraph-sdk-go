# \DeviceManagementRemoteAssistancePartnerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateRemoteAssistancePartners**](DeviceManagementRemoteAssistancePartnerApi.md#DeviceManagementCreateRemoteAssistancePartners) | **Post** /deviceManagement/remoteAssistancePartners | Create new navigation property to remoteAssistancePartners for deviceManagement
[**DeviceManagementGetRemoteAssistancePartners**](DeviceManagementRemoteAssistancePartnerApi.md#DeviceManagementGetRemoteAssistancePartners) | **Get** /deviceManagement/remoteAssistancePartners/{remoteAssistancePartner-id} | Get remoteAssistancePartners from deviceManagement
[**DeviceManagementListRemoteAssistancePartners**](DeviceManagementRemoteAssistancePartnerApi.md#DeviceManagementListRemoteAssistancePartners) | **Get** /deviceManagement/remoteAssistancePartners | Get remoteAssistancePartners from deviceManagement
[**DeviceManagementUpdateRemoteAssistancePartners**](DeviceManagementRemoteAssistancePartnerApi.md#DeviceManagementUpdateRemoteAssistancePartners) | **Patch** /deviceManagement/remoteAssistancePartners/{remoteAssistancePartner-id} | Update the navigation property remoteAssistancePartners in deviceManagement



## DeviceManagementCreateRemoteAssistancePartners

> MicrosoftGraphRemoteAssistancePartner DeviceManagementCreateRemoteAssistancePartners(ctx, microsoftGraphRemoteAssistancePartner)

Create new navigation property to remoteAssistancePartners for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphRemoteAssistancePartner** | [**MicrosoftGraphRemoteAssistancePartner**](MicrosoftGraphRemoteAssistancePartner.md)| New navigation property | 

### Return type

[**MicrosoftGraphRemoteAssistancePartner**](microsoft.graph.remoteAssistancePartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetRemoteAssistancePartners

> MicrosoftGraphRemoteAssistancePartner DeviceManagementGetRemoteAssistancePartners(ctx, remoteAssistancePartnerId, optional)

Get remoteAssistancePartners from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remoteAssistancePartnerId** | **string**| key: remoteAssistancePartner-id of remoteAssistancePartner | 
 **optional** | ***DeviceManagementGetRemoteAssistancePartnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetRemoteAssistancePartnersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphRemoteAssistancePartner**](microsoft.graph.remoteAssistancePartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListRemoteAssistancePartners

> CollectionOfRemoteAssistancePartner DeviceManagementListRemoteAssistancePartners(ctx, optional)

Get remoteAssistancePartners from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListRemoteAssistancePartnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListRemoteAssistancePartnersOpts struct


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

[**CollectionOfRemoteAssistancePartner**](Collection_of_remoteAssistancePartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateRemoteAssistancePartners

> DeviceManagementUpdateRemoteAssistancePartners(ctx, remoteAssistancePartnerId, microsoftGraphRemoteAssistancePartner)

Update the navigation property remoteAssistancePartners in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remoteAssistancePartnerId** | **string**| key: remoteAssistancePartner-id of remoteAssistancePartner | 
**microsoftGraphRemoteAssistancePartner** | [**MicrosoftGraphRemoteAssistancePartner**](MicrosoftGraphRemoteAssistancePartner.md)| New navigation property values | 

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

