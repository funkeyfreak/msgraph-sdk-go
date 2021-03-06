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

// MicrosoftGraphOrgContact struct for MicrosoftGraphOrgContact
type MicrosoftGraphOrgContact struct {
	Id                           string                          `json:"id,omitempty"`
	DeletedDateTime              time.Time                       `json:"deletedDateTime,omitempty"`
	Addresses                    []interface{}                   `json:"addresses,omitempty"`
	CompanyName                  string                          `json:"companyName,omitempty"`
	Department                   string                          `json:"department,omitempty"`
	DisplayName                  string                          `json:"displayName,omitempty"`
	GivenName                    string                          `json:"givenName,omitempty"`
	JobTitle                     string                          `json:"jobTitle,omitempty"`
	Mail                         string                          `json:"mail,omitempty"`
	MailNickname                 string                          `json:"mailNickname,omitempty"`
	OnPremisesLastSyncDateTime   time.Time                       `json:"onPremisesLastSyncDateTime,omitempty"`
	OnPremisesProvisioningErrors []interface{}                   `json:"onPremisesProvisioningErrors,omitempty"`
	OnPremisesSyncEnabled        bool                            `json:"onPremisesSyncEnabled,omitempty"`
	Phones                       []interface{}                   `json:"phones,omitempty"`
	ProxyAddresses               []string                        `json:"proxyAddresses,omitempty"`
	Surname                      string                          `json:"surname,omitempty"`
	DirectReports                []MicrosoftGraphDirectoryObject `json:"directReports,omitempty"`
	Manager                      interface{}                     `json:"manager,omitempty"`
	MemberOf                     []MicrosoftGraphDirectoryObject `json:"memberOf,omitempty"`
	TransitiveMemberOf           []MicrosoftGraphDirectoryObject `json:"transitiveMemberOf,omitempty"`
}
