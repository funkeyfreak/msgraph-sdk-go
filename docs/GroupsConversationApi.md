# \GroupsConversationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsConversationsCreateThreads**](GroupsConversationApi.md#GroupsConversationsCreateThreads) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads | Create new navigation property to threads for groups
[**GroupsConversationsGetThreads**](GroupsConversationApi.md#GroupsConversationsGetThreads) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id} | Get threads from groups
[**GroupsConversationsListThreads**](GroupsConversationApi.md#GroupsConversationsListThreads) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads | Get threads from groups
[**GroupsConversationsThreadsCreatePosts**](GroupsConversationApi.md#GroupsConversationsThreadsCreatePosts) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts | Create new navigation property to posts for groups
[**GroupsConversationsThreadsGetPosts**](GroupsConversationApi.md#GroupsConversationsThreadsGetPosts) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id} | Get posts from groups
[**GroupsConversationsThreadsListPosts**](GroupsConversationApi.md#GroupsConversationsThreadsListPosts) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts | Get posts from groups
[**GroupsConversationsThreadsPostsCreateAttachments**](GroupsConversationApi.md#GroupsConversationsThreadsPostsCreateAttachments) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/attachments | Create new navigation property to attachments for groups
[**GroupsConversationsThreadsPostsCreateExtensions**](GroupsConversationApi.md#GroupsConversationsThreadsPostsCreateExtensions) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/extensions | Create new navigation property to extensions for groups
[**GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for groups
[**GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for groups
[**GroupsConversationsThreadsPostsGetAttachments**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetAttachments) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Get attachments from groups
[**GroupsConversationsThreadsPostsGetExtensions**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetExtensions) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Get extensions from groups
[**GroupsConversationsThreadsPostsGetInReplyTo**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetInReplyTo) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Get inReplyTo from groups
[**GroupsConversationsThreadsPostsGetMultiValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetMultiValueExtendedProperties) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from groups
[**GroupsConversationsThreadsPostsGetSingleValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetSingleValueExtendedProperties) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from groups
[**GroupsConversationsThreadsPostsListAttachments**](GroupsConversationApi.md#GroupsConversationsThreadsPostsListAttachments) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/attachments | Get attachments from groups
[**GroupsConversationsThreadsPostsListExtensions**](GroupsConversationApi.md#GroupsConversationsThreadsPostsListExtensions) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/extensions | Get extensions from groups
[**GroupsConversationsThreadsPostsListMultiValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsListMultiValueExtendedProperties) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from groups
[**GroupsConversationsThreadsPostsListSingleValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsListSingleValueExtendedProperties) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from groups
[**GroupsConversationsThreadsPostsUpdateAttachments**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateAttachments) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Update the navigation property attachments in groups
[**GroupsConversationsThreadsPostsUpdateExtensions**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateExtensions) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Update the navigation property extensions in groups
[**GroupsConversationsThreadsPostsUpdateInReplyTo**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateInReplyTo) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Update the navigation property inReplyTo in groups
[**GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in groups
[**GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in groups
[**GroupsConversationsThreadsUpdatePosts**](GroupsConversationApi.md#GroupsConversationsThreadsUpdatePosts) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id} | Update the navigation property posts in groups
[**GroupsConversationsUpdateThreads**](GroupsConversationApi.md#GroupsConversationsUpdateThreads) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id} | Update the navigation property threads in groups
[**GroupsCreateConversations**](GroupsConversationApi.md#GroupsCreateConversations) | **Post** /groups/{group-id}/conversations | Create new navigation property to conversations for groups
[**GroupsGetConversations**](GroupsConversationApi.md#GroupsGetConversations) | **Get** /groups/{group-id}/conversations/{conversation-id} | Get conversations from groups
[**GroupsListConversations**](GroupsConversationApi.md#GroupsListConversations) | **Get** /groups/{group-id}/conversations | Get conversations from groups
[**GroupsUpdateConversations**](GroupsConversationApi.md#GroupsUpdateConversations) | **Patch** /groups/{group-id}/conversations/{conversation-id} | Update the navigation property conversations in groups



## GroupsConversationsCreateThreads

> MicrosoftGraphConversationThread GroupsConversationsCreateThreads(ctx, groupId, conversationId, microsoftGraphConversationThread)

Create new navigation property to threads for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**microsoftGraphConversationThread** | [**MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md)| New navigation property | 

### Return type

[**MicrosoftGraphConversationThread**](microsoft.graph.conversationThread.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsGetThreads

> MicrosoftGraphConversationThread GroupsConversationsGetThreads(ctx, groupId, conversationId, conversationThreadId, optional)

Get threads from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
 **optional** | ***GroupsConversationsGetThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsGetThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphConversationThread**](microsoft.graph.conversationThread.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsListThreads

> CollectionOfConversationThread GroupsConversationsListThreads(ctx, groupId, conversationId, optional)

Get threads from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
 **optional** | ***GroupsConversationsListThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsListThreadsOpts struct


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

[**CollectionOfConversationThread**](Collection_of_conversationThread.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsCreatePosts

> MicrosoftGraphPost GroupsConversationsThreadsCreatePosts(ctx, groupId, conversationId, conversationThreadId, microsoftGraphPost)

Create new navigation property to posts for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**microsoftGraphPost** | [**MicrosoftGraphPost**](MicrosoftGraphPost.md)| New navigation property | 

### Return type

[**MicrosoftGraphPost**](microsoft.graph.post.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsGetPosts

> MicrosoftGraphPost GroupsConversationsThreadsGetPosts(ctx, groupId, conversationId, conversationThreadId, postId, optional)

Get posts from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsConversationsThreadsGetPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsGetPostsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPost**](microsoft.graph.post.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsListPosts

> CollectionOfPost GroupsConversationsThreadsListPosts(ctx, groupId, conversationId, conversationThreadId, optional)

Get posts from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
 **optional** | ***GroupsConversationsThreadsListPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsListPostsOpts struct


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

[**CollectionOfPost**](Collection_of_post.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsCreateAttachments

> MicrosoftGraphAttachment GroupsConversationsThreadsPostsCreateAttachments(ctx, groupId, conversationId, conversationThreadId, postId, microsoftGraphAttachment)

Create new navigation property to attachments for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
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


## GroupsConversationsThreadsPostsCreateExtensions

> MicrosoftGraphExtension GroupsConversationsThreadsPostsCreateExtensions(ctx, groupId, conversationId, conversationThreadId, postId, microsoftGraphExtension)

Create new navigation property to extensions for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
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


## GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, microsoftGraphMultiValueLegacyExtendedProperty)

Create new navigation property to multiValueExtendedProperties for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
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


## GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, microsoftGraphSingleValueLegacyExtendedProperty)

Create new navigation property to singleValueExtendedProperties for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
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


## GroupsConversationsThreadsPostsGetAttachments

> MicrosoftGraphAttachment GroupsConversationsThreadsPostsGetAttachments(ctx, groupId, conversationId, conversationThreadId, postId, attachmentId, optional)

Get attachments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***GroupsConversationsThreadsPostsGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsPostsGetAttachmentsOpts struct


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


## GroupsConversationsThreadsPostsGetExtensions

> MicrosoftGraphExtension GroupsConversationsThreadsPostsGetExtensions(ctx, groupId, conversationId, conversationThreadId, postId, extensionId, optional)

Get extensions from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***GroupsConversationsThreadsPostsGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsPostsGetExtensionsOpts struct


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


## GroupsConversationsThreadsPostsGetInReplyTo

> MicrosoftGraphPost GroupsConversationsThreadsPostsGetInReplyTo(ctx, groupId, conversationId, conversationThreadId, postId, optional)

Get inReplyTo from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsConversationsThreadsPostsGetInReplyToOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsPostsGetInReplyToOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPost**](microsoft.graph.post.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsConversationsThreadsPostsGetMultiValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId, optional)

Get multiValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***GroupsConversationsThreadsPostsGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsPostsGetMultiValueExtendedPropertiesOpts struct


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


## GroupsConversationsThreadsPostsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsConversationsThreadsPostsGetSingleValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId, optional)

Get singleValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***GroupsConversationsThreadsPostsGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsPostsGetSingleValueExtendedPropertiesOpts struct


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


## GroupsConversationsThreadsPostsListAttachments

> CollectionOfAttachment GroupsConversationsThreadsPostsListAttachments(ctx, groupId, conversationId, conversationThreadId, postId, optional)

Get attachments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsConversationsThreadsPostsListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsPostsListAttachmentsOpts struct


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


## GroupsConversationsThreadsPostsListExtensions

> CollectionOfExtension GroupsConversationsThreadsPostsListExtensions(ctx, groupId, conversationId, conversationThreadId, postId, optional)

Get extensions from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsConversationsThreadsPostsListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsPostsListExtensionsOpts struct


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


## GroupsConversationsThreadsPostsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty GroupsConversationsThreadsPostsListMultiValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, optional)

Get multiValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsConversationsThreadsPostsListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsPostsListMultiValueExtendedPropertiesOpts struct


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


## GroupsConversationsThreadsPostsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty GroupsConversationsThreadsPostsListSingleValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, optional)

Get singleValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsConversationsThreadsPostsListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsThreadsPostsListSingleValueExtendedPropertiesOpts struct


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


## GroupsConversationsThreadsPostsUpdateAttachments

> GroupsConversationsThreadsPostsUpdateAttachments(ctx, groupId, conversationId, conversationThreadId, postId, attachmentId, microsoftGraphAttachment)

Update the navigation property attachments in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
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


## GroupsConversationsThreadsPostsUpdateExtensions

> GroupsConversationsThreadsPostsUpdateExtensions(ctx, groupId, conversationId, conversationThreadId, postId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
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


## GroupsConversationsThreadsPostsUpdateInReplyTo

> GroupsConversationsThreadsPostsUpdateInReplyTo(ctx, groupId, conversationId, conversationThreadId, postId, microsoftGraphPost)

Update the navigation property inReplyTo in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**microsoftGraphPost** | [**MicrosoftGraphPost**](MicrosoftGraphPost.md)| New navigation property values | 

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


## GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties

> GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)

Update the navigation property multiValueExtendedProperties in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
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


## GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties

> GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)

Update the navigation property singleValueExtendedProperties in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
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


## GroupsConversationsThreadsUpdatePosts

> GroupsConversationsThreadsUpdatePosts(ctx, groupId, conversationId, conversationThreadId, postId, microsoftGraphPost)

Update the navigation property posts in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**microsoftGraphPost** | [**MicrosoftGraphPost**](MicrosoftGraphPost.md)| New navigation property values | 

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


## GroupsConversationsUpdateThreads

> GroupsConversationsUpdateThreads(ctx, groupId, conversationId, conversationThreadId, microsoftGraphConversationThread)

Update the navigation property threads in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**microsoftGraphConversationThread** | [**MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md)| New navigation property values | 

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


## GroupsCreateConversations

> MicrosoftGraphConversation GroupsCreateConversations(ctx, groupId, microsoftGraphConversation)

Create new navigation property to conversations for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphConversation** | [**MicrosoftGraphConversation**](MicrosoftGraphConversation.md)| New navigation property | 

### Return type

[**MicrosoftGraphConversation**](microsoft.graph.conversation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetConversations

> MicrosoftGraphConversation GroupsGetConversations(ctx, groupId, conversationId, optional)

Get conversations from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
 **optional** | ***GroupsGetConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetConversationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphConversation**](microsoft.graph.conversation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsListConversations

> CollectionOfConversation GroupsListConversations(ctx, groupId, optional)

Get conversations from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsListConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsListConversationsOpts struct


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

[**CollectionOfConversation**](Collection_of_conversation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdateConversations

> GroupsUpdateConversations(ctx, groupId, conversationId, microsoftGraphConversation)

Update the navigation property conversations in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**microsoftGraphConversation** | [**MicrosoftGraphConversation**](MicrosoftGraphConversation.md)| New navigation property values | 

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

