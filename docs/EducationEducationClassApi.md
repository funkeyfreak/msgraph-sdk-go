# \EducationEducationClassApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationClassesGetGroup**](EducationEducationClassApi.md#EducationClassesGetGroup) | **Get** /education/classes/{educationClass-id}/group | Get group from education
[**EducationClassesGetMembers**](EducationEducationClassApi.md#EducationClassesGetMembers) | **Get** /education/classes/{educationClass-id}/members/{educationUser-id} | Get members from education
[**EducationClassesGetSchools**](EducationEducationClassApi.md#EducationClassesGetSchools) | **Get** /education/classes/{educationClass-id}/schools/{educationSchool-id} | Get schools from education
[**EducationClassesGetTeachers**](EducationEducationClassApi.md#EducationClassesGetTeachers) | **Get** /education/classes/{educationClass-id}/teachers/{educationUser-id} | Get teachers from education
[**EducationClassesListMembers**](EducationEducationClassApi.md#EducationClassesListMembers) | **Get** /education/classes/{educationClass-id}/members | Get members from education
[**EducationClassesListSchools**](EducationEducationClassApi.md#EducationClassesListSchools) | **Get** /education/classes/{educationClass-id}/schools | Get schools from education
[**EducationClassesListTeachers**](EducationEducationClassApi.md#EducationClassesListTeachers) | **Get** /education/classes/{educationClass-id}/teachers | Get teachers from education
[**EducationCreateClasses**](EducationEducationClassApi.md#EducationCreateClasses) | **Post** /education/classes | Create new navigation property to classes for education
[**EducationGetClasses**](EducationEducationClassApi.md#EducationGetClasses) | **Get** /education/classes/{educationClass-id} | Get classes from education
[**EducationListClasses**](EducationEducationClassApi.md#EducationListClasses) | **Get** /education/classes | Get classes from education
[**EducationUpdateClasses**](EducationEducationClassApi.md#EducationUpdateClasses) | **Patch** /education/classes/{educationClass-id} | Update the navigation property classes in education



## EducationClassesGetGroup

> MicrosoftGraphGroup EducationClassesGetGroup(ctx, educationClassId, optional)

Get group from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
 **optional** | ***EducationClassesGetGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationClassesGetGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphGroup**](microsoft.graph.group.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationClassesGetMembers

> MicrosoftGraphEducationUser EducationClassesGetMembers(ctx, educationClassId, educationUserId, optional)

Get members from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
**educationUserId** | **string**| key: educationUser-id of educationUser | 
 **optional** | ***EducationClassesGetMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationClassesGetMembersOpts struct


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


## EducationClassesGetSchools

> MicrosoftGraphEducationSchool EducationClassesGetSchools(ctx, educationClassId, educationSchoolId, optional)

Get schools from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
**educationSchoolId** | **string**| key: educationSchool-id of educationSchool | 
 **optional** | ***EducationClassesGetSchoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationClassesGetSchoolsOpts struct


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


## EducationClassesGetTeachers

> MicrosoftGraphEducationUser EducationClassesGetTeachers(ctx, educationClassId, educationUserId, optional)

Get teachers from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
**educationUserId** | **string**| key: educationUser-id of educationUser | 
 **optional** | ***EducationClassesGetTeachersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationClassesGetTeachersOpts struct


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


## EducationClassesListMembers

> CollectionOfEducationUser EducationClassesListMembers(ctx, educationClassId, optional)

Get members from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
 **optional** | ***EducationClassesListMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationClassesListMembersOpts struct


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


## EducationClassesListSchools

> CollectionOfEducationSchool EducationClassesListSchools(ctx, educationClassId, optional)

Get schools from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
 **optional** | ***EducationClassesListSchoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationClassesListSchoolsOpts struct


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


## EducationClassesListTeachers

> CollectionOfEducationUser EducationClassesListTeachers(ctx, educationClassId, optional)

Get teachers from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
 **optional** | ***EducationClassesListTeachersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationClassesListTeachersOpts struct


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


## EducationCreateClasses

> MicrosoftGraphEducationClass EducationCreateClasses(ctx, microsoftGraphEducationClass)

Create new navigation property to classes for education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphEducationClass** | [**MicrosoftGraphEducationClass**](MicrosoftGraphEducationClass.md)| New navigation property | 

### Return type

[**MicrosoftGraphEducationClass**](microsoft.graph.educationClass.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationGetClasses

> MicrosoftGraphEducationClass EducationGetClasses(ctx, educationClassId, optional)

Get classes from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
 **optional** | ***EducationGetClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationGetClassesOpts struct


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


## EducationListClasses

> CollectionOfEducationClass EducationListClasses(ctx, optional)

Get classes from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EducationListClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationListClassesOpts struct


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


## EducationUpdateClasses

> EducationUpdateClasses(ctx, educationClassId, microsoftGraphEducationClass)

Update the navigation property classes in education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
**microsoftGraphEducationClass** | [**MicrosoftGraphEducationClass**](MicrosoftGraphEducationClass.md)| New navigation property values | 

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

