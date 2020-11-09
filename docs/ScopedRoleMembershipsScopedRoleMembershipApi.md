# \ScopedRoleMembershipsScopedRoleMembershipApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership) | **Post** /scopedRoleMemberships | Add new entity to scopedRoleMemberships
[**ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership) | **Delete** /scopedRoleMemberships/{scopedRoleMembership-id} | Delete entity from scopedRoleMemberships
[**ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership) | **Get** /scopedRoleMemberships/{scopedRoleMembership-id} | Get entity from scopedRoleMemberships by key
[**ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership) | **Get** /scopedRoleMemberships | Get entities from scopedRoleMemberships
[**ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership) | **Patch** /scopedRoleMemberships/{scopedRoleMembership-id} | Update entity in scopedRoleMemberships



## ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership

> MicrosoftGraphScopedRoleMembership ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership(ctx, microsoftGraphScopedRoleMembership)

Add new entity to scopedRoleMemberships

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New entity | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](microsoft.graph.scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership

> ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership(ctx, scopedRoleMembershipId, optional)

Delete entity from scopedRoleMemberships

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
 **optional** | ***ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembershipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembershipOpts struct


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


## ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership

> MicrosoftGraphScopedRoleMembership ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership(ctx, scopedRoleMembershipId, optional)

Get entity from scopedRoleMemberships by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
 **optional** | ***ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembershipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembershipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](microsoft.graph.scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership

> CollectionOfScopedRoleMembership ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership(ctx, optional)

Get entities from scopedRoleMemberships

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembershipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembershipOpts struct


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

[**CollectionOfScopedRoleMembership**](Collection_of_scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership

> ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership(ctx, scopedRoleMembershipId, microsoftGraphScopedRoleMembership)

Update entity in scopedRoleMemberships

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New property values | 

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

