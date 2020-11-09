# \SitesFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesGetActivitiesByInterval53ee**](SitesFunctionsApi.md#SitesGetActivitiesByInterval53ee) | **Get** /sites/{site-id}/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;{startDateTime},endDateTime&#x3D;{endDateTime},interval&#x3D;{interval}) | Invoke function getActivitiesByInterval
[**SitesGetActivitiesByInterval96b0**](SitesFunctionsApi.md#SitesGetActivitiesByInterval96b0) | **Get** /sites/{site-id}/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**SitesGetByPath**](SitesFunctionsApi.md#SitesGetByPath) | **Get** /sites/{site-id}/microsoft.graph.getByPath(path&#x3D;{path}) | Invoke function getByPath
[**SitesListsItemsGetActivitiesByInterval53ee**](SitesFunctionsApi.md#SitesListsItemsGetActivitiesByInterval53ee) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;{startDateTime},endDateTime&#x3D;{endDateTime},interval&#x3D;{interval}) | Invoke function getActivitiesByInterval
[**SitesListsItemsGetActivitiesByInterval96b0**](SitesFunctionsApi.md#SitesListsItemsGetActivitiesByInterval96b0) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**SitesOnenoteNotebooksGetRecentNotebooks**](SitesFunctionsApi.md#SitesOnenoteNotebooksGetRecentNotebooks) | **Get** /sites/{site-id}/onenote/notebooks/microsoft.graph.getRecentNotebooks(includePersonalNotebooks&#x3D;{includePersonalNotebooks}) | Invoke function getRecentNotebooks
[**SitesOnenoteNotebooksSectionGroupsSectionsPagesPreview**](SitesFunctionsApi.md#SitesOnenoteNotebooksSectionGroupsSectionsPagesPreview) | **Get** /sites/{site-id}/onenote/notebooks/{notebook-id}/sectionGroups/{sectionGroup-id}/sections/{onenoteSection-id}/pages/{onenotePage-id}/microsoft.graph.preview() | Invoke function preview
[**SitesOnenoteNotebooksSectionsPagesPreview**](SitesFunctionsApi.md#SitesOnenoteNotebooksSectionsPagesPreview) | **Get** /sites/{site-id}/onenote/notebooks/{notebook-id}/sections/{onenoteSection-id}/pages/{onenotePage-id}/microsoft.graph.preview() | Invoke function preview
[**SitesOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview**](SitesFunctionsApi.md#SitesOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview) | **Get** /sites/{site-id}/onenote/pages/{onenotePage-id}/parentNotebook/sectionGroups/{sectionGroup-id}/sections/{onenoteSection-id}/pages/{onenotePage-id1}/microsoft.graph.preview() | Invoke function preview
[**SitesOnenotePagesParentNotebookSectionsPagesPreview**](SitesFunctionsApi.md#SitesOnenotePagesParentNotebookSectionsPagesPreview) | **Get** /sites/{site-id}/onenote/pages/{onenotePage-id}/parentNotebook/sections/{onenoteSection-id}/pages/{onenotePage-id1}/microsoft.graph.preview() | Invoke function preview
[**SitesOnenotePagesParentSectionPagesPreview**](SitesFunctionsApi.md#SitesOnenotePagesParentSectionPagesPreview) | **Get** /sites/{site-id}/onenote/pages/{onenotePage-id}/parentSection/pages/{onenotePage-id1}/microsoft.graph.preview() | Invoke function preview
[**SitesOnenotePagesPreview**](SitesFunctionsApi.md#SitesOnenotePagesPreview) | **Get** /sites/{site-id}/onenote/pages/{onenotePage-id}/microsoft.graph.preview() | Invoke function preview
[**SitesOnenoteSectionGroupsParentNotebookSectionsPagesPreview**](SitesFunctionsApi.md#SitesOnenoteSectionGroupsParentNotebookSectionsPagesPreview) | **Get** /sites/{site-id}/onenote/sectionGroups/{sectionGroup-id}/parentNotebook/sections/{onenoteSection-id}/pages/{onenotePage-id}/microsoft.graph.preview() | Invoke function preview
[**SitesOnenoteSectionGroupsSectionsPagesPreview**](SitesFunctionsApi.md#SitesOnenoteSectionGroupsSectionsPagesPreview) | **Get** /sites/{site-id}/onenote/sectionGroups/{sectionGroup-id}/sections/{onenoteSection-id}/pages/{onenotePage-id}/microsoft.graph.preview() | Invoke function preview
[**SitesOnenoteSectionsPagesPreview**](SitesFunctionsApi.md#SitesOnenoteSectionsPagesPreview) | **Get** /sites/{site-id}/onenote/sections/{onenoteSection-id}/pages/{onenotePage-id}/microsoft.graph.preview() | Invoke function preview



## SitesGetActivitiesByInterval53ee

> []interface{} SitesGetActivitiesByInterval53ee(ctx, siteId, startDateTime, endDateTime, interval)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**startDateTime** | **string**|  | 
**endDateTime** | **string**|  | 
**interval** | **string**|  | 

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetActivitiesByInterval96b0

> []interface{} SitesGetActivitiesByInterval96b0(ctx, siteId)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetByPath

> interface{} SitesGetByPath(ctx, siteId, path)

Invoke function getByPath

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**path** | **string**|  | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsGetActivitiesByInterval53ee

> []interface{} SitesListsItemsGetActivitiesByInterval53ee(ctx, siteId, listId, listItemId, startDateTime, endDateTime, interval)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
**startDateTime** | **string**|  | 
**endDateTime** | **string**|  | 
**interval** | **string**|  | 

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsGetActivitiesByInterval96b0

> []interface{} SitesListsItemsGetActivitiesByInterval96b0(ctx, siteId, listId, listItemId)

Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteNotebooksGetRecentNotebooks

> []interface{} SitesOnenoteNotebooksGetRecentNotebooks(ctx, siteId, includePersonalNotebooks)

Invoke function getRecentNotebooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**includePersonalNotebooks** | **bool**|  | [default to false]

### Return type

**[]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteNotebooksSectionGroupsSectionsPagesPreview

> interface{} SitesOnenoteNotebooksSectionGroupsSectionsPagesPreview(ctx, siteId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId)

Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteNotebooksSectionsPagesPreview

> interface{} SitesOnenoteNotebooksSectionsPagesPreview(ctx, siteId, notebookId, onenoteSectionId, onenotePageId)

Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**notebookId** | **string**| key: notebook-id of notebook | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview

> interface{} SitesOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview(ctx, siteId, onenotePageId, sectionGroupId, onenoteSectionId, onenotePageId1)

Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId1** | **string**| key: onenotePage-id of onenotePage | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenotePagesParentNotebookSectionsPagesPreview

> interface{} SitesOnenotePagesParentNotebookSectionsPagesPreview(ctx, siteId, onenotePageId, onenoteSectionId, onenotePageId1)

Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId1** | **string**| key: onenotePage-id of onenotePage | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenotePagesParentSectionPagesPreview

> interface{} SitesOnenotePagesParentSectionPagesPreview(ctx, siteId, onenotePageId, onenotePageId1)

Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
**onenotePageId1** | **string**| key: onenotePage-id of onenotePage | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenotePagesPreview

> interface{} SitesOnenotePagesPreview(ctx, siteId, onenotePageId)

Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteSectionGroupsParentNotebookSectionsPagesPreview

> interface{} SitesOnenoteSectionGroupsParentNotebookSectionsPagesPreview(ctx, siteId, sectionGroupId, onenoteSectionId, onenotePageId)

Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteSectionGroupsSectionsPagesPreview

> interface{} SitesOnenoteSectionGroupsSectionsPagesPreview(ctx, siteId, sectionGroupId, onenoteSectionId, onenotePageId)

Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteSectionsPagesPreview

> interface{} SitesOnenoteSectionsPagesPreview(ctx, siteId, onenoteSectionId, onenotePageId)

Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

**interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

