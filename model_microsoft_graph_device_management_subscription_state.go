/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphDeviceManagementSubscriptionState the model 'MicrosoftGraphDeviceManagementSubscriptionState'
type MicrosoftGraphDeviceManagementSubscriptionState string

// List of microsoft.graph.deviceManagementSubscriptionState
const (
	MICROSOFTGRAPHDEVICEMANAGEMENTSUBSCRIPTIONSTATE_PENDING    MicrosoftGraphDeviceManagementSubscriptionState = "pending"
	MICROSOFTGRAPHDEVICEMANAGEMENTSUBSCRIPTIONSTATE_ACTIVE     MicrosoftGraphDeviceManagementSubscriptionState = "active"
	MICROSOFTGRAPHDEVICEMANAGEMENTSUBSCRIPTIONSTATE_WARNING    MicrosoftGraphDeviceManagementSubscriptionState = "warning"
	MICROSOFTGRAPHDEVICEMANAGEMENTSUBSCRIPTIONSTATE_DISABLED   MicrosoftGraphDeviceManagementSubscriptionState = "disabled"
	MICROSOFTGRAPHDEVICEMANAGEMENTSUBSCRIPTIONSTATE_DELETED    MicrosoftGraphDeviceManagementSubscriptionState = "deleted"
	MICROSOFTGRAPHDEVICEMANAGEMENTSUBSCRIPTIONSTATE_BLOCKED    MicrosoftGraphDeviceManagementSubscriptionState = "blocked"
	MICROSOFTGRAPHDEVICEMANAGEMENTSUBSCRIPTIONSTATE_LOCKED_OUT MicrosoftGraphDeviceManagementSubscriptionState = "lockedOut"
)
