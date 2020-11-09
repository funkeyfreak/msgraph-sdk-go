# \GroupsConversationThreadApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateThreads**](GroupsConversationThreadApi.md#GroupsCreateThreads) | **Post** /groups/{group-id}/threads | Create new navigation property to threads for groups
[**GroupsGetThreads**](GroupsConversationThreadApi.md#GroupsGetThreads) | **Get** /groups/{group-id}/threads/{conversationThread-id} | Get threads from groups
[**GroupsListThreads**](GroupsConversationThreadApi.md#GroupsListThreads) | **Get** /groups/{group-id}/threads | Get threads from groups
[**GroupsThreadsCreatePosts**](GroupsConversationThreadApi.md#GroupsThreadsCreatePosts) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts | Create new navigation property to posts for groups
[**GroupsThreadsGetPosts**](GroupsConversationThreadApi.md#GroupsThreadsGetPosts) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id} | Get posts from groups
[**GroupsThreadsListPosts**](GroupsConversationThreadApi.md#GroupsThreadsListPosts) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts | Get posts from groups
[**GroupsThreadsPostsCreateAttachments**](GroupsConversationThreadApi.md#GroupsThreadsPostsCreateAttachments) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/attachments | Create new navigation property to attachments for groups
[**GroupsThreadsPostsCreateExtensions**](GroupsConversationThreadApi.md#GroupsThreadsPostsCreateExtensions) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/extensions | Create new navigation property to extensions for groups
[**GroupsThreadsPostsCreateMultiValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsCreateMultiValueExtendedProperties) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for groups
[**GroupsThreadsPostsCreateSingleValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsCreateSingleValueExtendedProperties) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for groups
[**GroupsThreadsPostsGetAttachments**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetAttachments) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Get attachments from groups
[**GroupsThreadsPostsGetExtensions**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetExtensions) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Get extensions from groups
[**GroupsThreadsPostsGetInReplyTo**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetInReplyTo) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Get inReplyTo from groups
[**GroupsThreadsPostsGetMultiValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetMultiValueExtendedProperties) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from groups
[**GroupsThreadsPostsGetSingleValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetSingleValueExtendedProperties) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from groups
[**GroupsThreadsPostsListAttachments**](GroupsConversationThreadApi.md#GroupsThreadsPostsListAttachments) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/attachments | Get attachments from groups
[**GroupsThreadsPostsListExtensions**](GroupsConversationThreadApi.md#GroupsThreadsPostsListExtensions) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/extensions | Get extensions from groups
[**GroupsThreadsPostsListMultiValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsListMultiValueExtendedProperties) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from groups
[**GroupsThreadsPostsListSingleValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsListSingleValueExtendedProperties) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from groups
[**GroupsThreadsPostsUpdateAttachments**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateAttachments) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Update the navigation property attachments in groups
[**GroupsThreadsPostsUpdateExtensions**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateExtensions) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Update the navigation property extensions in groups
[**GroupsThreadsPostsUpdateInReplyTo**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateInReplyTo) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Update the navigation property inReplyTo in groups
[**GroupsThreadsPostsUpdateMultiValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateMultiValueExtendedProperties) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in groups
[**GroupsThreadsPostsUpdateSingleValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateSingleValueExtendedProperties) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in groups
[**GroupsThreadsUpdatePosts**](GroupsConversationThreadApi.md#GroupsThreadsUpdatePosts) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id} | Update the navigation property posts in groups
[**GroupsUpdateThreads**](GroupsConversationThreadApi.md#GroupsUpdateThreads) | **Patch** /groups/{group-id}/threads/{conversationThread-id} | Update the navigation property threads in groups



## GroupsCreateThreads

> MicrosoftGraphConversationThread GroupsCreateThreads(ctx, groupId, microsoftGraphConversationThread)

Create new navigation property to threads for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsGetThreads

> MicrosoftGraphConversationThread GroupsGetThreads(ctx, groupId, conversationThreadId, optional)

Get threads from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
 **optional** | ***GroupsGetThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## GroupsListThreads

> CollectionOfConversationThread GroupsListThreads(ctx, groupId, optional)

Get threads from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsListThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsListThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## GroupsThreadsCreatePosts

> MicrosoftGraphPost GroupsThreadsCreatePosts(ctx, groupId, conversationThreadId, microsoftGraphPost)

Create new navigation property to posts for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsGetPosts

> MicrosoftGraphPost GroupsThreadsGetPosts(ctx, groupId, conversationThreadId, postId, optional)

Get posts from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsGetPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsGetPostsOpts struct


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


## GroupsThreadsListPosts

> CollectionOfPost GroupsThreadsListPosts(ctx, groupId, conversationThreadId, optional)

Get posts from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
 **optional** | ***GroupsThreadsListPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsListPostsOpts struct


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


## GroupsThreadsPostsCreateAttachments

> MicrosoftGraphAttachment GroupsThreadsPostsCreateAttachments(ctx, groupId, conversationThreadId, postId, microsoftGraphAttachment)

Create new navigation property to attachments for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsPostsCreateExtensions

> MicrosoftGraphExtension GroupsThreadsPostsCreateExtensions(ctx, groupId, conversationThreadId, postId, microsoftGraphExtension)

Create new navigation property to extensions for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsPostsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsThreadsPostsCreateMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, microsoftGraphMultiValueLegacyExtendedProperty)

Create new navigation property to multiValueExtendedProperties for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsPostsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsThreadsPostsCreateSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, microsoftGraphSingleValueLegacyExtendedProperty)

Create new navigation property to singleValueExtendedProperties for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsPostsGetAttachments

> MicrosoftGraphAttachment GroupsThreadsPostsGetAttachments(ctx, groupId, conversationThreadId, postId, attachmentId, optional)

Get attachments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***GroupsThreadsPostsGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsGetAttachmentsOpts struct


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


## GroupsThreadsPostsGetExtensions

> MicrosoftGraphExtension GroupsThreadsPostsGetExtensions(ctx, groupId, conversationThreadId, postId, extensionId, optional)

Get extensions from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***GroupsThreadsPostsGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsGetExtensionsOpts struct


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


## GroupsThreadsPostsGetInReplyTo

> MicrosoftGraphPost GroupsThreadsPostsGetInReplyTo(ctx, groupId, conversationThreadId, postId, optional)

Get inReplyTo from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsPostsGetInReplyToOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsGetInReplyToOpts struct


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


## GroupsThreadsPostsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsThreadsPostsGetMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId, optional)

Get multiValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***GroupsThreadsPostsGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsGetMultiValueExtendedPropertiesOpts struct


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


## GroupsThreadsPostsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsThreadsPostsGetSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId, optional)

Get singleValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***GroupsThreadsPostsGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsGetSingleValueExtendedPropertiesOpts struct


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


## GroupsThreadsPostsListAttachments

> CollectionOfAttachment GroupsThreadsPostsListAttachments(ctx, groupId, conversationThreadId, postId, optional)

Get attachments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsPostsListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsListAttachmentsOpts struct


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


## GroupsThreadsPostsListExtensions

> CollectionOfExtension GroupsThreadsPostsListExtensions(ctx, groupId, conversationThreadId, postId, optional)

Get extensions from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsPostsListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsListExtensionsOpts struct


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


## GroupsThreadsPostsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty GroupsThreadsPostsListMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, optional)

Get multiValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsPostsListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsListMultiValueExtendedPropertiesOpts struct


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


## GroupsThreadsPostsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty GroupsThreadsPostsListSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, optional)

Get singleValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsPostsListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsListSingleValueExtendedPropertiesOpts struct


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


## GroupsThreadsPostsUpdateAttachments

> GroupsThreadsPostsUpdateAttachments(ctx, groupId, conversationThreadId, postId, attachmentId, microsoftGraphAttachment)

Update the navigation property attachments in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsPostsUpdateExtensions

> GroupsThreadsPostsUpdateExtensions(ctx, groupId, conversationThreadId, postId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsPostsUpdateInReplyTo

> GroupsThreadsPostsUpdateInReplyTo(ctx, groupId, conversationThreadId, postId, microsoftGraphPost)

Update the navigation property inReplyTo in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsPostsUpdateMultiValueExtendedProperties

> GroupsThreadsPostsUpdateMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)

Update the navigation property multiValueExtendedProperties in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsPostsUpdateSingleValueExtendedProperties

> GroupsThreadsPostsUpdateSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)

Update the navigation property singleValueExtendedProperties in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsThreadsUpdatePosts

> GroupsThreadsUpdatePosts(ctx, groupId, conversationThreadId, postId, microsoftGraphPost)

Update the navigation property posts in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsUpdateThreads

> GroupsUpdateThreads(ctx, groupId, conversationThreadId, microsoftGraphConversationThread)

Update the navigation property threads in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

