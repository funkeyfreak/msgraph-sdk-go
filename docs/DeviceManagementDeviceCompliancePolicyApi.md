# \DeviceManagementDeviceCompliancePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceCompliancePolicies**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementCreateDeviceCompliancePolicies) | **Post** /deviceManagement/deviceCompliancePolicies | Create new navigation property to deviceCompliancePolicies for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesCreateAssignments**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesCreateAssignments) | **Post** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/assignments | Create new navigation property to assignments for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesCreateDeviceSettingStateSummaries**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesCreateDeviceSettingStateSummaries) | **Post** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceSettingStateSummaries | Create new navigation property to deviceSettingStateSummaries for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesCreateDeviceStatuses**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesCreateDeviceStatuses) | **Post** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceStatuses | Create new navigation property to deviceStatuses for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesCreateScheduledActionsForRule**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesCreateScheduledActionsForRule) | **Post** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/scheduledActionsForRule | Create new navigation property to scheduledActionsForRule for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesCreateUserStatuses**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesCreateUserStatuses) | **Post** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/userStatuses | Create new navigation property to userStatuses for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetAssignments**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesGetAssignments) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/assignments/{deviceCompliancePolicyAssignment-id} | Get assignments from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetDeviceSettingStateSummaries**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesGetDeviceSettingStateSummaries) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceSettingStateSummaries/{settingStateDeviceSummary-id} | Get deviceSettingStateSummaries from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverview**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverview) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceStatusOverview | Get deviceStatusOverview from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetDeviceStatuses**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesGetDeviceStatuses) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceStatuses/{deviceComplianceDeviceStatus-id} | Get deviceStatuses from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRule**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRule) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule-id} | Get scheduledActionsForRule from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetUserStatusOverview**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesGetUserStatusOverview) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/userStatusOverview | Get userStatusOverview from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetUserStatuses**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesGetUserStatuses) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/userStatuses/{deviceComplianceUserStatus-id} | Get userStatuses from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesListAssignments**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesListAssignments) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/assignments | Get assignments from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesListDeviceSettingStateSummaries**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesListDeviceSettingStateSummaries) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceSettingStateSummaries | Get deviceSettingStateSummaries from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesListDeviceStatuses**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesListDeviceStatuses) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceStatuses | Get deviceStatuses from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRule**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRule) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/scheduledActionsForRule | Get scheduledActionsForRule from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesListUserStatuses**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesListUserStatuses) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/userStatuses | Get userStatuses from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleCreateScheduledActionConfigurations**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleCreateScheduledActionConfigurations) | **Post** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule-id}/scheduledActionConfigurations | Create new navigation property to scheduledActionConfigurations for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurations**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurations) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule-id}/scheduledActionConfigurations/{deviceComplianceActionItem-id} | Get scheduledActionConfigurations from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurations**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurations) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule-id}/scheduledActionConfigurations | Get scheduledActionConfigurations from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleUpdateScheduledActionConfigurations**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleUpdateScheduledActionConfigurations) | **Patch** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule-id}/scheduledActionConfigurations/{deviceComplianceActionItem-id} | Update the navigation property scheduledActionConfigurations in deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateAssignments**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesUpdateAssignments) | **Patch** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/assignments/{deviceCompliancePolicyAssignment-id} | Update the navigation property assignments in deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateDeviceSettingStateSummaries**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesUpdateDeviceSettingStateSummaries) | **Patch** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceSettingStateSummaries/{settingStateDeviceSummary-id} | Update the navigation property deviceSettingStateSummaries in deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatusOverview**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatusOverview) | **Patch** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceStatusOverview | Update the navigation property deviceStatusOverview in deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatuses**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatuses) | **Patch** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/deviceStatuses/{deviceComplianceDeviceStatus-id} | Update the navigation property deviceStatuses in deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateScheduledActionsForRule**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesUpdateScheduledActionsForRule) | **Patch** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule-id} | Update the navigation property scheduledActionsForRule in deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateUserStatusOverview**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesUpdateUserStatusOverview) | **Patch** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/userStatusOverview | Update the navigation property userStatusOverview in deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateUserStatuses**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementDeviceCompliancePoliciesUpdateUserStatuses) | **Patch** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id}/userStatuses/{deviceComplianceUserStatus-id} | Update the navigation property userStatuses in deviceManagement
[**DeviceManagementGetDeviceCompliancePolicies**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementGetDeviceCompliancePolicies) | **Get** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id} | Get deviceCompliancePolicies from deviceManagement
[**DeviceManagementListDeviceCompliancePolicies**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementListDeviceCompliancePolicies) | **Get** /deviceManagement/deviceCompliancePolicies | Get deviceCompliancePolicies from deviceManagement
[**DeviceManagementUpdateDeviceCompliancePolicies**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementUpdateDeviceCompliancePolicies) | **Patch** /deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy-id} | Update the navigation property deviceCompliancePolicies in deviceManagement



