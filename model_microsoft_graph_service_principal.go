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

// MicrosoftGraphServicePrincipal struct for MicrosoftGraphServicePrincipal
type MicrosoftGraphServicePrincipal struct {
	Id                                 string                                            `json:"id,omitempty"`
	DeletedDateTime                    time.Time                                         `json:"deletedDateTime,omitempty"`
	AccountEnabled                     bool                                              `json:"accountEnabled,omitempty"`
	AddIns                             []MicrosoftGraphAddIn                             `json:"addIns,omitempty"`
	AlternativeNames                   []string                                          `json:"alternativeNames,omitempty"`
	AppDescription                     string                                            `json:"appDescription,omitempty"`
	AppDisplayName                     string                                            `json:"appDisplayName,omitempty"`
	AppId                              string                                            `json:"appId,omitempty"`
	ApplicationTemplateId              string                                            `json:"applicationTemplateId,omitempty"`
	AppOwnerOrganizationId             string                                            `json:"appOwnerOrganizationId,omitempty"`
	AppRoleAssignmentRequired          bool                                              `json:"appRoleAssignmentRequired,omitempty"`
	AppRoles                           []MicrosoftGraphAppRole                           `json:"appRoles,omitempty"`
	Description                        string                                            `json:"description,omitempty"`
	DisplayName                        string                                            `json:"displayName,omitempty"`
	Homepage                           string                                            `json:"homepage,omitempty"`
	Info                               interface{}                                       `json:"info,omitempty"`
	KeyCredentials                     []MicrosoftGraphKeyCredential                     `json:"keyCredentials,omitempty"`
	LoginUrl                           string                                            `json:"loginUrl,omitempty"`
	LogoutUrl                          string                                            `json:"logoutUrl,omitempty"`
	Notes                              string                                            `json:"notes,omitempty"`
	NotificationEmailAddresses         []string                                          `json:"notificationEmailAddresses,omitempty"`
	Oauth2PermissionScopes             []MicrosoftGraphPermissionScope                   `json:"oauth2PermissionScopes,omitempty"`
	PasswordCredentials                []MicrosoftGraphPasswordCredential                `json:"passwordCredentials,omitempty"`
	PreferredSingleSignOnMode          string                                            `json:"preferredSingleSignOnMode,omitempty"`
	PreferredTokenSigningKeyThumbprint string                                            `json:"preferredTokenSigningKeyThumbprint,omitempty"`
	ReplyUrls                          []string                                          `json:"replyUrls,omitempty"`
	SamlSingleSignOnSettings           interface{}                                       `json:"samlSingleSignOnSettings,omitempty"`
	ServicePrincipalNames              []string                                          `json:"servicePrincipalNames,omitempty"`
	ServicePrincipalType               string                                            `json:"servicePrincipalType,omitempty"`
	Tags                               []string                                          `json:"tags,omitempty"`
	TokenEncryptionKeyId               string                                            `json:"tokenEncryptionKeyId,omitempty"`
	AppRoleAssignedTo                  []MicrosoftGraphAppRoleAssignment                 `json:"appRoleAssignedTo,omitempty"`
	AppRoleAssignments                 []MicrosoftGraphAppRoleAssignment                 `json:"appRoleAssignments,omitempty"`
	ClaimsMappingPolicies              []MicrosoftGraphClaimsMappingPolicy               `json:"claimsMappingPolicies,omitempty"`
	CreatedObjects                     []MicrosoftGraphDirectoryObject                   `json:"createdObjects,omitempty"`
	DelegatedPermissionClassifications []MicrosoftGraphDelegatedPermissionClassification `json:"delegatedPermissionClassifications,omitempty"`
	Endpoints                          []MicrosoftGraphEndpoint                          `json:"endpoints,omitempty"`
	HomeRealmDiscoveryPolicies         []MicrosoftGraphHomeRealmDiscoveryPolicy          `json:"homeRealmDiscoveryPolicies,omitempty"`
	MemberOf                           []MicrosoftGraphDirectoryObject                   `json:"memberOf,omitempty"`
	Oauth2PermissionGrants             []MicrosoftGraphOAuth2PermissionGrant             `json:"oauth2PermissionGrants,omitempty"`
	OwnedObjects                       []MicrosoftGraphDirectoryObject                   `json:"ownedObjects,omitempty"`
	Owners                             []MicrosoftGraphDirectoryObject                   `json:"owners,omitempty"`
	TokenIssuancePolicies              []MicrosoftGraphTokenIssuancePolicy               `json:"tokenIssuancePolicies,omitempty"`
	TokenLifetimePolicies              []MicrosoftGraphTokenLifetimePolicy               `json:"tokenLifetimePolicies,omitempty"`
	TransitiveMemberOf                 []MicrosoftGraphDirectoryObject                   `json:"transitiveMemberOf,omitempty"`
}