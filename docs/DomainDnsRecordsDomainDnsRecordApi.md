# \DomainDnsRecordsDomainDnsRecordApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord) | **Post** /domainDnsRecords | Add new entity to domainDnsRecords
[**DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord) | **Delete** /domainDnsRecords/{domainDnsRecord-id} | Delete entity from domainDnsRecords
[**DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord) | **Get** /domainDnsRecords/{domainDnsRecord-id} | Get entity from domainDnsRecords by key
[**DomainDnsRecordsDomainDnsRecordListDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordListDomainDnsRecord) | **Get** /domainDnsRecords | Get entities from domainDnsRecords
[**DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord) | **Patch** /domainDnsRecords/{domainDnsRecord-id} | Update entity in domainDnsRecords



## DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord

> MicrosoftGraphDomainDnsRecord DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord(ctx, microsoftGraphDomainDnsRecord)

Add new entity to domainDnsRecords

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)| New entity | 

### Return type

[**MicrosoftGraphDomainDnsRecord**](microsoft.graph.domainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord

> DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord(ctx, domainDnsRecordId, optional)

Delete entity from domainDnsRecords

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainDnsRecordId** | **string**| key: domainDnsRecord-id of domainDnsRecord | 
 **optional** | ***DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecordOpts struct


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


## DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord

> MicrosoftGraphDomainDnsRecord DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord(ctx, domainDnsRecordId, optional)

Get entity from domainDnsRecords by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainDnsRecordId** | **string**| key: domainDnsRecord-id of domainDnsRecord | 
 **optional** | ***DomainDnsRecordsDomainDnsRecordGetDomainDnsRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainDnsRecordsDomainDnsRecordGetDomainDnsRecordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDomainDnsRecord**](microsoft.graph.domainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDnsRecordsDomainDnsRecordListDomainDnsRecord

> CollectionOfDomainDnsRecord DomainDnsRecordsDomainDnsRecordListDomainDnsRecord(ctx, optional)

Get entities from domainDnsRecords

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DomainDnsRecordsDomainDnsRecordListDomainDnsRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainDnsRecordsDomainDnsRecordListDomainDnsRecordOpts struct


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

[**CollectionOfDomainDnsRecord**](Collection_of_domainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord

> DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord(ctx, domainDnsRecordId, microsoftGraphDomainDnsRecord)

Update entity in domainDnsRecords

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainDnsRecordId** | **string**| key: domainDnsRecord-id of domainDnsRecord | 
**microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)| New property values | 

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

