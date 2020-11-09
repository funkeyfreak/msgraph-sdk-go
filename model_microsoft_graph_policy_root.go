/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphPolicyRoot struct for MicrosoftGraphPolicyRoot
type MicrosoftGraphPolicyRoot struct {
	Id                                        string                                     `json:"id,omitempty"`
	ActivityBasedTimeoutPolicies              []MicrosoftGraphActivityBasedTimeoutPolicy `json:"activityBasedTimeoutPolicies,omitempty"`
	ClaimsMappingPolicies                     []MicrosoftGraphClaimsMappingPolicy        `json:"claimsMappingPolicies,omitempty"`
	HomeRealmDiscoveryPolicies                []MicrosoftGraphHomeRealmDiscoveryPolicy   `json:"homeRealmDiscoveryPolicies,omitempty"`
	PermissionGrantPolicies                   []MicrosoftGraphPermissionGrantPolicy      `json:"permissionGrantPolicies,omitempty"`
	TokenIssuancePolicies                     []MicrosoftGraphTokenIssuancePolicy        `json:"tokenIssuancePolicies,omitempty"`
	TokenLifetimePolicies                     []MicrosoftGraphTokenLifetimePolicy        `json:"tokenLifetimePolicies,omitempty"`
	ConditionalAccessPolicies                 []MicrosoftGraphConditionalAccessPolicy    `json:"conditionalAccessPolicies,omitempty"`
	IdentitySecurityDefaultsEnforcementPolicy interface{}                                `json:"identitySecurityDefaultsEnforcementPolicy,omitempty"`
}
