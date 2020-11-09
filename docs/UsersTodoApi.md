# \UsersTodoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetTodo**](UsersTodoApi.md#UsersGetTodo) | **Get** /users/{user-id}/todo | Get todo from users
[**UsersTodoCreateLists**](UsersTodoApi.md#UsersTodoCreateLists) | **Post** /users/{user-id}/todo/lists | Create new navigation property to lists for users
[**UsersTodoGetLists**](UsersTodoApi.md#UsersTodoGetLists) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id} | Get lists from users
[**UsersTodoListLists**](UsersTodoApi.md#UsersTodoListLists) | **Get** /users/{user-id}/todo/lists | Get lists from users
[**UsersTodoListsCreateExtensions**](UsersTodoApi.md#UsersTodoListsCreateExtensions) | **Post** /users/{user-id}/todo/lists/{todoTaskList-id}/extensions | Create new navigation property to extensions for users
[**UsersTodoListsCreateTasks**](UsersTodoApi.md#UsersTodoListsCreateTasks) | **Post** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks | Create new navigation property to tasks for users
[**UsersTodoListsGetExtensions**](UsersTodoApi.md#UsersTodoListsGetExtensions) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Get extensions from users
[**UsersTodoListsGetTasks**](UsersTodoApi.md#UsersTodoListsGetTasks) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Get tasks from users
[**UsersTodoListsListExtensions**](UsersTodoApi.md#UsersTodoListsListExtensions) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/extensions | Get extensions from users
[**UsersTodoListsListTasks**](UsersTodoApi.md#UsersTodoListsListTasks) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks | Get tasks from users
[**UsersTodoListsTasksCreateExtensions**](UsersTodoApi.md#UsersTodoListsTasksCreateExtensions) | **Post** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions | Create new navigation property to extensions for users
[**UsersTodoListsTasksCreateLinkedResources**](UsersTodoApi.md#UsersTodoListsTasksCreateLinkedResources) | **Post** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources | Create new navigation property to linkedResources for users
[**UsersTodoListsTasksGetExtensions**](UsersTodoApi.md#UsersTodoListsTasksGetExtensions) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Get extensions from users
[**UsersTodoListsTasksGetLinkedResources**](UsersTodoApi.md#UsersTodoListsTasksGetLinkedResources) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Get linkedResources from users
[**UsersTodoListsTasksListExtensions**](UsersTodoApi.md#UsersTodoListsTasksListExtensions) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions | Get extensions from users
[**UsersTodoListsTasksListLinkedResources**](UsersTodoApi.md#UsersTodoListsTasksListLinkedResources) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources | Get linkedResources from users
[**UsersTodoListsTasksUpdateExtensions**](UsersTodoApi.md#UsersTodoListsTasksUpdateExtensions) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersTodoListsTasksUpdateLinkedResources**](UsersTodoApi.md#UsersTodoListsTasksUpdateLinkedResources) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Update the navigation property linkedResources in users
[**UsersTodoListsUpdateExtensions**](UsersTodoApi.md#UsersTodoListsUpdateExtensions) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersTodoListsUpdateTasks**](UsersTodoApi.md#UsersTodoListsUpdateTasks) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Update the navigation property tasks in users
[**UsersTodoUpdateLists**](UsersTodoApi.md#UsersTodoUpdateLists) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id} | Update the navigation property lists in users
[**UsersUpdateTodo**](UsersTodoApi.md#UsersUpdateTodo) | **Patch** /users/{user-id}/todo | Update the navigation property todo in users



## UsersGetTodo

> MicrosoftGraphTodo UsersGetTodo(ctx, userId, optional)

Get todo from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetTodoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetTodoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTodo**](microsoft.graph.todo.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoCreateLists

> MicrosoftGraphTodoTaskList UsersTodoCreateLists(ctx, userId, microsoftGraphTodoTaskList)

Create new navigation property to lists for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphTodoTaskList** | [**MicrosoftGraphTodoTaskList**](MicrosoftGraphTodoTaskList.md)| New navigation property | 

### Return type

[**MicrosoftGraphTodoTaskList**](microsoft.graph.todoTaskList.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoGetLists

> MicrosoftGraphTodoTaskList UsersTodoGetLists(ctx, userId, todoTaskListId, optional)

Get lists from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
 **optional** | ***UsersTodoGetListsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoGetListsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTodoTaskList**](microsoft.graph.todoTaskList.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoListLists

> CollectionOfTodoTaskList UsersTodoListLists(ctx, userId, optional)

Get lists from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersTodoListListsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoListListsOpts struct


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

[**CollectionOfTodoTaskList**](Collection_of_todoTaskList.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoListsCreateExtensions

> MicrosoftGraphExtension UsersTodoListsCreateExtensions(ctx, userId, todoTaskListId, microsoftGraphExtension)

Create new navigation property to extensions for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
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


## UsersTodoListsCreateTasks

> MicrosoftGraphTodoTask UsersTodoListsCreateTasks(ctx, userId, todoTaskListId, microsoftGraphTodoTask)

Create new navigation property to tasks for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**microsoftGraphTodoTask** | [**MicrosoftGraphTodoTask**](MicrosoftGraphTodoTask.md)| New navigation property | 

### Return type

[**MicrosoftGraphTodoTask**](microsoft.graph.todoTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoListsGetExtensions

> MicrosoftGraphExtension UsersTodoListsGetExtensions(ctx, userId, todoTaskListId, extensionId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***UsersTodoListsGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoListsGetExtensionsOpts struct


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


## UsersTodoListsGetTasks

> MicrosoftGraphTodoTask UsersTodoListsGetTasks(ctx, userId, todoTaskListId, todoTaskId, optional)

Get tasks from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
 **optional** | ***UsersTodoListsGetTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoListsGetTasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTodoTask**](microsoft.graph.todoTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoListsListExtensions

> CollectionOfExtension UsersTodoListsListExtensions(ctx, userId, todoTaskListId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
 **optional** | ***UsersTodoListsListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoListsListExtensionsOpts struct


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


## UsersTodoListsListTasks

> CollectionOfTodoTask UsersTodoListsListTasks(ctx, userId, todoTaskListId, optional)

Get tasks from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
 **optional** | ***UsersTodoListsListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoListsListTasksOpts struct


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

[**CollectionOfTodoTask**](Collection_of_todoTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoListsTasksCreateExtensions

> MicrosoftGraphExtension UsersTodoListsTasksCreateExtensions(ctx, userId, todoTaskListId, todoTaskId, microsoftGraphExtension)

Create new navigation property to extensions for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
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


## UsersTodoListsTasksCreateLinkedResources

> MicrosoftGraphLinkedResource UsersTodoListsTasksCreateLinkedResources(ctx, userId, todoTaskListId, todoTaskId, microsoftGraphLinkedResource)

Create new navigation property to linkedResources for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
**microsoftGraphLinkedResource** | [**MicrosoftGraphLinkedResource**](MicrosoftGraphLinkedResource.md)| New navigation property | 

### Return type

[**MicrosoftGraphLinkedResource**](microsoft.graph.linkedResource.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoListsTasksGetExtensions

> MicrosoftGraphExtension UsersTodoListsTasksGetExtensions(ctx, userId, todoTaskListId, todoTaskId, extensionId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***UsersTodoListsTasksGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoListsTasksGetExtensionsOpts struct


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


## UsersTodoListsTasksGetLinkedResources

> MicrosoftGraphLinkedResource UsersTodoListsTasksGetLinkedResources(ctx, userId, todoTaskListId, todoTaskId, linkedResourceId, optional)

Get linkedResources from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
**linkedResourceId** | **string**| key: linkedResource-id of linkedResource | 
 **optional** | ***UsersTodoListsTasksGetLinkedResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoListsTasksGetLinkedResourcesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphLinkedResource**](microsoft.graph.linkedResource.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoListsTasksListExtensions

> CollectionOfExtension UsersTodoListsTasksListExtensions(ctx, userId, todoTaskListId, todoTaskId, optional)

Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
 **optional** | ***UsersTodoListsTasksListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoListsTasksListExtensionsOpts struct


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


## UsersTodoListsTasksListLinkedResources

> CollectionOfLinkedResource UsersTodoListsTasksListLinkedResources(ctx, userId, todoTaskListId, todoTaskId, optional)

Get linkedResources from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
 **optional** | ***UsersTodoListsTasksListLinkedResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersTodoListsTasksListLinkedResourcesOpts struct


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

[**CollectionOfLinkedResource**](Collection_of_linkedResource.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTodoListsTasksUpdateExtensions

> UsersTodoListsTasksUpdateExtensions(ctx, userId, todoTaskListId, todoTaskId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
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


## UsersTodoListsTasksUpdateLinkedResources

> UsersTodoListsTasksUpdateLinkedResources(ctx, userId, todoTaskListId, todoTaskId, linkedResourceId, microsoftGraphLinkedResource)

Update the navigation property linkedResources in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
**linkedResourceId** | **string**| key: linkedResource-id of linkedResource | 
**microsoftGraphLinkedResource** | [**MicrosoftGraphLinkedResource**](MicrosoftGraphLinkedResource.md)| New navigation property values | 

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


## UsersTodoListsUpdateExtensions

> UsersTodoListsUpdateExtensions(ctx, userId, todoTaskListId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
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


## UsersTodoListsUpdateTasks

> UsersTodoListsUpdateTasks(ctx, userId, todoTaskListId, todoTaskId, microsoftGraphTodoTask)

Update the navigation property tasks in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
**microsoftGraphTodoTask** | [**MicrosoftGraphTodoTask**](MicrosoftGraphTodoTask.md)| New navigation property values | 

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


## UsersTodoUpdateLists

> UsersTodoUpdateLists(ctx, userId, todoTaskListId, microsoftGraphTodoTaskList)

Update the navigation property lists in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**microsoftGraphTodoTaskList** | [**MicrosoftGraphTodoTaskList**](MicrosoftGraphTodoTaskList.md)| New navigation property values | 

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


## UsersUpdateTodo

> UsersUpdateTodo(ctx, userId, microsoftGraphTodo)

Update the navigation property todo in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphTodo** | [**MicrosoftGraphTodo**](MicrosoftGraphTodo.md)| New navigation property values | 

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

