# \DeviceManagementTermsAndConditionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementCreateTermsAndConditions) | **Post** /deviceManagement/termsAndConditions | Create new navigation property to termsAndConditions for deviceManagement
[**DeviceManagementGetTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementGetTermsAndConditions) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id} | Get termsAndConditions from deviceManagement
[**DeviceManagementListTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementListTermsAndConditions) | **Get** /deviceManagement/termsAndConditions | Get termsAndConditions from deviceManagement
[**DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id}/termsAndConditions | Get termsAndConditions from deviceManagement
[**DeviceManagementTermsAndConditionsCreateAcceptanceStatuses**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsCreateAcceptanceStatuses) | **Post** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses | Create new navigation property to acceptanceStatuses for deviceManagement
[**DeviceManagementTermsAndConditionsCreateAssignments**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsCreateAssignments) | **Post** /deviceManagement/termsAndConditions/{termsAndConditions-id}/assignments | Create new navigation property to assignments for deviceManagement
[**DeviceManagementTermsAndConditionsGetAcceptanceStatuses**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsGetAcceptanceStatuses) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id} | Get acceptanceStatuses from deviceManagement
[**DeviceManagementTermsAndConditionsGetAssignments**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsGetAssignments) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/assignments/{termsAndConditionsAssignment-id} | Get assignments from deviceManagement
[**DeviceManagementTermsAndConditionsListAcceptanceStatuses**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsListAcceptanceStatuses) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses | Get acceptanceStatuses from deviceManagement
[**DeviceManagementTermsAndConditionsListAssignments**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsListAssignments) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/assignments | Get assignments from deviceManagement
[**DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses) | **Patch** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id} | Update the navigation property acceptanceStatuses in deviceManagement
[**DeviceManagementTermsAndConditionsUpdateAssignments**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsUpdateAssignments) | **Patch** /deviceManagement/termsAndConditions/{termsAndConditions-id}/assignments/{termsAndConditionsAssignment-id} | Update the navigation property assignments in deviceManagement
[**DeviceManagementUpdateTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementUpdateTermsAndConditions) | **Patch** /deviceManagement/termsAndConditions/{termsAndConditions-id} | Update the navigation property termsAndConditions in deviceManagement



## DeviceManagementCreateTermsAndConditions

> MicrosoftGraphTermsAndConditions DeviceManagementCreateTermsAndConditions(ctx, microsoftGraphTermsAndConditions)

Create new navigation property to termsAndConditions for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTermsAndConditions** | [**MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md)| New navigation property | 

### Return type

