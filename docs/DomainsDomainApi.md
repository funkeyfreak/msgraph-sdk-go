# \DomainsDomainApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsDomainCreateDomain**](DomainsDomainApi.md#DomainsDomainCreateDomain) | **Post** /domains | Add new entity to domains
[**DomainsDomainDeleteDomain**](DomainsDomainApi.md#DomainsDomainDeleteDomain) | **Delete** /domains/{domain-id} | Delete entity from domains
[**DomainsDomainGetDomain**](DomainsDomainApi.md#DomainsDomainGetDomain) | **Get** /domains/{domain-id} | Get entity from domains by key
[**DomainsDomainListDomain**](DomainsDomainApi.md#DomainsDomainListDomain) | **Get** /domains | Get entities from domains
[**DomainsDomainUpdateDomain**](DomainsDomainApi.md#DomainsDomainUpdateDomain) | **Patch** /domains/{domain-id} | Update entity in domains



## DomainsDomainCreateDomain

> MicrosoftGraphDomain DomainsDomainCreateDomain(ctx, microsoftGraphDomain)

Add new entity to domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDomain** | [**MicrosoftGraphDomain**](MicrosoftGraphDomain.md)| New entity | 

### Return type

[**MicrosoftGraphDomain**](microsoft.graph.domain.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainDeleteDomain

> DomainsDomainDeleteDomain(ctx, domainId, optional)

Delete entity from domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
 **optional** | ***DomainsDomainDeleteDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsDomainDeleteDomainOpts struct


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


## DomainsDomainGetDomain

> MicrosoftGraphDomain DomainsDomainGetDomain(ctx, domainId, optional)

Get entity from domains by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
 **optional** | ***DomainsDomainGetDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsDomainGetDomainOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDomain**](microsoft.graph.domain.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainListDomain

> CollectionOfDomain DomainsDomainListDomain(ctx, optional)

Get entities from domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DomainsDomainListDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsDomainListDomainOpts struct


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

[**CollectionOfDomain**](Collection_of_domain.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainUpdateDomain

> DomainsDomainUpdateDomain(ctx, domainId, microsoftGraphDomain)

Update entity in domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**microsoftGraphDomain** | [**MicrosoftGraphDomain**](MicrosoftGraphDomain.md)| New property values | 

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

