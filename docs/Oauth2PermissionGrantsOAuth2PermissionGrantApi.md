# \Oauth2PermissionGrantsOAuth2PermissionGrantApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2PermissionGrantsOAuth2PermissionGrantCreateOAuth2PermissionGrant**](Oauth2PermissionGrantsOAuth2PermissionGrantApi.md#Oauth2PermissionGrantsOAuth2PermissionGrantCreateOAuth2PermissionGrant) | **Post** /oauth2PermissionGrants | Add new entity to oauth2PermissionGrants
[**Oauth2PermissionGrantsOAuth2PermissionGrantDeleteOAuth2PermissionGrant**](Oauth2PermissionGrantsOAuth2PermissionGrantApi.md#Oauth2PermissionGrantsOAuth2PermissionGrantDeleteOAuth2PermissionGrant) | **Delete** /oauth2PermissionGrants/{oAuth2PermissionGrant-id} | Delete entity from oauth2PermissionGrants
[**Oauth2PermissionGrantsOAuth2PermissionGrantGetOAuth2PermissionGrant**](Oauth2PermissionGrantsOAuth2PermissionGrantApi.md#Oauth2PermissionGrantsOAuth2PermissionGrantGetOAuth2PermissionGrant) | **Get** /oauth2PermissionGrants/{oAuth2PermissionGrant-id} | Get entity from oauth2PermissionGrants by key
[**Oauth2PermissionGrantsOAuth2PermissionGrantListOAuth2PermissionGrant**](Oauth2PermissionGrantsOAuth2PermissionGrantApi.md#Oauth2PermissionGrantsOAuth2PermissionGrantListOAuth2PermissionGrant) | **Get** /oauth2PermissionGrants | Get entities from oauth2PermissionGrants
[**Oauth2PermissionGrantsOAuth2PermissionGrantUpdateOAuth2PermissionGrant**](Oauth2PermissionGrantsOAuth2PermissionGrantApi.md#Oauth2PermissionGrantsOAuth2PermissionGrantUpdateOAuth2PermissionGrant) | **Patch** /oauth2PermissionGrants/{oAuth2PermissionGrant-id} | Update entity in oauth2PermissionGrants



## Oauth2PermissionGrantsOAuth2PermissionGrantCreateOAuth2PermissionGrant

> MicrosoftGraphOAuth2PermissionGrant Oauth2PermissionGrantsOAuth2PermissionGrantCreateOAuth2PermissionGrant(ctx, microsoftGraphOAuth2PermissionGrant)

Add new entity to oauth2PermissionGrants

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOAuth2PermissionGrant** | [**MicrosoftGraphOAuth2PermissionGrant**](MicrosoftGraphOAuth2PermissionGrant.md)| New entity | 

### Return type

[**MicrosoftGraphOAuth2PermissionGrant**](microsoft.graph.oAuth2PermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2PermissionGrantsOAuth2PermissionGrantDeleteOAuth2PermissionGrant

> Oauth2PermissionGrantsOAuth2PermissionGrantDeleteOAuth2PermissionGrant(ctx, oAuth2PermissionGrantId, optional)

Delete entity from oauth2PermissionGrants

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oAuth2PermissionGrantId** | **string**| key: oAuth2PermissionGrant-id of oAuth2PermissionGrant | 
 **optional** | ***Oauth2PermissionGrantsOAuth2PermissionGrantDeleteOAuth2PermissionGrantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Oauth2PermissionGrantsOAuth2PermissionGrantDeleteOAuth2PermissionGrantOpts struct


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


## Oauth2PermissionGrantsOAuth2PermissionGrantGetOAuth2PermissionGrant

> MicrosoftGraphOAuth2PermissionGrant Oauth2PermissionGrantsOAuth2PermissionGrantGetOAuth2PermissionGrant(ctx, oAuth2PermissionGrantId, optional)

Get entity from oauth2PermissionGrants by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oAuth2PermissionGrantId** | **string**| key: oAuth2PermissionGrant-id of oAuth2PermissionGrant | 
 **optional** | ***Oauth2PermissionGrantsOAuth2PermissionGrantGetOAuth2PermissionGrantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Oauth2PermissionGrantsOAuth2PermissionGrantGetOAuth2PermissionGrantOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOAuth2PermissionGrant**](microsoft.graph.oAuth2PermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2PermissionGrantsOAuth2PermissionGrantListOAuth2PermissionGrant

> CollectionOfOAuth2PermissionGrant Oauth2PermissionGrantsOAuth2PermissionGrantListOAuth2PermissionGrant(ctx, optional)

Get entities from oauth2PermissionGrants

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Oauth2PermissionGrantsOAuth2PermissionGrantListOAuth2PermissionGrantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Oauth2PermissionGrantsOAuth2PermissionGrantListOAuth2PermissionGrantOpts struct


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

[**CollectionOfOAuth2PermissionGrant**](Collection_of_oAuth2PermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2PermissionGrantsOAuth2PermissionGrantUpdateOAuth2PermissionGrant

> Oauth2PermissionGrantsOAuth2PermissionGrantUpdateOAuth2PermissionGrant(ctx, oAuth2PermissionGrantId, microsoftGraphOAuth2PermissionGrant)

Update entity in oauth2PermissionGrants

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oAuth2PermissionGrantId** | **string**| key: oAuth2PermissionGrant-id of oAuth2PermissionGrant | 
**microsoftGraphOAuth2PermissionGrant** | [**MicrosoftGraphOAuth2PermissionGrant**](MicrosoftGraphOAuth2PermissionGrant.md)| New property values | 

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

