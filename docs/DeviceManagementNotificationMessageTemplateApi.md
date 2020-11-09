# \DeviceManagementNotificationMessageTemplateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateNotificationMessageTemplates**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementCreateNotificationMessageTemplates) | **Post** /deviceManagement/notificationMessageTemplates | Create new navigation property to notificationMessageTemplates for deviceManagement
[**DeviceManagementGetNotificationMessageTemplates**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementGetNotificationMessageTemplates) | **Get** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id} | Get notificationMessageTemplates from deviceManagement
[**DeviceManagementListNotificationMessageTemplates**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementListNotificationMessageTemplates) | **Get** /deviceManagement/notificationMessageTemplates | Get notificationMessageTemplates from deviceManagement
[**DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages) | **Post** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id}/localizedNotificationMessages | Create new navigation property to localizedNotificationMessages for deviceManagement
[**DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages) | **Get** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id}/localizedNotificationMessages/{localizedNotificationMessage-id} | Get localizedNotificationMessages from deviceManagement
[**DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages) | **Get** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id}/localizedNotificationMessages | Get localizedNotificationMessages from deviceManagement
[**DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages) | **Patch** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id}/localizedNotificationMessages/{localizedNotificationMessage-id} | Update the navigation property localizedNotificationMessages in deviceManagement
[**DeviceManagementUpdateNotificationMessageTemplates**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementUpdateNotificationMessageTemplates) | **Patch** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id} | Update the navigation property notificationMessageTemplates in deviceManagement



## DeviceManagementCreateNotificationMessageTemplates

> MicrosoftGraphNotificationMessageTemplate DeviceManagementCreateNotificationMessageTemplates(ctx, microsoftGraphNotificationMessageTemplate)

Create new navigation property to notificationMessageTemplates for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphNotificationMessageTemplate** | [**MicrosoftGraphNotificationMessageTemplate**](MicrosoftGraphNotificationMessageTemplate.md)| New navigation property | 

### Return type

[**MicrosoftGraphNotificationMessageTemplate**](microsoft.graph.notificationMessageTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetNotificationMessageTemplates

> MicrosoftGraphNotificationMessageTemplate DeviceManagementGetNotificationMessageTemplates(ctx, notificationMessageTemplateId, optional)

Get notificationMessageTemplates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
 **optional** | ***DeviceManagementGetNotificationMessageTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetNotificationMessageTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphNotificationMessageTemplate**](microsoft.graph.notificationMessageTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListNotificationMessageTemplates

> CollectionOfNotificationMessageTemplate DeviceManagementListNotificationMessageTemplates(ctx, optional)

Get notificationMessageTemplates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListNotificationMessageTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListNotificationMessageTemplatesOpts struct


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

[**CollectionOfNotificationMessageTemplate**](Collection_of_notificationMessageTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages

> MicrosoftGraphLocalizedNotificationMessage DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages(ctx, notificationMessageTemplateId, microsoftGraphLocalizedNotificationMessage)

Create new navigation property to localizedNotificationMessages for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
**microsoftGraphLocalizedNotificationMessage** | [**MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md)| New navigation property | 

### Return type

[**MicrosoftGraphLocalizedNotificationMessage**](microsoft.graph.localizedNotificationMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages

> MicrosoftGraphLocalizedNotificationMessage DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages(ctx, notificationMessageTemplateId, localizedNotificationMessageId, optional)

Get localizedNotificationMessages from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
**localizedNotificationMessageId** | **string**| key: localizedNotificationMessage-id of localizedNotificationMessage | 
 **optional** | ***DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphLocalizedNotificationMessage**](microsoft.graph.localizedNotificationMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages

> CollectionOfLocalizedNotificationMessage DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages(ctx, notificationMessageTemplateId, optional)

Get localizedNotificationMessages from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
 **optional** | ***DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessagesOpts struct


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

[**CollectionOfLocalizedNotificationMessage**](Collection_of_localizedNotificationMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages

> DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages(ctx, notificationMessageTemplateId, localizedNotificationMessageId, microsoftGraphLocalizedNotificationMessage)

Update the navigation property localizedNotificationMessages in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
**localizedNotificationMessageId** | **string**| key: localizedNotificationMessage-id of localizedNotificationMessage | 
**microsoftGraphLocalizedNotificationMessage** | [**MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md)| New navigation property values | 

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


## DeviceManagementUpdateNotificationMessageTemplates

> DeviceManagementUpdateNotificationMessageTemplates(ctx, notificationMessageTemplateId, microsoftGraphNotificationMessageTemplate)

Update the navigation property notificationMessageTemplates in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
**microsoftGraphNotificationMessageTemplate** | [**MicrosoftGraphNotificationMessageTemplate**](MicrosoftGraphNotificationMessageTemplate.md)| New navigation property values | 

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

