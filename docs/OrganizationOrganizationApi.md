# \OrganizationOrganizationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationOrganizationCreateOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationCreateOrganization) | **Post** /organization | Add new entity to organization
[**OrganizationOrganizationDeleteOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationDeleteOrganization) | **Delete** /organization/{organization-id} | Delete entity from organization
[**OrganizationOrganizationGetOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationGetOrganization) | **Get** /organization/{organization-id} | Get entity from organization by key
[**OrganizationOrganizationListOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationListOrganization) | **Get** /organization | Get entities from organization
[**OrganizationOrganizationUpdateOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationUpdateOrganization) | **Patch** /organization/{organization-id} | Update entity in organization



## OrganizationOrganizationCreateOrganization

> MicrosoftGraphOrganization OrganizationOrganizationCreateOrganization(ctx, microsoftGraphOrganization)

Add new entity to organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOrganization** | [**MicrosoftGraphOrganization**](MicrosoftGraphOrganization.md)| New entity | 

### Return type

[**MicrosoftGraphOrganization**](microsoft.graph.organization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationOrganizationDeleteOrganization

> OrganizationOrganizationDeleteOrganization(ctx, organizationId, optional)

Delete entity from organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| key: organization-id of organization | 
 **optional** | ***OrganizationOrganizationDeleteOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationOrganizationDeleteOrganizationOpts struct


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


## OrganizationOrganizationGetOrganization

> MicrosoftGraphOrganization OrganizationOrganizationGetOrganization(ctx, organizationId, optional)

Get entity from organization by key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| key: organization-id of organization | 
 **optional** | ***OrganizationOrganizationGetOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationOrganizationGetOrganizationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphOrganization**](microsoft.graph.organization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationOrganizationListOrganization

> CollectionOfOrganization OrganizationOrganizationListOrganization(ctx, optional)

Get entities from organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganizationOrganizationListOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationOrganizationListOrganizationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfOrganization**](Collection_of_organization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationOrganizationUpdateOrganization

> OrganizationOrganizationUpdateOrganization(ctx, organizationId, microsoftGraphOrganization)

Update entity in organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| key: organization-id of organization | 
**microsoftGraphOrganization** | [**MicrosoftGraphOrganization**](MicrosoftGraphOrganization.md)| New property values | 

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

