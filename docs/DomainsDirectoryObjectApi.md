# \DomainsDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsGetDomainNameReferences**](DomainsDirectoryObjectApi.md#DomainsGetDomainNameReferences) | **Get** /domains/{domain-id}/domainNameReferences/{directoryObject-id} | Get domainNameReferences from domains
[**DomainsListDomainNameReferences**](DomainsDirectoryObjectApi.md#DomainsListDomainNameReferences) | **Get** /domains/{domain-id}/domainNameReferences | Get domainNameReferences from domains



## DomainsGetDomainNameReferences

> MicrosoftGraphDirectoryObject DomainsGetDomainNameReferences(ctx, domainId, directoryObjectId, optional)

Get domainNameReferences from domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***DomainsGetDomainNameReferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsGetDomainNameReferencesOpts struct


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


## DomainsListDomainNameReferences

> CollectionOfDirectoryObject DomainsListDomainNameReferences(ctx, domainId, optional)

Get domainNameReferences from domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
 **optional** | ***DomainsListDomainNameReferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsListDomainNameReferencesOpts struct


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

