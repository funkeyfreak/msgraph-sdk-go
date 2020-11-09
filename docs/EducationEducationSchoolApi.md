# \EducationEducationSchoolApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationCreateSchools**](EducationEducationSchoolApi.md#EducationCreateSchools) | **Post** /education/schools | Create new navigation property to schools for education
[**EducationGetSchools**](EducationEducationSchoolApi.md#EducationGetSchools) | **Get** /education/schools/{educationSchool-id} | Get schools from education
[**EducationListSchools**](EducationEducationSchoolApi.md#EducationListSchools) | **Get** /education/schools | Get schools from education
[**EducationSchoolsGetClasses**](EducationEducationSchoolApi.md#EducationSchoolsGetClasses) | **Get** /education/schools/{educationSchool-id}/classes/{educationClass-id} | Get classes from education
[**EducationSchoolsGetUsers**](EducationEducationSchoolApi.md#EducationSchoolsGetUsers) | **Get** /education/schools/{educationSchool-id}/users/{educationUser-id} | Get users from education
[**EducationSchoolsListClasses**](EducationEducationSchoolApi.md#EducationSchoolsListClasses) | **Get** /education/schools/{educationSchool-id}/classes | Get classes from education
[**EducationSchoolsListUsers**](EducationEducationSchoolApi.md#EducationSchoolsListUsers) | **Get** /education/schools/{educationSchool-id}/users | Get users from education
[**EducationUpdateSchools**](EducationEducationSchoolApi.md#EducationUpdateSchools) | **Patch** /education/schools/{educationSchool-id} | Update the navigation property schools in education



## EducationCreateSchools

> MicrosoftGraphEducationSchool EducationCreateSchools(ctx, microsoftGraphEducationSchool)

Create new navigation property to schools for education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphEducationSchool** | [**MicrosoftGraphEducationSchool**](MicrosoftGraphEducationSchool.md)| New navigation property | 

### Return type

[**MicrosoftGraphEducationSchool**](microsoft.graph.educationSchool.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationGetSchools

> MicrosoftGraphEducationSchool EducationGetSchools(ctx, educationSchoolId, optional)

Get schools from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string**| key: educationSchool-id of educationSchool | 
 **optional** | ***EducationGetSchoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationGetSchoolsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEducationSchool**](microsoft.graph.educationSchool.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationListSchools

> CollectionOfEducationSchool EducationListSchools(ctx, optional)

Get schools from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EducationListSchoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationListSchoolsOpts struct


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

[**CollectionOfEducationSchool**](Collection_of_educationSchool.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsGetClasses

> MicrosoftGraphEducationClass EducationSchoolsGetClasses(ctx, educationSchoolId, educationClassId, optional)

Get classes from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string**| key: educationSchool-id of educationSchool | 
**educationClassId** | **string**| key: educationClass-id of educationClass | 
 **optional** | ***EducationSchoolsGetClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationSchoolsGetClassesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEducationClass**](microsoft.graph.educationClass.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsGetUsers

> MicrosoftGraphEducationUser EducationSchoolsGetUsers(ctx, educationSchoolId, educationUserId, optional)

Get users from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string**| key: educationSchool-id of educationSchool | 
**educationUserId** | **string**| key: educationUser-id of educationUser | 
 **optional** | ***EducationSchoolsGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationSchoolsGetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEducationUser**](microsoft.graph.educationUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsListClasses

> CollectionOfEducationClass EducationSchoolsListClasses(ctx, educationSchoolId, optional)

Get classes from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string**| key: educationSchool-id of educationSchool | 
 **optional** | ***EducationSchoolsListClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationSchoolsListClassesOpts struct


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

[**CollectionOfEducationClass**](Collection_of_educationClass.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsListUsers

> CollectionOfEducationUser EducationSchoolsListUsers(ctx, educationSchoolId, optional)

Get users from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string**| key: educationSchool-id of educationSchool | 
 **optional** | ***EducationSchoolsListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationSchoolsListUsersOpts struct


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

[**CollectionOfEducationUser**](Collection_of_educationUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationUpdateSchools

> EducationUpdateSchools(ctx, educationSchoolId, microsoftGraphEducationSchool)

Update the navigation property schools in education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string**| key: educationSchool-id of educationSchool | 
**microsoftGraphEducationSchool** | [**MicrosoftGraphEducationSchool**](MicrosoftGraphEducationSchool.md)| New navigation property values | 

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

