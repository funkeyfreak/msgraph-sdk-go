# MicrosoftGraphVppToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**AppleId** | **string** | The apple Id associated with the given Apple Volume Purchase Program Token. | [optional] 
**AutomaticallyUpdateApps** | **bool** | Whether or not apps for the VPP token will be automatically updated. | [optional] 
**CountryOrRegion** | **string** | Whether or not apps for the VPP token will be automatically updated. | [optional] 
**ExpirationDateTime** | [**time.Time**](time.Time.md) | The expiration date time of the Apple Volume Purchase Program Token. | [optional] 
**LastModifiedDateTime** | [**time.Time**](time.Time.md) | Last modification date time associated with the Apple Volume Purchase Program Token. | [optional] 
**LastSyncDateTime** | [**time.Time**](time.Time.md) | The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token. | [optional] 
**LastSyncStatus** | [**interface{}**](.md) | Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: &#x60;none&#x60;, &#x60;inProgress&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;. | [optional] 
**OrganizationName** | **string** | The organization associated with the Apple Volume Purchase Program Token | [optional] 
**State** | [**interface{}**](.md) | Current state of the Apple Volume Purchase Program Token. Possible values are: &#x60;unknown&#x60;, &#x60;valid&#x60;, &#x60;expired&#x60;, &#x60;invalid&#x60;, &#x60;assignedToExternalMDM&#x60;. | [optional] 
**Token** | **string** | The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program. | [optional] 
**VppTokenAccountType** | [**interface{}**](.md) | The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: &#x60;business&#x60;, &#x60;education&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


