# \PoliciesPermissionGrantPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreatePermissionGrantPolicies**](PoliciesPermissionGrantPolicyApi.md#PoliciesCreatePermissionGrantPolicies) | **Post** /policies/permissionGrantPolicies | Create new navigation property to permissionGrantPolicies for policies
[**PoliciesGetPermissionGrantPolicies**](PoliciesPermissionGrantPolicyApi.md#PoliciesGetPermissionGrantPolicies) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id} | Get permissionGrantPolicies from policies
[**PoliciesListPermissionGrantPolicies**](PoliciesPermissionGrantPolicyApi.md#PoliciesListPermissionGrantPolicies) | **Get** /policies/permissionGrantPolicies | Get permissionGrantPolicies from policies
[**PoliciesPermissionGrantPoliciesCreateExcludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesCreateExcludes) | **Post** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/excludes | Create new navigation property to excludes for policies
[**PoliciesPermissionGrantPoliciesCreateIncludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesCreateIncludes) | **Post** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/includes | Create new navigation property to includes for policies
[**PoliciesPermissionGrantPoliciesGetExcludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesGetExcludes) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/excludes/{permissionGrantConditionSet-id} | Get excludes from policies
[**PoliciesPermissionGrantPoliciesGetIncludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesGetIncludes) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/includes/{permissionGrantConditionSet-id} | Get includes from policies
[**PoliciesPermissionGrantPoliciesListExcludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesListExcludes) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/excludes | Get excludes from policies
[**PoliciesPermissionGrantPoliciesListIncludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesListIncludes) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/includes | Get includes from policies
[**PoliciesPermissionGrantPoliciesUpdateExcludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesUpdateExcludes) | **Patch** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/excludes/{permissionGrantConditionSet-id} | Update the navigation property excludes in policies
[**PoliciesPermissionGrantPoliciesUpdateIncludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesUpdateIncludes) | **Patch** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/includes/{permissionGrantConditionSet-id} | Update the navigation property includes in policies
[**PoliciesUpdatePermissionGrantPolicies**](PoliciesPermissionGrantPolicyApi.md#PoliciesUpdatePermissionGrantPolicies) | **Patch** /policies/permissionGrantPolicies/{permissionGrantPolicy-id} | Update the navigation property permissionGrantPolicies in policies



## PoliciesCreatePermissionGrantPolicies

> MicrosoftGraphPermissionGrantPolicy PoliciesCreatePermissionGrantPolicies(ctx, microsoftGraphPermissionGrantPolicy)

Create new navigation property to permissionGrantPolicies for policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphPermissionGrantPolicy** | [**MicrosoftGraphPermissionGrantPolicy**](MicrosoftGraphPermissionGrantPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphPermissionGrantPolicy**](microsoft.graph.permissionGrantPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetPermissionGrantPolicies

> MicrosoftGraphPermissionGrantPolicy PoliciesGetPermissionGrantPolicies(ctx, permissionGrantPolicyId, optional)

Get permissionGrantPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
 **optional** | ***PoliciesGetPermissionGrantPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesGetPermissionGrantPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPermissionGrantPolicy**](microsoft.graph.permissionGrantPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListPermissionGrantPolicies

> CollectionOfPermissionGrantPolicy PoliciesListPermissionGrantPolicies(ctx, optional)

Get permissionGrantPolicies from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesListPermissionGrantPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesListPermissionGrantPoliciesOpts struct


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

[**CollectionOfPermissionGrantPolicy**](Collection_of_permissionGrantPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesCreateExcludes

> MicrosoftGraphPermissionGrantConditionSet PoliciesPermissionGrantPoliciesCreateExcludes(ctx, permissionGrantPolicyId, microsoftGraphPermissionGrantConditionSet)

Create new navigation property to excludes for policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
**microsoftGraphPermissionGrantConditionSet** | [**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md)| New navigation property | 

### Return type

[**MicrosoftGraphPermissionGrantConditionSet**](microsoft.graph.permissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesCreateIncludes

> MicrosoftGraphPermissionGrantConditionSet PoliciesPermissionGrantPoliciesCreateIncludes(ctx, permissionGrantPolicyId, microsoftGraphPermissionGrantConditionSet)

Create new navigation property to includes for policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
**microsoftGraphPermissionGrantConditionSet** | [**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md)| New navigation property | 

### Return type

[**MicrosoftGraphPermissionGrantConditionSet**](microsoft.graph.permissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesGetExcludes

> MicrosoftGraphPermissionGrantConditionSet PoliciesPermissionGrantPoliciesGetExcludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId, optional)

Get excludes from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string**| key: permissionGrantConditionSet-id of permissionGrantConditionSet | 
 **optional** | ***PoliciesPermissionGrantPoliciesGetExcludesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPermissionGrantPoliciesGetExcludesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPermissionGrantConditionSet**](microsoft.graph.permissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesGetIncludes

> MicrosoftGraphPermissionGrantConditionSet PoliciesPermissionGrantPoliciesGetIncludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId, optional)

Get includes from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string**| key: permissionGrantConditionSet-id of permissionGrantConditionSet | 
 **optional** | ***PoliciesPermissionGrantPoliciesGetIncludesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPermissionGrantPoliciesGetIncludesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPermissionGrantConditionSet**](microsoft.graph.permissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesListExcludes

> CollectionOfPermissionGrantConditionSet PoliciesPermissionGrantPoliciesListExcludes(ctx, permissionGrantPolicyId, optional)

Get excludes from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
 **optional** | ***PoliciesPermissionGrantPoliciesListExcludesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPermissionGrantPoliciesListExcludesOpts struct


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

[**CollectionOfPermissionGrantConditionSet**](Collection_of_permissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesListIncludes

> CollectionOfPermissionGrantConditionSet PoliciesPermissionGrantPoliciesListIncludes(ctx, permissionGrantPolicyId, optional)

Get includes from policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
 **optional** | ***PoliciesPermissionGrantPoliciesListIncludesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPermissionGrantPoliciesListIncludesOpts struct


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

[**CollectionOfPermissionGrantConditionSet**](Collection_of_permissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesUpdateExcludes

> PoliciesPermissionGrantPoliciesUpdateExcludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId, microsoftGraphPermissionGrantConditionSet)

Update the navigation property excludes in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string**| key: permissionGrantConditionSet-id of permissionGrantConditionSet | 
**microsoftGraphPermissionGrantConditionSet** | [**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md)| New navigation property values | 

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


## PoliciesPermissionGrantPoliciesUpdateIncludes

> PoliciesPermissionGrantPoliciesUpdateIncludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId, microsoftGraphPermissionGrantConditionSet)

Update the navigation property includes in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string**| key: permissionGrantConditionSet-id of permissionGrantConditionSet | 
**microsoftGraphPermissionGrantConditionSet** | [**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md)| New navigation property values | 

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


## PoliciesUpdatePermissionGrantPolicies

> PoliciesUpdatePermissionGrantPolicies(ctx, permissionGrantPolicyId, microsoftGraphPermissionGrantPolicy)

Update the navigation property permissionGrantPolicies in policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string**| key: permissionGrantPolicy-id of permissionGrantPolicy | 
**microsoftGraphPermissionGrantPolicy** | [**MicrosoftGraphPermissionGrantPolicy**](MicrosoftGraphPermissionGrantPolicy.md)| New navigation property values | 

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

