# \TeamsConversationMemberApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsCreateMembers**](TeamsConversationMemberApi.md#TeamsCreateMembers) | **Post** /teams/{team-id}/members | Create new navigation property to members for teams
[**TeamsGetMembers**](TeamsConversationMemberApi.md#TeamsGetMembers) | **Get** /teams/{team-id}/members/{conversationMember-id} | Get members from teams
[**TeamsListMembers**](TeamsConversationMemberApi.md#TeamsListMembers) | **Get** /teams/{team-id}/members | Get members from teams
[**TeamsUpdateMembers**](TeamsConversationMemberApi.md#TeamsUpdateMembers) | **Patch** /teams/{team-id}/members/{conversationMember-id} | Update the navigation property members in teams



## TeamsCreateMembers

> MicrosoftGraphConversationMember TeamsCreateMembers(ctx, teamId, microsoftGraphConversationMember)

Create new navigation property to members for teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**microsoftGraphConversationMember** | [**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md)| New navigation property | 

### Return type

[**MicrosoftGraphConversationMember**](microsoft.graph.conversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetMembers

> MicrosoftGraphConversationMember TeamsGetMembers(ctx, teamId, conversationMemberId, optional)

Get members from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**conversationMemberId** | **string**| key: conversationMember-id of conversationMember | 
 **optional** | ***TeamsGetMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsGetMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphConversationMember**](microsoft.graph.conversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListMembers

> CollectionOfConversationMember TeamsListMembers(ctx, teamId, optional)

Get members from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
 **optional** | ***TeamsListMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsListMembersOpts struct


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

[**CollectionOfConversationMember**](Collection_of_conversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateMembers

> TeamsUpdateMembers(ctx, teamId, conversationMemberId, microsoftGraphConversationMember)

Update the navigation property members in teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**conversationMemberId** | **string**| key: conversationMember-id of conversationMember | 
**microsoftGraphConversationMember** | [**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md)| New navigation property values | 

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

