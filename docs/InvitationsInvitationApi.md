# \InvitationsInvitationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvitationsInvitationCreateInvitation**](InvitationsInvitationApi.md#InvitationsInvitationCreateInvitation) | **Post** /invitations | Add new entity to invitations
[**InvitationsInvitationDeleteInvitation**](InvitationsInvitationApi.md#InvitationsInvitationDeleteInvitation) | **Delete** /invitations/{invitation-id} | Delete entity from invitations
[**InvitationsInvitationGetInvitation**](InvitationsInvitationApi.md#InvitationsInvitationGetInvitation) | **Get** /invitations/{invitation-id} | Get entity from invitations by key
[**InvitationsInvitationListInvitation**](InvitationsInvitationApi.md#InvitationsInvitationListInvitation) | **Get** /invitations | Get entities from invitations
[**InvitationsInvitationUpdateInvitation**](InvitationsInvitationApi.md#InvitationsInvitationUpdateInvitation) | **Patch** /invitations/{invitation-id} | Update entity in invitations



## InvitationsInvitationCreateInvitation

> MicrosoftGraphInvitation InvitationsInvitationCreateInvitation(ctx, microsoftGraphInvitation)

Add new entity to invitations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphInvitation** | [**MicrosoftGraphInvitation**](MicrosoftGraphInvitation.md)| New entity | 

### Return type

[**MicrosoftGraphInvitation**](microsoft.graph.invitation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsInvitationDeleteInvitation

> InvitationsInvitationDeleteInvitation(ctx, invitationId, optional)

Delete entity from invitations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string**| key: invitation-id of invitation | 
 **optional** | ***InvitationsInvitationDeleteInvitationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InvitationsInvitationDeleteInvitationOpts struct


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


## InvitationsInvitationGetInvitation

> MicrosoftGraphInvitation InvitationsInvitationGetInvitation(ctx, invitationId, optional)

Get entity from invitations by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string**| key: invitation-id of invitation | 
 **optional** | ***InvitationsInvitationGetInvitationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InvitationsInvitationGetInvitationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphInvitation**](microsoft.graph.invitation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsInvitationListInvitation

> CollectionOfInvitation InvitationsInvitationListInvitation(ctx, optional)

Get entities from invitations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvitationsInvitationListInvitationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InvitationsInvitationListInvitationOpts struct


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

[**CollectionOfInvitation**](Collection_of_invitation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsInvitationUpdateInvitation

> InvitationsInvitationUpdateInvitation(ctx, invitationId, microsoftGraphInvitation)

Update entity in invitations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string**| key: invitation-id of invitation | 
**microsoftGraphInvitation** | [**MicrosoftGraphInvitation**](MicrosoftGraphInvitation.md)| New property values | 

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

