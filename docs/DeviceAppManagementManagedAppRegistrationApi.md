# \DeviceAppManagementManagedAppRegistrationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateManagedAppRegistrations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementCreateManagedAppRegistrations) | **Post** /deviceAppManagement/managedAppRegistrations | Create new navigation property to managedAppRegistrations for deviceAppManagement
[**DeviceAppManagementGetManagedAppRegistrations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementGetManagedAppRegistrations) | **Get** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id} | Get managedAppRegistrations from deviceAppManagement
[**DeviceAppManagementListManagedAppRegistrations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementListManagedAppRegistrations) | **Get** /deviceAppManagement/managedAppRegistrations | Get managedAppRegistrations from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsCreateAppliedPolicies**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsCreateAppliedPolicies) | **Post** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/appliedPolicies | Create new navigation property to appliedPolicies for deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsCreateIntendedPolicies**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsCreateIntendedPolicies) | **Post** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/intendedPolicies | Create new navigation property to intendedPolicies for deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsCreateOperations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsCreateOperations) | **Post** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/operations | Create new navigation property to operations for deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsGetAppliedPolicies**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsGetAppliedPolicies) | **Get** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/appliedPolicies/{managedAppPolicy-id} | Get appliedPolicies from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsGetIntendedPolicies**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsGetIntendedPolicies) | **Get** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/intendedPolicies/{managedAppPolicy-id} | Get intendedPolicies from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsGetOperations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsGetOperations) | **Get** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/operations/{managedAppOperation-id} | Get operations from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsListAppliedPolicies**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsListAppliedPolicies) | **Get** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/appliedPolicies | Get appliedPolicies from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsListIntendedPolicies**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsListIntendedPolicies) | **Get** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/intendedPolicies | Get intendedPolicies from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsListOperations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsListOperations) | **Get** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/operations | Get operations from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsUpdateAppliedPolicies**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsUpdateAppliedPolicies) | **Patch** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/appliedPolicies/{managedAppPolicy-id} | Update the navigation property appliedPolicies in deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsUpdateIntendedPolicies**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsUpdateIntendedPolicies) | **Patch** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/intendedPolicies/{managedAppPolicy-id} | Update the navigation property intendedPolicies in deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsUpdateOperations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementManagedAppRegistrationsUpdateOperations) | **Patch** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id}/operations/{managedAppOperation-id} | Update the navigation property operations in deviceAppManagement
[**DeviceAppManagementUpdateManagedAppRegistrations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementUpdateManagedAppRegistrations) | **Patch** /deviceAppManagement/managedAppRegistrations/{managedAppRegistration-id} | Update the navigation property managedAppRegistrations in deviceAppManagement



## DeviceAppManagementCreateManagedAppRegistrations

> MicrosoftGraphManagedAppRegistration DeviceAppManagementCreateManagedAppRegistrations(ctx, microsoftGraphManagedAppRegistration)

