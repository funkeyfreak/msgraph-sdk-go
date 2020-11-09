# \ContactsOrgContactApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactsOrgContactCreateOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactCreateOrgContact) | **Post** /contacts | Add new entity to contacts
[**ContactsOrgContactDeleteOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactDeleteOrgContact) | **Delete** /contacts/{orgContact-id} | Delete entity from contacts
[**ContactsOrgContactGetOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactGetOrgContact) | **Get** /contacts/{orgContact-id} | Get entity from contacts by key
[**ContactsOrgContactListOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactListOrgContact) | **Get** /contacts | Get entities from contacts
[**ContactsOrgContactUpdateOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactUpdateOrgContact) | **Patch** /contacts/{orgContact-id} | Update entity in contacts



## ContactsOrgContactCreateOrgContact

> MicrosoftGraphOrgContact ContactsOrgContactCreateOrgContact(ctx, microsoftGraphOrgContact)

Add new entity to contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOrgContact** | [**MicrosoftGraphOrgContact**](MicrosoftGraphOrgContact.md)| New entity | 

### Return type

[**MicrosoftGraphOrgContact**](microsoft.graph.orgContact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactDeleteOrgContact

> ContactsOrgContactDeleteOrgContact(ctx, orgContactId, optional)

Delete entity from contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
 **optional** | ***ContactsOrgContactDeleteOrgContactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsOrgContactDeleteOrgContactOpts struct


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


## ContactsOrgContactGetOrgContact

> MicrosoftGraphOrgContact ContactsOrgContactGetOrgContact(ctx, orgContactId, optional)

Get entity from contacts by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
 **optional** | ***ContactsOrgContactGetOrgContactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsOrgContactGetOrgContactOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphOrgContact**](microsoft.graph.orgContact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactListOrgContact

> CollectionOfOrgContact ContactsOrgContactListOrgContact(ctx, optional)

Get entities from contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContactsOrgContactListOrgContactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactsOrgContactListOrgContactOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfOrgContact**](Collection_of_orgContact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactUpdateOrgContact

> ContactsOrgContactUpdateOrgContact(ctx, orgContactId, microsoftGraphOrgContact)

Update entity in contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string**| key: orgContact-id of orgContact | 
**microsoftGraphOrgContact** | [**MicrosoftGraphOrgContact**](MicrosoftGraphOrgContact.md)| New property values | 

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

