# MicrosoftGraphWindowsInformationProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**CreatedDateTime** | [**time.Time**](time.Time.md) | The date and time the policy was created. | [optional] 
**Description** | **string** | The policy&#39;s description. | [optional] 
**DisplayName** | **string** | Policy display name. | [optional] 
**LastModifiedDateTime** | [**time.Time**](time.Time.md) | Last time the policy was modified. | [optional] 
**Version** | **string** | Version of the entity. | [optional] 
**AzureRightsManagementServicesAllowed** | **bool** | Specifies whether to allow Azure RMS encryption for WIP | [optional] 
**DataRecoveryCertificate** | [**interface{}**](.md) | Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data recovery agent(DRA) certificate for encrypting file system(EFS) | [optional] 
**EnforcementLevel** | [**interface{}**](.md) | WIP enforcement level.See the Enum definition for supported values | [optional] 
**EnterpriseDomain** | **string** | Primary enterprise domain | [optional] 
**EnterpriseInternalProxyServers** | **[]interface{}** | This is the comma-separated list of internal proxy servers. For example, \&quot;157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59\&quot;. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies | [optional] 
**EnterpriseIPRanges** | **[]interface{}** | Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to | [optional] 
**EnterpriseIPRangesAreAuthoritative** | **bool** | Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false | [optional] 
**EnterpriseNetworkDomainNames** | **[]interface{}** | This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected These locations will be considered a safe destination for enterprise data to be shared to | [optional] 
**EnterpriseProtectedDomainNames** | **[]interface{}** | List of enterprise domains to be protected | [optional] 
**EnterpriseProxiedDomains** | **[]interface{}** | Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy | [optional] 
**EnterpriseProxyServers** | **[]interface{}** | This is a list of proxy servers. Any server not on this list is considered non-enterprise | [optional] 
**EnterpriseProxyServersAreAuthoritative** | **bool** | Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false | [optional] 
**ExemptApps** | **[]interface{}** | Exempt applications can also access enterprise data, but the data handled by those applications are not protected. This is because some critical enterprise applications may have compatibility problems with encrypted data. | [optional] 
**IconsVisible** | **bool** | Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in the title bar of a WIP-protected app | [optional] 
**IndexingEncryptedStoresOrItemsBlocked** | **bool** | This switch is for the Windows Search Indexer, to allow or disallow indexing of items | [optional] 
**IsAssigned** | **bool** | Indicates if the policy is deployed to any inclusion groups or not. | [optional] 
**NeutralDomainResources** | **[]interface{}** | List of domain names that can used for work or personal resource | [optional] 
**ProtectedApps** | **[]interface{}** | Protected applications can access enterprise data and the data handled by those applications are protected with encryption | [optional] 
**ProtectionUnderLockConfigRequired** | **bool** | Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured | [optional] 
**RevokeOnUnenrollDisabled** | **bool** | This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1 (Don&#39;t revoke keys), the keys will not be revoked and the user will continue to have access to protected files after unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently. | [optional] 
**RightsManagementServicesTemplateId** | **string** | TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who has access to RMS-protected file and how long they have access | [optional] 
**SmbAutoEncryptedFileExtensions** | **[]interface{}** | Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share within the corporate boundary | [optional] 
**Assignments** | [**[]MicrosoftGraphTargetedManagedAppPolicyAssignment**](microsoft.graph.targetedManagedAppPolicyAssignment.md) |  | [optional] 
**ExemptAppLockerFiles** | [**[]MicrosoftGraphWindowsInformationProtectionAppLockerFile**](microsoft.graph.windowsInformationProtectionAppLockerFile.md) |  | [optional] 
**ProtectedAppLockerFiles** | [**[]MicrosoftGraphWindowsInformationProtectionAppLockerFile**](microsoft.graph.windowsInformationProtectionAppLockerFile.md) |  | [optional] 
**DaysWithoutContactBeforeUnenroll** | **int32** | Offline interval before app data is wiped (days)  | [optional] 
**MdmEnrollmentUrl** | **string** | Enrollment url for the MDM | [optional] 
**MinutesOfInactivityBeforeDeviceLock** | **int32** | Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 &lt;&#x3D; X &lt;&#x3D; 999. | [optional] 
**NumberOfPastPinsRemembered** | **int32** | Integer value that specifies the number of past PINs that can be associated to a user account that can&#39;t be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**PasswordMaximumAttemptCount** | **int32** | The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 &lt;&#x3D; X &lt;&#x3D; 16 for desktop and 0 &lt;&#x3D; X &lt;&#x3D; 999 for mobile devices. | [optional] 
**PinExpirationDays** | **int32** | Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user&#39;s PIN will never expire. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**PinLowercaseLetters** | [**interface{}**](.md) | Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow. | [optional] 
**PinMinimumLength** | **int32** | Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest. | [optional] 
**PinSpecialCharacters** | [**interface{}**](.md) | Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! \&quot; # $ % &amp; &#39; ( ) * + , - . / : ; &lt; &#x3D; &gt; ? @ [ \\ ] ^ _ &#x60; { | } ~. Default is NotAllow. | [optional] 
**PinUppercaseLetters** | [**interface{}**](.md) | Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow. | [optional] 
**RevokeOnMdmHandoffDisabled** | **bool** | New property in RS2, pending documentation | [optional] 
**WindowsHelloForBusinessBlocked** | **bool** | Boolean value that sets Windows Hello for Business as a method for signing into Windows. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


