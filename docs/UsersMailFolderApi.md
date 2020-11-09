# \UsersMailFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateMailFolders**](UsersMailFolderApi.md#UsersCreateMailFolders) | **Post** /users/{user-id}/mailFolders | Create new navigation property to mailFolders for users
[**UsersGetMailFolders**](UsersMailFolderApi.md#UsersGetMailFolders) | **Get** /users/{user-id}/mailFolders/{mailFolder-id} | Get mailFolders from users
[**UsersListMailFolders**](UsersMailFolderApi.md#UsersListMailFolders) | **Get** /users/{user-id}/mailFolders | Get mailFolders from users
[**UsersMailFoldersCreateChildFolders**](UsersMailFolderApi.md#UsersMailFoldersCreateChildFolders) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/childFolders | Create new navigation property to childFolders for users
[**UsersMailFoldersCreateMessageRules**](UsersMailFolderApi.md#UsersMailFoldersCreateMessageRules) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messageRules | Create new navigation property to messageRules for users
[**UsersMailFoldersCreateMessages**](UsersMailFolderApi.md#UsersMailFoldersCreateMessages) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages | Create new navigation property to messages for users
[**UsersMailFoldersCreateMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersCreateMultiValueExtendedProperties) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersMailFoldersCreateSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersCreateSingleValueExtendedProperties) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersMailFoldersGetChildFolders**](UsersMailFolderApi.md#UsersMailFoldersGetChildFolders) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Get childFolders from users
[**UsersMailFoldersGetMessageRules**](UsersMailFolderApi.md#UsersMailFoldersGetMessageRules) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Get messageRules from users
[**UsersMailFoldersGetMessages**](UsersMailFolderApi.md#UsersMailFoldersGetMessages) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id} | Get messages from users
[**UsersMailFoldersGetMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersGetMultiValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from users
[**UsersMailFoldersGetSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersGetSingleValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from users
[**UsersMailFoldersListChildFolders**](UsersMailFolderApi.md#UsersMailFoldersListChildFolders) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/childFolders | Get childFolders from users
[**UsersMailFoldersListMessageRules**](UsersMailFolderApi.md#UsersMailFoldersListMessageRules) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messageRules | Get messageRules from users
[**UsersMailFoldersListMessages**](UsersMailFolderApi.md#UsersMailFoldersListMessages) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages | Get messages from users
[**UsersMailFoldersListMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersListMultiValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersMailFoldersListSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersListSingleValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersMailFoldersMessagesCreateAttachments**](UsersMailFolderApi.md#UsersMailFoldersMessagesCreateAttachments) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/attachments | Create new navigation property to attachments for users
[**UsersMailFoldersMessagesCreateExtensions**](UsersMailFolderApi.md#UsersMailFoldersMessagesCreateExtensions) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions | Create new navigation property to extensions for users
[**UsersMailFoldersMessagesCreateMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesCreateMultiValueExtendedProperties) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersMailFoldersMessagesCreateSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesCreateSingleValueExtendedProperties) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersMailFoldersMessagesGetAttachments**](UsersMailFolderApi.md#UsersMailFoldersMessagesGetAttachments) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Get attachments from users
[**UsersMailFoldersMessagesGetExtensions**](UsersMailFolderApi.md#UsersMailFoldersMessagesGetExtensions) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Get extensions from users
[**UsersMailFoldersMessagesGetMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesGetMultiValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from users
[**UsersMailFoldersMessagesGetSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesGetSingleValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from users
[**UsersMailFoldersMessagesListAttachments**](UsersMailFolderApi.md#UsersMailFoldersMessagesListAttachments) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/attachments | Get attachments from users
[**UsersMailFoldersMessagesListExtensions**](UsersMailFolderApi.md#UsersMailFoldersMessagesListExtensions) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions | Get extensions from users
[**UsersMailFoldersMessagesListMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesListMultiValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersMailFoldersMessagesListSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesListSingleValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersMailFoldersMessagesUpdateAttachments**](UsersMailFolderApi.md#UsersMailFoldersMessagesUpdateAttachments) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Update the navigation property attachments in users
[**UsersMailFoldersMessagesUpdateExtensions**](UsersMailFolderApi.md#UsersMailFoldersMessagesUpdateExtensions) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersMailFoldersMessagesUpdateMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesUpdateMultiValueExtendedProperties) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in users
[**UsersMailFoldersMessagesUpdateSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesUpdateSingleValueExtendedProperties) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in users
[**UsersMailFoldersUpdateChildFolders**](UsersMailFolderApi.md#UsersMailFoldersUpdateChildFolders) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Update the navigation property childFolders in users
[**UsersMailFoldersUpdateMessageRules**](UsersMailFolderApi.md#UsersMailFoldersUpdateMessageRules) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Update the navigation property messageRules in users
[**UsersMailFoldersUpdateMessages**](UsersMailFolderApi.md#UsersMailFoldersUpdateMessages) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id} | Update the navigation property messages in users
[**UsersMailFoldersUpdateMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersUpdateMultiValueExtendedProperties) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in users
[**UsersMailFoldersUpdateSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersUpdateSingleValueExtendedProperties) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in users
[**UsersUpdateMailFolders**](UsersMailFolderApi.md#UsersUpdateMailFolders) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id} | Update the navigation property mailFolders in users



## UsersCreateMailFolders

> MicrosoftGraphMailFolder UsersCreateMailFolders(ctx, userId, microsoftGraphMailFolder)

Create new navigation property to mailFolders for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)| New navigation property | 

### Return type

[**MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetMailFolders

> MicrosoftGraphMailFolder UsersGetMailFolders(ctx, userId, mailFolderId, optional)

Get mailFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersGetMailFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetMailFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListMailFolders

> CollectionOfMailFolder UsersListMailFolders(ctx, userId, optional)

Get mailFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListMailFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListMailFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfMailFolder**](Collection_of_mailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersCreateChildFolders

> MicrosoftGraphMailFolder UsersMailFoldersCreateChildFolders(ctx, userId, mailFolderId, microsoftGraphMailFolder)

Create new navigation property to childFolders for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)| New navigation property | 

### Return type

[**MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersCreateMessageRules

> MicrosoftGraphMessageRule UsersMailFoldersCreateMessageRules(ctx, userId, mailFolderId, microsoftGraphMessageRule)

Create new navigation property to messageRules for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**microsoftGraphMessageRule** | [**MicrosoftGraphMessageRule**](MicrosoftGraphMessageRule.md)| New navigation property | 

### Return type

[**MicrosoftGraphMessageRule**](microsoft.graph.messageRule.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersCreateMessages

> MicrosoftGraphMessage UsersMailFoldersCreateMessages(ctx, userId, mailFolderId, microsoftGraphMessage)

Create new navigation property to messages for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersCreateMultiValueExtendedProperties(ctx, userId, mailFolderId, microsoftGraphMultiValueLegacyExtendedProperty)

Create new navigation property to multiValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersCreateSingleValueExtendedProperties(ctx, userId, mailFolderId, microsoftGraphSingleValueLegacyExtendedProperty)

Create new navigation property to singleValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersGetChildFolders

> MicrosoftGraphMailFolder UsersMailFoldersGetChildFolders(ctx, userId, mailFolderId, mailFolderId1, optional)

Get childFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**mailFolderId1** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersGetChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersGetChildFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersGetMessageRules

> MicrosoftGraphMessageRule UsersMailFoldersGetMessageRules(ctx, userId, mailFolderId, messageRuleId, optional)

Get messageRules from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageRuleId** | **string**| key: messageRule-id of messageRule | 
 **optional** | ***UsersMailFoldersGetMessageRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersGetMessageRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphMessageRule**](microsoft.graph.messageRule.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersGetMessages

> MicrosoftGraphMessage UsersMailFoldersGetMessages(ctx, userId, mailFolderId, messageId, optional)

Get messages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMailFoldersGetMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersGetMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

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


## UsersMailFoldersGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersGetMultiValueExtendedProperties(ctx, userId, mailFolderId, multiValueLegacyExtendedPropertyId, optional)

Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***UsersMailFoldersGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersGetMultiValueExtendedPropertiesOpts struct


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


## UsersMailFoldersGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersGetSingleValueExtendedProperties(ctx, userId, mailFolderId, singleValueLegacyExtendedPropertyId, optional)

Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***UsersMailFoldersGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersGetSingleValueExtendedPropertiesOpts struct


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


## UsersMailFoldersListChildFolders

> CollectionOfMailFolder UsersMailFoldersListChildFolders(ctx, userId, mailFolderId, optional)

Get childFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersListChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersListChildFoldersOpts struct


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

[**CollectionOfMailFolder**](Collection_of_mailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersListMessageRules

> CollectionOfMessageRule UsersMailFoldersListMessageRules(ctx, userId, mailFolderId, optional)

Get messageRules from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersListMessageRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersListMessageRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfMessageRule**](Collection_of_messageRule.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersListMessages

> CollectionOfMessage UsersMailFoldersListMessages(ctx, userId, mailFolderId, optional)

Get messages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersListMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersListMessagesOpts struct


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

[**CollectionOfMessage**](Collection_of_message.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersMailFoldersListMultiValueExtendedProperties(ctx, userId, mailFolderId, optional)

Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersListMultiValueExtendedPropertiesOpts struct


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


## UsersMailFoldersListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersMailFoldersListSingleValueExtendedProperties(ctx, userId, mailFolderId, optional)

Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersListSingleValueExtendedPropertiesOpts struct


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


## UsersMailFoldersMessagesCreateAttachments

> MicrosoftGraphAttachment UsersMailFoldersMessagesCreateAttachments(ctx, userId, mailFolderId, messageId, microsoftGraphAttachment)

Create new navigation property to attachments for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersMessagesCreateExtensions

> MicrosoftGraphExtension UsersMailFoldersMessagesCreateExtensions(ctx, userId, mailFolderId, messageId, microsoftGraphExtension)

Create new navigation property to extensions for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersMessagesCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersMessagesCreateMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, microsoftGraphMultiValueLegacyExtendedProperty)

Create new navigation property to multiValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersMessagesCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersMessagesCreateSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, microsoftGraphSingleValueLegacyExtendedProperty)

Create new navigation property to singleValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersMessagesGetAttachments

> MicrosoftGraphAttachment UsersMailFoldersMessagesGetAttachments(ctx, userId, mailFolderId, messageId, attachmentId, optional)

Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***UsersMailFoldersMessagesGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesGetAttachmentsOpts struct


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


## UsersMailFoldersMessagesGetExtensions

> MicrosoftGraphExtension UsersMailFoldersMessagesGetExtensions(ctx, userId, mailFolderId, messageId, extensionId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***UsersMailFoldersMessagesGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesGetExtensionsOpts struct


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


## UsersMailFoldersMessagesGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersMessagesGetMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId, optional)

Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***UsersMailFoldersMessagesGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesGetMultiValueExtendedPropertiesOpts struct


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


## UsersMailFoldersMessagesGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersMessagesGetSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId, optional)

Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***UsersMailFoldersMessagesGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesGetSingleValueExtendedPropertiesOpts struct


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


## UsersMailFoldersMessagesListAttachments

> CollectionOfAttachment UsersMailFoldersMessagesListAttachments(ctx, userId, mailFolderId, messageId, optional)

Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMailFoldersMessagesListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesListAttachmentsOpts struct


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


## UsersMailFoldersMessagesListExtensions

> CollectionOfExtension UsersMailFoldersMessagesListExtensions(ctx, userId, mailFolderId, messageId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMailFoldersMessagesListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesListExtensionsOpts struct


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


## UsersMailFoldersMessagesListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersMailFoldersMessagesListMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, optional)

Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMailFoldersMessagesListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesListMultiValueExtendedPropertiesOpts struct


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


## UsersMailFoldersMessagesListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersMailFoldersMessagesListSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, optional)

Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMailFoldersMessagesListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesListSingleValueExtendedPropertiesOpts struct


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


## UsersMailFoldersMessagesUpdateAttachments

> UsersMailFoldersMessagesUpdateAttachments(ctx, userId, mailFolderId, messageId, attachmentId, microsoftGraphAttachment)

Update the navigation property attachments in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersMessagesUpdateExtensions

> UsersMailFoldersMessagesUpdateExtensions(ctx, userId, mailFolderId, messageId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersMessagesUpdateMultiValueExtendedProperties

> UsersMailFoldersMessagesUpdateMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)

Update the navigation property multiValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersMessagesUpdateSingleValueExtendedProperties

> UsersMailFoldersMessagesUpdateSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)

Update the navigation property singleValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersUpdateChildFolders

> UsersMailFoldersUpdateChildFolders(ctx, userId, mailFolderId, mailFolderId1, microsoftGraphMailFolder)

Update the navigation property childFolders in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**mailFolderId1** | **string**| key: mailFolder-id of mailFolder | 
**microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)| New navigation property values | 

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


## UsersMailFoldersUpdateMessageRules

> UsersMailFoldersUpdateMessageRules(ctx, userId, mailFolderId, messageRuleId, microsoftGraphMessageRule)

Update the navigation property messageRules in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageRuleId** | **string**| key: messageRule-id of messageRule | 
**microsoftGraphMessageRule** | [**MicrosoftGraphMessageRule**](MicrosoftGraphMessageRule.md)| New navigation property values | 

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


## UsersMailFoldersUpdateMessages

> UsersMailFoldersUpdateMessages(ctx, userId, mailFolderId, messageId, microsoftGraphMessage)

Update the navigation property messages in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersUpdateMultiValueExtendedProperties

> UsersMailFoldersUpdateMultiValueExtendedProperties(ctx, userId, mailFolderId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)

Update the navigation property multiValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersMailFoldersUpdateSingleValueExtendedProperties

> UsersMailFoldersUpdateSingleValueExtendedProperties(ctx, userId, mailFolderId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)

Update the navigation property singleValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## UsersUpdateMailFolders

> UsersUpdateMailFolders(ctx, userId, mailFolderId, microsoftGraphMailFolder)

Update the navigation property mailFolders in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)| New navigation property values | 

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

