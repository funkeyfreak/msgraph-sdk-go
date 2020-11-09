# \DeviceManagementDeviceManagementPartnerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceManagementPartners**](DeviceManagementDeviceManagementPartnerApi.md#DeviceManagementCreateDeviceManagementPartners) | **Post** /deviceManagement/deviceManagementPartners | Create new navigation property to deviceManagementPartners for deviceManagement
[**DeviceManagementGetDeviceManagementPartners**](DeviceManagementDeviceManagementPartnerApi.md#DeviceManagementGetDeviceManagementPartners) | **Get** /deviceManagement/deviceManagementPartners/{deviceManagementPartner-id} | Get deviceManagementPartners from deviceManagement
[**DeviceManagementListDeviceManagementPartners**](DeviceManagementDeviceManagementPartnerApi.md#DeviceManagementListDeviceManagementPartners) | **Get** /deviceManagement/deviceManagementPartners | Get deviceManagementPartners from deviceManagement
[**DeviceManagementUpdateDeviceManagementPartners**](DeviceManagementDeviceManagementPartnerApi.md#DeviceManagementUpdateDeviceManagementPartners) | **Patch** /deviceManagement/deviceManagementPartners/{deviceManagementPartner-id} | Update the navigation property deviceManagementPartners in deviceManagement



## DeviceManagementCreateDeviceManagementPartners

> MicrosoftGraphDeviceManagementPartner DeviceManagementCreateDeviceManagementPartners(ctx, microsoftGraphDeviceManagementPartner)

Create new navigation property to deviceManagementPartners for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceManagementPartner** | [**MicrosoftGraphDeviceManagementPartner**](MicrosoftGraphDeviceManagementPartner.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceManagementPartner**](microsoft.graph.deviceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetDeviceManagementPartners

> MicrosoftGraphDeviceManagementPartner DeviceManagementGetDeviceManagementPartners(ctx, deviceManagementPartnerId, optional)

Get deviceManagementPartners from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementPartnerId** | **string**| key: deviceManagementPartner-id of deviceManagementPartner | 
 **optional** | ***DeviceManagementGetDeviceManagementPartnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDeviceManagementPartnersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagementPartner**](microsoft.graph.deviceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceManagementPartners

> CollectionOfDeviceManagementPartner DeviceManagementListDeviceManagementPartners(ctx, optional)

Get deviceManagementPartners from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListDeviceManagementPartnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListDeviceManagementPartnersOpts struct


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

[**CollectionOfDeviceManagementPartner**](Collection_of_deviceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceManagementPartners

> DeviceManagementUpdateDeviceManagementPartners(ctx, deviceManagementPartnerId, microsoftGraphDeviceManagementPartner)

Update the navigation property deviceManagementPartners in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementPartnerId** | **string**| key: deviceManagementPartner-id of deviceManagementPartner | 
**microsoftGraphDeviceManagementPartner** | [**MicrosoftGraphDeviceManagementPartner**](MicrosoftGraphDeviceManagementPartner.md)| New navigation property values | 

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

