# \UsersMessageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateMessages**](UsersMessageApi.md#UsersCreateMessages) | **Post** /users/{user-id}/messages | Create new navigation property to messages for users
[**UsersGetMessages**](UsersMessageApi.md#UsersGetMessages) | **Get** /users/{user-id}/messages/{message-id} | Get messages from users
[**UsersListMessages**](UsersMessageApi.md#UsersListMessages) | **Get** /users/{user-id}/messages | Get messages from users
[**UsersMessagesCreateAttachments**](UsersMessageApi.md#UsersMessagesCreateAttachments) | **Post** /users/{user-id}/messages/{message-id}/attachments | Create new navigation property to attachments for users
[**UsersMessagesCreateExtensions**](UsersMessageApi.md#UsersMessagesCreateExtensions) | **Post** /users/{user-id}/messages/{message-id}/extensions | Create new navigation property to extensions for users
[**UsersMessagesCreateMultiValueExtendedProperties**](UsersMessageApi.md#UsersMessagesCreateMultiValueExtendedProperties) | **Post** /users/{user-id}/messages/{message-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersMessagesCreateSingleValueExtendedProperties**](UsersMessageApi.md#UsersMessagesCreateSingleValueExtendedProperties) | **Post** /users/{user-id}/messages/{message-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersMessagesGetAttachments**](UsersMessageApi.md#UsersMessagesGetAttachments) | **Get** /users/{user-id}/messages/{message-id}/attachments/{attachment-id} | Get attachments from users
[**UsersMessagesGetExtensions**](UsersMessageApi.md#UsersMessagesGetExtensions) | **Get** /users/{user-id}/messages/{message-id}/extensions/{extension-id} | Get extensions from users
[**UsersMessagesGetMultiValueExtendedProperties**](UsersMessageApi.md#UsersMessagesGetMultiValueExtendedProperties) | **Get** /users/{user-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from users
[**UsersMessagesGetSingleValueExtendedProperties**](UsersMessageApi.md#UsersMessagesGetSingleValueExtendedProperties) | **Get** /users/{user-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from users
[**UsersMessagesListAttachments**](UsersMessageApi.md#UsersMessagesListAttachments) | **Get** /users/{user-id}/messages/{message-id}/attachments | Get attachments from users
[**UsersMessagesListExtensions**](UsersMessageApi.md#UsersMessagesListExtensions) | **Get** /users/{user-id}/messages/{message-id}/extensions | Get extensions from users
[**UsersMessagesListMultiValueExtendedProperties**](UsersMessageApi.md#UsersMessagesListMultiValueExtendedProperties) | **Get** /users/{user-id}/messages/{message-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersMessagesListSingleValueExtendedProperties**](UsersMessageApi.md#UsersMessagesListSingleValueExtendedProperties) | **Get** /users/{user-id}/messages/{message-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersMessagesUpdateAttachments**](UsersMessageApi.md#UsersMessagesUpdateAttachments) | **Patch** /users/{user-id}/messages/{message-id}/attachments/{attachment-id} | Update the navigation property attachments in users
[**UsersMessagesUpdateExtensions**](UsersMessageApi.md#UsersMessagesUpdateExtensions) | **Patch** /users/{user-id}/messages/{message-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersMessagesUpdateMultiValueExtendedProperties**](UsersMessageApi.md#UsersMessagesUpdateMultiValueExtendedProperties) | **Patch** /users/{user-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in users
[**UsersMessagesUpdateSingleValueExtendedProperties**](UsersMessageApi.md#UsersMessagesUpdateSingleValueExtendedProperties) | **Patch** /users/{user-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in users
[**UsersUpdateMessages**](UsersMessageApi.md#UsersUpdateMessages) | **Patch** /users/{user-id}/messages/{message-id} | Update the navigation property messages in users



## UsersCreateMessages

> MicrosoftGraphMessage UsersCreateMessages(ctx, userId, microsoftGraphMessage)

Create new navigation property to messages for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphMessage** | [**MicrosoftGraphMessage**](MicrosoftGraphMessage.md)| New navigation property | 

### Return type

[**MicrosoftGraphMessage**](microsoft.graph.message.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetMessages

> MicrosoftGraphMessage UsersGetMessages(ctx, userId, messageId, optional)

Get messages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersGetMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphMessage**](microsoft.graph.message.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListMessages

> CollectionOfMessage UsersListMessages(ctx, userId, optional)

Get messages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListMessagesOpts struct


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

[**CollectionOfMessage**](Collection_of_message.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMessagesCreateAttachments

> MicrosoftGraphAttachment UsersMessagesCreateAttachments(ctx, userId, messageId, microsoftGraphAttachment)

Create new navigation property to attachments for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
**microsoftGraphAttachment** | [**MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md)| New navigation property | 

### Return type

[**MicrosoftGraphAttachment**](microsoft.graph.attachment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMessagesCreateExtensions

> MicrosoftGraphExtension UsersMessagesCreateExtensions(ctx, userId, messageId, microsoftGraphExtension)

Create new navigation property to extensions for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
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


## UsersMessagesCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMessagesCreateMultiValueExtendedProperties(ctx, userId, messageId, microsoftGraphMultiValueLegacyExtendedProperty)

Create new navigation property to multiValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
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


## UsersMessagesCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMessagesCreateSingleValueExtendedProperties(ctx, userId, messageId, microsoftGraphSingleValueLegacyExtendedProperty)

Create new navigation property to singleValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
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


## UsersMessagesGetAttachments

> MicrosoftGraphAttachment UsersMessagesGetAttachments(ctx, userId, messageId, attachmentId, optional)

Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***UsersMessagesGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMessagesGetAttachmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAttachment**](microsoft.graph.attachment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMessagesGetExtensions

> MicrosoftGraphExtension UsersMessagesGetExtensions(ctx, userId, messageId, extensionId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***UsersMessagesGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMessagesGetExtensionsOpts struct


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


## UsersMessagesGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMessagesGetMultiValueExtendedProperties(ctx, userId, messageId, multiValueLegacyExtendedPropertyId, optional)

Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***UsersMessagesGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMessagesGetMultiValueExtendedPropertiesOpts struct


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


## UsersMessagesGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMessagesGetSingleValueExtendedProperties(ctx, userId, messageId, singleValueLegacyExtendedPropertyId, optional)

Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***UsersMessagesGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMessagesGetSingleValueExtendedPropertiesOpts struct


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


## UsersMessagesListAttachments

> CollectionOfAttachment UsersMessagesListAttachments(ctx, userId, messageId, optional)

Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMessagesListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMessagesListAttachmentsOpts struct


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

[**CollectionOfAttachment**](Collection_of_attachment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMessagesListExtensions

> CollectionOfExtension UsersMessagesListExtensions(ctx, userId, messageId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMessagesListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMessagesListExtensionsOpts struct


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


## UsersMessagesListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersMessagesListMultiValueExtendedProperties(ctx, userId, messageId, optional)

Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMessagesListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMessagesListMultiValueExtendedPropertiesOpts struct


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


## UsersMessagesListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersMessagesListSingleValueExtendedProperties(ctx, userId, messageId, optional)

Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMessagesListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMessagesListSingleValueExtendedPropertiesOpts struct


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


## UsersMessagesUpdateAttachments

> UsersMessagesUpdateAttachments(ctx, userId, messageId, attachmentId, microsoftGraphAttachment)

Update the navigation property attachments in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
**attachmentId** | **string**| key: attachment-id of attachment | 
**microsoftGraphAttachment** | [**MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md)| New navigation property values | 

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


## UsersMessagesUpdateExtensions

> UsersMessagesUpdateExtensions(ctx, userId, messageId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
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


## UsersMessagesUpdateMultiValueExtendedProperties

> UsersMessagesUpdateMultiValueExtendedProperties(ctx, userId, messageId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)

Update the navigation property multiValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
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


## UsersMessagesUpdateSingleValueExtendedProperties

> UsersMessagesUpdateSingleValueExtendedProperties(ctx, userId, messageId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)

Update the navigation property singleValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
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


## UsersUpdateMessages

> UsersUpdateMessages(ctx, userId, messageId, microsoftGraphMessage)

Update the navigation property messages in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**messageId** | **string**| key: message-id of message | 
**microsoftGraphMessage** | [**MicrosoftGraphMessage**](MicrosoftGraphMessage.md)| New navigation property values | 

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

