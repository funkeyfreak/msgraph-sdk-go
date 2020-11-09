# \UsersContactApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersContactsCreateExtensions**](UsersContactApi.md#UsersContactsCreateExtensions) | **Post** /users/{user-id}/contacts/{contact-id}/extensions | Create new navigation property to extensions for users
[**UsersContactsCreateMultiValueExtendedProperties**](UsersContactApi.md#UsersContactsCreateMultiValueExtendedProperties) | **Post** /users/{user-id}/contacts/{contact-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersContactsCreateSingleValueExtendedProperties**](UsersContactApi.md#UsersContactsCreateSingleValueExtendedProperties) | **Post** /users/{user-id}/contacts/{contact-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersContactsGetExtensions**](UsersContactApi.md#UsersContactsGetExtensions) | **Get** /users/{user-id}/contacts/{contact-id}/extensions/{extension-id} | Get extensions from users
[**UsersContactsGetMultiValueExtendedProperties**](UsersContactApi.md#UsersContactsGetMultiValueExtendedProperties) | **Get** /users/{user-id}/contacts/{contact-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from users
[**UsersContactsGetPhoto**](UsersContactApi.md#UsersContactsGetPhoto) | **Get** /users/{user-id}/contacts/{contact-id}/photo | Get photo from users
[**UsersContactsGetSingleValueExtendedProperties**](UsersContactApi.md#UsersContactsGetSingleValueExtendedProperties) | **Get** /users/{user-id}/contacts/{contact-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from users
[**UsersContactsListExtensions**](UsersContactApi.md#UsersContactsListExtensions) | **Get** /users/{user-id}/contacts/{contact-id}/extensions | Get extensions from users
[**UsersContactsListMultiValueExtendedProperties**](UsersContactApi.md#UsersContactsListMultiValueExtendedProperties) | **Get** /users/{user-id}/contacts/{contact-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersContactsListSingleValueExtendedProperties**](UsersContactApi.md#UsersContactsListSingleValueExtendedProperties) | **Get** /users/{user-id}/contacts/{contact-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersContactsUpdateExtensions**](UsersContactApi.md#UsersContactsUpdateExtensions) | **Patch** /users/{user-id}/contacts/{contact-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersContactsUpdateMultiValueExtendedProperties**](UsersContactApi.md#UsersContactsUpdateMultiValueExtendedProperties) | **Patch** /users/{user-id}/contacts/{contact-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in users
[**UsersContactsUpdatePhoto**](UsersContactApi.md#UsersContactsUpdatePhoto) | **Patch** /users/{user-id}/contacts/{contact-id}/photo | Update the navigation property photo in users
[**UsersContactsUpdateSingleValueExtendedProperties**](UsersContactApi.md#UsersContactsUpdateSingleValueExtendedProperties) | **Patch** /users/{user-id}/contacts/{contact-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in users
[**UsersCreateContacts**](UsersContactApi.md#UsersCreateContacts) | **Post** /users/{user-id}/contacts | Create new navigation property to contacts for users
[**UsersGetContacts**](UsersContactApi.md#UsersGetContacts) | **Get** /users/{user-id}/contacts/{contact-id} | Get contacts from users
[**UsersListContacts**](UsersContactApi.md#UsersListContacts) | **Get** /users/{user-id}/contacts | Get contacts from users
[**UsersUpdateContacts**](UsersContactApi.md#UsersUpdateContacts) | **Patch** /users/{user-id}/contacts/{contact-id} | Update the navigation property contacts in users



## UsersContactsCreateExtensions

> MicrosoftGraphExtension UsersContactsCreateExtensions(ctx, userId, contactId, microsoftGraphExtension)

Create new navigation property to extensions for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)| New navigation property | 

### Return type

[**MicrosoftGraphExtension**](microsoft.graph.extension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersContactsCreateMultiValueExtendedProperties(ctx, userId, contactId, microsoftGraphMultiValueLegacyExtendedProperty)

Create new navigation property to multiValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)| New navigation property | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersContactsCreateSingleValueExtendedProperties(ctx, userId, contactId, microsoftGraphSingleValueLegacyExtendedProperty)

Create new navigation property to singleValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)| New navigation property | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsGetExtensions

> MicrosoftGraphExtension UsersContactsGetExtensions(ctx, userId, contactId, extensionId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***UsersContactsGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactsGetExtensionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphExtension**](microsoft.graph.extension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersContactsGetMultiValueExtendedProperties(ctx, userId, contactId, multiValueLegacyExtendedPropertyId, optional)

Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***UsersContactsGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactsGetMultiValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsGetPhoto

> MicrosoftGraphProfilePhoto UsersContactsGetPhoto(ctx, userId, contactId, optional)

Get photo from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersContactsGetPhotoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactsGetPhotoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphProfilePhoto**](microsoft.graph.profilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersContactsGetSingleValueExtendedProperties(ctx, userId, contactId, singleValueLegacyExtendedPropertyId, optional)

Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***UsersContactsGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactsGetSingleValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsListExtensions

> CollectionOfExtension UsersContactsListExtensions(ctx, userId, contactId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersContactsListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactsListExtensionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**CollectionOfExtension**](Collection_of_extension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersContactsListMultiValueExtendedProperties(ctx, userId, contactId, optional)

Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersContactsListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactsListMultiValueExtendedPropertiesOpts struct


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

[**CollectionOfMultiValueLegacyExtendedProperty**](Collection_of_multiValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersContactsListSingleValueExtendedProperties(ctx, userId, contactId, optional)

Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersContactsListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactsListSingleValueExtendedPropertiesOpts struct


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

[**CollectionOfSingleValueLegacyExtendedProperty**](Collection_of_singleValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsUpdateExtensions

> UsersContactsUpdateExtensions(ctx, userId, contactId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**extensionId** | **string**| key: extension-id of extension | 
**microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)| New navigation property values | 

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


## UsersContactsUpdateMultiValueExtendedProperties

> UsersContactsUpdateMultiValueExtendedProperties(ctx, userId, contactId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)

Update the navigation property multiValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
**microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)| New navigation property values | 

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


## UsersContactsUpdatePhoto

> UsersContactsUpdatePhoto(ctx, userId, contactId, microsoftGraphProfilePhoto)

Update the navigation property photo in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)| New navigation property values | 

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


## UsersContactsUpdateSingleValueExtendedProperties

> UsersContactsUpdateSingleValueExtendedProperties(ctx, userId, contactId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)

Update the navigation property singleValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
**microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)| New navigation property values | 

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


## UsersCreateContacts

> MicrosoftGraphContact UsersCreateContacts(ctx, userId, microsoftGraphContact)

Create new navigation property to contacts for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphContact** | [**MicrosoftGraphContact**](MicrosoftGraphContact.md)| New navigation property | 

### Return type

[**MicrosoftGraphContact**](microsoft.graph.contact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetContacts

> MicrosoftGraphContact UsersGetContacts(ctx, userId, contactId, optional)

Get contacts from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersGetContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphContact**](microsoft.graph.contact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListContacts

> CollectionOfContact UsersListContacts(ctx, userId, optional)

Get contacts from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListContactsOpts struct


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

[**CollectionOfContact**](Collection_of_contact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateContacts

> UsersUpdateContacts(ctx, userId, contactId, microsoftGraphContact)

Update the navigation property contacts in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphContact** | [**MicrosoftGraphContact**](MicrosoftGraphContact.md)| New navigation property values | 

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

