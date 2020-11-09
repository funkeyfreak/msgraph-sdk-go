# \SecurityAlertApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityCreateAlerts**](SecurityAlertApi.md#SecurityCreateAlerts) | **Post** /Security/alerts | Create new navigation property to alerts for Security
[**SecurityGetAlerts**](SecurityAlertApi.md#SecurityGetAlerts) | **Get** /Security/alerts/{alert-id} | Get alerts from Security
[**SecurityListAlerts**](SecurityAlertApi.md#SecurityListAlerts) | **Get** /Security/alerts | Get alerts from Security
[**SecurityUpdateAlerts**](SecurityAlertApi.md#SecurityUpdateAlerts) | **Patch** /Security/alerts/{alert-id} | Update the navigation property alerts in Security



## SecurityCreateAlerts

> MicrosoftGraphAlert SecurityCreateAlerts(ctx, microsoftGraphAlert)

Create new navigation property to alerts for Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphAlert** | [**MicrosoftGraphAlert**](MicrosoftGraphAlert.md)| New navigation property | 

### Return type

[**MicrosoftGraphAlert**](microsoft.graph.alert.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityGetAlerts

> MicrosoftGraphAlert SecurityGetAlerts(ctx, alertId, optional)

Get alerts from Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string**| key: alert-id of alert | 
 **optional** | ***SecurityGetAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityGetAlertsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAlert**](microsoft.graph.alert.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityListAlerts

> CollectionOfAlert SecurityListAlerts(ctx, optional)

Get alerts from Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityListAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityListAlertsOpts struct


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

[**CollectionOfAlert**](Collection_of_alert.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityUpdateAlerts

> SecurityUpdateAlerts(ctx, alertId, microsoftGraphAlert)

Update the navigation property alerts in Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string**| key: alert-id of alert | 
**microsoftGraphAlert** | [**MicrosoftGraphAlert**](MicrosoftGraphAlert.md)| New navigation property values | 

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

