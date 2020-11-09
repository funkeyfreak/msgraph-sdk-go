# \MeMailFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateMailFolders**](MeMailFolderApi.md#MeCreateMailFolders) | **Post** /me/mailFolders | Create new navigation property to mailFolders for me
[**MeGetMailFolders**](MeMailFolderApi.md#MeGetMailFolders) | **Get** /me/mailFolders/{mailFolder-id} | Get mailFolders from me
[**MeListMailFolders**](MeMailFolderApi.md#MeListMailFolders) | **Get** /me/mailFolders | Get mailFolders from me
[**MeMailFoldersCreateChildFolders**](MeMailFolderApi.md#MeMailFoldersCreateChildFolders) | **Post** /me/mailFolders/{mailFolder-id}/childFolders | Create new navigation property to childFolders for me
[**MeMailFoldersCreateMessageRules**](MeMailFolderApi.md#MeMailFoldersCreateMessageRules) | **Post** /me/mailFolders/{mailFolder-id}/messageRules | Create new navigation property to messageRules for me
[**MeMailFoldersCreateMessages**](MeMailFolderApi.md#MeMailFoldersCreateMessages) | **Post** /me/mailFolders/{mailFolder-id}/messages | Create new navigation property to messages for me
[**MeMailFoldersCreateMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersCreateMultiValueExtendedProperties) | **Post** /me/mailFolders/{mailFolder-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for me
[**MeMailFoldersCreateSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersCreateSingleValueExtendedProperties) | **Post** /me/mailFolders/{mailFolder-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for me
[**MeMailFoldersGetChildFolders**](MeMailFolderApi.md#MeMailFoldersGetChildFolders) | **Get** /me/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Get childFolders from me
[**MeMailFoldersGetMessageRules**](MeMailFolderApi.md#MeMailFoldersGetMessageRules) | **Get** /me/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Get messageRules from me
[**MeMailFoldersGetMessages**](MeMailFolderApi.md#MeMailFoldersGetMessages) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id} | Get messages from me
[**MeMailFoldersGetMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersGetMultiValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from me
[**MeMailFoldersGetSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersGetSingleValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from me
[**MeMailFoldersListChildFolders**](MeMailFolderApi.md#MeMailFoldersListChildFolders) | **Get** /me/mailFolders/{mailFolder-id}/childFolders | Get childFolders from me
[**MeMailFoldersListMessageRules**](MeMailFolderApi.md#MeMailFoldersListMessageRules) | **Get** /me/mailFolders/{mailFolder-id}/messageRules | Get messageRules from me
[**MeMailFoldersListMessages**](MeMailFolderApi.md#MeMailFoldersListMessages) | **Get** /me/mailFolders/{mailFolder-id}/messages | Get messages from me
[**MeMailFoldersListMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersListMultiValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from me
[**MeMailFoldersListSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersListSingleValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from me
[**MeMailFoldersMessagesCreateAttachments**](MeMailFolderApi.md#MeMailFoldersMessagesCreateAttachments) | **Post** /me/mailFolders/{mailFolder-id}/messages/{message-id}/attachments | Create new navigation property to attachments for me
[**MeMailFoldersMessagesCreateExtensions**](MeMailFolderApi.md#MeMailFoldersMessagesCreateExtensions) | **Post** /me/mailFolders/{mailFolder-id}/messages/{message-id}/extensions | Create new navigation property to extensions for me
[**MeMailFoldersMessagesCreateMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesCreateMultiValueExtendedProperties) | **Post** /me/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for me
[**MeMailFoldersMessagesCreateSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesCreateSingleValueExtendedProperties) | **Post** /me/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for me
[**MeMailFoldersMessagesGetAttachments**](MeMailFolderApi.md#MeMailFoldersMessagesGetAttachments) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Get attachments from me
[**MeMailFoldersMessagesGetExtensions**](MeMailFolderApi.md#MeMailFoldersMessagesGetExtensions) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Get extensions from me
[**MeMailFoldersMessagesGetMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesGetMultiValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from me
[**MeMailFoldersMessagesGetSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesGetSingleValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from me
[**MeMailFoldersMessagesListAttachments**](MeMailFolderApi.md#MeMailFoldersMessagesListAttachments) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/attachments | Get attachments from me
[**MeMailFoldersMessagesListExtensions**](MeMailFolderApi.md#MeMailFoldersMessagesListExtensions) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/extensions | Get extensions from me
[**MeMailFoldersMessagesListMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesListMultiValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from me
[**MeMailFoldersMessagesListSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesListSingleValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from me
[**MeMailFoldersMessagesUpdateAttachments**](MeMailFolderApi.md#MeMailFoldersMessagesUpdateAttachments) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Update the navigation property attachments in me
[**MeMailFoldersMessagesUpdateExtensions**](MeMailFolderApi.md#MeMailFoldersMessagesUpdateExtensions) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Update the navigation property extensions in me
[**MeMailFoldersMessagesUpdateMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesUpdateMultiValueExtendedProperties) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in me
[**MeMailFoldersMessagesUpdateSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesUpdateSingleValueExtendedProperties) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in me
[**MeMailFoldersUpdateChildFolders**](MeMailFolderApi.md#MeMailFoldersUpdateChildFolders) | **Patch** /me/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Update the navigation property childFolders in me
[**MeMailFoldersUpdateMessageRules**](MeMailFolderApi.md#MeMailFoldersUpdateMessageRules) | **Patch** /me/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Update the navigation property messageRules in me
[**MeMailFoldersUpdateMessages**](MeMailFolderApi.md#MeMailFoldersUpdateMessages) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id} | Update the navigation property messages in me
[**MeMailFoldersUpdateMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersUpdateMultiValueExtendedProperties) | **Patch** /me/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in me
[**MeMailFoldersUpdateSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersUpdateSingleValueExtendedProperties) | **Patch** /me/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in me
[**MeUpdateMailFolders**](MeMailFolderApi.md#MeUpdateMailFolders) | **Patch** /me/mailFolders/{mailFolder-id} | Update the navigation property mailFolders in me



## MeCreateMailFolders

> MicrosoftGraphMailFolder MeCreateMailFolders(ctx, microsoftGraphMailFolder)

Create new navigation property to mailFolders for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeGetMailFolders

> MicrosoftGraphMailFolder MeGetMailFolders(ctx, mailFolderId, optional)

Get mailFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeGetMailFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetMailFoldersOpts struct


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


## MeListMailFolders

> CollectionOfMailFolder MeListMailFolders(ctx, optional)

Get mailFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListMailFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListMailFoldersOpts struct


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


## MeMailFoldersCreateChildFolders

> MicrosoftGraphMailFolder MeMailFoldersCreateChildFolders(ctx, mailFolderId, microsoftGraphMailFolder)

Create new navigation property to childFolders for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersCreateMessageRules

> MicrosoftGraphMessageRule MeMailFoldersCreateMessageRules(ctx, mailFolderId, microsoftGraphMessageRule)

Create new navigation property to messageRules for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersCreateMessages

> MicrosoftGraphMessage MeMailFoldersCreateMessages(ctx, mailFolderId, microsoftGraphMessage)

Create new navigation property to messages for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMailFoldersCreateMultiValueExtendedProperties(ctx, mailFolderId, microsoftGraphMultiValueLegacyExtendedProperty)

Create new navigation property to multiValueExtendedProperties for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMailFoldersCreateSingleValueExtendedProperties(ctx, mailFolderId, microsoftGraphSingleValueLegacyExtendedProperty)

Create new navigation property to singleValueExtendedProperties for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersGetChildFolders

> MicrosoftGraphMailFolder MeMailFoldersGetChildFolders(ctx, mailFolderId, mailFolderId1, optional)

Get childFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**mailFolderId1** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersGetChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersGetChildFoldersOpts struct


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


## MeMailFoldersGetMessageRules

> MicrosoftGraphMessageRule MeMailFoldersGetMessageRules(ctx, mailFolderId, messageRuleId, optional)

Get messageRules from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageRuleId** | **string**| key: messageRule-id of messageRule | 
 **optional** | ***MeMailFoldersGetMessageRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersGetMessageRulesOpts struct


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


## MeMailFoldersGetMessages

> MicrosoftGraphMessage MeMailFoldersGetMessages(ctx, mailFolderId, messageId, optional)

Get messages from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***MeMailFoldersGetMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersGetMessagesOpts struct


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


## MeMailFoldersGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMailFoldersGetMultiValueExtendedProperties(ctx, mailFolderId, multiValueLegacyExtendedPropertyId, optional)

Get multiValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***MeMailFoldersGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersGetMultiValueExtendedPropertiesOpts struct


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


## MeMailFoldersGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMailFoldersGetSingleValueExtendedProperties(ctx, mailFolderId, singleValueLegacyExtendedPropertyId, optional)

Get singleValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***MeMailFoldersGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersGetSingleValueExtendedPropertiesOpts struct


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


## MeMailFoldersListChildFolders

> CollectionOfMailFolder MeMailFoldersListChildFolders(ctx, mailFolderId, optional)

Get childFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersListChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersListChildFoldersOpts struct


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


## MeMailFoldersListMessageRules

> CollectionOfMessageRule MeMailFoldersListMessageRules(ctx, mailFolderId, optional)

Get messageRules from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersListMessageRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersListMessageRulesOpts struct


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


## MeMailFoldersListMessages

> CollectionOfMessage MeMailFoldersListMessages(ctx, mailFolderId, optional)

Get messages from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersListMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersListMessagesOpts struct


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


## MeMailFoldersListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty MeMailFoldersListMultiValueExtendedProperties(ctx, mailFolderId, optional)

Get multiValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersListMultiValueExtendedPropertiesOpts struct


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


## MeMailFoldersListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty MeMailFoldersListSingleValueExtendedProperties(ctx, mailFolderId, optional)

Get singleValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersListSingleValueExtendedPropertiesOpts struct


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


## MeMailFoldersMessagesCreateAttachments

> MicrosoftGraphAttachment MeMailFoldersMessagesCreateAttachments(ctx, mailFolderId, messageId, microsoftGraphAttachment)

Create new navigation property to attachments for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersMessagesCreateExtensions

> MicrosoftGraphExtension MeMailFoldersMessagesCreateExtensions(ctx, mailFolderId, messageId, microsoftGraphExtension)

Create new navigation property to extensions for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersMessagesCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMailFoldersMessagesCreateMultiValueExtendedProperties(ctx, mailFolderId, messageId, microsoftGraphMultiValueLegacyExtendedProperty)

Create new navigation property to multiValueExtendedProperties for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersMessagesCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMailFoldersMessagesCreateSingleValueExtendedProperties(ctx, mailFolderId, messageId, microsoftGraphSingleValueLegacyExtendedProperty)

Create new navigation property to singleValueExtendedProperties for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersMessagesGetAttachments

> MicrosoftGraphAttachment MeMailFoldersMessagesGetAttachments(ctx, mailFolderId, messageId, attachmentId, optional)

Get attachments from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***MeMailFoldersMessagesGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesGetAttachmentsOpts struct


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


## MeMailFoldersMessagesGetExtensions

> MicrosoftGraphExtension MeMailFoldersMessagesGetExtensions(ctx, mailFolderId, messageId, extensionId, optional)

Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***MeMailFoldersMessagesGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesGetExtensionsOpts struct


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


## MeMailFoldersMessagesGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMailFoldersMessagesGetMultiValueExtendedProperties(ctx, mailFolderId, messageId, multiValueLegacyExtendedPropertyId, optional)

Get multiValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***MeMailFoldersMessagesGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesGetMultiValueExtendedPropertiesOpts struct


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


## MeMailFoldersMessagesGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMailFoldersMessagesGetSingleValueExtendedProperties(ctx, mailFolderId, messageId, singleValueLegacyExtendedPropertyId, optional)

Get singleValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***MeMailFoldersMessagesGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesGetSingleValueExtendedPropertiesOpts struct


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


## MeMailFoldersMessagesListAttachments

> CollectionOfAttachment MeMailFoldersMessagesListAttachments(ctx, mailFolderId, messageId, optional)

Get attachments from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***MeMailFoldersMessagesListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesListAttachmentsOpts struct


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


## MeMailFoldersMessagesListExtensions

> CollectionOfExtension MeMailFoldersMessagesListExtensions(ctx, mailFolderId, messageId, optional)

Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***MeMailFoldersMessagesListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesListExtensionsOpts struct


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


## MeMailFoldersMessagesListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty MeMailFoldersMessagesListMultiValueExtendedProperties(ctx, mailFolderId, messageId, optional)

Get multiValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***MeMailFoldersMessagesListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesListMultiValueExtendedPropertiesOpts struct


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


## MeMailFoldersMessagesListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty MeMailFoldersMessagesListSingleValueExtendedProperties(ctx, mailFolderId, messageId, optional)

Get singleValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***MeMailFoldersMessagesListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesListSingleValueExtendedPropertiesOpts struct


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


## MeMailFoldersMessagesUpdateAttachments

> MeMailFoldersMessagesUpdateAttachments(ctx, mailFolderId, messageId, attachmentId, microsoftGraphAttachment)

Update the navigation property attachments in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersMessagesUpdateExtensions

> MeMailFoldersMessagesUpdateExtensions(ctx, mailFolderId, messageId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersMessagesUpdateMultiValueExtendedProperties

> MeMailFoldersMessagesUpdateMultiValueExtendedProperties(ctx, mailFolderId, messageId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)

Update the navigation property multiValueExtendedProperties in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersMessagesUpdateSingleValueExtendedProperties

> MeMailFoldersMessagesUpdateSingleValueExtendedProperties(ctx, mailFolderId, messageId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)

Update the navigation property singleValueExtendedProperties in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersUpdateChildFolders

> MeMailFoldersUpdateChildFolders(ctx, mailFolderId, mailFolderId1, microsoftGraphMailFolder)

Update the navigation property childFolders in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersUpdateMessageRules

> MeMailFoldersUpdateMessageRules(ctx, mailFolderId, messageRuleId, microsoftGraphMessageRule)

Update the navigation property messageRules in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersUpdateMessages

> MeMailFoldersUpdateMessages(ctx, mailFolderId, messageId, microsoftGraphMessage)

Update the navigation property messages in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersUpdateMultiValueExtendedProperties

> MeMailFoldersUpdateMultiValueExtendedProperties(ctx, mailFolderId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)

Update the navigation property multiValueExtendedProperties in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersUpdateSingleValueExtendedProperties

> MeMailFoldersUpdateSingleValueExtendedProperties(ctx, mailFolderId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)

Update the navigation property singleValueExtendedProperties in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeUpdateMailFolders

> MeUpdateMailFolders(ctx, mailFolderId, microsoftGraphMailFolder)

Update the navigation property mailFolders in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

