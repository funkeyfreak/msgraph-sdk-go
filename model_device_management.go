/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// DeviceManagement Singleton entity that acts as a container for all device management functionality.
type DeviceManagement struct {
	// Intune Account Id for given tenant
	IntuneAccountId string `json:"intuneAccountId,omitempty"`
	// Account level settings.
	Settings interface{} `json:"settings,omitempty"`
	// intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
	IntuneBrand interface{} `json:"intuneBrand,omitempty"`
	// Tenant mobile device management subscription state.
	SubscriptionState                                    interface{}                                                        `json:"subscriptionState,omitempty"`
	TermsAndConditions                                   []MicrosoftGraphTermsAndConditions                                 `json:"termsAndConditions,omitempty"`
	DeviceCompliancePolicies                             []MicrosoftGraphDeviceCompliancePolicy                             `json:"deviceCompliancePolicies,omitempty"`
	DeviceCompliancePolicyDeviceStateSummary             interface{}                                                        `json:"deviceCompliancePolicyDeviceStateSummary,omitempty"`
	DeviceCompliancePolicySettingStateSummaries          []MicrosoftGraphDeviceCompliancePolicySettingStateSummary          `json:"deviceCompliancePolicySettingStateSummaries,omitempty"`
	DeviceConfigurationDeviceStateSummaries              interface{}                                                        `json:"deviceConfigurationDeviceStateSummaries,omitempty"`
	DeviceConfigurations                                 []MicrosoftGraphDeviceConfiguration                                `json:"deviceConfigurations,omitempty"`
	IosUpdateStatuses                                    []MicrosoftGraphIosUpdateDeviceStatus                              `json:"iosUpdateStatuses,omitempty"`
	SoftwareUpdateStatusSummary                          interface{}                                                        `json:"softwareUpdateStatusSummary,omitempty"`
	ComplianceManagementPartners                         []MicrosoftGraphComplianceManagementPartner                        `json:"complianceManagementPartners,omitempty"`
	ConditionalAccessSettings                            interface{}                                                        `json:"conditionalAccessSettings,omitempty"`
	DeviceCategories                                     []MicrosoftGraphDeviceCategory                                     `json:"deviceCategories,omitempty"`
	DeviceEnrollmentConfigurations                       []MicrosoftGraphDeviceEnrollmentConfiguration                      `json:"deviceEnrollmentConfigurations,omitempty"`
	DeviceManagementPartners                             []MicrosoftGraphDeviceManagementPartner                            `json:"deviceManagementPartners,omitempty"`
	ExchangeConnectors                                   []MicrosoftGraphDeviceManagementExchangeConnector                  `json:"exchangeConnectors,omitempty"`
	MobileThreatDefenseConnectors                        []MicrosoftGraphMobileThreatDefenseConnector                       `json:"mobileThreatDefenseConnectors,omitempty"`
	ApplePushNotificationCertificate                     interface{}                                                        `json:"applePushNotificationCertificate,omitempty"`
	DetectedApps                                         []MicrosoftGraphDetectedApp                                        `json:"detectedApps,omitempty"`
	ManagedDeviceOverview                                interface{}                                                        `json:"managedDeviceOverview,omitempty"`
	ManagedDevices                                       []MicrosoftGraphManagedDevice                                      `json:"managedDevices,omitempty"`
	NotificationMessageTemplates                         []MicrosoftGraphNotificationMessageTemplate                        `json:"notificationMessageTemplates,omitempty"`
	ResourceOperations                                   []MicrosoftGraphResourceOperation                                  `json:"resourceOperations,omitempty"`
	RoleAssignments                                      []MicrosoftGraphDeviceAndAppManagementRoleAssignment               `json:"roleAssignments,omitempty"`
	RoleDefinitions                                      []MicrosoftGraphRoleDefinition                                     `json:"roleDefinitions,omitempty"`
	RemoteAssistancePartners                             []MicrosoftGraphRemoteAssistancePartner                            `json:"remoteAssistancePartners,omitempty"`
	TelecomExpenseManagementPartners                     []MicrosoftGraphTelecomExpenseManagementPartner                    `json:"telecomExpenseManagementPartners,omitempty"`
	TroubleshootingEvents                                []MicrosoftGraphDeviceManagementTroubleshootingEvent               `json:"troubleshootingEvents,omitempty"`
	WindowsInformationProtectionAppLearningSummaries     []MicrosoftGraphWindowsInformationProtectionAppLearningSummary     `json:"windowsInformationProtectionAppLearningSummaries,omitempty"`
	WindowsInformationProtectionNetworkLearningSummaries []MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary `json:"windowsInformationProtectionNetworkLearningSummaries,omitempty"`
}