## DeviceManagementCreateDeviceCompliancePolicies

> MicrosoftGraphDeviceCompliancePolicy DeviceManagementCreateDeviceCompliancePolicies(ctx, microsoftGraphDeviceCompliancePolicy)

Create new navigation property to deviceCompliancePolicies for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceCompliancePolicy** | [**MicrosoftGraphDeviceCompliancePolicy**](MicrosoftGraphDeviceCompliancePolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicy**](microsoft.graph.deviceCompliancePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesCreateAssignments

> MicrosoftGraphDeviceCompliancePolicyAssignment DeviceManagementDeviceCompliancePoliciesCreateAssignments(ctx, deviceCompliancePolicyId, microsoftGraphDeviceCompliancePolicyAssignment)

Create new navigation property to assignments for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceCompliancePolicyAssignment** | [**MicrosoftGraphDeviceCompliancePolicyAssignment**](MicrosoftGraphDeviceCompliancePolicyAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyAssignment**](microsoft.graph.deviceCompliancePolicyAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesCreateDeviceSettingStateSummaries

> MicrosoftGraphSettingStateDeviceSummary DeviceManagementDeviceCompliancePoliciesCreateDeviceSettingStateSummaries(ctx, deviceCompliancePolicyId, microsoftGraphSettingStateDeviceSummary)

Create new navigation property to deviceSettingStateSummaries for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphSettingStateDeviceSummary** | [**MicrosoftGraphSettingStateDeviceSummary**](MicrosoftGraphSettingStateDeviceSummary.md)| New navigation property | 

### Return type

[**MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesCreateDeviceStatuses

> MicrosoftGraphDeviceComplianceDeviceStatus DeviceManagementDeviceCompliancePoliciesCreateDeviceStatuses(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceDeviceStatus)

Create new navigation property to deviceStatuses for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceDeviceStatus** | [**MicrosoftGraphDeviceComplianceDeviceStatus**](MicrosoftGraphDeviceComplianceDeviceStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceDeviceStatus**](microsoft.graph.deviceComplianceDeviceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesCreateScheduledActionsForRule

> MicrosoftGraphDeviceComplianceScheduledActionForRule DeviceManagementDeviceCompliancePoliciesCreateScheduledActionsForRule(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceScheduledActionForRule)

Create new navigation property to scheduledActionsForRule for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceScheduledActionForRule** | [**MicrosoftGraphDeviceComplianceScheduledActionForRule**](MicrosoftGraphDeviceComplianceScheduledActionForRule.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceScheduledActionForRule**](microsoft.graph.deviceComplianceScheduledActionForRule.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesCreateUserStatuses

> MicrosoftGraphDeviceComplianceUserStatus DeviceManagementDeviceCompliancePoliciesCreateUserStatuses(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceUserStatus)

Create new navigation property to userStatuses for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceUserStatus** | [**MicrosoftGraphDeviceComplianceUserStatus**](MicrosoftGraphDeviceComplianceUserStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceUserStatus**](microsoft.graph.deviceComplianceUserStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetAssignments

> MicrosoftGraphDeviceCompliancePolicyAssignment DeviceManagementDeviceCompliancePoliciesGetAssignments(ctx, deviceCompliancePolicyId, deviceCompliancePolicyAssignmentId, optional)

Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceCompliancePolicyAssignmentId** | **string**| key: deviceCompliancePolicyAssignment-id of deviceCompliancePolicyAssignment | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyAssignment**](microsoft.graph.deviceCompliancePolicyAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetDeviceSettingStateSummaries

> MicrosoftGraphSettingStateDeviceSummary DeviceManagementDeviceCompliancePoliciesGetDeviceSettingStateSummaries(ctx, deviceCompliancePolicyId, settingStateDeviceSummaryId, optional)

Get deviceSettingStateSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**settingStateDeviceSummaryId** | **string**| key: settingStateDeviceSummary-id of settingStateDeviceSummary | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetDeviceSettingStateSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetDeviceSettingStateSummariesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverview

> MicrosoftGraphDeviceComplianceDeviceOverview DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverview(ctx, deviceCompliancePolicyId, optional)

Get deviceStatusOverview from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceDeviceOverview**](microsoft.graph.deviceComplianceDeviceOverview.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetDeviceStatuses

> MicrosoftGraphDeviceComplianceDeviceStatus DeviceManagementDeviceCompliancePoliciesGetDeviceStatuses(ctx, deviceCompliancePolicyId, deviceComplianceDeviceStatusId, optional)

Get deviceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceDeviceStatusId** | **string**| key: deviceComplianceDeviceStatus-id of deviceComplianceDeviceStatus | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetDeviceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetDeviceStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceDeviceStatus**](microsoft.graph.deviceComplianceDeviceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRule

> MicrosoftGraphDeviceComplianceScheduledActionForRule DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRule(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, optional)

Get scheduledActionsForRule from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceScheduledActionForRule**](microsoft.graph.deviceComplianceScheduledActionForRule.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetUserStatusOverview

> MicrosoftGraphDeviceComplianceUserOverview DeviceManagementDeviceCompliancePoliciesGetUserStatusOverview(ctx, deviceCompliancePolicyId, optional)

Get userStatusOverview from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetUserStatusOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetUserStatusOverviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceUserOverview**](microsoft.graph.deviceComplianceUserOverview.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetUserStatuses

> MicrosoftGraphDeviceComplianceUserStatus DeviceManagementDeviceCompliancePoliciesGetUserStatuses(ctx, deviceCompliancePolicyId, deviceComplianceUserStatusId, optional)

Get userStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceUserStatusId** | **string**| key: deviceComplianceUserStatus-id of deviceComplianceUserStatus | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetUserStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetUserStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceUserStatus**](microsoft.graph.deviceComplianceUserStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesListAssignments

> CollectionOfDeviceCompliancePolicyAssignment DeviceManagementDeviceCompliancePoliciesListAssignments(ctx, deviceCompliancePolicyId, optional)

Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesListAssignmentsOpts struct


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

[**CollectionOfDeviceCompliancePolicyAssignment**](Collection_of_deviceCompliancePolicyAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesListDeviceSettingStateSummaries

> CollectionOfSettingStateDeviceSummary DeviceManagementDeviceCompliancePoliciesListDeviceSettingStateSummaries(ctx, deviceCompliancePolicyId, optional)

Get deviceSettingStateSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesListDeviceSettingStateSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesListDeviceSettingStateSummariesOpts struct


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

[**CollectionOfSettingStateDeviceSummary**](Collection_of_settingStateDeviceSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesListDeviceStatuses

> CollectionOfDeviceComplianceDeviceStatus DeviceManagementDeviceCompliancePoliciesListDeviceStatuses(ctx, deviceCompliancePolicyId, optional)

Get deviceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesListDeviceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesListDeviceStatusesOpts struct


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

[**CollectionOfDeviceComplianceDeviceStatus**](Collection_of_deviceComplianceDeviceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRule

> CollectionOfDeviceComplianceScheduledActionForRule DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRule(ctx, deviceCompliancePolicyId, optional)

Get scheduledActionsForRule from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRuleOpts struct


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

[**CollectionOfDeviceComplianceScheduledActionForRule**](Collection_of_deviceComplianceScheduledActionForRule.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesListUserStatuses

> CollectionOfDeviceComplianceUserStatus DeviceManagementDeviceCompliancePoliciesListUserStatuses(ctx, deviceCompliancePolicyId, optional)

Get userStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesListUserStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesListUserStatusesOpts struct


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

[**CollectionOfDeviceComplianceUserStatus**](Collection_of_deviceComplianceUserStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleCreateScheduledActionConfigurations

> MicrosoftGraphDeviceComplianceActionItem DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleCreateScheduledActionConfigurations(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, microsoftGraphDeviceComplianceActionItem)

Create new navigation property to scheduledActionConfigurations for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
**microsoftGraphDeviceComplianceActionItem** | [**MicrosoftGraphDeviceComplianceActionItem**](MicrosoftGraphDeviceComplianceActionItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceActionItem**](microsoft.graph.deviceComplianceActionItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurations

> MicrosoftGraphDeviceComplianceActionItem DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurations(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, deviceComplianceActionItemId, optional)

Get scheduledActionConfigurations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
**deviceComplianceActionItemId** | **string**| key: deviceComplianceActionItem-id of deviceComplianceActionItem | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceActionItem**](microsoft.graph.deviceComplianceActionItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurations

> CollectionOfDeviceComplianceActionItem DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurations(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, optional)

Get scheduledActionConfigurations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurationsOpts struct


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

[**CollectionOfDeviceComplianceActionItem**](Collection_of_deviceComplianceActionItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleUpdateScheduledActionConfigurations

> DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleUpdateScheduledActionConfigurations(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, deviceComplianceActionItemId, microsoftGraphDeviceComplianceActionItem)

Update the navigation property scheduledActionConfigurations in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
**deviceComplianceActionItemId** | **string**| key: deviceComplianceActionItem-id of deviceComplianceActionItem | 
**microsoftGraphDeviceComplianceActionItem** | [**MicrosoftGraphDeviceComplianceActionItem**](MicrosoftGraphDeviceComplianceActionItem.md)| New navigation property values | 

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


## DeviceManagementDeviceCompliancePoliciesUpdateAssignments

> DeviceManagementDeviceCompliancePoliciesUpdateAssignments(ctx, deviceCompliancePolicyId, deviceCompliancePolicyAssignmentId, microsoftGraphDeviceCompliancePolicyAssignment)

Update the navigation property assignments in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceCompliancePolicyAssignmentId** | **string**| key: deviceCompliancePolicyAssignment-id of deviceCompliancePolicyAssignment | 
**microsoftGraphDeviceCompliancePolicyAssignment** | [**MicrosoftGraphDeviceCompliancePolicyAssignment**](MicrosoftGraphDeviceCompliancePolicyAssignment.md)| New navigation property values | 

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


## DeviceManagementDeviceCompliancePoliciesUpdateDeviceSettingStateSummaries

> DeviceManagementDeviceCompliancePoliciesUpdateDeviceSettingStateSummaries(ctx, deviceCompliancePolicyId, settingStateDeviceSummaryId, microsoftGraphSettingStateDeviceSummary)

Update the navigation property deviceSettingStateSummaries in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**settingStateDeviceSummaryId** | **string**| key: settingStateDeviceSummary-id of settingStateDeviceSummary | 
**microsoftGraphSettingStateDeviceSummary** | [**MicrosoftGraphSettingStateDeviceSummary**](MicrosoftGraphSettingStateDeviceSummary.md)| New navigation property values | 

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


## DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatusOverview

> DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatusOverview(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceDeviceOverview)

Update the navigation property deviceStatusOverview in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceDeviceOverview** | [**MicrosoftGraphDeviceComplianceDeviceOverview**](MicrosoftGraphDeviceComplianceDeviceOverview.md)| New navigation property values | 

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


## DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatuses

> DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatuses(ctx, deviceCompliancePolicyId, deviceComplianceDeviceStatusId, microsoftGraphDeviceComplianceDeviceStatus)

Update the navigation property deviceStatuses in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceDeviceStatusId** | **string**| key: deviceComplianceDeviceStatus-id of deviceComplianceDeviceStatus | 
**microsoftGraphDeviceComplianceDeviceStatus** | [**MicrosoftGraphDeviceComplianceDeviceStatus**](MicrosoftGraphDeviceComplianceDeviceStatus.md)| New navigation property values | 

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


## DeviceManagementDeviceCompliancePoliciesUpdateScheduledActionsForRule

> DeviceManagementDeviceCompliancePoliciesUpdateScheduledActionsForRule(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, microsoftGraphDeviceComplianceScheduledActionForRule)

Update the navigation property scheduledActionsForRule in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
**microsoftGraphDeviceComplianceScheduledActionForRule** | [**MicrosoftGraphDeviceComplianceScheduledActionForRule**](MicrosoftGraphDeviceComplianceScheduledActionForRule.md)| New navigation property values | 

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


## DeviceManagementDeviceCompliancePoliciesUpdateUserStatusOverview

> DeviceManagementDeviceCompliancePoliciesUpdateUserStatusOverview(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceUserOverview)

Update the navigation property userStatusOverview in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceUserOverview** | [**MicrosoftGraphDeviceComplianceUserOverview**](MicrosoftGraphDeviceComplianceUserOverview.md)| New navigation property values | 

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


## DeviceManagementDeviceCompliancePoliciesUpdateUserStatuses

> DeviceManagementDeviceCompliancePoliciesUpdateUserStatuses(ctx, deviceCompliancePolicyId, deviceComplianceUserStatusId, microsoftGraphDeviceComplianceUserStatus)

Update the navigation property userStatuses in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceUserStatusId** | **string**| key: deviceComplianceUserStatus-id of deviceComplianceUserStatus | 
**microsoftGraphDeviceComplianceUserStatus** | [**MicrosoftGraphDeviceComplianceUserStatus**](MicrosoftGraphDeviceComplianceUserStatus.md)| New navigation property values | 

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


## DeviceManagementGetDeviceCompliancePolicies

> MicrosoftGraphDeviceCompliancePolicy DeviceManagementGetDeviceCompliancePolicies(ctx, deviceCompliancePolicyId, optional)

Get deviceCompliancePolicies from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementGetDeviceCompliancePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDeviceCompliancePoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicy**](microsoft.graph.deviceCompliancePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceCompliancePolicies

> CollectionOfDeviceCompliancePolicy DeviceManagementListDeviceCompliancePolicies(ctx, optional)

Get deviceCompliancePolicies from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListDeviceCompliancePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListDeviceCompliancePoliciesOpts struct


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

[**CollectionOfDeviceCompliancePolicy**](Collection_of_deviceCompliancePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceCompliancePolicies

> DeviceManagementUpdateDeviceCompliancePolicies(ctx, deviceCompliancePolicyId, microsoftGraphDeviceCompliancePolicy)

Update the navigation property deviceCompliancePolicies in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceCompliancePolicy** | [**MicrosoftGraphDeviceCompliancePolicy**](MicrosoftGraphDeviceCompliancePolicy.md)| New navigation property values | 

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