Create new navigation property to managedAppRegistrations for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphManagedAppRegistration** | [**MicrosoftGraphManagedAppRegistration**](MicrosoftGraphManagedAppRegistration.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppRegistration**](microsoft.graph.managedAppRegistration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetManagedAppRegistrations

> MicrosoftGraphManagedAppRegistration DeviceAppManagementGetManagedAppRegistrations(ctx, managedAppRegistrationId, optional)

Get managedAppRegistrations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
 **optional** | ***DeviceAppManagementGetManagedAppRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetManagedAppRegistrationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppRegistration**](microsoft.graph.managedAppRegistration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListManagedAppRegistrations

> CollectionOfManagedAppRegistration DeviceAppManagementListManagedAppRegistrations(ctx, optional)

Get managedAppRegistrations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListManagedAppRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListManagedAppRegistrationsOpts struct


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

[**CollectionOfManagedAppRegistration**](Collection_of_managedAppRegistration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsCreateAppliedPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementManagedAppRegistrationsCreateAppliedPolicies(ctx, managedAppRegistrationId, microsoftGraphManagedAppPolicy)

Create new navigation property to appliedPolicies for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsCreateIntendedPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementManagedAppRegistrationsCreateIntendedPolicies(ctx, managedAppRegistrationId, microsoftGraphManagedAppPolicy)

Create new navigation property to intendedPolicies for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsCreateOperations

> MicrosoftGraphManagedAppOperation DeviceAppManagementManagedAppRegistrationsCreateOperations(ctx, managedAppRegistrationId, microsoftGraphManagedAppOperation)

Create new navigation property to operations for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**microsoftGraphManagedAppOperation** | [**MicrosoftGraphManagedAppOperation**](MicrosoftGraphManagedAppOperation.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppOperation**](microsoft.graph.managedAppOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsGetAppliedPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementManagedAppRegistrationsGetAppliedPolicies(ctx, managedAppRegistrationId, managedAppPolicyId, optional)

Get appliedPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsGetAppliedPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsGetAppliedPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsGetIntendedPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementManagedAppRegistrationsGetIntendedPolicies(ctx, managedAppRegistrationId, managedAppPolicyId, optional)

Get intendedPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsGetIntendedPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsGetIntendedPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsGetOperations

> MicrosoftGraphManagedAppOperation DeviceAppManagementManagedAppRegistrationsGetOperations(ctx, managedAppRegistrationId, managedAppOperationId, optional)

Get operations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppOperationId** | **string**| key: managedAppOperation-id of managedAppOperation | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsGetOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsGetOperationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppOperation**](microsoft.graph.managedAppOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsListAppliedPolicies

> CollectionOfManagedAppPolicy DeviceAppManagementManagedAppRegistrationsListAppliedPolicies(ctx, managedAppRegistrationId, optional)

Get appliedPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsListAppliedPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsListAppliedPoliciesOpts struct


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

[**CollectionOfManagedAppPolicy**](Collection_of_managedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsListIntendedPolicies

> CollectionOfManagedAppPolicy DeviceAppManagementManagedAppRegistrationsListIntendedPolicies(ctx, managedAppRegistrationId, optional)

Get intendedPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsListIntendedPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsListIntendedPoliciesOpts struct


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

[**CollectionOfManagedAppPolicy**](Collection_of_managedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsListOperations

> CollectionOfManagedAppOperation DeviceAppManagementManagedAppRegistrationsListOperations(ctx, managedAppRegistrationId, optional)

Get operations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsListOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsListOperationsOpts struct


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

[**CollectionOfManagedAppOperation**](Collection_of_managedAppOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsUpdateAppliedPolicies

> DeviceAppManagementManagedAppRegistrationsUpdateAppliedPolicies(ctx, managedAppRegistrationId, managedAppPolicyId, microsoftGraphManagedAppPolicy)

Update the navigation property appliedPolicies in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property values | 

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


## DeviceAppManagementManagedAppRegistrationsUpdateIntendedPolicies

> DeviceAppManagementManagedAppRegistrationsUpdateIntendedPolicies(ctx, managedAppRegistrationId, managedAppPolicyId, microsoftGraphManagedAppPolicy)

Update the navigation property intendedPolicies in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property values | 

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


## DeviceAppManagementManagedAppRegistrationsUpdateOperations

> DeviceAppManagementManagedAppRegistrationsUpdateOperations(ctx, managedAppRegistrationId, managedAppOperationId, microsoftGraphManagedAppOperation)

Update the navigation property operations in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppOperationId** | **string**| key: managedAppOperation-id of managedAppOperation | 
**microsoftGraphManagedAppOperation** | [**MicrosoftGraphManagedAppOperation**](MicrosoftGraphManagedAppOperation.md)| New navigation property values | 

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


## DeviceAppManagementUpdateManagedAppRegistrations

> DeviceAppManagementUpdateManagedAppRegistrations(ctx, managedAppRegistrationId, microsoftGraphManagedAppRegistration)

Update the navigation property managedAppRegistrations in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**microsoftGraphManagedAppRegistration** | [**MicrosoftGraphManagedAppRegistration**](MicrosoftGraphManagedAppRegistration.md)| New navigation property values | 

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

