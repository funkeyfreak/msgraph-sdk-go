# \ContactsDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactsGetDirectReports**](ContactsDirectoryObjectApi.md#ContactsGetDirectReports) | **Get** /contacts/{orgContact-id}/directReports/{directoryObject-id} | Get directReports from contacts
[**ContactsGetManager**](ContactsDirectoryObjectApi.md#ContactsGetManager) | **Get** /contacts/{orgContact-id}/manager | Get manager from contacts
[**ContactsGetMemberOf**](ContactsDirectoryObjectApi.md#ContactsGetMemberOf) | **Get** /contacts/{orgContact-id}/memberOf/{directoryObject-id} | Get memberOf from contacts
[**ContactsGetTransitiveMemberOf**](ContactsDirectoryObjectApi.md#ContactsGetTransitiveMemberOf) | **Get** /contacts/{orgContact-id}/transitiveMemberOf/{directoryObject-id} | Get transitiveMemberOf from contacts
[**ContactsListDirectReports**](ContactsDirectoryObjectApi.md#ContactsListDirectReports) | **Get** /contacts/{orgContact-id}/directReports | Get directReports from contacts
[**ContactsListMemberOf**](ContactsDirectoryObjectApi.md#ContactsListMemberOf) | **Get** /contacts/{orgContact-id}/memberOf | Get memberOf from contacts
[**ContactsListTransitiveMemberOf**](ContactsDirectoryObjectApi.md#ContactsListTransitiveMemberOf) | **Get** /contacts/{orgContact-id}/transitiveMemberOf | Get transitiveMemberOf from contacts



## ContactsGetDirectReports

> MicrosoftGraphDirectoryObject ContactsGetDirectReports(ctx, orgContactId, directoryObjectId, optional)

Get directReports from contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***ContactsGetDirectReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsGetDirectReportsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsGetManager

> MicrosoftGraphDirectoryObject ContactsGetManager(ctx, orgContactId, optional)

Get manager from contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
 **optional** | ***ContactsGetManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsGetManagerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsGetMemberOf

> MicrosoftGraphDirectoryObject ContactsGetMemberOf(ctx, orgContactId, directoryObjectId, optional)

Get memberOf from contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***ContactsGetMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsGetMemberOfOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsGetTransitiveMemberOf

> MicrosoftGraphDirectoryObject ContactsGetTransitiveMemberOf(ctx, orgContactId, directoryObjectId, optional)

Get transitiveMemberOf from contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***ContactsGetTransitiveMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsGetTransitiveMemberOfOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsListDirectReports

> CollectionOfDirectoryObject ContactsListDirectReports(ctx, orgContactId, optional)

Get directReports from contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
 **optional** | ***ContactsListDirectReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsListDirectReportsOpts struct


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

[**CollectionOfDirectoryObject**](Collection_of_directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsListMemberOf

> CollectionOfDirectoryObject ContactsListMemberOf(ctx, orgContactId, optional)

Get memberOf from contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
 **optional** | ***ContactsListMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsListMemberOfOpts struct


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

[**CollectionOfDirectoryObject**](Collection_of_directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsListTransitiveMemberOf

> CollectionOfDirectoryObject ContactsListTransitiveMemberOf(ctx, orgContactId, optional)

Get transitiveMemberOf from contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
 **optional** | ***ContactsListTransitiveMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsListTransitiveMemberOfOpts struct


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

[**CollectionOfDirectoryObject**](Collection_of_directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

