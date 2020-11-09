# DeviceCompliancePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | [**time.Time**](time.Time.md) | DateTime the object was created. | [optional] 
**Description** | **string** | Admin provided description of the Device Configuration. | [optional] 
**DisplayName** | **string** | Admin provided name of the device configuration. | [optional] 
**LastModifiedDateTime** | [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**Version** | **int32** | Version of the device configuration. | [optional] 
**Assignments** | [**[]MicrosoftGraphDeviceCompliancePolicyAssignment**](microsoft.graph.deviceCompliancePolicyAssignment.md) |  | [optional] 
**DeviceSettingStateSummaries** | [**[]MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md) |  | [optional] 
**DeviceStatuses** | [**[]MicrosoftGraphDeviceComplianceDeviceStatus**](microsoft.graph.deviceComplianceDeviceStatus.md) |  | [optional] 
**DeviceStatusOverview** | [**interface{}**](.md) |  | [optional] 
**ScheduledActionsForRule** | [**[]MicrosoftGraphDeviceComplianceScheduledActionForRule**](microsoft.graph.deviceComplianceScheduledActionForRule.md) |  | [optional] 
**UserStatuses** | [**[]MicrosoftGraphDeviceComplianceUserStatus**](microsoft.graph.deviceComplianceUserStatus.md) |  | [optional] 
**UserStatusOverview** | [**interface{}**](.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

