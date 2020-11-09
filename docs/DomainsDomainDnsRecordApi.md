# \DomainsDomainDnsRecordApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsCreateServiceConfigurationRecords**](DomainsDomainDnsRecordApi.md#DomainsCreateServiceConfigurationRecords) | **Post** /domains/{domain-id}/serviceConfigurationRecords | Create new navigation property to serviceConfigurationRecords for domains
[**DomainsCreateVerificationDnsRecords**](DomainsDomainDnsRecordApi.md#DomainsCreateVerificationDnsRecords) | **Post** /domains/{domain-id}/verificationDnsRecords | Create new navigation property to verificationDnsRecords for domains
[**DomainsGetServiceConfigurationRecords**](DomainsDomainDnsRecordApi.md#DomainsGetServiceConfigurationRecords) | **Get** /domains/{domain-id}/serviceConfigurationRecords/{domainDnsRecord-id} | Get serviceConfigurationRecords from domains
[**DomainsGetVerificationDnsRecords**](DomainsDomainDnsRecordApi.md#DomainsGetVerificationDnsRecords) | **Get** /domains/{domain-id}/verificationDnsRecords/{domainDnsRecord-id} | Get verificationDnsRecords from domains
[**DomainsListServiceConfigurationRecords**](DomainsDomainDnsRecordApi.md#DomainsListServiceConfigurationRecords) | **Get** /domains/{domain-id}/serviceConfigurationRecords | Get serviceConfigurationRecords from domains
[**DomainsListVerificationDnsRecords**](DomainsDomainDnsRecordApi.md#DomainsListVerificationDnsRecords) | **Get** /domains/{domain-id}/verificationDnsRecords | Get verificationDnsRecords from domains
[**DomainsUpdateServiceConfigurationRecords**](DomainsDomainDnsRecordApi.md#DomainsUpdateServiceConfigurationRecords) | **Patch** /domains/{domain-id}/serviceConfigurationRecords/{domainDnsRecord-id} | Update the navigation property serviceConfigurationRecords in domains
[**DomainsUpdateVerificationDnsRecords**](DomainsDomainDnsRecordApi.md#DomainsUpdateVerificationDnsRecords) | **Patch** /domains/{domain-id}/verificationDnsRecords/{domainDnsRecord-id} | Update the navigation property verificationDnsRecords in domains



## DomainsCreateServiceConfigurationRecords

> MicrosoftGraphDomainDnsRecord DomainsCreateServiceConfigurationRecords(ctx, domainId, microsoftGraphDomainDnsRecord)

Create new navigation property to serviceConfigurationRecords for domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)| New navigation property | 

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


## DomainsCreateVerificationDnsRecords

> MicrosoftGraphDomainDnsRecord DomainsCreateVerificationDnsRecords(ctx, domainId, microsoftGraphDomainDnsRecord)

Create new navigation property to verificationDnsRecords for domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)| New navigation property | 

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


## DomainsGetServiceConfigurationRecords

> MicrosoftGraphDomainDnsRecord DomainsGetServiceConfigurationRecords(ctx, domainId, domainDnsRecordId, optional)

Get serviceConfigurationRecords from domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**domainDnsRecordId** | **string**| key: domainDnsRecord-id of domainDnsRecord | 
 **optional** | ***DomainsGetServiceConfigurationRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsGetServiceConfigurationRecordsOpts struct


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


## DomainsGetVerificationDnsRecords

> MicrosoftGraphDomainDnsRecord DomainsGetVerificationDnsRecords(ctx, domainId, domainDnsRecordId, optional)

Get verificationDnsRecords from domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**domainDnsRecordId** | **string**| key: domainDnsRecord-id of domainDnsRecord | 
 **optional** | ***DomainsGetVerificationDnsRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsGetVerificationDnsRecordsOpts struct


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


## DomainsListServiceConfigurationRecords

> CollectionOfDomainDnsRecord DomainsListServiceConfigurationRecords(ctx, domainId, optional)

Get serviceConfigurationRecords from domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
 **optional** | ***DomainsListServiceConfigurationRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsListServiceConfigurationRecordsOpts struct


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


## DomainsListVerificationDnsRecords

> CollectionOfDomainDnsRecord DomainsListVerificationDnsRecords(ctx, domainId, optional)

Get verificationDnsRecords from domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
 **optional** | ***DomainsListVerificationDnsRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsListVerificationDnsRecordsOpts struct


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


## DomainsUpdateServiceConfigurationRecords

> DomainsUpdateServiceConfigurationRecords(ctx, domainId, domainDnsRecordId, microsoftGraphDomainDnsRecord)

Update the navigation property serviceConfigurationRecords in domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**domainDnsRecordId** | **string**| key: domainDnsRecord-id of domainDnsRecord | 
**microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)| New navigation property values | 

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


## DomainsUpdateVerificationDnsRecords

> DomainsUpdateVerificationDnsRecords(ctx, domainId, domainDnsRecordId, microsoftGraphDomainDnsRecord)

Update the navigation property verificationDnsRecords in domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**domainDnsRecordId** | **string**| key: domainDnsRecord-id of domainDnsRecord | 
**microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)| New navigation property values | 

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

