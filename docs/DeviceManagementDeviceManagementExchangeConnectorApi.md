# \DeviceManagementDeviceManagementExchangeConnectorApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateExchangeConnectors**](DeviceManagementDeviceManagementExchangeConnectorApi.md#DeviceManagementCreateExchangeConnectors) | **Post** /deviceManagement/exchangeConnectors | Create new navigation property to exchangeConnectors for deviceManagement
[**DeviceManagementGetExchangeConnectors**](DeviceManagementDeviceManagementExchangeConnectorApi.md#DeviceManagementGetExchangeConnectors) | **Get** /deviceManagement/exchangeConnectors/{deviceManagementExchangeConnector-id} | Get exchangeConnectors from deviceManagement
[**DeviceManagementListExchangeConnectors**](DeviceManagementDeviceManagementExchangeConnectorApi.md#DeviceManagementListExchangeConnectors) | **Get** /deviceManagement/exchangeConnectors | Get exchangeConnectors from deviceManagement
[**DeviceManagementUpdateExchangeConnectors**](DeviceManagementDeviceManagementExchangeConnectorApi.md#DeviceManagementUpdateExchangeConnectors) | **Patch** /deviceManagement/exchangeConnectors/{deviceManagementExchangeConnector-id} | Update the navigation property exchangeConnectors in deviceManagement



## DeviceManagementCreateExchangeConnectors

> MicrosoftGraphDeviceManagementExchangeConnector DeviceManagementCreateExchangeConnectors(ctx, microsoftGraphDeviceManagementExchangeConnector)

Create new navigation property to exchangeConnectors for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceManagementExchangeConnector** | [**MicrosoftGraphDeviceManagementExchangeConnector**](MicrosoftGraphDeviceManagementExchangeConnector.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceManagementExchangeConnector**](microsoft.graph.deviceManagementExchangeConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetExchangeConnectors

> MicrosoftGraphDeviceManagementExchangeConnector DeviceManagementGetExchangeConnectors(ctx, deviceManagementExchangeConnectorId, optional)

Get exchangeConnectors from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementExchangeConnectorId** | **string**| key: deviceManagementExchangeConnector-id of deviceManagementExchangeConnector | 
 **optional** | ***DeviceManagementGetExchangeConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetExchangeConnectorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagementExchangeConnector**](microsoft.graph.deviceManagementExchangeConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListExchangeConnectors

> CollectionOfDeviceManagementExchangeConnector DeviceManagementListExchangeConnectors(ctx, optional)

Get exchangeConnectors from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListExchangeConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListExchangeConnectorsOpts struct


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

[**CollectionOfDeviceManagementExchangeConnector**](Collection_of_deviceManagementExchangeConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateExchangeConnectors

> DeviceManagementUpdateExchangeConnectors(ctx, deviceManagementExchangeConnectorId, microsoftGraphDeviceManagementExchangeConnector)

Update the navigation property exchangeConnectors in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementExchangeConnectorId** | **string**| key: deviceManagementExchangeConnector-id of deviceManagementExchangeConnector | 
**microsoftGraphDeviceManagementExchangeConnector** | [**MicrosoftGraphDeviceManagementExchangeConnector**](MicrosoftGraphDeviceManagementExchangeConnector.md)| New navigation property values | 

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

