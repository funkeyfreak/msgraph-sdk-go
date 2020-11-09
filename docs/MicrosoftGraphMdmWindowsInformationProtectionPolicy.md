# MicrosoftGraphMdmWindowsInformationProtectionPolicy

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


