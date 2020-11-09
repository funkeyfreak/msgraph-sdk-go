/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// Domain struct for Domain
type Domain struct {
	AuthenticationType               string                          `json:"authenticationType,omitempty"`
	AvailabilityStatus               string                          `json:"availabilityStatus,omitempty"`
	IsAdminManaged                   bool                            `json:"isAdminManaged,omitempty"`
	IsDefault                        bool                            `json:"isDefault,omitempty"`
	IsInitial                        bool                            `json:"isInitial,omitempty"`
	IsRoot                           bool                            `json:"isRoot,omitempty"`
	IsVerified                       bool                            `json:"isVerified,omitempty"`
	Manufacturer                     string                          `json:"manufacturer,omitempty"`
	Model                            string                          `json:"model,omitempty"`
	PasswordNotificationWindowInDays int32                           `json:"passwordNotificationWindowInDays,omitempty"`
	PasswordValidityPeriodInDays     int32                           `json:"passwordValidityPeriodInDays,omitempty"`
	State                            interface{}                     `json:"state,omitempty"`
	SupportedServices                []string                        `json:"supportedServices,omitempty"`
	DomainNameReferences             []MicrosoftGraphDirectoryObject `json:"domainNameReferences,omitempty"`
	ServiceConfigurationRecords      []MicrosoftGraphDomainDnsRecord `json:"serviceConfigurationRecords,omitempty"`
	VerificationDnsRecords           []MicrosoftGraphDomainDnsRecord `json:"verificationDnsRecords,omitempty"`
}