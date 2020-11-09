# MicrosoftGraphTargetedManagedAppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**CreatedDateTime** | [**time.Time**](time.Time.md) | The date and time the policy was created. | [optional] 
**Description** | **string** | The policy&#39;s description. | [optional] 
**DisplayName** | **string** | Policy display name. | [optional] 
**LastModifiedDateTime** | [**time.Time**](time.Time.md) | Last time the policy was modified. | [optional] 
**Version** | **string** | Version of the entity. | [optional] 
**CustomSettings** | [**[]MicrosoftGraphKeyValuePair**](microsoft.graph.keyValuePair.md) | A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service | [optional] 
**DeployedAppCount** | **int32** | Count of apps to which the current policy is deployed. | [optional] 
**IsAssigned** | **bool** | Indicates if the policy is deployed to any inclusion groups or not. | [optional] 
**Apps** | [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**Assignments** | [**[]MicrosoftGraphTargetedManagedAppPolicyAssignment**](microsoft.graph.targetedManagedAppPolicyAssignment.md) |  | [optional] 
**DeploymentSummary** | [**interface{}**](.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


