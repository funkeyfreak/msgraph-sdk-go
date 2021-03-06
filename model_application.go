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

// Application Represents an Azure Active Directory object. The directoryObject type is the base type for many other directory entity types.
type Application struct {
	AddIns                     []MicrosoftGraphAddIn                    `json:"addIns,omitempty"`
	Api                        interface{}                              `json:"api,omitempty"`
	AppId                      string                                   `json:"appId,omitempty"`
	ApplicationTemplateId      string                                   `json:"applicationTemplateId,omitempty"`
	AppRoles                   []MicrosoftGraphAppRole                  `json:"appRoles,omitempty"`
	CreatedDateTime            time.Time                                `json:"createdDateTime,omitempty"`
	Description                string                                   `json:"description,omitempty"`
	DisplayName                string                                   `json:"displayName,omitempty"`
	GroupMembershipClaims      string                                   `json:"groupMembershipClaims,omitempty"`
	IdentifierUris             []string                                 `json:"identifierUris,omitempty"`
	Info                       interface{}                              `json:"info,omitempty"`
	IsDeviceOnlyAuthSupported  bool                                     `json:"isDeviceOnlyAuthSupported,omitempty"`
	IsFallbackPublicClient     bool                                     `json:"isFallbackPublicClient,omitempty"`
	KeyCredentials             []MicrosoftGraphKeyCredential            `json:"keyCredentials,omitempty"`
	Logo                       string                                   `json:"logo,omitempty"`
	Notes                      string                                   `json:"notes,omitempty"`
	Oauth2RequirePostResponse  bool                                     `json:"oauth2RequirePostResponse,omitempty"`
	OptionalClaims             interface{}                              `json:"optionalClaims,omitempty"`
	ParentalControlSettings    interface{}                              `json:"parentalControlSettings,omitempty"`
	PasswordCredentials        []MicrosoftGraphPasswordCredential       `json:"passwordCredentials,omitempty"`
	PublicClient               interface{}                              `json:"publicClient,omitempty"`
	PublisherDomain            string                                   `json:"publisherDomain,omitempty"`
	RequiredResourceAccess     []MicrosoftGraphRequiredResourceAccess   `json:"requiredResourceAccess,omitempty"`
	SignInAudience             string                                   `json:"signInAudience,omitempty"`
	Tags                       []string                                 `json:"tags,omitempty"`
	TokenEncryptionKeyId       string                                   `json:"tokenEncryptionKeyId,omitempty"`
	Web                        interface{}                              `json:"web,omitempty"`
	CreatedOnBehalfOf          interface{}                              `json:"createdOnBehalfOf,omitempty"`
	ExtensionProperties        []MicrosoftGraphExtensionProperty        `json:"extensionProperties,omitempty"`
	HomeRealmDiscoveryPolicies []MicrosoftGraphHomeRealmDiscoveryPolicy `json:"homeRealmDiscoveryPolicies,omitempty"`
	Owners                     []MicrosoftGraphDirectoryObject          `json:"owners,omitempty"`
	TokenIssuancePolicies      []MicrosoftGraphTokenIssuancePolicy      `json:"tokenIssuancePolicies,omitempty"`
	TokenLifetimePolicies      []MicrosoftGraphTokenLifetimePolicy      `json:"tokenLifetimePolicies,omitempty"`
}
