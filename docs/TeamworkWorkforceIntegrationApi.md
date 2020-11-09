# \TeamworkWorkforceIntegrationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamworkCreateWorkforceIntegrations**](TeamworkWorkforceIntegrationApi.md#TeamworkCreateWorkforceIntegrations) | **Post** /teamwork/workforceIntegrations | Create new navigation property to workforceIntegrations for teamwork
[**TeamworkGetWorkforceIntegrations**](TeamworkWorkforceIntegrationApi.md#TeamworkGetWorkforceIntegrations) | **Get** /teamwork/workforceIntegrations/{workforceIntegration-id} | Get workforceIntegrations from teamwork
[**TeamworkListWorkforceIntegrations**](TeamworkWorkforceIntegrationApi.md#TeamworkListWorkforceIntegrations) | **Get** /teamwork/workforceIntegrations | Get workforceIntegrations from teamwork
[**TeamworkUpdateWorkforceIntegrations**](TeamworkWorkforceIntegrationApi.md#TeamworkUpdateWorkforceIntegrations) | **Patch** /teamwork/workforceIntegrations/{workforceIntegration-id} | Update the navigation property workforceIntegrations in teamwork



## TeamworkCreateWorkforceIntegrations

> MicrosoftGraphWorkforceIntegration TeamworkCreateWorkforceIntegrations(ctx, microsoftGraphWorkforceIntegration)

Create new navigation property to workforceIntegrations for teamwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphWorkforceIntegration** | [**MicrosoftGraphWorkforceIntegration**](MicrosoftGraphWorkforceIntegration.md)| New navigation property | 

### Return type

[**MicrosoftGraphWorkforceIntegration**](microsoft.graph.workforceIntegration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamworkGetWorkforceIntegrations

> MicrosoftGraphWorkforceIntegration TeamworkGetWorkforceIntegrations(ctx, workforceIntegrationId, optional)

Get workforceIntegrations from teamwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workforceIntegrationId** | **string**| key: workforceIntegration-id of workforceIntegration | 
 **optional** | ***TeamworkGetWorkforceIntegrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamworkGetWorkforceIntegrationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkforceIntegration**](microsoft.graph.workforceIntegration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamworkListWorkforceIntegrations

> CollectionOfWorkforceIntegration TeamworkListWorkforceIntegrations(ctx, optional)

Get workforceIntegrations from teamwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamworkListWorkforceIntegrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamworkListWorkforceIntegrationsOpts struct


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

[**CollectionOfWorkforceIntegration**](Collection_of_workforceIntegration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamworkUpdateWorkforceIntegrations

> TeamworkUpdateWorkforceIntegrations(ctx, workforceIntegrationId, microsoftGraphWorkforceIntegration)

Update the navigation property workforceIntegrations in teamwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workforceIntegrationId** | **string**| key: workforceIntegration-id of workforceIntegration | 
**microsoftGraphWorkforceIntegration** | [**MicrosoftGraphWorkforceIntegration**](MicrosoftGraphWorkforceIntegration.md)| New navigation property values | 

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

