# \ServicePrincipalsAppRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsCreateAppRoleAssignedTo**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsCreateAppRoleAssignedTo) | **Post** /servicePrincipals/{servicePrincipal-id}/appRoleAssignedTo | Create new navigation property to appRoleAssignedTo for servicePrincipals
[**ServicePrincipalsCreateAppRoleAssignments**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsCreateAppRoleAssignments) | **Post** /servicePrincipals/{servicePrincipal-id}/appRoleAssignments | Create new navigation property to appRoleAssignments for servicePrincipals
[**ServicePrincipalsGetAppRoleAssignedTo**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsGetAppRoleAssignedTo) | **Get** /servicePrincipals/{servicePrincipal-id}/appRoleAssignedTo/{appRoleAssignment-id} | Get appRoleAssignedTo from servicePrincipals
[**ServicePrincipalsGetAppRoleAssignments**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsGetAppRoleAssignments) | **Get** /servicePrincipals/{servicePrincipal-id}/appRoleAssignments/{appRoleAssignment-id} | Get appRoleAssignments from servicePrincipals
[**ServicePrincipalsListAppRoleAssignedTo**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsListAppRoleAssignedTo) | **Get** /servicePrincipals/{servicePrincipal-id}/appRoleAssignedTo | Get appRoleAssignedTo from servicePrincipals
[**ServicePrincipalsListAppRoleAssignments**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsListAppRoleAssignments) | **Get** /servicePrincipals/{servicePrincipal-id}/appRoleAssignments | Get appRoleAssignments from servicePrincipals
[**ServicePrincipalsUpdateAppRoleAssignedTo**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsUpdateAppRoleAssignedTo) | **Patch** /servicePrincipals/{servicePrincipal-id}/appRoleAssignedTo/{appRoleAssignment-id} | Update the navigation property appRoleAssignedTo in servicePrincipals
[**ServicePrincipalsUpdateAppRoleAssignments**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsUpdateAppRoleAssignments) | **Patch** /servicePrincipals/{servicePrincipal-id}/appRoleAssignments/{appRoleAssignment-id} | Update the navigation property appRoleAssignments in servicePrincipals



## ServicePrincipalsCreateAppRoleAssignedTo

> MicrosoftGraphAppRoleAssignment ServicePrincipalsCreateAppRoleAssignedTo(ctx, servicePrincipalId, microsoftGraphAppRoleAssignment)

Create new navigation property to appRoleAssignedTo for servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](microsoft.graph.appRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsCreateAppRoleAssignments

> MicrosoftGraphAppRoleAssignment ServicePrincipalsCreateAppRoleAssignments(ctx, servicePrincipalId, microsoftGraphAppRoleAssignment)

Create new navigation property to appRoleAssignments for servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](microsoft.graph.appRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsGetAppRoleAssignedTo

> MicrosoftGraphAppRoleAssignment ServicePrincipalsGetAppRoleAssignedTo(ctx, servicePrincipalId, appRoleAssignmentId, optional)

Get appRoleAssignedTo from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**appRoleAssignmentId** | **string**| key: appRoleAssignment-id of appRoleAssignment | 
 **optional** | ***ServicePrincipalsGetAppRoleAssignedToOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetAppRoleAssignedToOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](microsoft.graph.appRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsGetAppRoleAssignments

> MicrosoftGraphAppRoleAssignment ServicePrincipalsGetAppRoleAssignments(ctx, servicePrincipalId, appRoleAssignmentId, optional)

Get appRoleAssignments from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**appRoleAssignmentId** | **string**| key: appRoleAssignment-id of appRoleAssignment | 
 **optional** | ***ServicePrincipalsGetAppRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsGetAppRoleAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](microsoft.graph.appRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListAppRoleAssignedTo

> CollectionOfAppRoleAssignment ServicePrincipalsListAppRoleAssignedTo(ctx, servicePrincipalId, optional)

Get appRoleAssignedTo from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListAppRoleAssignedToOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListAppRoleAssignedToOpts struct


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

[**CollectionOfAppRoleAssignment**](Collection_of_appRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListAppRoleAssignments

> CollectionOfAppRoleAssignment ServicePrincipalsListAppRoleAssignments(ctx, servicePrincipalId, optional)

Get appRoleAssignments from servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
 **optional** | ***ServicePrincipalsListAppRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServicePrincipalsListAppRoleAssignmentsOpts struct


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

[**CollectionOfAppRoleAssignment**](Collection_of_appRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsUpdateAppRoleAssignedTo

> ServicePrincipalsUpdateAppRoleAssignedTo(ctx, servicePrincipalId, appRoleAssignmentId, microsoftGraphAppRoleAssignment)

Update the navigation property appRoleAssignedTo in servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**appRoleAssignmentId** | **string**| key: appRoleAssignment-id of appRoleAssignment | 
**microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)| New navigation property values | 

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


## ServicePrincipalsUpdateAppRoleAssignments

> ServicePrincipalsUpdateAppRoleAssignments(ctx, servicePrincipalId, appRoleAssignmentId, microsoftGraphAppRoleAssignment)

Update the navigation property appRoleAssignments in servicePrincipals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string**| key: servicePrincipal-id of servicePrincipal | 
**appRoleAssignmentId** | **string**| key: appRoleAssignment-id of appRoleAssignment | 
**microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)| New navigation property values | 

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

