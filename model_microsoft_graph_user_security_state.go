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

// MicrosoftGraphUserSecurityState struct for MicrosoftGraphUserSecurityState
type MicrosoftGraphUserSecurityState struct {
	AadUserId                    string      `json:"aadUserId,omitempty"`
	AccountName                  string      `json:"accountName,omitempty"`
	DomainName                   string      `json:"domainName,omitempty"`
	EmailRole                    interface{} `json:"emailRole,omitempty"`
	IsVpn                        bool        `json:"isVpn,omitempty"`
	LogonDateTime                time.Time   `json:"logonDateTime,omitempty"`
	LogonId                      string      `json:"logonId,omitempty"`
	LogonIp                      string      `json:"logonIp,omitempty"`
	LogonLocation                string      `json:"logonLocation,omitempty"`
	LogonType                    interface{} `json:"logonType,omitempty"`
	OnPremisesSecurityIdentifier string      `json:"onPremisesSecurityIdentifier,omitempty"`
	RiskScore                    string      `json:"riskScore,omitempty"`
	UserAccountType              interface{} `json:"userAccountType,omitempty"`
	UserPrincipalName            string      `json:"userPrincipalName,omitempty"`
}
