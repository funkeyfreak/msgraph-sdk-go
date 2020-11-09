# \UsersPersonApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreatePeople**](UsersPersonApi.md#UsersCreatePeople) | **Post** /users/{user-id}/people | Create new navigation property to people for users
[**UsersGetPeople**](UsersPersonApi.md#UsersGetPeople) | **Get** /users/{user-id}/people/{person-id} | Get people from users
[**UsersListPeople**](UsersPersonApi.md#UsersListPeople) | **Get** /users/{user-id}/people | Get people from users
[**UsersUpdatePeople**](UsersPersonApi.md#UsersUpdatePeople) | **Patch** /users/{user-id}/people/{person-id} | Update the navigation property people in users



## UsersCreatePeople

> MicrosoftGraphPerson UsersCreatePeople(ctx, userId, microsoftGraphPerson)

Create new navigation property to people for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphPerson** | [**MicrosoftGraphPerson**](MicrosoftGraphPerson.md)| New navigation property | 

### Return type

[**MicrosoftGraphPerson**](microsoft.graph.person.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetPeople

> MicrosoftGraphPerson UsersGetPeople(ctx, userId, personId, optional)

Get people from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**personId** | **string**| key: person-id of person | 
 **optional** | ***UsersGetPeopleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetPeopleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphPerson**](microsoft.graph.person.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListPeople

> CollectionOfPerson UsersListPeople(ctx, userId, optional)

Get people from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListPeopleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListPeopleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfPerson**](Collection_of_person.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdatePeople

> UsersUpdatePeople(ctx, userId, personId, microsoftGraphPerson)

Update the navigation property people in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**personId** | **string**| key: person-id of person | 
**microsoftGraphPerson** | [**MicrosoftGraphPerson**](MicrosoftGraphPerson.md)| New navigation property values | 

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

