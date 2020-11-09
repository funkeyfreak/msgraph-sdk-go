# \IdentityConditionalAccessRootApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityConditionalAccessCreateNamedLocations**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessCreateNamedLocations) | **Post** /identity/conditionalAccess/namedLocations | Create new navigation property to namedLocations for identity
[**IdentityConditionalAccessCreatePolicies**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessCreatePolicies) | **Post** /identity/conditionalAccess/policies | Create new navigation property to policies for identity
[**IdentityConditionalAccessGetNamedLocations**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessGetNamedLocations) | **Get** /identity/conditionalAccess/namedLocations/{namedLocation-id} | Get namedLocations from identity
[**IdentityConditionalAccessGetPolicies**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessGetPolicies) | **Get** /identity/conditionalAccess/policies/{conditionalAccessPolicy-id} | Get policies from identity
[**IdentityConditionalAccessListNamedLocations**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessListNamedLocations) | **Get** /identity/conditionalAccess/namedLocations | Get namedLocations from identity
[**IdentityConditionalAccessListPolicies**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessListPolicies) | **Get** /identity/conditionalAccess/policies | Get policies from identity
[**IdentityConditionalAccessUpdateNamedLocations**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessUpdateNamedLocations) | **Patch** /identity/conditionalAccess/namedLocations/{namedLocation-id} | Update the navigation property namedLocations in identity
[**IdentityConditionalAccessUpdatePolicies**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessUpdatePolicies) | **Patch** /identity/conditionalAccess/policies/{conditionalAccessPolicy-id} | Update the navigation property policies in identity
[**IdentityGetConditionalAccess**](IdentityConditionalAccessRootApi.md#IdentityGetConditionalAccess) | **Get** /identity/conditionalAccess | Get conditionalAccess from identity
[**IdentityUpdateConditionalAccess**](IdentityConditionalAccessRootApi.md#IdentityUpdateConditionalAccess) | **Patch** /identity/conditionalAccess | Update the navigation property conditionalAccess in identity



## IdentityConditionalAccessCreateNamedLocations

> MicrosoftGraphNamedLocation IdentityConditionalAccessCreateNamedLocations(ctx, microsoftGraphNamedLocation)

Create new navigation property to namedLocations for identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphNamedLocation** | [**MicrosoftGraphNamedLocation**](MicrosoftGraphNamedLocation.md)| New navigation property | 

### Return type

[**MicrosoftGraphNamedLocation**](microsoft.graph.namedLocation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityConditionalAccessCreatePolicies

> MicrosoftGraphConditionalAccessPolicy IdentityConditionalAccessCreatePolicies(ctx, microsoftGraphConditionalAccessPolicy)

Create new navigation property to policies for identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphConditionalAccessPolicy** | [**MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphConditionalAccessPolicy**](microsoft.graph.conditionalAccessPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityConditionalAccessGetNamedLocations

> MicrosoftGraphNamedLocation IdentityConditionalAccessGetNamedLocations(ctx, namedLocationId, optional)

Get namedLocations from identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namedLocationId** | **string**| key: namedLocation-id of namedLocation | 
 **optional** | ***IdentityConditionalAccessGetNamedLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IdentityConditionalAccessGetNamedLocationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphNamedLocation**](microsoft.graph.namedLocation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityConditionalAccessGetPolicies

> MicrosoftGraphConditionalAccessPolicy IdentityConditionalAccessGetPolicies(ctx, conditionalAccessPolicyId, optional)

Get policies from identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionalAccessPolicyId** | **string**| key: conditionalAccessPolicy-id of conditionalAccessPolicy | 
 **optional** | ***IdentityConditionalAccessGetPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IdentityConditionalAccessGetPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphConditionalAccessPolicy**](microsoft.graph.conditionalAccessPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityConditionalAccessListNamedLocations

> CollectionOfNamedLocation IdentityConditionalAccessListNamedLocations(ctx, optional)

Get namedLocations from identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IdentityConditionalAccessListNamedLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IdentityConditionalAccessListNamedLocationsOpts struct


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

[**CollectionOfNamedLocation**](Collection_of_namedLocation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityConditionalAccessListPolicies

> CollectionOfConditionalAccessPolicy IdentityConditionalAccessListPolicies(ctx, optional)

Get policies from identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IdentityConditionalAccessListPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IdentityConditionalAccessListPoliciesOpts struct


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

[**CollectionOfConditionalAccessPolicy**](Collection_of_conditionalAccessPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityConditionalAccessUpdateNamedLocations

> IdentityConditionalAccessUpdateNamedLocations(ctx, namedLocationId, microsoftGraphNamedLocation)

Update the navigation property namedLocations in identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namedLocationId** | **string**| key: namedLocation-id of namedLocation | 
**microsoftGraphNamedLocation** | [**MicrosoftGraphNamedLocation**](MicrosoftGraphNamedLocation.md)| New navigation property values | 

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


## IdentityConditionalAccessUpdatePolicies

> IdentityConditionalAccessUpdatePolicies(ctx, conditionalAccessPolicyId, microsoftGraphConditionalAccessPolicy)

Update the navigation property policies in identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionalAccessPolicyId** | **string**| key: conditionalAccessPolicy-id of conditionalAccessPolicy | 
**microsoftGraphConditionalAccessPolicy** | [**MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md)| New navigation property values | 

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


## IdentityGetConditionalAccess

> MicrosoftGraphConditionalAccessRoot IdentityGetConditionalAccess(ctx, optional)

Get conditionalAccess from identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IdentityGetConditionalAccessOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IdentityGetConditionalAccessOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphConditionalAccessRoot**](microsoft.graph.conditionalAccessRoot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityUpdateConditionalAccess

> IdentityUpdateConditionalAccess(ctx, microsoftGraphConditionalAccessRoot)

Update the navigation property conditionalAccess in identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphConditionalAccessRoot** | [**MicrosoftGraphConditionalAccessRoot**](MicrosoftGraphConditionalAccessRoot.md)| New navigation property values | 

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

