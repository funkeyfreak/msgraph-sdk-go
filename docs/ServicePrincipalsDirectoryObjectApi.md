# \ServicePrincipalsDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsGetCreatedObjects**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsGetCreatedObjects) | **Get** /servicePrincipals/{servicePrincipal-id}/createdObjects/{directoryObject-id} | Get createdObjects from servicePrincipals
[**ServicePrincipalsGetMemberOf**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsGetMemberOf) | **Get** /servicePrincipals/{servicePrincipal-id}/memberOf/{directoryObject-id} | Get memberOf from servicePrincipals
[**ServicePrincipalsGetOwnedObjects**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsGetOwnedObjects) | **Get** /servicePrincipals/{servicePrincipal-id}/ownedObjects/{directoryObject-id} | Get ownedObjects from servicePrincipals
[**ServicePrincipalsGetOwners**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsGetOwners) | **Get** /servicePrincipals/{servicePrincipal-id}/owners/{directoryObject-id} | Get owners from servicePrincipals
[**ServicePrincipalsGetTransitiveMemberOf**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsGetTransitiveMemberOf) | **Get** /servicePrincipals/{servicePrincipal-id}/transitiveMemberOf/{directoryObject-id} | Get transitiveMemberOf from servicePrincipals
[**ServicePrincipalsListCreatedObjects**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsListCreatedObjects) | **Get** /servicePrincipals/{servicePrincipal-id}/createdObjects | Get createdObjects from servicePrincipals
[**ServicePrincipalsListMemberOf**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsListMemberOf) | **Get** /servicePrincipals/{servicePrincipal-id}/memberOf | Get memberOf from servicePrincipals
[**ServicePrincipalsListOwnedObjects**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsListOwnedObjects) | **Get** /servicePrincipals/{servicePrincipal-id}/ownedObjects | Get ownedObjects from servicePrincipals
[**ServicePrincipalsListOwners**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsListOwners) | **Get** /servicePrincipals/{servicePrincipal-id}/owners | Get owners from servicePrincipals
[**ServicePrincipalsListTransitiveMemberOf**](ServicePrincipalsDirectoryObjectApi.md#ServicePrincipalsListTransitiveMemberOf) | **Get** /servicePrincipals/{servicePrincipal-id}/transitiveMemberOf | Get transitiveMemberOf from servicePrincipals



## ServicePrincipalsGetCreatedObjects

> MicrosoftGraphDirectoryObject ServicePrincipalsGetCreatedObjects(ctx, servicePrincipalId, directoryObjectId, optional)

Get createdObjects from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***ServicePrincipalsGetCreatedObjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetCreatedObjectsOpts struct


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


## ServicePrincipalsGetMemberOf

> MicrosoftGraphDirectoryObject ServicePrincipalsGetMemberOf(ctx, servicePrincipalId, directoryObjectId, optional)

Get memberOf from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***ServicePrincipalsGetMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetMemberOfOpts struct


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


## ServicePrincipalsGetOwnedObjects

> MicrosoftGraphDirectoryObject ServicePrincipalsGetOwnedObjects(ctx, servicePrincipalId, directoryObjectId, optional)

Get ownedObjects from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***ServicePrincipalsGetOwnedObjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetOwnedObjectsOpts struct


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


## ServicePrincipalsGetOwners

> MicrosoftGraphDirectoryObject ServicePrincipalsGetOwners(ctx, servicePrincipalId, directoryObjectId, optional)

Get owners from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***ServicePrincipalsGetOwnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetOwnersOpts struct


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


## ServicePrincipalsGetTransitiveMemberOf

> MicrosoftGraphDirectoryObject ServicePrincipalsGetTransitiveMemberOf(ctx, servicePrincipalId, directoryObjectId, optional)

Get transitiveMemberOf from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***ServicePrincipalsGetTransitiveMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetTransitiveMemberOfOpts struct


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


## ServicePrincipalsListCreatedObjects

> CollectionOfDirectoryObject ServicePrincipalsListCreatedObjects(ctx, servicePrincipalId, optional)

Get createdObjects from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListCreatedObjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListCreatedObjectsOpts struct


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


## ServicePrincipalsListMemberOf

> CollectionOfDirectoryObject ServicePrincipalsListMemberOf(ctx, servicePrincipalId, optional)

Get memberOf from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListMemberOfOpts struct


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


## ServicePrincipalsListOwnedObjects

> CollectionOfDirectoryObject ServicePrincipalsListOwnedObjects(ctx, servicePrincipalId, optional)

Get ownedObjects from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListOwnedObjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListOwnedObjectsOpts struct


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


## ServicePrincipalsListOwners

> CollectionOfDirectoryObject ServicePrincipalsListOwners(ctx, servicePrincipalId, optional)

Get owners from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListOwnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListOwnersOpts struct


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


## ServicePrincipalsListTransitiveMemberOf

> CollectionOfDirectoryObject ServicePrincipalsListTransitiveMemberOf(ctx, servicePrincipalId, optional)

Get transitiveMemberOf from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListTransitiveMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListTransitiveMemberOfOpts struct


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

