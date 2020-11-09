# \SchemaExtensionsSchemaExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemaExtensionsSchemaExtensionCreateSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionCreateSchemaExtension) | **Post** /schemaExtensions | Add new entity to schemaExtensions
[**SchemaExtensionsSchemaExtensionDeleteSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionDeleteSchemaExtension) | **Delete** /schemaExtensions/{schemaExtension-id} | Delete entity from schemaExtensions
[**SchemaExtensionsSchemaExtensionGetSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionGetSchemaExtension) | **Get** /schemaExtensions/{schemaExtension-id} | Get entity from schemaExtensions by key
[**SchemaExtensionsSchemaExtensionListSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionListSchemaExtension) | **Get** /schemaExtensions | Get entities from schemaExtensions
[**SchemaExtensionsSchemaExtensionUpdateSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionUpdateSchemaExtension) | **Patch** /schemaExtensions/{schemaExtension-id} | Update entity in schemaExtensions



## SchemaExtensionsSchemaExtensionCreateSchemaExtension

> MicrosoftGraphSchemaExtension SchemaExtensionsSchemaExtensionCreateSchemaExtension(ctx, microsoftGraphSchemaExtension)

Add new entity to schemaExtensions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSchemaExtension** | [**MicrosoftGraphSchemaExtension**](MicrosoftGraphSchemaExtension.md)| New entity | 

### Return type

[**MicrosoftGraphSchemaExtension**](microsoft.graph.schemaExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaExtensionsSchemaExtensionDeleteSchemaExtension

> SchemaExtensionsSchemaExtensionDeleteSchemaExtension(ctx, schemaExtensionId, optional)

Delete entity from schemaExtensions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaExtensionId** | **string**| key: schemaExtension-id of schemaExtension | 
 **optional** | ***SchemaExtensionsSchemaExtensionDeleteSchemaExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SchemaExtensionsSchemaExtensionDeleteSchemaExtensionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| ETag | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaExtensionsSchemaExtensionGetSchemaExtension

> MicrosoftGraphSchemaExtension SchemaExtensionsSchemaExtensionGetSchemaExtension(ctx, schemaExtensionId, optional)

Get entity from schemaExtensions by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaExtensionId** | **string**| key: schemaExtension-id of schemaExtension | 
 **optional** | ***SchemaExtensionsSchemaExtensionGetSchemaExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SchemaExtensionsSchemaExtensionGetSchemaExtensionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSchemaExtension**](microsoft.graph.schemaExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaExtensionsSchemaExtensionListSchemaExtension

> CollectionOfSchemaExtension SchemaExtensionsSchemaExtensionListSchemaExtension(ctx, optional)

Get entities from schemaExtensions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SchemaExtensionsSchemaExtensionListSchemaExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SchemaExtensionsSchemaExtensionListSchemaExtensionOpts struct


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

[**CollectionOfSchemaExtension**](Collection_of_schemaExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaExtensionsSchemaExtensionUpdateSchemaExtension

> SchemaExtensionsSchemaExtensionUpdateSchemaExtension(ctx, schemaExtensionId, microsoftGraphSchemaExtension)

Update entity in schemaExtensions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaExtensionId** | **string**| key: schemaExtension-id of schemaExtension | 
**microsoftGraphSchemaExtension** | [**MicrosoftGraphSchemaExtension**](MicrosoftGraphSchemaExtension.md)| New property values | 

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

