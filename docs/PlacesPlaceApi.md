# \PlacesPlaceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlacesPlaceCreatePlace**](PlacesPlaceApi.md#PlacesPlaceCreatePlace) | **Post** /places | Add new entity to places
[**PlacesPlaceDeletePlace**](PlacesPlaceApi.md#PlacesPlaceDeletePlace) | **Delete** /places/{place-id} | Delete entity from places
[**PlacesPlaceGetPlace**](PlacesPlaceApi.md#PlacesPlaceGetPlace) | **Get** /places/{place-id} | Get entity from places by key
[**PlacesPlaceListPlace**](PlacesPlaceApi.md#PlacesPlaceListPlace) | **Get** /places | Get entities from places
[**PlacesPlaceUpdatePlace**](PlacesPlaceApi.md#PlacesPlaceUpdatePlace) | **Patch** /places/{place-id} | Update entity in places



## PlacesPlaceCreatePlace

> MicrosoftGraphPlace PlacesPlaceCreatePlace(ctx, microsoftGraphPlace)

Add new entity to places

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphPlace** | [**MicrosoftGraphPlace**](MicrosoftGraphPlace.md)| New entity | 

### Return type

[**MicrosoftGraphPlace**](microsoft.graph.place.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlacesPlaceDeletePlace

> PlacesPlaceDeletePlace(ctx, placeId, optional)

Delete entity from places

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeId** | **string**| key: place-id of place | 
 **optional** | ***PlacesPlaceDeletePlaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlacesPlaceDeletePlaceOpts struct


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


## PlacesPlaceGetPlace

> MicrosoftGraphPlace PlacesPlaceGetPlace(ctx, placeId, optional)

Get entity from places by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeId** | **string**| key: place-id of place | 
 **optional** | ***PlacesPlaceGetPlaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlacesPlaceGetPlaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlace**](microsoft.graph.place.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlacesPlaceListPlace

> CollectionOfPlace PlacesPlaceListPlace(ctx, optional)

Get entities from places

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlacesPlaceListPlaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlacesPlaceListPlaceOpts struct


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

[**CollectionOfPlace**](Collection_of_place.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlacesPlaceUpdatePlace

> PlacesPlaceUpdatePlace(ctx, placeId, microsoftGraphPlace)

Update entity in places

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeId** | **string**| key: place-id of place | 
**microsoftGraphPlace** | [**MicrosoftGraphPlace**](MicrosoftGraphPlace.md)| New property values | 

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

