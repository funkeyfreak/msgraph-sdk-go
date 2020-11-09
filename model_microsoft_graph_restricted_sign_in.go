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

// MicrosoftGraphRestrictedSignIn struct for MicrosoftGraphRestrictedSignIn
type MicrosoftGraphRestrictedSignIn struct {
	Id                               string        `json:"id,omitempty"`
	AppDisplayName                   string        `json:"appDisplayName,omitempty"`
	AppId                            string        `json:"appId,omitempty"`
	AppliedConditionalAccessPolicies []interface{} `json:"appliedConditionalAccessPolicies,omitempty"`
	ClientAppUsed                    string        `json:"clientAppUsed,omitempty"`
	ConditionalAccessStatus          interface{}   `json:"conditionalAccessStatus,omitempty"`
	CorrelationId                    string        `json:"correlationId,omitempty"`
	CreatedDateTime                  time.Time     `json:"createdDateTime,omitempty"`
	DeviceDetail                     interface{}   `json:"deviceDetail,omitempty"`
	IpAddress                        string        `json:"ipAddress,omitempty"`
	IsInteractive                    bool          `json:"isInteractive,omitempty"`
	Location                         interface{}   `json:"location,omitempty"`
	ResourceDisplayName              string        `json:"resourceDisplayName,omitempty"`
	ResourceId                       string        `json:"resourceId,omitempty"`
	RiskDetail                       interface{}   `json:"riskDetail,omitempty"`
	RiskEventTypes                   []interface{} `json:"riskEventTypes,omitempty"`
	RiskEventTypesV2                 []string      `json:"riskEventTypes_v2,omitempty"`
	RiskLevelAggregated              interface{}   `json:"riskLevelAggregated,omitempty"`
	RiskLevelDuringSignIn            interface{}   `json:"riskLevelDuringSignIn,omitempty"`
	RiskState                        interface{}   `json:"riskState,omitempty"`
	Status                           interface{}   `json:"status,omitempty"`
	UserDisplayName                  string        `json:"userDisplayName,omitempty"`
	UserId                           string        `json:"userId,omitempty"`
	UserPrincipalName                string        `json:"userPrincipalName,omitempty"`
	TargetTenantId                   string        `json:"targetTenantId,omitempty"`
}
