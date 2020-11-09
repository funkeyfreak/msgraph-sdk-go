/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphSubscribedSku struct for MicrosoftGraphSubscribedSku
type MicrosoftGraphSubscribedSku struct {
	Id               string                          `json:"id,omitempty"`
	AppliesTo        string                          `json:"appliesTo,omitempty"`
	CapabilityStatus string                          `json:"capabilityStatus,omitempty"`
	ConsumedUnits    int32                           `json:"consumedUnits,omitempty"`
	PrepaidUnits     interface{}                     `json:"prepaidUnits,omitempty"`
	ServicePlans     []MicrosoftGraphServicePlanInfo `json:"servicePlans,omitempty"`
	SkuId            string                          `json:"skuId,omitempty"`
	SkuPartNumber    string                          `json:"skuPartNumber,omitempty"`
}