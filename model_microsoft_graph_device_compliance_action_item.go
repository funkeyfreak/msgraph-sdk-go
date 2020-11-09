/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphDeviceComplianceActionItem struct for MicrosoftGraphDeviceComplianceActionItem
type MicrosoftGraphDeviceComplianceActionItem struct {
	Id string `json:"id,omitempty"`
	// What action to take
	ActionType interface{} `json:"actionType,omitempty"`
	// Number of hours to wait till the action will be enforced. Valid values 0 to 8760
	GracePeriodHours int32 `json:"gracePeriodHours,omitempty"`
	// A list of group IDs to speicify who to CC this notification message to.
	NotificationMessageCCList []string `json:"notificationMessageCCList,omitempty"`
	// What notification Message template to use
	NotificationTemplateId string `json:"notificationTemplateId,omitempty"`
}
