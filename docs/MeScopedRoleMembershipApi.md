# \MeScopedRoleMembershipApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateScopedRoleMemberOf**](MeScopedRoleMembershipApi.md#MeCreateScopedRoleMemberOf) | **Post** /me/scopedRoleMemberOf | Create new navigation property to scopedRoleMemberOf for me
[**MeGetScopedRoleMemberOf**](MeScopedRoleMembershipApi.md#MeGetScopedRoleMemberOf) | **Get** /me/scopedRoleMemberOf/{scopedRoleMembership-id} | Get scopedRoleMemberOf from me
[**MeListScopedRoleMemberOf**](MeScopedRoleMembershipApi.md#MeListScopedRoleMemberOf) | **Get** /me/scopedRoleMemberOf | Get scopedRoleMemberOf from me
[**MeUpdateScopedRoleMemberOf**](MeScopedRoleMembershipApi.md#MeUpdateScopedRoleMemberOf) | **Patch** /me/scopedRoleMemberOf/{scopedRoleMembership-id} | Update the navigation property scopedRoleMemberOf in me



## MeCreateScopedRoleMemberOf

> MicrosoftGraphScopedRoleMembership MeCreateScopedRoleMemberOf(ctx, microsoftGraphScopedRoleMembership)

Create new navigation property to scopedRoleMemberOf for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New navigation property | 

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


## MeGetScopedRoleMemberOf

> MicrosoftGraphScopedRoleMembership MeGetScopedRoleMemberOf(ctx, scopedRoleMembershipId, optional)

Get scopedRoleMemberOf from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
 **optional** | ***MeGetScopedRoleMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetScopedRoleMemberOfOpts struct


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


## MeListScopedRoleMemberOf

> CollectionOfScopedRoleMembership MeListScopedRoleMemberOf(ctx, optional)

Get scopedRoleMemberOf from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListScopedRoleMemberOfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListScopedRoleMemberOfOpts struct


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


## MeUpdateScopedRoleMemberOf

> MeUpdateScopedRoleMemberOf(ctx, scopedRoleMembershipId, microsoftGraphScopedRoleMembership)

Update the navigation property scopedRoleMemberOf in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New navigation property values | 

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

