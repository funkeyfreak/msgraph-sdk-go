# MicrosoftGraphRoleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Description** | **string** | Description of the Role definition. | [optional] 
**DisplayName** | **string** | Display Name of the Role definition. | [optional] 
**IsBuiltIn** | **bool** | Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition. | [optional] 
**RolePermissions** | **[]interface{}** | List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission. | [optional] 
**RoleAssignments** | [**[]MicrosoftGraphRoleAssignment**](microsoft.graph.roleAssignment.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


