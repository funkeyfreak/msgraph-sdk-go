# MicrosoftGraphManagedDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**ActivationLockBypassCode** | **string** | Code that allows the Activation Lock on a device to be bypassed. This property is read-only. | [optional] 
**AndroidSecurityPatchLevel** | **string** | Android security patch level. This property is read-only. | [optional] 
**AzureADDeviceId** | **string** | The unique identifier for the Azure Active Directory device. Read only. This property is read-only. | [optional] 
**AzureADRegistered** | **bool** | Whether the device is Azure Active Directory registered. This property is read-only. | [optional] 
**ComplianceGracePeriodExpirationDateTime** | [**time.Time**](time.Time.md) | The DateTime when device compliance grace period expires. This property is read-only. | [optional] 
**ComplianceState** | [**interface{}**](.md) | Compliance state of the device. This property is read-only. | [optional] 
**ConfigurationManagerClientEnabledFeatures** | [**interface{}**](.md) | ConfigrMgr client enabled features. This property is read-only. | [optional] 
**DeviceActionResults** | **[]interface{}** | List of ComplexType deviceActionResult objects. This property is read-only. | [optional] 
**DeviceCategoryDisplayName** | **string** | Device category display name. This property is read-only. | [optional] 
**DeviceEnrollmentType** | [**interface{}**](.md) | Enrollment type of the device. This property is read-only. | [optional] 
**DeviceHealthAttestationState** | [**interface{}**](.md) | The device health attestation state. This property is read-only. | [optional] 
**DeviceName** | **string** | Name of the device. This property is read-only. | [optional] 
**DeviceRegistrationState** | [**interface{}**](.md) | Device registration state. This property is read-only. | [optional] 
**EasActivated** | **bool** | Whether the device is Exchange ActiveSync activated. This property is read-only. | [optional] 
**EasActivationDateTime** | [**time.Time**](time.Time.md) | Exchange ActivationSync activation time of the device. This property is read-only. | [optional] 
**EasDeviceId** | **string** | Exchange ActiveSync Id of the device. This property is read-only. | [optional] 
**EmailAddress** | **string** | Email(s) for the user associated with the device. This property is read-only. | [optional] 
**EnrolledDateTime** | [**time.Time**](time.Time.md) | Enrollment time of the device. This property is read-only. | [optional] 
**ExchangeAccessState** | [**interface{}**](.md) | The Access State of the device in Exchange. This property is read-only. | [optional] 
**ExchangeAccessStateReason** | [**interface{}**](.md) | The reason for the device&#39;s access state in Exchange. This property is read-only. | [optional] 
**ExchangeLastSuccessfulSyncDateTime** | [**time.Time**](time.Time.md) | Last time the device contacted Exchange. This property is read-only. | [optional] 
**FreeStorageSpaceInBytes** | **int64** | Free Storage in Bytes. This property is read-only. | [optional] 
**Imei** | **string** | IMEI. This property is read-only. | [optional] 
**IsEncrypted** | **bool** | Device encryption status. This property is read-only. | [optional] 
**IsSupervised** | **bool** | Device supervised status. This property is read-only. | [optional] 
**JailBroken** | **string** | whether the device is jail broken or rooted. This property is read-only. | [optional] 
**LastSyncDateTime** | [**time.Time**](time.Time.md) | The date and time that the device last completed a successful sync with Intune. This property is read-only. | [optional] 
**ManagedDeviceName** | **string** | Automatically generated name to identify a device. Can be overwritten to a user friendly name. | [optional] 
**ManagedDeviceOwnerType** | [**interface{}**](.md) | Ownership of the device. Can be &#39;company&#39; or &#39;personal&#39; | [optional] 
**ManagementAgent** | [**interface{}**](.md) | Management channel of the device. Intune, EAS, etc. This property is read-only. | [optional] 
**Manufacturer** | **string** | Manufacturer of the device. This property is read-only. | [optional] 
**Meid** | **string** | MEID. This property is read-only. | [optional] 
**Model** | **string** | Model of the device. This property is read-only. | [optional] 
**OperatingSystem** | **string** | Operating system of the device. Windows, iOS, etc. This property is read-only. | [optional] 
**OsVersion** | **string** | Operating system version of the device. This property is read-only. | [optional] 
**PartnerReportedThreatState** | [**interface{}**](.md) | Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. | [optional] 
**PhoneNumber** | **string** | Phone number of the device. This property is read-only. | [optional] 
**RemoteAssistanceSessionErrorDetails** | **string** | An error string that identifies issues when creating Remote Assistance session objects. This property is read-only. | [optional] 
**RemoteAssistanceSessionUrl** | **string** | Url that allows a Remote Assistance session to be established with the device. This property is read-only. | [optional] 
**SerialNumber** | **string** | SerialNumber. This property is read-only. | [optional] 
**SubscriberCarrier** | **string** | Subscriber Carrier. This property is read-only. | [optional] 
**TotalStorageSpaceInBytes** | **int64** | Total Storage in Bytes. This property is read-only. | [optional] 
**UserDisplayName** | **string** | User display name. This property is read-only. | [optional] 
**UserId** | **string** | Unique Identifier for the user associated with the device. This property is read-only. | [optional] 
**UserPrincipalName** | **string** | Device user principal name. This property is read-only. | [optional] 
**WiFiMacAddress** | **string** | Wi-Fi MAC. This property is read-only. | [optional] 
**DeviceCompliancePolicyStates** | [**[]MicrosoftGraphDeviceCompliancePolicyState**](microsoft.graph.deviceCompliancePolicyState.md) |  | [optional] 
**DeviceConfigurationStates** | [**[]MicrosoftGraphDeviceConfigurationState**](microsoft.graph.deviceConfigurationState.md) |  | [optional] 
**DeviceCategory** | [**interface{}**](.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


