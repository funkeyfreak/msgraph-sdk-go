# MicrosoftGraphDeviceAppManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**IsEnabledForMicrosoftStoreForBusiness** | **bool** | Whether the account is enabled for syncing applications from the Microsoft Store for Business. | [optional] 
**MicrosoftStoreForBusinessLanguage** | **string** | The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is &lt;languagecode2&gt;-&lt;country/regioncode2&gt;, where &lt;languagecode2&gt; is a lowercase two-letter code derived from ISO 639-1 and &lt;country/regioncode2&gt; is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture. | [optional] 
**MicrosoftStoreForBusinessLastCompletedApplicationSyncTime** | [**time.Time**](time.Time.md) | The last time an application sync from the Microsoft Store for Business was completed. | [optional] 
**MicrosoftStoreForBusinessLastSuccessfulSyncDateTime** | [**time.Time**](time.Time.md) | The last time the apps from the Microsoft Store for Business were synced successfully for the account. | [optional] 
**ManagedEBooks** | [**[]MicrosoftGraphManagedEBook**](microsoft.graph.managedEBook.md) |  | [optional] 
**MobileAppCategories** | [**[]MicrosoftGraphMobileAppCategory**](microsoft.graph.mobileAppCategory.md) |  | [optional] 
**MobileAppConfigurations** | [**[]MicrosoftGraphManagedDeviceMobileAppConfiguration**](microsoft.graph.managedDeviceMobileAppConfiguration.md) |  | [optional] 
**MobileApps** | [**[]MicrosoftGraphMobileApp**](microsoft.graph.mobileApp.md) |  | [optional] 
**VppTokens** | [**[]MicrosoftGraphVppToken**](microsoft.graph.vppToken.md) |  | [optional] 
**AndroidManagedAppProtections** | [**[]MicrosoftGraphAndroidManagedAppProtection**](microsoft.graph.androidManagedAppProtection.md) |  | [optional] 
**DefaultManagedAppProtections** | [**[]MicrosoftGraphDefaultManagedAppProtection**](microsoft.graph.defaultManagedAppProtection.md) |  | [optional] 
**IosManagedAppProtections** | [**[]MicrosoftGraphIosManagedAppProtection**](microsoft.graph.iosManagedAppProtection.md) |  | [optional] 
**ManagedAppPolicies** | [**[]MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md) |  | [optional] 
**ManagedAppRegistrations** | [**[]MicrosoftGraphManagedAppRegistration**](microsoft.graph.managedAppRegistration.md) |  | [optional] 
**ManagedAppStatuses** | [**[]MicrosoftGraphManagedAppStatus**](microsoft.graph.managedAppStatus.md) |  | [optional] 
**MdmWindowsInformationProtectionPolicies** | [**[]MicrosoftGraphMdmWindowsInformationProtectionPolicy**](microsoft.graph.mdmWindowsInformationProtectionPolicy.md) |  | [optional] 
**TargetedManagedAppConfigurations** | [**[]MicrosoftGraphTargetedManagedAppConfiguration**](microsoft.graph.targetedManagedAppConfiguration.md) |  | [optional] 
**WindowsInformationProtectionPolicies** | [**[]MicrosoftGraphWindowsInformationProtectionPolicy**](microsoft.graph.windowsInformationProtectionPolicy.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


