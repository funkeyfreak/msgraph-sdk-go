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

// Group Represents an Azure Active Directory object. The directoryObject type is the base type for many other directory entity types.
type Group struct {
	AssignedLabels                []interface{}                        `json:"assignedLabels,omitempty"`
	AssignedLicenses              []interface{}                        `json:"assignedLicenses,omitempty"`
	Classification                string                               `json:"classification,omitempty"`
	CreatedDateTime               time.Time                            `json:"createdDateTime,omitempty"`
	Description                   string                               `json:"description,omitempty"`
	DisplayName                   string                               `json:"displayName,omitempty"`
	ExpirationDateTime            time.Time                            `json:"expirationDateTime,omitempty"`
	GroupTypes                    []string                             `json:"groupTypes,omitempty"`
	HasMembersWithLicenseErrors   bool                                 `json:"hasMembersWithLicenseErrors,omitempty"`
	LicenseProcessingState        interface{}                          `json:"licenseProcessingState,omitempty"`
	Mail                          string                               `json:"mail,omitempty"`
	MailEnabled                   bool                                 `json:"mailEnabled,omitempty"`
	MailNickname                  string                               `json:"mailNickname,omitempty"`
	MembershipRule                string                               `json:"membershipRule,omitempty"`
	MembershipRuleProcessingState string                               `json:"membershipRuleProcessingState,omitempty"`
	OnPremisesDomainName          string                               `json:"onPremisesDomainName,omitempty"`
	OnPremisesLastSyncDateTime    time.Time                            `json:"onPremisesLastSyncDateTime,omitempty"`
	OnPremisesNetBiosName         string                               `json:"onPremisesNetBiosName,omitempty"`
	OnPremisesProvisioningErrors  []interface{}                        `json:"onPremisesProvisioningErrors,omitempty"`
	OnPremisesSamAccountName      string                               `json:"onPremisesSamAccountName,omitempty"`
	OnPremisesSecurityIdentifier  string                               `json:"onPremisesSecurityIdentifier,omitempty"`
	OnPremisesSyncEnabled         bool                                 `json:"onPremisesSyncEnabled,omitempty"`
	PreferredDataLocation         string                               `json:"preferredDataLocation,omitempty"`
	PreferredLanguage             string                               `json:"preferredLanguage,omitempty"`
	ProxyAddresses                []string                             `json:"proxyAddresses,omitempty"`
	RenewedDateTime               time.Time                            `json:"renewedDateTime,omitempty"`
	SecurityEnabled               bool                                 `json:"securityEnabled,omitempty"`
	SecurityIdentifier            string                               `json:"securityIdentifier,omitempty"`
	Theme                         string                               `json:"theme,omitempty"`
	Visibility                    string                               `json:"visibility,omitempty"`
	AllowExternalSenders          bool                                 `json:"allowExternalSenders,omitempty"`
	AutoSubscribeNewMembers       bool                                 `json:"autoSubscribeNewMembers,omitempty"`
	HideFromAddressLists          bool                                 `json:"hideFromAddressLists,omitempty"`
	HideFromOutlookClients        bool                                 `json:"hideFromOutlookClients,omitempty"`
	IsSubscribedByMail            bool                                 `json:"isSubscribedByMail,omitempty"`
	UnseenCount                   int32                                `json:"unseenCount,omitempty"`
	IsArchived                    bool                                 `json:"isArchived,omitempty"`
	AppRoleAssignments            []MicrosoftGraphAppRoleAssignment    `json:"appRoleAssignments,omitempty"`
	CreatedOnBehalfOf             interface{}                          `json:"createdOnBehalfOf,omitempty"`
	MemberOf                      []MicrosoftGraphDirectoryObject      `json:"memberOf,omitempty"`
	Members                       []MicrosoftGraphDirectoryObject      `json:"members,omitempty"`
	MembersWithLicenseErrors      []MicrosoftGraphDirectoryObject      `json:"membersWithLicenseErrors,omitempty"`
	Owners                        []MicrosoftGraphDirectoryObject      `json:"owners,omitempty"`
	Settings                      []MicrosoftGraphGroupSetting         `json:"settings,omitempty"`
	TransitiveMemberOf            []MicrosoftGraphDirectoryObject      `json:"transitiveMemberOf,omitempty"`
	TransitiveMembers             []MicrosoftGraphDirectoryObject      `json:"transitiveMembers,omitempty"`
	AcceptedSenders               []MicrosoftGraphDirectoryObject      `json:"acceptedSenders,omitempty"`
	Calendar                      interface{}                          `json:"calendar,omitempty"`
	CalendarView                  []MicrosoftGraphEvent                `json:"calendarView,omitempty"`
	Conversations                 []MicrosoftGraphConversation         `json:"conversations,omitempty"`
	Events                        []MicrosoftGraphEvent                `json:"events,omitempty"`
	Photo                         interface{}                          `json:"photo,omitempty"`
	Photos                        []MicrosoftGraphProfilePhoto         `json:"photos,omitempty"`
	RejectedSenders               []MicrosoftGraphDirectoryObject      `json:"rejectedSenders,omitempty"`
	Threads                       []MicrosoftGraphConversationThread   `json:"threads,omitempty"`
	Drive                         interface{}                          `json:"drive,omitempty"`
	Drives                        []MicrosoftGraphDrive                `json:"drives,omitempty"`
	Sites                         []MicrosoftGraphSite                 `json:"sites,omitempty"`
	Extensions                    []MicrosoftGraphExtension            `json:"extensions,omitempty"`
	GroupLifecyclePolicies        []MicrosoftGraphGroupLifecyclePolicy `json:"groupLifecyclePolicies,omitempty"`
	Planner                       interface{}                          `json:"planner,omitempty"`
	Onenote                       interface{}                          `json:"onenote,omitempty"`
	Team                          interface{}                          `json:"team,omitempty"`
}