[**MicrosoftGraphTermsAndConditions**](microsoft.graph.termsAndConditions.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetTermsAndConditions

> MicrosoftGraphTermsAndConditions DeviceManagementGetTermsAndConditions(ctx, termsAndConditionsId, optional)

Get termsAndConditions from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
 **optional** | ***DeviceManagementGetTermsAndConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetTermsAndConditionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditions**](microsoft.graph.termsAndConditions.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListTermsAndConditions

> CollectionOfTermsAndConditions DeviceManagementListTermsAndConditions(ctx, optional)

Get termsAndConditions from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListTermsAndConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListTermsAndConditionsOpts struct


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

[**CollectionOfTermsAndConditions**](Collection_of_termsAndConditions.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions

> MicrosoftGraphTermsAndConditions DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId, optional)

Get termsAndConditions from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string**| key: termsAndConditionsAcceptanceStatus-id of termsAndConditionsAcceptanceStatus | 
 **optional** | ***DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditions**](microsoft.graph.termsAndConditions.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsCreateAcceptanceStatuses

> MicrosoftGraphTermsAndConditionsAcceptanceStatus DeviceManagementTermsAndConditionsCreateAcceptanceStatuses(ctx, termsAndConditionsId, microsoftGraphTermsAndConditionsAcceptanceStatus)

Create new navigation property to acceptanceStatuses for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**microsoftGraphTermsAndConditionsAcceptanceStatus** | [**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](MicrosoftGraphTermsAndConditionsAcceptanceStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](microsoft.graph.termsAndConditionsAcceptanceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsCreateAssignments

> MicrosoftGraphTermsAndConditionsAssignment DeviceManagementTermsAndConditionsCreateAssignments(ctx, termsAndConditionsId, microsoftGraphTermsAndConditionsAssignment)

Create new navigation property to assignments for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**microsoftGraphTermsAndConditionsAssignment** | [**MicrosoftGraphTermsAndConditionsAssignment**](MicrosoftGraphTermsAndConditionsAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphTermsAndConditionsAssignment**](microsoft.graph.termsAndConditionsAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsGetAcceptanceStatuses

> MicrosoftGraphTermsAndConditionsAcceptanceStatus DeviceManagementTermsAndConditionsGetAcceptanceStatuses(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId, optional)

Get acceptanceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string**| key: termsAndConditionsAcceptanceStatus-id of termsAndConditionsAcceptanceStatus | 
 **optional** | ***DeviceManagementTermsAndConditionsGetAcceptanceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsGetAcceptanceStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](microsoft.graph.termsAndConditionsAcceptanceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsGetAssignments

> MicrosoftGraphTermsAndConditionsAssignment DeviceManagementTermsAndConditionsGetAssignments(ctx, termsAndConditionsId, termsAndConditionsAssignmentId, optional)

Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAssignmentId** | **string**| key: termsAndConditionsAssignment-id of termsAndConditionsAssignment | 
 **optional** | ***DeviceManagementTermsAndConditionsGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditionsAssignment**](microsoft.graph.termsAndConditionsAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsListAcceptanceStatuses

> CollectionOfTermsAndConditionsAcceptanceStatus DeviceManagementTermsAndConditionsListAcceptanceStatuses(ctx, termsAndConditionsId, optional)

Get acceptanceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
 **optional** | ***DeviceManagementTermsAndConditionsListAcceptanceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsListAcceptanceStatusesOpts struct


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

[**CollectionOfTermsAndConditionsAcceptanceStatus**](Collection_of_termsAndConditionsAcceptanceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsListAssignments

> CollectionOfTermsAndConditionsAssignment DeviceManagementTermsAndConditionsListAssignments(ctx, termsAndConditionsId, optional)

Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
 **optional** | ***DeviceManagementTermsAndConditionsListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsListAssignmentsOpts struct


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

[**CollectionOfTermsAndConditionsAssignment**](Collection_of_termsAndConditionsAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses

> DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId, microsoftGraphTermsAndConditionsAcceptanceStatus)

Update the navigation property acceptanceStatuses in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string**| key: termsAndConditionsAcceptanceStatus-id of termsAndConditionsAcceptanceStatus | 
**microsoftGraphTermsAndConditionsAcceptanceStatus** | [**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](MicrosoftGraphTermsAndConditionsAcceptanceStatus.md)| New navigation property values | 

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


## DeviceManagementTermsAndConditionsUpdateAssignments

> DeviceManagementTermsAndConditionsUpdateAssignments(ctx, termsAndConditionsId, termsAndConditionsAssignmentId, microsoftGraphTermsAndConditionsAssignment)

Update the navigation property assignments in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAssignmentId** | **string**| key: termsAndConditionsAssignment-id of termsAndConditionsAssignment | 
**microsoftGraphTermsAndConditionsAssignment** | [**MicrosoftGraphTermsAndConditionsAssignment**](MicrosoftGraphTermsAndConditionsAssignment.md)| New navigation property values | 

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


## DeviceManagementUpdateTermsAndConditions

> DeviceManagementUpdateTermsAndConditions(ctx, termsAndConditionsId, microsoftGraphTermsAndConditions)

Update the navigation property termsAndConditions in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**microsoftGraphTermsAndConditions** | [**MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md)| New navigation property values | 

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

