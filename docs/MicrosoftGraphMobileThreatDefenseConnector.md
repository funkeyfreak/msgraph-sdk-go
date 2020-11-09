# MicrosoftGraphMobileThreatDefenseConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**AndroidDeviceBlockedOnMissingPartnerData** | **bool** | For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant | [optional] 
**AndroidEnabled** | **bool** | For Android, set whether data from the data sync partner should be used during compliance evaluations | [optional] 
**IosDeviceBlockedOnMissingPartnerData** | **bool** | For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant | [optional] 
**IosEnabled** | **bool** | For IOS, get or set whether data from the data sync partner should be used during compliance evaluations | [optional] 
**LastHeartbeatDateTime** | [**time.Time**](time.Time.md) | DateTime of last Heartbeat recieved from the Data Sync Partner | [optional] 
**PartnerState** | [**interface{}**](.md) | Data Sync Partner state for this account | [optional] 
**PartnerUnresponsivenessThresholdInDays** | **int32** | Get or Set days the per tenant tolerance to unresponsiveness for this partner integration | [optional] 
**PartnerUnsupportedOsVersionBlocked** | **bool** | Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


