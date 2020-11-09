# \DeviceManagementMobileThreatDefenseConnectorApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateMobileThreatDefenseConnectors**](DeviceManagementMobileThreatDefenseConnectorApi.md#DeviceManagementCreateMobileThreatDefenseConnectors) | **Post** /deviceManagement/mobileThreatDefenseConnectors | Create new navigation property to mobileThreatDefenseConnectors for deviceManagement
[**DeviceManagementGetMobileThreatDefenseConnectors**](DeviceManagementMobileThreatDefenseConnectorApi.md#DeviceManagementGetMobileThreatDefenseConnectors) | **Get** /deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector-id} | Get mobileThreatDefenseConnectors from deviceManagement
[**DeviceManagementListMobileThreatDefenseConnectors**](DeviceManagementMobileThreatDefenseConnectorApi.md#DeviceManagementListMobileThreatDefenseConnectors) | **Get** /deviceManagement/mobileThreatDefenseConnectors | Get mobileThreatDefenseConnectors from deviceManagement
[**DeviceManagementUpdateMobileThreatDefenseConnectors**](DeviceManagementMobileThreatDefenseConnectorApi.md#DeviceManagementUpdateMobileThreatDefenseConnectors) | **Patch** /deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector-id} | Update the navigation property mobileThreatDefenseConnectors in deviceManagement



## DeviceManagementCreateMobileThreatDefenseConnectors

> MicrosoftGraphMobileThreatDefenseConnector DeviceManagementCreateMobileThreatDefenseConnectors(ctx, microsoftGraphMobileThreatDefenseConnector)

Create new navigation property to mobileThreatDefenseConnectors for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphMobileThreatDefenseConnector** | [**MicrosoftGraphMobileThreatDefenseConnector**](MicrosoftGraphMobileThreatDefenseConnector.md)| New navigation property | 

### Return type

[**MicrosoftGraphMobileThreatDefenseConnector**](microsoft.graph.mobileThreatDefenseConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetMobileThreatDefenseConnectors

> MicrosoftGraphMobileThreatDefenseConnector DeviceManagementGetMobileThreatDefenseConnectors(ctx, mobileThreatDefenseConnectorId, optional)

Get mobileThreatDefenseConnectors from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileThreatDefenseConnectorId** | **string**| key: mobileThreatDefenseConnector-id of mobileThreatDefenseConnector | 
 **optional** | ***DeviceManagementGetMobileThreatDefenseConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetMobileThreatDefenseConnectorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMobileThreatDefenseConnector**](microsoft.graph.mobileThreatDefenseConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListMobileThreatDefenseConnectors

> CollectionOfMobileThreatDefenseConnector DeviceManagementListMobileThreatDefenseConnectors(ctx, optional)

Get mobileThreatDefenseConnectors from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListMobileThreatDefenseConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListMobileThreatDefenseConnectorsOpts struct


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

[**CollectionOfMobileThreatDefenseConnector**](Collection_of_mobileThreatDefenseConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateMobileThreatDefenseConnectors

> DeviceManagementUpdateMobileThreatDefenseConnectors(ctx, mobileThreatDefenseConnectorId, microsoftGraphMobileThreatDefenseConnector)

Update the navigation property mobileThreatDefenseConnectors in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileThreatDefenseConnectorId** | **string**| key: mobileThreatDefenseConnector-id of mobileThreatDefenseConnector | 
**microsoftGraphMobileThreatDefenseConnector** | [**MicrosoftGraphMobileThreatDefenseConnector**](MicrosoftGraphMobileThreatDefenseConnector.md)| New navigation property values | 

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

