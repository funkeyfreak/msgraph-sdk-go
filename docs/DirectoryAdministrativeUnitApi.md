# \DirectoryAdministrativeUnitApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryAdministrativeUnitsCreateExtensions**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsCreateExtensions) | **Post** /directory/administrativeUnits/{administrativeUnit-id}/extensions | Create new navigation property to extensions for directory
[**DirectoryAdministrativeUnitsCreateScopedRoleMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsCreateScopedRoleMembers) | **Post** /directory/administrativeUnits/{administrativeUnit-id}/scopedRoleMembers | Create new navigation property to scopedRoleMembers for directory
[**DirectoryAdministrativeUnitsGetExtensions**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsGetExtensions) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/extensions/{extension-id} | Get extensions from directory
[**DirectoryAdministrativeUnitsGetMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsGetMembers) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/members/{directoryObject-id} | Get members from directory
[**DirectoryAdministrativeUnitsGetScopedRoleMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsGetScopedRoleMembers) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/scopedRoleMembers/{scopedRoleMembership-id} | Get scopedRoleMembers from directory
[**DirectoryAdministrativeUnitsListExtensions**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsListExtensions) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/extensions | Get extensions from directory
[**DirectoryAdministrativeUnitsListMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsListMembers) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/members | Get members from directory
[**DirectoryAdministrativeUnitsListScopedRoleMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsListScopedRoleMembers) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/scopedRoleMembers | Get scopedRoleMembers from directory
[**DirectoryAdministrativeUnitsUpdateExtensions**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsUpdateExtensions) | **Patch** /directory/administrativeUnits/{administrativeUnit-id}/extensions/{extension-id} | Update the navigation property extensions in directory
[**DirectoryAdministrativeUnitsUpdateScopedRoleMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsUpdateScopedRoleMembers) | **Patch** /directory/administrativeUnits/{administrativeUnit-id}/scopedRoleMembers/{scopedRoleMembership-id} | Update the navigation property scopedRoleMembers in directory
[**DirectoryCreateAdministrativeUnits**](DirectoryAdministrativeUnitApi.md#DirectoryCreateAdministrativeUnits) | **Post** /directory/administrativeUnits | Create new navigation property to administrativeUnits for directory
[**DirectoryGetAdministrativeUnits**](DirectoryAdministrativeUnitApi.md#DirectoryGetAdministrativeUnits) | **Get** /directory/administrativeUnits/{administrativeUnit-id} | Get administrativeUnits from directory
[**DirectoryListAdministrativeUnits**](DirectoryAdministrativeUnitApi.md#DirectoryListAdministrativeUnits) | **Get** /directory/administrativeUnits | Get administrativeUnits from directory
[**DirectoryUpdateAdministrativeUnits**](DirectoryAdministrativeUnitApi.md#DirectoryUpdateAdministrativeUnits) | **Patch** /directory/administrativeUnits/{administrativeUnit-id} | Update the navigation property administrativeUnits in directory



## DirectoryAdministrativeUnitsCreateExtensions

> MicrosoftGraphExtension DirectoryAdministrativeUnitsCreateExtensions(ctx, administrativeUnitId, microsoftGraphExtension)

Create new navigation property to extensions for directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
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


## DirectoryAdministrativeUnitsCreateScopedRoleMembers

> MicrosoftGraphScopedRoleMembership DirectoryAdministrativeUnitsCreateScopedRoleMembers(ctx, administrativeUnitId, microsoftGraphScopedRoleMembership)

Create new navigation property to scopedRoleMembers for directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New navigation property | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](microsoft.graph.scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryAdministrativeUnitsGetExtensions

> MicrosoftGraphExtension DirectoryAdministrativeUnitsGetExtensions(ctx, administrativeUnitId, extensionId, optional)

Get extensions from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***DirectoryAdministrativeUnitsGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryAdministrativeUnitsGetExtensionsOpts struct


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


## DirectoryAdministrativeUnitsGetMembers

> MicrosoftGraphDirectoryObject DirectoryAdministrativeUnitsGetMembers(ctx, administrativeUnitId, directoryObjectId, optional)

Get members from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
 **optional** | ***DirectoryAdministrativeUnitsGetMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryAdministrativeUnitsGetMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryAdministrativeUnitsGetScopedRoleMembers

> MicrosoftGraphScopedRoleMembership DirectoryAdministrativeUnitsGetScopedRoleMembers(ctx, administrativeUnitId, scopedRoleMembershipId, optional)

Get scopedRoleMembers from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
 **optional** | ***DirectoryAdministrativeUnitsGetScopedRoleMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryAdministrativeUnitsGetScopedRoleMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](microsoft.graph.scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryAdministrativeUnitsListExtensions

> CollectionOfExtension DirectoryAdministrativeUnitsListExtensions(ctx, administrativeUnitId, optional)

Get extensions from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
 **optional** | ***DirectoryAdministrativeUnitsListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryAdministrativeUnitsListExtensionsOpts struct


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

[**CollectionOfExtension**](Collection_of_extension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryAdministrativeUnitsListMembers

> CollectionOfDirectoryObject DirectoryAdministrativeUnitsListMembers(ctx, administrativeUnitId, optional)

Get members from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
 **optional** | ***DirectoryAdministrativeUnitsListMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryAdministrativeUnitsListMembersOpts struct


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

[**CollectionOfDirectoryObject**](Collection_of_directoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryAdministrativeUnitsListScopedRoleMembers

> CollectionOfScopedRoleMembership DirectoryAdministrativeUnitsListScopedRoleMembers(ctx, administrativeUnitId, optional)

Get scopedRoleMembers from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
 **optional** | ***DirectoryAdministrativeUnitsListScopedRoleMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryAdministrativeUnitsListScopedRoleMembersOpts struct


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

[**CollectionOfScopedRoleMembership**](Collection_of_scopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryAdministrativeUnitsUpdateExtensions

> DirectoryAdministrativeUnitsUpdateExtensions(ctx, administrativeUnitId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
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


## DirectoryAdministrativeUnitsUpdateScopedRoleMembers

> DirectoryAdministrativeUnitsUpdateScopedRoleMembers(ctx, administrativeUnitId, scopedRoleMembershipId, microsoftGraphScopedRoleMembership)

Update the navigation property scopedRoleMembers in directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
**scopedRoleMembershipId** | **string**| key: scopedRoleMembership-id of scopedRoleMembership | 
**microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)| New navigation property values | 

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


## DirectoryCreateAdministrativeUnits

> MicrosoftGraphAdministrativeUnit DirectoryCreateAdministrativeUnits(ctx, microsoftGraphAdministrativeUnit)

Create new navigation property to administrativeUnits for directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphAdministrativeUnit** | [**MicrosoftGraphAdministrativeUnit**](MicrosoftGraphAdministrativeUnit.md)| New navigation property | 

### Return type

[**MicrosoftGraphAdministrativeUnit**](microsoft.graph.administrativeUnit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryGetAdministrativeUnits

> MicrosoftGraphAdministrativeUnit DirectoryGetAdministrativeUnits(ctx, administrativeUnitId, optional)

Get administrativeUnits from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
 **optional** | ***DirectoryGetAdministrativeUnitsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryGetAdministrativeUnitsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAdministrativeUnit**](microsoft.graph.administrativeUnit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryListAdministrativeUnits

> CollectionOfAdministrativeUnit DirectoryListAdministrativeUnits(ctx, optional)

Get administrativeUnits from directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryListAdministrativeUnitsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryListAdministrativeUnitsOpts struct


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

[**CollectionOfAdministrativeUnit**](Collection_of_administrativeUnit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryUpdateAdministrativeUnits

> DirectoryUpdateAdministrativeUnits(ctx, administrativeUnitId, microsoftGraphAdministrativeUnit)

Update the navigation property administrativeUnits in directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string**| key: administrativeUnit-id of administrativeUnit | 
**microsoftGraphAdministrativeUnit** | [**MicrosoftGraphAdministrativeUnit**](MicrosoftGraphAdministrativeUnit.md)| New navigation property values | 

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

