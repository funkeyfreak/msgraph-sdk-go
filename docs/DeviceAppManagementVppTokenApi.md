# \DeviceAppManagementVppTokenApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateVppTokens**](DeviceAppManagementVppTokenApi.md#DeviceAppManagementCreateVppTokens) | **Post** /deviceAppManagement/vppTokens | Create new navigation property to vppTokens for deviceAppManagement
[**DeviceAppManagementGetVppTokens**](DeviceAppManagementVppTokenApi.md#DeviceAppManagementGetVppTokens) | **Get** /deviceAppManagement/vppTokens/{vppToken-id} | Get vppTokens from deviceAppManagement
[**DeviceAppManagementListVppTokens**](DeviceAppManagementVppTokenApi.md#DeviceAppManagementListVppTokens) | **Get** /deviceAppManagement/vppTokens | Get vppTokens from deviceAppManagement
[**DeviceAppManagementUpdateVppTokens**](DeviceAppManagementVppTokenApi.md#DeviceAppManagementUpdateVppTokens) | **Patch** /deviceAppManagement/vppTokens/{vppToken-id} | Update the navigation property vppTokens in deviceAppManagement



## DeviceAppManagementCreateVppTokens

> MicrosoftGraphVppToken DeviceAppManagementCreateVppTokens(ctx, microsoftGraphVppToken)

Create new navigation property to vppTokens for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphVppToken** | [**MicrosoftGraphVppToken**](MicrosoftGraphVppToken.md)| New navigation property | 

### Return type

[**MicrosoftGraphVppToken**](microsoft.graph.vppToken.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetVppTokens

> MicrosoftGraphVppToken DeviceAppManagementGetVppTokens(ctx, vppTokenId, optional)

Get vppTokens from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vppTokenId** | **string**| key: vppToken-id of vppToken | 
 **optional** | ***DeviceAppManagementGetVppTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetVppTokensOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphVppToken**](microsoft.graph.vppToken.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListVppTokens

> CollectionOfVppToken DeviceAppManagementListVppTokens(ctx, optional)

Get vppTokens from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListVppTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListVppTokensOpts struct


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

[**CollectionOfVppToken**](Collection_of_vppToken.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateVppTokens

> DeviceAppManagementUpdateVppTokens(ctx, vppTokenId, microsoftGraphVppToken)

Update the navigation property vppTokens in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vppTokenId** | **string**| key: vppToken-id of vppToken | 
**microsoftGraphVppToken** | [**MicrosoftGraphVppToken**](MicrosoftGraphVppToken.md)| New navigation property values | 

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

