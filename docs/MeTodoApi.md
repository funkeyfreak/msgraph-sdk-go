# \MeTodoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeGetTodo**](MeTodoApi.md#MeGetTodo) | **Get** /me/todo | Get todo from me
[**MeTodoCreateLists**](MeTodoApi.md#MeTodoCreateLists) | **Post** /me/todo/lists | Create new navigation property to lists for me
[**MeTodoGetLists**](MeTodoApi.md#MeTodoGetLists) | **Get** /me/todo/lists/{todoTaskList-id} | Get lists from me
[**MeTodoListLists**](MeTodoApi.md#MeTodoListLists) | **Get** /me/todo/lists | Get lists from me
[**MeTodoListsCreateExtensions**](MeTodoApi.md#MeTodoListsCreateExtensions) | **Post** /me/todo/lists/{todoTaskList-id}/extensions | Create new navigation property to extensions for me
[**MeTodoListsCreateTasks**](MeTodoApi.md#MeTodoListsCreateTasks) | **Post** /me/todo/lists/{todoTaskList-id}/tasks | Create new navigation property to tasks for me
[**MeTodoListsGetExtensions**](MeTodoApi.md#MeTodoListsGetExtensions) | **Get** /me/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Get extensions from me
[**MeTodoListsGetTasks**](MeTodoApi.md#MeTodoListsGetTasks) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Get tasks from me
[**MeTodoListsListExtensions**](MeTodoApi.md#MeTodoListsListExtensions) | **Get** /me/todo/lists/{todoTaskList-id}/extensions | Get extensions from me
[**MeTodoListsListTasks**](MeTodoApi.md#MeTodoListsListTasks) | **Get** /me/todo/lists/{todoTaskList-id}/tasks | Get tasks from me
[**MeTodoListsTasksCreateExtensions**](MeTodoApi.md#MeTodoListsTasksCreateExtensions) | **Post** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions | Create new navigation property to extensions for me
[**MeTodoListsTasksCreateLinkedResources**](MeTodoApi.md#MeTodoListsTasksCreateLinkedResources) | **Post** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources | Create new navigation property to linkedResources for me
[**MeTodoListsTasksGetExtensions**](MeTodoApi.md#MeTodoListsTasksGetExtensions) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Get extensions from me
[**MeTodoListsTasksGetLinkedResources**](MeTodoApi.md#MeTodoListsTasksGetLinkedResources) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Get linkedResources from me
[**MeTodoListsTasksListExtensions**](MeTodoApi.md#MeTodoListsTasksListExtensions) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions | Get extensions from me
[**MeTodoListsTasksListLinkedResources**](MeTodoApi.md#MeTodoListsTasksListLinkedResources) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources | Get linkedResources from me
[**MeTodoListsTasksUpdateExtensions**](MeTodoApi.md#MeTodoListsTasksUpdateExtensions) | **Patch** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Update the navigation property extensions in me
[**MeTodoListsTasksUpdateLinkedResources**](MeTodoApi.md#MeTodoListsTasksUpdateLinkedResources) | **Patch** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Update the navigation property linkedResources in me
[**MeTodoListsUpdateExtensions**](MeTodoApi.md#MeTodoListsUpdateExtensions) | **Patch** /me/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Update the navigation property extensions in me
[**MeTodoListsUpdateTasks**](MeTodoApi.md#MeTodoListsUpdateTasks) | **Patch** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Update the navigation property tasks in me
[**MeTodoUpdateLists**](MeTodoApi.md#MeTodoUpdateLists) | **Patch** /me/todo/lists/{todoTaskList-id} | Update the navigation property lists in me
[**MeUpdateTodo**](MeTodoApi.md#MeUpdateTodo) | **Patch** /me/todo | Update the navigation property todo in me



## MeGetTodo

> MicrosoftGraphTodo MeGetTodo(ctx, optional)

Get todo from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeGetTodoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetTodoOpts struct


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


## MeTodoCreateLists

> MicrosoftGraphTodoTaskList MeTodoCreateLists(ctx, microsoftGraphTodoTaskList)

Create new navigation property to lists for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTodoGetLists

> MicrosoftGraphTodoTaskList MeTodoGetLists(ctx, todoTaskListId, optional)

Get lists from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
 **optional** | ***MeTodoGetListsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoGetListsOpts struct


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


## MeTodoListLists

> CollectionOfTodoTaskList MeTodoListLists(ctx, optional)

Get lists from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeTodoListListsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoListListsOpts struct


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


## MeTodoListsCreateExtensions

> MicrosoftGraphExtension MeTodoListsCreateExtensions(ctx, todoTaskListId, microsoftGraphExtension)

Create new navigation property to extensions for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTodoListsCreateTasks

> MicrosoftGraphTodoTask MeTodoListsCreateTasks(ctx, todoTaskListId, microsoftGraphTodoTask)

Create new navigation property to tasks for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTodoListsGetExtensions

> MicrosoftGraphExtension MeTodoListsGetExtensions(ctx, todoTaskListId, extensionId, optional)

Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***MeTodoListsGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoListsGetExtensionsOpts struct


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


## MeTodoListsGetTasks

> MicrosoftGraphTodoTask MeTodoListsGetTasks(ctx, todoTaskListId, todoTaskId, optional)

Get tasks from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
 **optional** | ***MeTodoListsGetTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoListsGetTasksOpts struct


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


## MeTodoListsListExtensions

> CollectionOfExtension MeTodoListsListExtensions(ctx, todoTaskListId, optional)

Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
 **optional** | ***MeTodoListsListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoListsListExtensionsOpts struct


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


## MeTodoListsListTasks

> CollectionOfTodoTask MeTodoListsListTasks(ctx, todoTaskListId, optional)

Get tasks from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
 **optional** | ***MeTodoListsListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoListsListTasksOpts struct


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


## MeTodoListsTasksCreateExtensions

> MicrosoftGraphExtension MeTodoListsTasksCreateExtensions(ctx, todoTaskListId, todoTaskId, microsoftGraphExtension)

Create new navigation property to extensions for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTodoListsTasksCreateLinkedResources

> MicrosoftGraphLinkedResource MeTodoListsTasksCreateLinkedResources(ctx, todoTaskListId, todoTaskId, microsoftGraphLinkedResource)

Create new navigation property to linkedResources for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTodoListsTasksGetExtensions

> MicrosoftGraphExtension MeTodoListsTasksGetExtensions(ctx, todoTaskListId, todoTaskId, extensionId, optional)

Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***MeTodoListsTasksGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoListsTasksGetExtensionsOpts struct


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


## MeTodoListsTasksGetLinkedResources

> MicrosoftGraphLinkedResource MeTodoListsTasksGetLinkedResources(ctx, todoTaskListId, todoTaskId, linkedResourceId, optional)

Get linkedResources from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
**linkedResourceId** | **string**| key: linkedResource-id of linkedResource | 
 **optional** | ***MeTodoListsTasksGetLinkedResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoListsTasksGetLinkedResourcesOpts struct


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


## MeTodoListsTasksListExtensions

> CollectionOfExtension MeTodoListsTasksListExtensions(ctx, todoTaskListId, todoTaskId, optional)

Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
 **optional** | ***MeTodoListsTasksListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoListsTasksListExtensionsOpts struct


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


## MeTodoListsTasksListLinkedResources

> CollectionOfLinkedResource MeTodoListsTasksListLinkedResources(ctx, todoTaskListId, todoTaskId, optional)

Get linkedResources from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string**| key: todoTaskList-id of todoTaskList | 
**todoTaskId** | **string**| key: todoTask-id of todoTask | 
 **optional** | ***MeTodoListsTasksListLinkedResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeTodoListsTasksListLinkedResourcesOpts struct


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


## MeTodoListsTasksUpdateExtensions

> MeTodoListsTasksUpdateExtensions(ctx, todoTaskListId, todoTaskId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTodoListsTasksUpdateLinkedResources

> MeTodoListsTasksUpdateLinkedResources(ctx, todoTaskListId, todoTaskId, linkedResourceId, microsoftGraphLinkedResource)

Update the navigation property linkedResources in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTodoListsUpdateExtensions

> MeTodoListsUpdateExtensions(ctx, todoTaskListId, extensionId, microsoftGraphExtension)

Update the navigation property extensions in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTodoListsUpdateTasks

> MeTodoListsUpdateTasks(ctx, todoTaskListId, todoTaskId, microsoftGraphTodoTask)

Update the navigation property tasks in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeTodoUpdateLists

> MeTodoUpdateLists(ctx, todoTaskListId, microsoftGraphTodoTaskList)

Update the navigation property lists in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeUpdateTodo

> MeUpdateTodo(ctx, microsoftGraphTodo)

Update the navigation property todo in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

