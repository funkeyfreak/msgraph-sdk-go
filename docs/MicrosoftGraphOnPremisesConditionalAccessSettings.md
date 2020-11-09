# MicrosoftGraphOnPremisesConditionalAccessSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Enabled** | **bool** | Indicates if on premises conditional access is enabled for this organization | [optional] 
**ExcludedGroups** | **[]string** | User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy. | [optional] 
**IncludedGroups** | **[]string** | User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access. | [optional] 
**OverrideDefaultRule** | **bool** | Override the default access rule when allowing a device to ensure access is granted. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


