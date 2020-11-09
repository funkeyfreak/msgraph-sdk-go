# MicrosoftGraphTermsAndConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**AcceptanceStatement** | **string** | Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&amp;C policy. This is shown to the user on prompts to accept the T&amp;C policy. | [optional] 
**BodyText** | **string** | Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&amp;C policy. | [optional] 
**CreatedDateTime** | [**time.Time**](time.Time.md) | DateTime the object was created. | [optional] 
**Description** | **string** | Administrator-supplied description of the T&amp;C policy. | [optional] 
**DisplayName** | **string** | Administrator-supplied name for the T&amp;C policy.  | [optional] 
**LastModifiedDateTime** | [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**Title** | **string** | Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&amp;C policy. | [optional] 
**Version** | **int32** | Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&amp;C policy. | [optional] 
**AcceptanceStatuses** | [**[]MicrosoftGraphTermsAndConditionsAcceptanceStatus**](microsoft.graph.termsAndConditionsAcceptanceStatus.md) |  | [optional] 
**Assignments** | [**[]MicrosoftGraphTermsAndConditionsAssignment**](microsoft.graph.termsAndConditionsAssignment.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


