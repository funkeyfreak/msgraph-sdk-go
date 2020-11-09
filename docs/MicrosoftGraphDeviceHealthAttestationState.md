# MicrosoftGraphDeviceHealthAttestationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttestationIdentityKey** | **string** | TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate. | [optional] 
**BitLockerStatus** | **string** | On or Off of BitLocker Drive Encryption | [optional] 
**BootAppSecurityVersion** | **string** | The security version number of the Boot Application | [optional] 
**BootDebugging** | **string** | When bootDebugging is enabled, the device is used in development and testing | [optional] 
**BootManagerSecurityVersion** | **string** | The security version number of the Boot Application | [optional] 
**BootManagerVersion** | **string** | The version of the Boot Manager | [optional] 
**BootRevisionListInfo** | **string** | The Boot Revision List that was loaded during initial boot on the attested device | [optional] 
**CodeIntegrity** | **string** |  When code integrity is enabled, code execution is restricted to integrity verified code | [optional] 
**CodeIntegrityCheckVersion** | **string** | The version of the Boot Manager | [optional] 
**CodeIntegrityPolicy** | **string** | The Code Integrity policy that is controlling the security of the boot environment | [optional] 
**ContentNamespaceUrl** | **string** | The DHA report version. (Namespace version) | [optional] 
**ContentVersion** | **string** | The HealthAttestation state schema version | [optional] 
**DataExcutionPolicy** | **string** | DEP Policy defines a set of hardware and software technologies that perform additional checks on memory  | [optional] 
**DeviceHealthAttestationStatus** | **string** | The DHA report version. (Namespace version) | [optional] 
**EarlyLaunchAntiMalwareDriverProtection** | **string** | ELAM provides protection for the computers in your network when they start up | [optional] 
**HealthAttestationSupportedStatus** | **string** | This attribute indicates if DHA is supported for the device | [optional] 
**HealthStatusMismatchInfo** | **string** | This attribute appears if DHA-Service detects an integrity issue | [optional] 
**IssuedDateTime** | [**time.Time**](time.Time.md) | The DateTime when device was evaluated or issued to MDM | [optional] 
**LastUpdateDateTime** | **string** | The Timestamp of the last update. | [optional] 
**OperatingSystemKernelDebugging** | **string** | When operatingSystemKernelDebugging is enabled, the device is used in development and testing | [optional] 
**OperatingSystemRevListInfo** | **string** | The Operating System Revision List that was loaded during initial boot on the attested device | [optional] 
**Pcr0** | **string** | The measurement that is captured in PCR[0] | [optional] 
**PcrHashAlgorithm** | **string** | Informational attribute that identifies the HASH algorithm that was used by TPM | [optional] 
**ResetCount** | **int64** | The number of times a PC device has hibernated or resumed | [optional] 
**RestartCount** | **int64** | The number of times a PC device has rebooted | [optional] 
**SafeMode** | **string** | Safe mode is a troubleshooting option for Windows that starts your computer in a limited state | [optional] 
**SecureBoot** | **string** | When Secure Boot is enabled, the core components must have the correct cryptographic signatures | [optional] 
**SecureBootConfigurationPolicyFingerPrint** | **string** | Fingerprint of the Custom Secure Boot Configuration Policy | [optional] 
**TestSigning** | **string** | When test signing is allowed, the device does not enforce signature validation during boot | [optional] 
**TpmVersion** | **string** | The security version number of the Boot Application | [optional] 
**VirtualSecureMode** | **string** | VSM is a container that protects high value assets from a compromised kernel | [optional] 
**WindowsPE** | **string** | Operating system running with limited services that is used to prepare a computer for Windows | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


