/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

import (
	"time"
)

// MicrosoftGraphUser struct for MicrosoftGraphUser
type MicrosoftGraphUser struct {
	Id                              string                          `json:"id,omitempty"`
	DeletedDateTime                 time.Time                       `json:"deletedDateTime,omitempty"`
	AccountEnabled                  bool                            `json:"accountEnabled,omitempty"`
	AgeGroup                        string                          `json:"ageGroup,omitempty"`
	AssignedLicenses                []MicrosoftGraphAssignedLicense `json:"assignedLicenses,omitempty"`
	AssignedPlans                   []MicrosoftGraphAssignedPlan    `json:"assignedPlans,omitempty"`
	BusinessPhones                  []string                        `json:"businessPhones,omitempty"`
	City                            string                          `json:"city,omitempty"`
	CompanyName                     string                          `json:"companyName,omitempty"`
	ConsentProvidedForMinor         string                          `json:"consentProvidedForMinor,omitempty"`
	Country                         string                          `json:"country,omitempty"`
	CreatedDateTime                 time.Time                       `json:"createdDateTime,omitempty"`
	CreationType                    string                          `json:"creationType,omitempty"`
	Department                      string                          `json:"department,omitempty"`
	DisplayName                     string                          `json:"displayName,omitempty"`
	EmployeeId                      string                          `json:"employeeId,omitempty"`
	ExternalUserState               string                          `json:"externalUserState,omitempty"`
	ExternalUserStateChangeDateTime time.Time                       `json:"externalUserStateChangeDateTime,omitempty"`
	FaxNumber                       string                          `json:"faxNumber,omitempty"`
	GivenName                       string                          `json:"givenName,omitempty"`
	Identities                      []interface{}                   `json:"identities,omitempty"`
	ImAddresses                     []string                        `json:"imAddresses,omitempty"`
	IsResourceAccount               bool                            `json:"isResourceAccount,omitempty"`
	JobTitle                        string                          `json:"jobTitle,omitempty"`
	LastPasswordChangeDateTime      time.Time                       `json:"lastPasswordChangeDateTime,omitempty"`
	LegalAgeGroupClassification     string                          `json:"legalAgeGroupClassification,omitempty"`
	LicenseAssignmentStates         []interface{}                   `json:"licenseAssignmentStates,omitempty"`
	Mail                            string                          `json:"mail,omitempty"`
	MailNickname                    string                          `json:"mailNickname,omitempty"`
	MobilePhone                     string                          `json:"mobilePhone,omitempty"`
	OfficeLocation                  string                          `json:"officeLocation,omitempty"`
	OnPremisesDistinguishedName     string                          `json:"onPremisesDistinguishedName,omitempty"`
	OnPremisesDomainName            string                          `json:"onPremisesDomainName,omitempty"`
	OnPremisesExtensionAttributes   interface{}                     `json:"onPremisesExtensionAttributes,omitempty"`
	OnPremisesImmutableId           string                          `json:"onPremisesImmutableId,omitempty"`
	OnPremisesLastSyncDateTime      time.Time                       `json:"onPremisesLastSyncDateTime,omitempty"`
	OnPremisesProvisioningErrors    []interface{}                   `json:"onPremisesProvisioningErrors,omitempty"`
	OnPremisesSamAccountName        string                          `json:"onPremisesSamAccountName,omitempty"`
	OnPremisesSecurityIdentifier    string                          `json:"onPremisesSecurityIdentifier,omitempty"`
	OnPremisesSyncEnabled           bool                            `json:"onPremisesSyncEnabled,omitempty"`
	OnPremisesUserPrincipalName     string                          `json:"onPremisesUserPrincipalName,omitempty"`
	OtherMails                      []string                        `json:"otherMails,omitempty"`
	PasswordPolicies                string                          `json:"passwordPolicies,omitempty"`
	PasswordProfile                 interface{}                     `json:"passwordProfile,omitempty"`
	PostalCode                      string                          `json:"postalCode,omitempty"`
	PreferredLanguage               string                          `json:"preferredLanguage,omitempty"`
	ProvisionedPlans                []MicrosoftGraphProvisionedPlan `json:"provisionedPlans,omitempty"`
	ProxyAddresses                  []string                        `json:"proxyAddresses,omitempty"`
	ShowInAddressList               bool                            `json:"showInAddressList,omitempty"`
	SignInSessionsValidFromDateTime time.Time                       `json:"signInSessionsValidFromDateTime,omitempty"`
	State                           string                          `json:"state,omitempty"`
	StreetAddress                   string                          `json:"streetAddress,omitempty"`
	Surname                         string                          `json:"surname,omitempty"`
	UsageLocation                   string                          `json:"usageLocation,omitempty"`
	UserPrincipalName               string                          `json:"userPrincipalName,omitempty"`
	UserType                        string                          `json:"userType,omitempty"`
	MailboxSettings                 interface{}                     `json:"mailboxSettings,omitempty"`
	// The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
	DeviceEnrollmentLimit                 int32                                                `json:"deviceEnrollmentLimit,omitempty"`
	AboutMe                               string                                               `json:"aboutMe,omitempty"`
	Birthday                              time.Time                                            `json:"birthday,omitempty"`
	HireDate                              time.Time                                            `json:"hireDate,omitempty"`
	Interests                             []string                                             `json:"interests,omitempty"`
	MySite                                string                                               `json:"mySite,omitempty"`
	PastProjects                          []string                                             `json:"pastProjects,omitempty"`
	PreferredName                         string                                               `json:"preferredName,omitempty"`
	Responsibilities                      []string                                             `json:"responsibilities,omitempty"`
	Schools                               []string                                             `json:"schools,omitempty"`
	Skills                                []string                                             `json:"skills,omitempty"`
	AppRoleAssignments                    []MicrosoftGraphAppRoleAssignment                    `json:"appRoleAssignments,omitempty"`
	CreatedObjects                        []MicrosoftGraphDirectoryObject                      `json:"createdObjects,omitempty"`
	DirectReports                         []MicrosoftGraphDirectoryObject                      `json:"directReports,omitempty"`
	LicenseDetails                        []MicrosoftGraphLicenseDetails                       `json:"licenseDetails,omitempty"`
	Manager                               interface{}                                          `json:"manager,omitempty"`
	MemberOf                              []MicrosoftGraphDirectoryObject                      `json:"memberOf,omitempty"`
	Oauth2PermissionGrants                []MicrosoftGraphOAuth2PermissionGrant                `json:"oauth2PermissionGrants,omitempty"`
	OwnedDevices                          []MicrosoftGraphDirectoryObject                      `json:"ownedDevices,omitempty"`
	OwnedObjects                          []MicrosoftGraphDirectoryObject                      `json:"ownedObjects,omitempty"`
	RegisteredDevices                     []MicrosoftGraphDirectoryObject                      `json:"registeredDevices,omitempty"`
	ScopedRoleMemberOf                    []MicrosoftGraphScopedRoleMembership                 `json:"scopedRoleMemberOf,omitempty"`
	TransitiveMemberOf                    []MicrosoftGraphDirectoryObject                      `json:"transitiveMemberOf,omitempty"`
	Calendar                              interface{}                                          `json:"calendar,omitempty"`
	CalendarGroups                        []MicrosoftGraphCalendarGroup                        `json:"calendarGroups,omitempty"`
	Calendars                             []MicrosoftGraphCalendar                             `json:"calendars,omitempty"`
	CalendarView                          []MicrosoftGraphEvent                                `json:"calendarView,omitempty"`
	ContactFolders                        []MicrosoftGraphContactFolder                        `json:"contactFolders,omitempty"`
	Contacts                              []MicrosoftGraphContact                              `json:"contacts,omitempty"`
	Events                                []MicrosoftGraphEvent                                `json:"events,omitempty"`
	InferenceClassification               interface{}                                          `json:"inferenceClassification,omitempty"`
	MailFolders                           []MicrosoftGraphMailFolder                           `json:"mailFolders,omitempty"`
	Messages                              []MicrosoftGraphMessage                              `json:"messages,omitempty"`
	Outlook                               interface{}                                          `json:"outlook,omitempty"`
	People                                []MicrosoftGraphPerson                               `json:"people,omitempty"`
	Photo                                 interface{}                                          `json:"photo,omitempty"`
	Photos                                []MicrosoftGraphProfilePhoto                         `json:"photos,omitempty"`
	Drive                                 interface{}                                          `json:"drive,omitempty"`
	Drives                                []MicrosoftGraphDrive                                `json:"drives,omitempty"`
	FollowedSites                         []MicrosoftGraphSite                                 `json:"followedSites,omitempty"`
	Extensions                            []MicrosoftGraphExtension                            `json:"extensions,omitempty"`
	ManagedDevices                        []MicrosoftGraphManagedDevice                        `json:"managedDevices,omitempty"`
	ManagedAppRegistrations               []MicrosoftGraphManagedAppRegistration               `json:"managedAppRegistrations,omitempty"`
	DeviceManagementTroubleshootingEvents []MicrosoftGraphDeviceManagementTroubleshootingEvent `json:"deviceManagementTroubleshootingEvents,omitempty"`
	Planner                               interface{}                                          `json:"planner,omitempty"`
	Insights                              interface{}                                          `json:"insights,omitempty"`
	Settings                              interface{}                                          `json:"settings,omitempty"`
	Onenote                               interface{}                                          `json:"onenote,omitempty"`
	Activities                            []MicrosoftGraphUserActivity                         `json:"activities,omitempty"`
	OnlineMeetings                        []MicrosoftGraphOnlineMeeting                        `json:"onlineMeetings,omitempty"`
	JoinedTeams                           []MicrosoftGraphTeam                                 `json:"joinedTeams,omitempty"`
	Teamwork                              interface{}                                          `json:"teamwork,omitempty"`
	Todo                                  interface{}                                          `json:"todo,omitempty"`
}