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

// MicrosoftGraphEducationUser struct for MicrosoftGraphEducationUser
type MicrosoftGraphEducationUser struct {
	Id                             string                          `json:"id,omitempty"`
	AccountEnabled                 bool                            `json:"accountEnabled,omitempty"`
	AssignedLicenses               []MicrosoftGraphAssignedLicense `json:"assignedLicenses,omitempty"`
	AssignedPlans                  []MicrosoftGraphAssignedPlan    `json:"assignedPlans,omitempty"`
	BusinessPhones                 []string                        `json:"businessPhones,omitempty"`
	CreatedBy                      interface{}                     `json:"createdBy,omitempty"`
	Department                     string                          `json:"department,omitempty"`
	DisplayName                    string                          `json:"displayName,omitempty"`
	ExternalSource                 interface{}                     `json:"externalSource,omitempty"`
	GivenName                      string                          `json:"givenName,omitempty"`
	Mail                           string                          `json:"mail,omitempty"`
	MailingAddress                 interface{}                     `json:"mailingAddress,omitempty"`
	MailNickname                   string                          `json:"mailNickname,omitempty"`
	MiddleName                     string                          `json:"middleName,omitempty"`
	MobilePhone                    string                          `json:"mobilePhone,omitempty"`
	OfficeLocation                 string                          `json:"officeLocation,omitempty"`
	PasswordPolicies               string                          `json:"passwordPolicies,omitempty"`
	PasswordProfile                interface{}                     `json:"passwordProfile,omitempty"`
	PreferredLanguage              string                          `json:"preferredLanguage,omitempty"`
	PrimaryRole                    interface{}                     `json:"primaryRole,omitempty"`
	ProvisionedPlans               []MicrosoftGraphProvisionedPlan `json:"provisionedPlans,omitempty"`
	RefreshTokensValidFromDateTime time.Time                       `json:"refreshTokensValidFromDateTime,omitempty"`
	ResidenceAddress               interface{}                     `json:"residenceAddress,omitempty"`
	ShowInAddressList              bool                            `json:"showInAddressList,omitempty"`
	Student                        interface{}                     `json:"student,omitempty"`
	Surname                        string                          `json:"surname,omitempty"`
	Teacher                        interface{}                     `json:"teacher,omitempty"`
	UsageLocation                  string                          `json:"usageLocation,omitempty"`
	UserPrincipalName              string                          `json:"userPrincipalName,omitempty"`
	UserType                       string                          `json:"userType,omitempty"`
	Classes                        []MicrosoftGraphEducationClass  `json:"classes,omitempty"`
	Schools                        []MicrosoftGraphEducationSchool `json:"schools,omitempty"`
	User                           interface{}                     `json:"user,omitempty"`
}
