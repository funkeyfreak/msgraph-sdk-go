# \DirectoryObjectsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryObjectsCheckMemberGroups**](DirectoryObjectsActionsApi.md#DirectoryObjectsCheckMemberGroups) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**DirectoryObjectsCheckMemberObjects**](DirectoryObjectsActionsApi.md#DirectoryObjectsCheckMemberObjects) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**DirectoryObjectsGetAvailableExtensionProperties**](DirectoryObjectsActionsApi.md#DirectoryObjectsGetAvailableExtensionProperties) | **Post** /directoryObjects/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**DirectoryObjectsGetByIds**](DirectoryObjectsActionsApi.md#DirectoryObjectsGetByIds) | **Post** /directoryObjects/microsoft.graph.getByIds | Invoke action getByIds
[**DirectoryObjectsGetMemberGroups**](DirectoryObjectsActionsApi.md#DirectoryObjectsGetMemberGroups) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**DirectoryObjectsGetMemberObjects**](DirectoryObjectsActionsApi.md#DirectoryObjectsGetMemberObjects) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**DirectoryObjectsRestore**](DirectoryObjectsActionsApi.md#DirectoryObjectsRestore) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.restore | Invoke action restore
[**DirectoryObjectsValidateProperties**](DirectoryObjectsActionsApi.md#DirectoryObjectsValidateProperties) | **Post** /directoryObjects/microsoft.graph.validateProperties | Invoke action validateProperties



## DirectoryObjectsCheckMemberGroups

> []string DirectoryObjectsCheckMemberGroups(ctx, directoryObjectId, inlineObject44)

Invoke action checkMemberGroups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
**inlineObject44** | [**InlineObject44**](InlineObject44.md)|  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsCheckMemberObjects

> []string DirectoryObjectsCheckMemberObjects(ctx, directoryObjectId, inlineObject45)

Invoke action checkMemberObjects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
**inlineObject45** | [**InlineObject45**](InlineObject45.md)|  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty DirectoryObjectsGetAvailableExtensionProperties(ctx, inlineObject48)

Invoke action getAvailableExtensionProperties

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject48** | [**InlineObject48**](InlineObject48.md)|  | 

### Return type

[**[]MicrosoftGraphExtensionProperty**](microsoft.graph.extensionProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsGetByIds

> []MicrosoftGraphDirectoryObject DirectoryObjectsGetByIds(ctx, inlineObject49)

Invoke action getByIds

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject49** | [**InlineObject49**](InlineObject49.md)|  | 

### Return type

[**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsGetMemberGroups

> []string DirectoryObjectsGetMemberGroups(ctx, directoryObjectId, inlineObject46)

Invoke action getMemberGroups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
**inlineObject46** | [**InlineObject46**](InlineObject46.md)|  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsGetMemberObjects

> []string DirectoryObjectsGetMemberObjects(ctx, directoryObjectId, inlineObject47)

Invoke action getMemberObjects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
**inlineObject47** | [**InlineObject47**](InlineObject47.md)|  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsRestore

> interface{} DirectoryObjectsRestore(ctx, directoryObjectId)

Invoke action restore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsValidateProperties

> DirectoryObjectsValidateProperties(ctx, inlineObject50)

Invoke action validateProperties

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject50** | [**InlineObject50**](InlineObject50.md)|  | 

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

